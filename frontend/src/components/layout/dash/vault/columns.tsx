import { createColumnHelper } from "@tanstack/react-table"
import type { DataTableFeatures } from "@/components/layout/data-table-features.ts"
import type { Entry } from "@bindings/github.com/newstatue/evorsio/internal/vault"

const columnHelper = createColumnHelper<DataTableFeatures, Entry>()

export const columns = columnHelper.columns([
    columnHelper.accessor("Payload.title", {
        header: "标题",
    }),

    columnHelper.accessor("Payload.username", {
        header: "用户名",
    }),

    columnHelper.accessor("Payload.url", {
        header: "网址",
    }),
])