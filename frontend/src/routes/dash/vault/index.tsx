import { createFileRoute } from "@tanstack/react-router"
import {DataTable} from "@/components/layout/data-table.tsx";
import {columns} from "@/components/layout/dash/vault/columns.tsx";
import {ListEntries} from "@bindings/github.com/newstatue/evorsio/internal/vault/service.ts";
import {useCursorDataTable} from "@/components/layout/cursor-data-table.tsx";
import {useState} from "react";
import type {Entry} from "@bindings/github.com/newstatue/evorsio/internal/vault";
import {
    Sheet,
    SheetClose,
    SheetContent,
    SheetFooter,
    SheetHeader,
    SheetTitle
} from "@/components/ui/sheet.tsx";
import {
    Field,
    FieldGroup,
    FieldLabel,
} from "@/components/ui/field.tsx";

import {Button} from "@/components/ui/button.tsx";

import {CopyInput} from "@/components/layout/copy-input.tsx";
import {PasswordInput} from "@/components/layout/copy-password-input.tsx";



export const Route = createFileRoute("/dash/vault/")({
    component: RouteComponent,
    staticData: {
        breadcrumb: "密码",
    },
})

function RouteComponent() {

    const [selectedEntry, setSelectedEntry] = useState<Entry | null>(null)
    const [open, setOpen] = useState(false)

    // const create = async () => {
    //   await CreateEntry({
    //     Title:"hello", Notes: "beizhu", Password: "pass", URL: "https://github.com", Username: "wyx"
    //   })
    // }
    //
    //
    //
    const {table, canNextPage, goToNextPage} = useCursorDataTable({
        columns,
        queryKey: "vault-entries",
        queryFn: ({ cursor, size }) =>
            ListEntries({
                Cursor: cursor,
                Host: "",
                Size: size,
            }),
    })

    return (
        <div className="flex min-h-0 flex-1 flex-col">
            <DataTable
                table={table}
                canNextPage={canNextPage}
                goToNextPage={goToNextPage}
                onRowClick={(entry) => {
                    setSelectedEntry(entry)
                    setOpen(true)
                }}
            />
            <Sheet open={open} onOpenChange={setOpen}>
                <SheetContent showCloseButton={false}>
                    <SheetHeader>
                        <SheetTitle>查看密码</SheetTitle>
                    </SheetHeader>
                    <FieldGroup className="px-6">
                        <Field>
                            <FieldLabel htmlFor="title">标题</FieldLabel>
                            <CopyInput id="title" value={selectedEntry?.Payload?.title} label="标题"/>
                        </Field>
                        <Field>
                            <FieldLabel htmlFor="url">URL</FieldLabel>
                            <CopyInput id="url" value={selectedEntry?.Payload?.url} label="URL"/>
                        </Field>
                        <Field>
                            <FieldLabel htmlFor="username">用户名</FieldLabel>
                            <CopyInput id="username" value={selectedEntry?.Payload?.username} label="用户名"/>
                        </Field>
                        <Field>
                            <FieldLabel htmlFor="password">密码</FieldLabel>
                            <PasswordInput id="password" value={selectedEntry?.Payload?.password} label="密码"/>
                        </Field>
                        <Field>
                            <FieldLabel htmlFor="notes">备注</FieldLabel>
                            <CopyInput id="notes" value={selectedEntry?.Payload?.notes} label="备注"/>
                        </Field>
                    </FieldGroup>
                    <SheetFooter>
                        <Button type="submit">保存</Button>
                        <SheetClose render={<Button variant="outline">取消</Button>} />
                    </SheetFooter>
                </SheetContent>
            </Sheet>

        </div>
    )
}
