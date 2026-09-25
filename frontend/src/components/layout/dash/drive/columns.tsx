import { createColumnHelper } from "@tanstack/react-table"
import type { DataTableFeatures } from "@/components/layout/data-table-features.ts"
import type { Entry } from "@bindings/github.com/newstatue/evorsio/internal/drive"
import {getDriveIcon} from "@/components/layout/dash/drive/drive-icon.tsx";

const columnHelper = createColumnHelper<DataTableFeatures, Entry>()

export const columns = columnHelper.columns([
  columnHelper.accessor("Mime", {
    header: "",
    size: 40,
    cell: ({ row }) => {
      const entry = row.original
      const Icon = getDriveIcon(entry.Mime, entry.Type)

      return <Icon className="size-4" />
    },
  }),

  columnHelper.accessor("Name", {
    header: "名称",
    size: 240,
  }),

])