"use client"

import { type RowData, type ReactTable } from "@tanstack/react-table"

import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table"

import { type DataTableFeatures } from "./data-table-features"
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
import { System } from "@wailsio/runtime"

interface DataTableProps<TData extends RowData> {
  table: ReactTable<DataTableFeatures, TData>
  canNextPage: boolean
  goToNextPage: () => Promise<void>
}

export function DataTable<TData extends RowData>({
  table,
  canNextPage,
  goToNextPage,
}: DataTableProps<TData>) {
  return (
    <div className="flex min-h-0 flex-1 flex-col">
      <div className="min-h-0 flex-1 overflow-auto rounded-lg border">
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
      <div className="flex justify-end gap-2 py-4">
        <Field orientation="horizontal" className="w-fit">
          {System.IsWindows() ? (
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
    </div>
  )
}
