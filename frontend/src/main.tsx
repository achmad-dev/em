/*
--- MIT License (c) 2024 achmad
--- See LICENSE for more details
*/
import { createRoot } from 'react-dom/client'
import { BrowserRouter } from 'react-router'
import { QueryClient, QueryClientProvider } from 'react-query';
import './index.css'
import AppRoutes from './pages/route.tsx'

const queryClient = new QueryClient();


createRoot(document.getElementById('root')!).render(
  <QueryClientProvider client={queryClient}>
    <BrowserRouter>
      <AppRoutes />
    </BrowserRouter>
  </QueryClientProvider>,
)
