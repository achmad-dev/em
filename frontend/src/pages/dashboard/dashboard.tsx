/*
--- MIT License (c) 2024 achmad
--- See LICENSE for more details
*/

import HrDashboard from "@/components/template/dashboard";
import useStore from "@/store/useStore";
import { Toaster } from "@/components/ui/toaster";

export default function Dashboard() {
    const { role } = useStore();
    return (
        <>
            <HrDashboard role={role ? role : ""}/>
            <Toaster />
        </>
    );
}