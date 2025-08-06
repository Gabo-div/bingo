package user

import (
	"context"

	"connectrpc.com/connect"
	"github.com/Gabo-div/bingo/apps/backend-main/internal/auth"
	user "github.com/Gabo-div/bingo/packages/protobuf/go/proto/user"
	"github.com/Gabo-div/bingo/packages/protobuf/go/proto/user/userconnect"
	"github.com/go-chi/chi/v5"
)

type server struct{}

func (s *server) GetUser(
	ctx context.Context, req *connect.Request[user.Empty],
) (*connect.Response[user.User], error) {
	userData, err := auth.UserFromHeader(ctx, req.Header())

	if err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}

	res := connect.NewResponse(&user.User{
		Id:            userData.ID,
		Name:          userData.Name,
		Email:         userData.Email,
		EmailVerified: userData.EmailVerified,
		Role:          userData.Role,
		Banned:        userData.Banned,
		BanReason:     userData.BanReason,
		BanExpires:    userData.BanExpires,
		Image:         userData.Image,
		CreatedAt:     userData.CreatedAt,
		UpdatedAt:     userData.UpdatedAt,
	})

	return res, nil
}

func Register(r *chi.Mux) {
	path, handler := userconnect.NewUserServiceHandler(&server{})
	r.Mount(path, handler)
}
