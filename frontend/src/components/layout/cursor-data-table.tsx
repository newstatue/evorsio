import { useState } from "react"
import {
    type ColumnDef,
    type PaginationState, type RowData,
    type SortingState,
    useTable,
} from "@tanstack/react-table"
import { useInfiniteQuery } from "@tanstack/react-query"

import { features } from "@/components/layout/data-table-features.ts"
import type {PageResult} from "@bindings/github.com/newstatue/evorsio/internal/common";

type UseCursorDataTableOptions<TData extends RowData> = {
    queryKey: string
    columns: ColumnDef<typeof features, TData, any>[]

    queryFn: (params: {
        cursor: string
        size: number
        sorting: SortingState
        globalFilter: string
    }) => Promise<PageResult<TData|null>>

    initialPageSize?: number
}


export function useCursorDataTable<TData extends RowData>({
                                              queryKey,
                                              columns,
                                              queryFn,
                                              initialPageSize = 10,
                                          }: UseCursorDataTableOptions<TData>) {
    const [globalFilter, setGlobalFilter] = useState("")
    const [sorting, setSorting] = useState<SortingState>([])
    const [pagination, setPagination] = useState<PaginationState>({
        pageIndex: 0,
        pageSize: initialPageSize,
    })

    const dataQuery = useInfiniteQuery({
        queryKey: [queryKey, pagination.pageSize, sorting, globalFilter],
        queryFn: ({ pageParam }) =>
            queryFn({
                cursor: pageParam,
                size: pagination.pageSize,
                sorting,
                globalFilter,
            }),
        initialPageParam: "",
        getNextPageParam: (lastPage) => lastPage.NextCursor,
    })

    const currentPage = dataQuery.data?.pages[pagination.pageIndex]
    const hasCachedNextPage = Boolean(
        dataQuery.data?.pages[pagination.pageIndex + 1]
    )
    const canNextPage = hasCachedNextPage || Boolean(currentPage?.HasMore)

    const table = useTable(
        {
            features,
            columns,
            data: currentPage?.Items?.filter((item) => item !== null) ?? [],
            pageCount: -1,
            state: { sorting, globalFilter, pagination },
            onSortingChange: (updater) => {
                setSorting(updater)
                setPagination((previous) => ({ ...previous, pageIndex: 0 }))
            },
            onGlobalFilterChange: (updater) => {
                setGlobalFilter(updater)
                setPagination((previous) => ({ ...previous, pageIndex: 0 }))
            },
            onPaginationChange: setPagination,
            manualFiltering: true,
            manualSorting: true,
            manualPagination: true,
        },
        (state) => state
    )

    async function goToNextPage() {
        const nextPageIndex = pagination.pageIndex + 1

        if (!dataQuery.data?.pages[nextPageIndex]) {
            const result = await dataQuery.fetchNextPage()
            if (!result.data?.pages[nextPageIndex]) return
        }

        table.nextPage()
    }

    return {
        table,
        dataQuery,
        currentPage,
        canNextPage,
        goToNextPage,
        globalFilter,
        sorting,
        pagination,
    }
}