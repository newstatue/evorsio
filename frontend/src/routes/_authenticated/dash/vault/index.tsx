import { createFileRoute } from "@tanstack/react-router"
import {DataTable, DataTablePagination, DataTablePanel, DataTableToolbar} from "@/components/layout/data-table.tsx";
import {columns} from "@/components/layout/dash/vault/columns.tsx";
import {CreateEntry, ListEntries} from "@bindings/github.com/newstatue/evorsio/internal/vault/service.ts";
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
    Field, FieldGroup,
    FieldLabel,
} from "@/components/ui/field.tsx";

import {Button} from "@/components/ui/button.tsx";

import {CopyInput} from "@/components/layout/copy-input.tsx";
import {PasswordInput} from "@/components/layout/copy-password-input.tsx";
import {Input} from "@/components/ui/input.tsx";
import {Plus} from "lucide-react";
import {toast} from "@/components/ui/toast.tsx";

export const Route = createFileRoute("/_authenticated/dash/vault/")({
    component: RouteComponent,
    staticData: {
        breadcrumb: "密码",
    },
})

function RouteComponent() {
    type SheetMode =  "edit" | "create"

    const [mode, setMode] = useState<SheetMode>("edit")
    const [form, setForm] = useState({
        title: "",
        url: "",
        username: "",
        password: "",
        notes: "",
    })

    const [selectedEntry, setSelectedEntry] = useState<Entry | null>(null)
    const [open, setOpen] = useState(false)

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

    const handleSubmit = async () => {
        switch (mode) {
            case "edit":
                if(!selectedEntry) return
                break
            case "create":
                try {
                    await CreateEntry({
                        Title: form.title,
                        Notes: form.notes,
                        Password: form.password,
                        URL: form.url,
                        Username: form.username,
                    })

                    toast.add({
                        title: "创建成功",
                        description: "密码已保存",
                        type: "success",
                    })
                    setOpen(false)
                } catch (err) {
                    toast.add({
                        title: "创建失败",
                        description: (err as Error).message,
                        type: "error",
                    })
                    setOpen(false)
                }
                break
        }
    }


    return (
        <div>
            <DataTable>
                <DataTableToolbar>
                    <div className="flex flex-1 items-center gap-2">
                        <Input
                            placeholder="搜索密码..."
                            className="max-w-sm"
                        />
                    </div>

                    <div className="flex items-center gap-2">
                        <Button variant="outline">
                            筛选
                        </Button>

                        <Button onClick={() => {
                            setSelectedEntry(null)
                            setForm({
                                title: "",
                                url: "",
                                username: "",
                                password: "",
                                notes: "",
                            })
                            setMode("create")
                            setOpen(true)
                        }}>
                            新建
                            <Plus data-icon="inline-end"/>
                        </Button>
                    </div>
                </DataTableToolbar>
                <DataTablePanel table={table} onRowClick={(entry) => {
                    setSelectedEntry(entry)
                    setForm({
                        title: entry.Payload?.title ?? "",
                        url: entry.Payload?.url ?? "",
                        username: entry.Payload?.username ?? "",
                        password: entry.Payload?.password ?? "",
                        notes: entry.Payload?.notes ?? "",
                    })
                    setMode("edit")
                    setOpen(true)
                }}/>
                <DataTablePagination table={table} canNextPage={canNextPage} goToNextPage={goToNextPage}/>
            </DataTable>
            <Sheet open={open} onOpenChange={setOpen}>
                <SheetContent showCloseButton={false}>
                    <form onSubmit={(e)=>{
                        e.preventDefault()
                        handleSubmit()
                    }}>
                        <SheetHeader>
                            <SheetTitle>
                                {mode === "create" ? "新建密码" : "编辑密码"}
                            </SheetTitle>
                        </SheetHeader>
                        <FieldGroup className="px-6">
                            <Field>
                                <FieldLabel htmlFor="title">标题</FieldLabel>
                                <CopyInput id="title" value={form.title} onChange={(v) =>
                                    setForm((prev) => ({
                                        ...prev,
                                        title: v,
                                    }))
                                } label="标题"/>
                            </Field>
                            <Field>
                                <FieldLabel htmlFor="url">URL</FieldLabel>
                                <CopyInput id="url" value={form.url} onChange={(v) =>
                                    setForm((prev) => ({
                                        ...prev,
                                        url:v,
                                    }))
                                } label="URL"/>
                            </Field>
                            <Field>
                                <FieldLabel htmlFor="username">用户名</FieldLabel>
                                <CopyInput id="username" value={form.username} onChange={(v) =>
                                    setForm((prev) => ({
                                        ...prev,
                                        username:v,
                                    }))
                                } label="用户名"/>
                            </Field>
                            <Field>
                                <FieldLabel htmlFor="password">密码</FieldLabel>
                                <PasswordInput id="password" value={form.password}  onChange={(v) =>
                                    setForm((prev) => ({
                                        ...prev,
                                        password:v,
                                    }))
                                } label="密码"/>
                            </Field>
                            <Field>
                                <FieldLabel htmlFor="notes">备注</FieldLabel>
                                <CopyInput id="notes" value={form.notes} onChange={(v) =>
                                    setForm((prev) => ({
                                        ...prev,
                                        notes:v,
                                    }))
                                } label="备注"/>
                            </Field>
                        </FieldGroup>
                        <SheetFooter>
                            <Button type="submit">保存</Button>
                            <SheetClose render={<Button variant="outline">取消</Button>} />
                        </SheetFooter>
                    </form>
                </SheetContent>
            </Sheet>
        </div>
    )
}
