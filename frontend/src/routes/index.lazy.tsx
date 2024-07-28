import { createLazyFileRoute } from '@tanstack/react-router'
import { Button } from "@/components/ui/button"

export const Route = createLazyFileRoute('/')({
  component: Index,
})

export function Index() {
  return (
    <>
      <div className="flex items-center">
        <h1 className="text-lg font-semibold md:text-2xl">Dashboard</h1>
      </div>
    </>
  )
}
