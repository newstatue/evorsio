import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/dash/vault/')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/dash/vault/"!</div>
}
