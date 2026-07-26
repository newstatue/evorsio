import {useState} from "react";
import { useNavigate } from "react-router-dom";
import { toast } from "sonner";

import { Button } from "@/components/ui/button";
import {
    Card,
    CardContent,
    CardDescription,
    CardHeader,
    CardTitle,
} from "@/components/ui/card";
import {
    Field,
    FieldDescription,
    FieldGroup,
    FieldLabel,
} from "@/components/ui/field";
import { Input } from "@/components/ui/input";
import {
    InputOTP,
    InputOTPGroup,
    InputOTPSlot,
} from "@/components/ui/input-otp";

import {
    usePostAuthLogin,
    usePostAuthSendCode,
} from "@/api/default";
import {useAuthStore} from "@/store/useAuthStore.ts";


export function AuthForm() {
    const navigate = useNavigate();

    const [email, setEmail] = useState("");
    const [code, setCode] = useState("");

    const {
        mutate: sendCode,
        isPending: isSendingCode,
    } = usePostAuthSendCode({
        mutation: {
            onSuccess: () => {
                toast.success("Verification code sent");
            },
            onError: (error) => {
                console.error(error);
                toast.error("Failed to send verification code");
            },
        },
    });

    const {
        mutate: login,
        isPending: isLoggingIn,
    } = usePostAuthLogin({
        mutation: {
            onSuccess: (response) => {
                if (response.status !== 200) {
                    toast.error(response.data.detail)
                    return;
                }
                useAuthStore.getState().login(response.data.token)
                navigate("/")
            },
            onError: (error) => {
                console.error(error);
                toast.error("Invalid email or verification code");
            },
        },
    });

    function handleSendCode() {
        const normalizedEmail = email.trim();

        if (!normalizedEmail) {
            toast.error("Please enter your email");
            return;
        }

        sendCode({
            data: {
                email: normalizedEmail,
            },
        });
    }

    function handleSubmit() {
        const normalizedEmail = email.trim();

        if (!normalizedEmail) {
            toast.error("Please enter your email");
            return;
        }

        if (code.length !== 5) {
            toast.error("Please enter the 5-digit verification code");
            return;
        }

        login({
            data: {
                email: normalizedEmail,
                code,
            },
        });
    }

    return (
        <div className="flex flex-col gap-6">
            <Card>
                <CardHeader>
                    <CardTitle>Login to your account</CardTitle>

                    <CardDescription>
                        Enter your email below to login to your account
                    </CardDescription>
                </CardHeader>

                <CardContent>
                    <form onSubmit={event => {
                        event.preventDefault();
                        handleSubmit();
                    }}>
                        <FieldGroup>
                            <Field>
                                <FieldLabel htmlFor="email">Email</FieldLabel>

                                <Input
                                    id="email"
                                    type="email"
                                    placeholder="m@example.com"
                                    value={email}
                                    onChange={(event) => setEmail(event.target.value)}
                                    disabled={isLoggingIn}
                                    required
                                />
                            </Field>

                            <Field>
                                <FieldLabel htmlFor="code">Code</FieldLabel>

                                <div className="flex items-center gap-3">
                                    <InputOTP
                                        id="code"
                                        maxLength={5}
                                        value={code}
                                        onChange={setCode}
                                        disabled={isLoggingIn}
                                    >
                                        <InputOTPGroup>
                                            <InputOTPSlot index={0} />
                                            <InputOTPSlot index={1} />
                                            <InputOTPSlot index={2} />
                                            <InputOTPSlot index={3} />
                                            <InputOTPSlot index={4} />
                                        </InputOTPGroup>
                                    </InputOTP>

                                    <Button
                                        className="flex-1"
                                        type="button"
                                        onClick={handleSendCode}
                                        disabled={!email.trim() || isSendingCode || isLoggingIn}
                                    >
                                        {isSendingCode ? "Sending..." : "Send"}
                                    </Button>
                                </div>
                            </Field>

                            <Field>
                                <Button
                                    type="submit"
                                    disabled={
                                        isLoggingIn ||
                                        isSendingCode ||
                                        !email.trim() ||
                                        code.length !== 5
                                    }
                                >
                                    {isLoggingIn ? "Logging in..." : "Login"}
                                </Button>

                                <FieldDescription className="text-center">
                                    Don&apos;t have an account?{" "}
                                    <span>
                    We&apos;ll create one automatically when you sign in.
                  </span>
                                </FieldDescription>
                            </Field>
                        </FieldGroup>
                    </form>
                </CardContent>
            </Card>
        </div>
    );
}