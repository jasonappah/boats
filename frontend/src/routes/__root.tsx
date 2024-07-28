import { createRootRoute, Link, LinkProps, Outlet } from '@tanstack/react-router'
import { TanStackRouterDevtools } from '@tanstack/router-devtools'
import {
  Menu,
  Ship,
} from "lucide-react"
import { Button } from "@/components/ui/button"
import { clsx } from 'clsx'
import { Sheet, SheetContent, SheetTrigger } from "@/components/ui/sheet"
import dynamicIconImports from 'lucide-react/dynamicIconImports';
import { lazy, Suspense } from 'react'
import { Toaster } from "@/components/ui/sonner"
import { toast } from "sonner"
import {
  QueryClient,
  QueryClientProvider,
} from '@tanstack/react-query'
import { ThemeProvider } from '@/components/theme-provider'
import { ModeToggle } from '@/components/theme-toggle'

const queryClient = new QueryClient()
queryClient.setDefaultOptions({
  mutations: {
    onSuccess: (data, variables, context) => {
      toast.success('Success!')
    },
    onError: (error, variables, context) => {
      toast.error(error.message)
    },
  }
})

export const Route = createRootRoute({
  component: () => (
    <ThemeProvider
      attribute="class"
      defaultTheme="system"
      enableSystem
    >
      <QueryClientProvider client={queryClient}>
        <div className="grid min-h-screen w-full md:grid-cols-[220px_1fr] lg:grid-cols-[280px_1fr]">
          <div className="hidden border-r bg-muted/40 md:block">
            <div className="flex h-full max-h-screen flex-col gap-2">
              <div className="flex h-14 items-center border-b px-4 lg:h-[60px] lg:px-6">
                <a href="/" className="flex items-center gap-2 font-semibold">
                  <Ship className="h-6 w-6" />
                  <span className="">Boats</span>
                </a>
              </div>
              <div className="flex-1">
                <NavItems/>
              </div>
            </div>
          </div>
          <div className="flex flex-col">
            <header className="flex h-14 items-center gap-4 border-b bg-muted/40 px-4 lg:h-[60px] lg:px-6">
              <Sheet>
                <SheetTrigger asChild>
                  <Button
                    variant="outline"
                    size="icon"
                    className="shrink-0 md:hidden"
                  >
                    <Menu className="h-5 w-5" />
                    <span className="sr-only">Toggle navigation menu</span>
                  </Button>
                </SheetTrigger>
                <SheetContent side="left" className="flex flex-col">
                  <NavItems/>
                </SheetContent>
              </Sheet>
              <ModeToggle />
            </header>
            <main className="flex flex-1 flex-col gap-4 p-4 lg:gap-6 lg:p-6">
              <Outlet />
            </main>
          </div>
        </div>  
        <Toaster richColors/>
        <TanStackRouterDevtools />
      </QueryClientProvider>
    </ThemeProvider>
  ),
})

function NavItem(props: { name: string, icon: (keyof (typeof dynamicIconImports)), to: LinkProps["to"] }) {
  const {name, icon, to} = props
  const LucideIcon = lazy(dynamicIconImports[icon]);
  const iconClasses = "sm:h-5 h-6 sm:w-5 w-6"
  return (
    <Link
        to={to}
        className={clsx("mx-[-0.65rem] flex items-center gap-4 sm:gap-3 rounded-xl sm:rounded-lg px-3 py-2 text-muted-foreground hover:text-foreground sm:hover:text-primary transition-all [&.active]:bg-muted")}
    >
      <Suspense fallback={<span className={iconClasses} />}>
        <LucideIcon className={iconClasses} />
      </Suspense>
      {name}
    </Link>
  )
}

function NavItems() {
    return (
      <nav className="grid gap-2 text-lg sm:text-sm font-medium sm:items-start sm:px-2 lg:px-4">
        <Link
            to="/"
            className="sm:hidden flex items-center gap-2 text-lg font-semibold"
        >
            <Ship className="h-6 w-6" />
            <span className="sr-only">Boats</span>
        </Link>
        <NavItem name="Dashboard" to="/" icon="house" />
        <NavItem name="Links" to="/links" icon="link" />
    </nav>
  )
}

