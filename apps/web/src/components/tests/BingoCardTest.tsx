import { useState } from "react"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { gameClient } from "@/lib/grpc"
import { Input } from "@/components/ui/input"
import type { GameBoard } from "@repo/protobuf/ts/proto/game/game_pb.js"

const parseCards = (cards: GameBoard[]) => {
  return cards.map((card) => ({
    seed: card.seed,
    numbers: card.cols.map((col) => col.numbers.map((n) => Number(n)))
  }))
}

export function BingoCardTest() {
  const [response, setResponse] = useState<ReturnType<typeof parseCards>>([])

  const [apiLoading, setApiLoading] = useState(false)
  const [number, setNumber] = useState(1)

  const testGoAPI = async () => {
    setApiLoading(true)

    try {
      const res = await gameClient.generateBoards({ number })
      setResponse(parseCards(res.boards))
    } catch (error) {
      console.log(error)
      setResponse([])
    }

    setApiLoading(false)
  }

  return (
    <Card className="bg-zinc-900 border-zinc-800 md:col-span-2">
      <CardHeader>
        <CardTitle className="text-white">Generate Boards</CardTitle>
        <CardDescription className="text-zinc-400">
          Generate random bingo boards
        </CardDescription>
      </CardHeader>
      <CardContent className="space-y-4">
        <Input placeholder="Number" value={number} onChange={(e) => setNumber(Number(e.currentTarget.value))} type="number" />

        <Button
          onClick={testGoAPI}
          disabled={apiLoading}
          className="w-full bg-blue-600 hover:bg-blue-700 text-white"
        >
          {apiLoading ? "Testing..." : "Test"}
        </Button>


        <div className="p-4 bg-zinc-800 rounded-lg grid grid-cols-[repeat(auto-fill,minmax(250px,1fr))] gap-2">
          {response.map((board) => (
            <div key={board.seed} className="grid grid-cols-5 bg-zinc-700 p-2 rounded-md gap-2">
              <div className="col-span-5 text-center text-sm text-zinc-200 font-bold">
                {board.seed}
              </div>
              <div className="w-full aspect-square bg-zinc-800 rounded-full flex items-center justify-center font-bold">B</div>
              <div className="w-full aspect-square bg-zinc-800 rounded-full flex items-center justify-center font-bold">I</div>
              <div className="w-full aspect-square bg-zinc-800 rounded-full flex items-center justify-center font-bold">N</div>
              <div className="w-full aspect-square bg-zinc-800 rounded-full flex items-center justify-center font-bold">G</div>
              <div className="w-full aspect-square bg-zinc-800 rounded-full flex items-center justify-center font-bold">O</div>
              {board.numbers.flatMap((col) => col).map((n, i) => (<div key={i} className="w-full aspect-square bg-zinc-600 rounded-full flex items-center justify-center">{n}</div>))}
            </div>
          ))}

        </div>
      </CardContent>
    </Card>
  )
}
