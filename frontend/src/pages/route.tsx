/*
--- MIT License (c) 2024 achmad
--- See LICENSE for more details
*/
import SignInPage from "./signin/signin";
import { Routes, Route, useNavigate } from "react-router";
import Dashboard from "./dashboard/dashboard";
import useStore from "@/store/useStore";
import { useEffect } from "react";


const AppRoutes = () => {
    const { token } = useStore();
    const navigate = useNavigate();

    useEffect(() => {
        if (token.length === 0) {
            console.log("Token not found");
            navigate("/signin");
        } else {
            console.log("Token found");
            navigate("/");
        }
    }, [token, navigate]);

    return (
        <Routes>
            <Route path="/signin" element={<SignInPage />} />
            <Route path="/" element={<Dashboard />} />
            <Route path="*" element={<h1>404</h1>} />
        </Routes>
    );
}

export default AppRoutes;