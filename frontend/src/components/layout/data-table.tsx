"use client"

import {type RowData, type ReactTable} from "@tanstack/react-table"

import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table"

import {type DataTableFeatures} from "./data-table-features"
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select.tsx"
import { Field, FieldLabel } from "@/components/ui/field.tsx"
import {
  Pagination,
  PaginationContent,
  PaginationItem,
  PaginationNext,
  PaginationPrevious,
} from "@/components/ui/pagination.tsx"
import {
  NativeSelect,
  NativeSelectOption,
} from "@/components/ui/native-select.tsx"
import {platform} from "@/lib/platform.ts";
import type {ReactNode} from "react";

export function DataTable({ children }: { children: ReactNode }) {
  return (
      <div className="flex flex-col gap-4">
        {children}
      </div>
  )
}


export function DataTableToolbar({ children }: { children?: ReactNode }) {
  return (
      <div className="flex items-center justify-between gap-2">
        {children}
      </div>
  )
}

interface DataTableProps<TData extends RowData> {
  table: ReactTable<DataTableFeatures, TData>
  onRowClick?: (data: TData) => void
}

export function DataTablePanel<TData extends RowData>({
                                                        table,
                                                        onRowClick
                                                      }: DataTableProps<TData>) {
  return (
        <div className="overflow-x-auto rounded-lg border">
          <Table className="table-fixed">
            <TableHeader className="sticky top-0 z-10 bg-muted">
              {table.getHeaderGroups().map((headerGroup) => (
                  <TableRow key={headerGroup.id}>
                    {headerGroup.headers.map((header) => {
                      return (
                          <TableHead
                              key={header.id}
                              style={{
                                width: header.getSize(),
                              }}
                          >
                            {header.isPlaceholder ? null : (
                                <table.FlexRender header={header} />
                            )}
                          </TableHead>
                      )
                    })}
                  </TableRow>
              ))}
            </TableHeader>
            <TableBody>
              {table.getRowModel().rows?.length ? (
                  table.getRowModel().rows.map((row) => (
                      <TableRow
                          key={row.id}
                          data-state={row.getIsSelected() && "selected"}
                          className={onRowClick ? "cursor-pointer" : undefined}
                          onClick={() => onRowClick?.(row.original)}
                      >
                        {row.getVisibleCells().map((cell) => (
                            <TableCell
                                className="truncate"
                                key={cell.id}
                                style={{
                                  width: cell.column.getSize(),
                                }}
                            >
                              <table.FlexRender cell={cell} />
                            </TableCell>
                        ))}
                      </TableRow>
                  ))
              ) : (
                  <TableRow>
                    <TableCell
                        colSpan={table.getVisibleLeafColumns().length}
                        className="h-24 text-center"
                    >
                      No results.
                    </TableCell>
                  </TableRow>
              )}
            </TableBody>
          </Table>
        </div>
  )
}

interface DataTablePaginationProps<TData extends RowData> {
  table: ReactTable<DataTableFeatures, TData>
  canNextPage: boolean
  goToNextPage: () => Promise<void>
}

export function DataTablePagination<TData extends RowData>({
                                                             table,
                                                             canNextPage,
                                                             goToNextPage,
                                                           }: DataTablePaginationProps<TData>) {
  return (
      <div className="flex justify-end gap-2 py-4">
        <Field orientation="horizontal" className="w-fit">
          {platform.isWindowsLike() ? (
              <Select
                  value={String(table.state.pagination.pageSize)}
                  onValueChange={(value) => {
                    table.setPageSize(Number(value))
                  }}
              >
                <SelectTrigger id="select-rows-per-page">
                  <SelectValue />
                </SelectTrigger>
                <SelectContent align="start">
                  <SelectGroup>
                    <SelectItem value="10">10</SelectItem>
                    <SelectItem value="25">25</SelectItem>
                    <SelectItem value="50">50</SelectItem>
                    <SelectItem value="100">100</SelectItem>
                  </SelectGroup>
                </SelectContent>
                <FieldLabel htmlFor="select-rows-per-page">/页</FieldLabel>
              </Select>
          ) : (
              <>
                <NativeSelect
                    value={String(table.state.pagination.pageSize)}
                    onChange={(event) => {
                      table.setPageSize(Number(event.target.value))
                    }}
                >
                  <NativeSelectOption value="10">10</NativeSelectOption>
                  <NativeSelectOption value="25">25</NativeSelectOption>
                  <NativeSelectOption value="50">50</NativeSelectOption>
                  <NativeSelectOption value="100">100</NativeSelectOption>
                </NativeSelect>
                <FieldLabel htmlFor="select-rows-per-page">/页</FieldLabel>
              </>
          )}
        </Field>
        <Pagination className="mx-0 w-auto">
          <PaginationContent>
            <PaginationItem>
              <PaginationPrevious
                  onClick={() => table.previousPage()}
                  aria-disabled={!table.getCanPreviousPage()}
                  className={
                    !table.getCanPreviousPage()
                        ? "pointer-events-none opacity-50"
                        : undefined
                  }
                  text="上一页"
              />
            </PaginationItem>
            <PaginationItem>
              <PaginationNext
                  onClick={goToNextPage}
                  aria-disabled={!canNextPage}
                  className={
                    !canNextPage ? "pointer-events-none opacity-50" : undefined
                  }
                  text="下一页"
              />
            </PaginationItem>
          </PaginationContent>
        </Pagination>
      </div>
  )}
