import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/dash/drive/')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/dash/drive/"!</div>
}
