import { authClient } from "@/lib/auth"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Link } from "@tanstack/react-router"

export function AuthStatus() {
  const { data: session, isPending } = authClient.useSession()

  const handleSignOut = async () => {
    await authClient.signOut()
    window.location.reload()
  }

  return (
    <Card className="bg-zinc-900 border-zinc-800">
      <CardHeader>
        <CardTitle className="text-white">Better-Auth Status</CardTitle>
        <CardDescription className="text-zinc-400">
          {isPending ? "Checking authentication..." : session ? "You are signed in" : "You are not signed in"}
        </CardDescription>
      </CardHeader>
      <CardContent className="space-y-4">
        {isPending ? (
          <div className="text-zinc-400">Loading...</div>
        ) : session ? (
          <div className="space-y-4">
            <div className="p-4 bg-zinc-800 rounded-lg">
              <h3 className="font-semibold text-white mb-2">User Info</h3>
              <p className="text-zinc-300">Email: {session.user.email}</p>
              <p className="text-zinc-300">Name: {session.user.name}</p>
              <p className="text-zinc-300">ID: {session.user.id}</p>
            </div>
            <Button
              onClick={handleSignOut}
              variant="outline"
              className="w-full bg-zinc-800 border-zinc-700 text-white hover:bg-zinc-700"
            >
              Sign Out
            </Button>
          </div>
        ) : (
          <div className="space-y-3">
            <Link to="/login" className="block">
              <Button className="w-full bg-white text-zinc-900 hover:bg-zinc-100">
                Sign In
              </Button>
            </Link>
            <Link to="/signup" className="block">
              <Button
                variant="outline"
                className="w-full bg-zinc-800 border-zinc-700 text-white hover:bg-zinc-700"
              >
                Create Account
              </Button>
            </Link>
          </div>
        )}
      </CardContent>
    </Card>
  )
}
