import { createFileRoute } from '@tanstack/react-router'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { AuthStatus } from "@/components/AuthStatus"
import { GoAPITest } from "@/components/GoApiTest"
import { GoAuthTest } from '@/components/GoAuthTest'

export const Route = createFileRoute('/')({
  component: RouteComponent,
})

function RouteComponent() {
  return (
    <div className="min-h-screen bg-zinc-950 text-white">
      <div className="container mx-auto px-4 py-8">
        <div className="max-w-4xl mx-auto">
          <div className="text-center mb-8">
            <h1 className="text-4xl font-bold mb-4">Authly</h1>
            <p className="text-zinc-400 text-lg">
              Authentication demo with Better Auth, Drizzle, and Go API integration
            </p>
          </div>

          <div className="grid gap-6 md:grid-cols-2">
            <AuthStatus />
            <GoAPITest />
            <GoAuthTest />
          </div>

          <div className="mt-8 text-center">
            <Card className="bg-zinc-900 border-zinc-800">
              <CardHeader>
                <CardTitle className="text-white">Quick Start</CardTitle>
                <CardDescription className="text-zinc-400">
                  Get started with authentication
                </CardDescription>
              </CardHeader>
              <CardContent>
                <div className="grid gap-4 md:grid-cols-3">
                  <div className="text-center">
                    <div className="text-2xl mb-2">🔐</div>
                    <h3 className="font-semibold text-white">Sign Up</h3>
                    <p className="text-zinc-400 text-sm">Create a new account</p>
                  </div>
                  <div className="text-center">
                    <div className="text-2xl mb-2">🔑</div>
                    <h3 className="font-semibold text-white">Sign In</h3>
                    <p className="text-zinc-400 text-sm">Access your account</p>
                  </div>
                  <div className="text-center">
                    <div className="text-2xl mb-2">🚀</div>
                    <h3 className="font-semibold text-white">Test API</h3>
                    <p className="text-zinc-400 text-sm">Verify with Go server</p>
                  </div>
                </div>
              </CardContent>
            </Card>
          </div>
        </div>
      </div>
    </div>
  )
}
