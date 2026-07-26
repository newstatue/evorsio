import { StrictMode } from "react"
import { createRoot } from "react-dom/client"

import "./index.css"
import { ThemeProvider } from "@/components/theme-provider.tsx"

import { router } from "./router";
import {RouterProvider} from "react-router-dom";
import {QueryClient, QueryClientProvider} from "@tanstack/react-query";
import { Toaster } from "sonner";

const queryClient = new QueryClient()


createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <ThemeProvider>
        <QueryClientProvider client={queryClient}>
            <RouterProvider router={router} />
            <Toaster/>
        </QueryClientProvider>
    </ThemeProvider>
  </StrictMode>
)
