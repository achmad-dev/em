/*
--- MIT License (c) 2024 achmad
--- See LICENSE for more details
*/

import { LoginForm } from "@/components/login-form";
import { Toaster } from "@/components/ui/toaster";

const SignInPage = () => {
    return (
        <div className="flex items-center justify-center min-h-screen">
            <LoginForm />
            <Toaster />
        </div>
    )
}

export default SignInPage