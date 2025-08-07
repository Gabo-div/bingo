import { createFileRoute } from '@tanstack/react-router'
import { AuthStatus } from "@/components/tests/AuthStatus"
import { GoAPITest } from "@/components/tests/GoApiTest"
import { GoAuthTest } from '@/components/tests/GoAuthTest'

export const Route = createFileRoute('/tests')({
  component: RouteComponent,
})

function RouteComponent() {
  return (
    <div className="min-h-screen bg-zinc-950 text-white">
      <div className="container mx-auto px-4 py-8">
        <div className="max-w-4xl mx-auto">
          <div className="text-center mb-8">
            <h1 className="text-4xl font-bold mb-4">Tests</h1>
            <p className="text-zinc-400 text-lg">
              Route to test Better Auth Server and Go gRPC Server
            </p>
          </div>
          <div className="grid gap-6 md:grid-cols-2">
            <AuthStatus />
            <GoAPITest />
            <GoAuthTest />
          </div>
        </div>
      </div>
    </div>
  )
}
