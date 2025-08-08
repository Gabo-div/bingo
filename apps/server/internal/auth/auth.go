package auth

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

const tokenHeader = "Authorization"

type User struct {
	ID            string
	Name          string
	Email         string
	EmailVerified bool
	Role          string
	Banned        bool
	BanReason     string
	BanExpires    string
	Image         string
	CreatedAt     string
	UpdatedAt     string
}

func userFromToken(userID string, token jwt.Token) User {
	var name string
	token.Get("name", &name)

	var email string
	token.Get("email", &email)

	var emailVerified bool
	token.Get("emailVerified", &emailVerified)

	var role string
	token.Get("role", &role)

	var banned bool
	token.Get("banned", &banned)

	var banReason string
	token.Get("banReason", &banReason)

	var banExpires string
	token.Get("banExpires", &banExpires)

	var image string
	token.Get("image", &image)

	var createdAt string
	token.Get("createdAt", &createdAt)

	var updatedAt string
	token.Get("updatedAt", &updatedAt)

	return User{
		ID:            userID,
		Name:          name,
		Email:         email,
		EmailVerified: emailVerified,
		Role:          role,
		Banned:        banned,
		BanReason:     banReason,
		BanExpires:    banExpires,
		Image:         image,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}
}

func NewAuthInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			if !req.Spec().IsClient && req.Header().Get(tokenHeader) == "" {
				// Check token in handlers.
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					errors.New("no token provided"),
				)
			}

			keyset, err := jwk.Fetch(ctx, "http://localhost:3000/api/auth/jwks")
			if err != nil {
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					errors.New("no token provided"),
				)
			}

			token, err := jwt.ParseHeader(req.Header(), "Authorization", jwt.WithKeySet(keyset))
			if err != nil {
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					errors.New("no token provided"),
				)
			}

			userID, exist := token.Subject()
			if !exist {
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					errors.New("missing user id"),
				)
			}

			userObj := userFromToken(userID, token)
			newCtx := context.WithValue(ctx, "user", userObj)
			return next(newCtx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
