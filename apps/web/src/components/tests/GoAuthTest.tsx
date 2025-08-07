import { useState } from "react"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { userClient } from "@/lib/grpc"

export function GoAuthTest() {
  const [apiResponse, setApiResponse] = useState<string>("")
  const [apiLoading, setApiLoading] = useState(false)

  const testGoAPI = async () => {
    setApiLoading(true)
    setApiResponse("")

    try {
      const res = await userClient.getUser({})
      setApiResponse(`✅ Go API Success: ${JSON.stringify(res, null, 2)}`)
    } catch (error) {
      setApiResponse(`❌ Error: ${error}`)
    }

    setApiLoading(false)
  }

  return (
    <Card className="bg-zinc-900 border-zinc-800 md:col-span-2">
      <CardHeader>
        <CardTitle className="text-white">Go Auth</CardTitle>
        <CardDescription className="text-zinc-400">
          Test against a Go Auth
        </CardDescription>
      </CardHeader>
      <CardContent className="space-y-4">
        <Button
          onClick={testGoAPI}
          disabled={apiLoading}
          className="w-full bg-blue-600 hover:bg-blue-700 text-white"
        >
          {apiLoading ? "Testing..." : "Test"}
        </Button>

        {apiResponse && (
          <div className="p-4 bg-zinc-800 rounded-lg">
            <h3 className="font-semibold text-white mb-2">API Response</h3>
            <pre className="text-xs text-zinc-300 whitespace-pre-wrap">
              {apiResponse}
            </pre>
          </div>
        )}
      </CardContent>
    </Card>
  )
}
