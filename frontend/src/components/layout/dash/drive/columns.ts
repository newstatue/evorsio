import { createColumnHelper } from "@tanstack/react-table"
import type { DataTableFeatures } from "@/components/layout/data-table-features.ts"
import type { Entry } from "@bindings/github.com/newstatue/evorsio/internal/drive"

const columnHelper = createColumnHelper<DataTableFeatures, Entry>()

export const columns = columnHelper.columns([
  columnHelper.accessor("Name", {
    header: "名称",
  }),
  columnHelper.accessor("Type", {
    header: "类别",
  }),
  columnHelper.accessor("Path", {
    header: "路径",
  }),
])
