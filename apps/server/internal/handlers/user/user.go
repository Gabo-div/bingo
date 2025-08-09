package user

import (
	"context"

	"connectrpc.com/connect"
	user "github.com/Gabo-div/bingo/packages/protobuf/go/proto/user"
	"github.com/Gabo-div/bingo/packages/protobuf/go/proto/user/userconnect"
	"github.com/go-chi/chi/v5"

	"github.com/Gabo-div/bingo/apps/backend-main/internal/auth"
)

type server struct{}

func (s *server) GetUser(
	ctx context.Context, req *connect.Request[user.Empty],
) (*connect.Response[user.User], error) {
	userData := ctx.Value("user").(auth.User)

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
	interceptors := connect.WithInterceptors(auth.NewAuthInterceptor())
	path, handler := userconnect.NewUserServiceHandler(&server{}, interceptors)
	r.Mount(path, handler)
}
