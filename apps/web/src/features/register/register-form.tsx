import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Field, FieldDescription, FieldGroup, FieldLabel } from "@/components/ui/field"
import { Input } from "@/components/ui/input"
import { type ComponentProps, useState } from "react"
import { authClient } from "@/lib/auth-client.ts"
import { toast } from "@/components/ui/toast.tsx"
import { Link, useNavigate } from "@tanstack/react-router"

export function RegisterForm({ className, ...props }: ComponentProps<"div">) {
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const [confirmPassword, setConfirmPassword] = useState("")
  const navigate = useNavigate()

  async function handleRegister() {
    if (password !== confirmPassword) {
      toast.add({
        type: "error",
        title: "Passwords do not match",
      })
      return
    }

    const { data, error } = await authClient.signUp.email({
      name: email.split("@")[0],
      email,
      password,
    })

    console.log(data)
    if (error) {
      toast.add({
        type: "error",
        description: error.message,
      })
      return
    }

    toast.add({
      type: "success",
      description: "Sign up successfully",
    })
    await navigate({ to: "/dashboard" })
  }

  return (
    <div className={cn("flex flex-col gap-6", className)} {...props}>
      <Card>
        <CardHeader className="text-center">
          <CardTitle className="text-xl">Welcome</CardTitle>
          <CardDescription>Register with your Email account</CardDescription>
        </CardHeader>
        <CardContent>
          <form
            onSubmit={async (e) => {
              e.preventDefault()
              await handleRegister()
            }}
          >
            <FieldGroup>
              <Field>
                <FieldLabel htmlFor="email">Email</FieldLabel>
                <Input
                  id="email"
                  type="email"
                  placeholder="m@example.com"
                  required
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                />
              </Field>
              <Field>
                <FieldLabel htmlFor="password">Password</FieldLabel>
                <Input
                  id="password"
                  type="password"
                  required
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                />
              </Field>
              <Field>
                <FieldLabel htmlFor="password">Confirm Password</FieldLabel>
                <Input
                  id="confirm-password"
                  type="password"
                  required
                  value={confirmPassword}
                  onChange={(e) => setConfirmPassword(e.target.value)}
                />
              </Field>
              <Field>
                <Button type="submit">Register</Button>
                <FieldDescription className="text-center">
                  Already have an account? <Link to="/login">Sign In</Link>
                </FieldDescription>
              </Field>
            </FieldGroup>
          </form>
        </CardContent>
      </Card>
      {/*
      <FieldDescription className="px-6 text-center">
        By clicking continue, you agree to our <a href="#">Terms of Service</a> and <a href="#">Privacy Policy</a>.
      </FieldDescription>
      */}
    </div>
  )
}
