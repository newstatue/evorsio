import { createFileRoute } from "@tanstack/react-router"
import { DataTable } from "@/components/layout/data-table.tsx"
import { columns } from "@/components/layout/dash/drive/columns"
import { ListEntries } from "@bindings/github.com/newstatue/evorsio/internal/drive/service.ts"
import { useState } from "react"
import {
  type PaginationState,
  type SortingState,
  useTable,
} from "@tanstack/react-table"
import { useInfiniteQuery } from "@tanstack/react-query"
import { features } from "@/components/layout/data-table-features.ts"

export const Route = createFileRoute("/dash/drive/")({
  component: RouteComponent,
  staticData: {
    breadcrumb: "存储",
  },
})

function RouteComponent() {
  const [globalFilter, setGlobalFilter] = useState("")
  const [sorting, setSorting] = useState<SortingState>([])
  const [pagination, setPagination] = useState<PaginationState>({
    pageIndex: 0,
    pageSize: 10,
  })

  const dataQuery = useInfiniteQuery({
    queryKey: ["drive-entries", pagination.pageSize, sorting, globalFilter],
    queryFn: async ({ pageParam }) => {
      return await ListEntries({
        Dir: "/",
        Name: globalFilter,
        Cursor: pageParam,
        Size: pagination.pageSize,
      })
    },
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

  console.log({
    currentPage,
    hasCachedNextPage,
    canNextPage,
  })

  return (
    <div className="flex min-h-0 flex-1 flex-col">
      <DataTable
        table={table}
        canNextPage={canNextPage}
        goToNextPage={goToNextPage}
      />
    </div>
  )
}
