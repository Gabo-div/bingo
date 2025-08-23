import type { EchoService } from "@repo/protobuf/ts/proto/echo/echo_pb.js"
import { Input } from "../ui/input"
import { Label } from "../ui/label"

type Props = {
  field: typeof EchoService.methods[0]["input"]["fields"][0]
  value: any
  onChange: (value: any) => void
}

export default function MethodInputField(props: Props) {
  if (props.field.fieldKind === "scalar") {
    return <MethodScalarField {...props} />
  }

  return null
}


function MethodScalarField({ field, value, onChange }: Props) {
  return <div className="flex items-center gap-4">
    <Label>{field.name}</Label>
    <Input placeholder={field.name} value={value} onChange={(e) => onChange(e.currentTarget.value)} />
  </div>
}
