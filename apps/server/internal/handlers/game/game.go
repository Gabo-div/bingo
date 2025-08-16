package game

import (
	"context"

	"connectrpc.com/connect"
	"github.com/Gabo-div/bingo/apps/server/internal/game"
	gamerpc "github.com/Gabo-div/bingo/packages/protobuf/go/proto/game"
	"github.com/Gabo-div/bingo/packages/protobuf/go/proto/game/gameconnect"
	"github.com/go-chi/chi/v5"
)

type server struct{}

func (s *server) GenerateBoards(
	ctx context.Context, req *connect.Request[gamerpc.GenerateBoardsRequest],
) (*connect.Response[gamerpc.GenerateBoardsResponse], error) {
	boards := make([]*gamerpc.GameBoard, req.Msg.Number)

	for i := range req.Msg.Number {
		board, seed := game.GenerateBoard(0)

		boards[i] = &gamerpc.GameBoard{
			Seed: seed,
			Cols: make([]*gamerpc.GameBoard_Row, 5),
		}

		for y := range 5 {
			boards[i].Cols[y] = &gamerpc.GameBoard_Row{
				Numbers: make([]uint32, 5),
			}
			for x := range 5 {
				boards[i].Cols[y].Numbers[x] = uint32(board[y][x])
			}
		}
	}

	res := connect.NewResponse(&gamerpc.GenerateBoardsResponse{
		Boards: boards,
	})

	return res, nil
}

func Register(r *chi.Mux) {
	path, handler := gameconnect.NewGameServiceHandler(&server{})
	r.Mount(path, handler)
}
