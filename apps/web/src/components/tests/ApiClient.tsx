import { useEffect, useMemo, useState } from "react"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { GrpcMethods, GrpcServices, transport } from "@/lib/grpc"
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
  SelectGroup,
  SelectLabel
} from "@/components/ui/select"
import { useForm } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import * as z from "zod"
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import MethodInputField from "./MethodInputField"

const formSchema = z.object({
  method: z.string(),
  timeout: z.string(),
  inputs: z.record(z.string(), z.any())
})

export function ApiClient() {
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      timeout: "0",
      inputs: {}
    }
  })

  const methodName = form.watch("method")
  const selectedMethod = useMemo(() => GrpcMethods.find((m) => m.name === methodName), [methodName])

  const [apiResponse, setApiResponse] = useState<string>("")
  const [apiLoading, setApiLoading] = useState(false)

  useEffect(() => {
    setApiResponse("")
    setApiLoading(false)
    form.setValue("inputs", {})
  }, [form, methodName])

  const onSubmit = async ({ timeout, inputs }: z.infer<typeof formSchema>) => {
    if (apiLoading || !selectedMethod) {
      return
    }

    setApiLoading(true)
    setApiResponse("")

    try {
      const abortController = new AbortController()
      const header = undefined
      const contextValues = undefined

      const parsedInputs: Record<string, any> = {}

      Object.keys(inputs).forEach((fieldName) => {
        const field = selectedMethod.input.fields.find((f) => f.localName === fieldName)

        if (!field) {
          return
        }

        let value = inputs[fieldName]

        if (field.scalar === 8) {
          value = value === "true"
        }

        parsedInputs[fieldName] = value
      })

      console.log("Sending: ", parsedInputs)

      const res = await transport.unary(selectedMethod as any, abortController.signal, Number(timeout), header, parsedInputs, contextValues)

      setApiResponse(`✅ Success: ${JSON.stringify(res.message, null, 2)}`)
    } catch (error) {
      setApiResponse(`❌ Error: ${error}`)
    }

    setApiLoading(false)
  }

  return (
    <Card className="bg-zinc-900 border-zinc-800">
      <CardHeader>
        <CardTitle className="text-white">Api Client</CardTitle>
        <CardDescription className="text-zinc-400">
          Test against a Go API gRPC server
        </CardDescription>
      </CardHeader>
      <CardContent className="space-y-4">
        <Form {...form}>
          <form
            onSubmit={form.handleSubmit(onSubmit)}
            className="flex items-end gap-4"
          >
            <FormField
              control={form.control}
              name="method"
              render={({ field }) => (
                <FormItem className="flex-1">
                  <FormLabel>
                    Method
                  </FormLabel>
                  <Select defaultValue={field.value} onValueChange={field.onChange}>
                    <FormControl>
                      <SelectTrigger className="w-full">
                        <SelectValue placeholder="Method" />
                      </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                      {GrpcServices.map((service) => (
                        <SelectGroup key={service.name}>
                          <SelectLabel>
                            {service.name}
                          </SelectLabel>
                          {service.methods.map((method) => (
                            <SelectItem key={method.name} value={method.name}>
                              {method.name}
                            </SelectItem>
                          ))}
                        </SelectGroup>
                      ))}
                    </SelectContent>
                  </Select>
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="timeout"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Timeout</FormLabel>
                  <FormControl>
                    <Input placeholder="0" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            <Button
              disabled={apiLoading}
              className="bg-blue-600 hover:bg-blue-700 text-white"
            >
              {apiLoading ? "Testing..." : "Test"}
            </Button>
          </form>
        </Form>

        {selectedMethod ? (<div className="bg-zinc-800 py-6 px-4 rounded-md border">
          {
            selectedMethod.input.fields.length ? (selectedMethod.input.fields.map((methodField) => (
              <MethodInputField key={methodField.toString()}
                field={methodField}
                value={form.watch(`inputs.${methodField.localName}`)}
                onChange={(value) => {
                  form.setValue(`inputs.${methodField.localName}`, value)
                }} />
            ))) : <div className="text-center w-full text-sm text-zinc-300 font-bold">No fields</div>
          }
        </div>) : null}

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
