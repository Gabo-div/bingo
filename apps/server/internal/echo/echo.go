package echo

import (
	"context"

	"connectrpc.com/connect"
	echo "github.com/Gabo-div/bingo/packages/protobuf/go/proto/echo"
	"github.com/Gabo-div/bingo/packages/protobuf/go/proto/echo/echoconnect"
	"github.com/go-chi/chi/v5"
)

type server struct{}

func (s *server) Echo(
	ctx context.Context, req *connect.Request[echo.EchoRequest],
) (*connect.Response[echo.EchoResponse], error) {
	res := connect.NewResponse(&echo.EchoResponse{
		Message: req.Msg.Message,
	})
	return res, nil
}

func Register(r *chi.Mux) {
	path, handler := echoconnect.NewEchoServiceHandler(&server{})
	r.Mount(path, handler)
}
