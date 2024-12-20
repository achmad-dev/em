/*
--- MIT License (c) 2024 achmad
--- See LICENSE for more details
*/

import Dashboard from "@/components/template/dashboard";
import useStore from "@/store/useStore";
import { Toaster } from "@/components/ui/toaster";

export default function DashboardLayout() {
    const { role } = useStore();
    return (
        <>
            <Dashboard role={role ? role : ""}/>
            <Toaster />
        </>
    );
}