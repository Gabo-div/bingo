package game

import (
	"context"
	"time"

	"connectrpc.com/connect"
	"github.com/Gabo-div/bingo/apps/server/internal/converters"
	"github.com/Gabo-div/bingo/apps/server/internal/database"
	"github.com/Gabo-div/bingo/apps/server/internal/game"
	queries "github.com/Gabo-div/bingo/packages/database/repositories/go"
	gamerpc "github.com/Gabo-div/bingo/packages/protobuf/go/proto/game"
	"github.com/Gabo-div/bingo/packages/protobuf/go/proto/game/gameconnect"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgtype"
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

func (s *server) CreateGame(
	ctx context.Context, req *connect.Request[gamerpc.CreateGameRequest],
) (*connect.Response[gamerpc.Game], error) {
	start, err := converters.StringToTimestamp(req.Msg.Start)

	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	cardPrice, err := converters.Float64ToNumeric(float64(req.Msg.CardPrice))

	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	currentTime := time.Now()

	q := database.GetQueries()

	game, err := q.CreateGame(ctx, queries.CreateGameParams{
		Name:       req.Msg.Name,
		Open:       req.Msg.Open,
		CardPrice:  cardPrice,
		MaxPlayers: int32(req.Msg.MaxPlayers),
		GameTypeId: req.Msg.GameTypeId,
		Start:      start,
		CreatedAt: pgtype.Timestamp{
			Time:  currentTime,
			Valid: true,
		},
		UpdatedAt: pgtype.Timestamp{
			Time:  currentTime,
			Valid: true,
		},
	})

	if err != nil {
		return nil, err
	}

	cardPriceFloat, err := game.CardPrice.Float64Value()

	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&gamerpc.Game{
		Id:         game.ID,
		Name:       game.Name,
		Start:      game.Start.Time.String(),
		Open:       game.Open,
		CardPrice:  float32(cardPriceFloat.Float64),
		MaxPlayers: uint32(game.MaxPlayers),
		GameTypeId: game.GameTypeId,
		CreatedAt:  game.CreatedAt.Time.String(),
		UpdatedAt:  game.UpdatedAt.Time.String(),
	})

	return res, nil
}

func (s *server) EditGame(
	ctx context.Context, req *connect.Request[gamerpc.EditGameRequest],
) (*connect.Response[gamerpc.Game], error) {
	q := database.GetQueries()

	name := pgtype.Text{}

	if req.Msg.Name != nil {
		name = pgtype.Text{
			String: *req.Msg.Name,
			Valid:  true,
		}
	}

	open := pgtype.Bool{}

	if req.Msg.Open != nil {
		open = pgtype.Bool{
			Bool:  *req.Msg.Open,
			Valid: true,
		}
	}

	cardPrice := pgtype.Numeric{}

	if req.Msg.CardPrice != nil {
		value, err := converters.Float64ToNumeric(float64(*req.Msg.CardPrice))

		if err != nil {
			return nil, connect.NewError(connect.CodeInvalidArgument, err)
		}

		cardPrice = value
	}

	maxPlayers := pgtype.Int4{}

	if req.Msg.MaxPlayers != nil {
		maxPlayers = pgtype.Int4{
			Int32: int32(*req.Msg.MaxPlayers),
			Valid: true,
		}
	}

	gameTypeId := pgtype.Int8{}

	if req.Msg.GameTypeId != nil {
		gameTypeId = pgtype.Int8{
			Int64: *req.Msg.GameTypeId,
			Valid: true,
		}
	}

	start := pgtype.Timestamp{}

	if req.Msg.Start != nil {
		value, err := converters.StringToTimestamp(*req.Msg.Start)

		if err != nil {
			return nil, connect.NewError(connect.CodeInvalidArgument, err)
		}

		start = value
	}

	game, err := q.EditGame(ctx, queries.EditGameParams{
		ID:         req.Msg.Id,
		Name:       name,
		Open:       open,
		CardPrice:  cardPrice,
		MaxPlayers: maxPlayers,
		GameTypeId: gameTypeId,
		Start:      start,
		CreatedAt:  pgtype.Timestamp{},
		UpdatedAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
	})

	if err != nil {
		return nil, err
	}

	cardPriceFloat, err := game.CardPrice.Float64Value()

	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&gamerpc.Game{
		Id:         game.ID,
		Name:       game.Name,
		Start:      game.Start.Time.String(),
		Open:       game.Open,
		CardPrice:  float32(cardPriceFloat.Float64),
		MaxPlayers: uint32(game.MaxPlayers),
		GameTypeId: game.GameTypeId,
		CreatedAt:  game.CreatedAt.Time.String(),
		UpdatedAt:  game.UpdatedAt.Time.String(),
	})

	return res, nil
}

func (s *server) DeleteGame(
	ctx context.Context, req *connect.Request[gamerpc.DeleteGameRequest],
) (*connect.Response[gamerpc.Game], error) {
	q := database.GetQueries()

	game, err := q.DeleteGame(ctx, req.Msg.Id)

	if err != nil {
		return nil, err
	}

	if game.ID == 0 {
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	cardPriceFloat, err := game.CardPrice.Float64Value()

	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&gamerpc.Game{
		Id:         game.ID,
		Name:       game.Name,
		Start:      game.Start.Time.String(),
		Open:       game.Open,
		CardPrice:  float32(cardPriceFloat.Float64),
		MaxPlayers: uint32(game.MaxPlayers),
		GameTypeId: game.GameTypeId,
		CreatedAt:  game.CreatedAt.Time.String(),
		UpdatedAt:  game.UpdatedAt.Time.String(),
	})

	return res, nil
}

func (s *server) GetGame(
	ctx context.Context, req *connect.Request[gamerpc.GetGameRequest],
) (*connect.Response[gamerpc.Game], error) {
	q := database.GetQueries()

	game, err := q.GetGame(ctx, req.Msg.Id)

	if err != nil {
		return nil, err
	}

	if game.ID == 0 {
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	cardPriceFloat, err := game.CardPrice.Float64Value()

	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&gamerpc.Game{
		Id:         game.ID,
		Name:       game.Name,
		Start:      game.Start.Time.String(),
		Open:       game.Open,
		CardPrice:  float32(cardPriceFloat.Float64),
		MaxPlayers: uint32(game.MaxPlayers),
		GameTypeId: game.GameTypeId,
		CreatedAt:  game.CreatedAt.Time.String(),
		UpdatedAt:  game.UpdatedAt.Time.String(),
	})

	return res, nil
}

func Register(r *chi.Mux) {
	path, handler := gameconnect.NewGameServiceHandler(&server{})
	r.Mount(path, handler)
}
