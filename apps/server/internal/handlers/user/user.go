package user

import (
	"context"
	"strconv"
	"google.golang.org/protobuf/types/known/timestamppb"

	"connectrpc.com/connect"
	queries "github.com/Gabo-div/bingo/packages/database/repositories/go"
	user "github.com/Gabo-div/bingo/packages/protobuf/go/proto/user"
	"github.com/Gabo-div/bingo/packages/protobuf/go/proto/user/userconnect"
	"github.com/go-chi/chi/v5"

	"github.com/Gabo-div/bingo/apps/server/internal/auth"
	"github.com/Gabo-div/bingo/apps/server/internal/database"
)

type server struct{}

var q = database.GetQueries()

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

// rpc AddPaymentMethod(PaymentMethodReq) returns (Empty);
func (s *server) AddPaymentMethod(
	ctx context.Context, req *connect.Request[user.PaymentMethodReq],
) (*connect.Response[user.Empty], error) {
	userData := ctx.Value("user").(auth.User)

	userIDNum, err := strconv.ParseInt(userData.ID, 10, 64)
	if err != nil {
		res := connect.NewResponse(&user.Empty{})
		return res, err
	}
	methodTypeId, err := strconv.ParseInt(req.Msg.MethodType, 10, 64)
	if err != nil {
		res := connect.NewResponse(&user.Empty{})
		return res, err
	}

	err = q.AddPaymentMethodWithUserID(ctx, queries.AddPaymentMethodWithUserIDParams{
		UserId: userIDNum,
		MethodTypeId: methodTypeId,
		Data: req.Msg.Data,
	})
	if err != nil {
		res := connect.NewResponse(&user.Empty{})
		return res, err
	}

	res := connect.NewResponse(&user.Empty{})
	return res, nil
}

// rpc GetPaymentMethods(Empty) returns (PaymentMethodList);
func (s *server) GetPaymentMethods(
	ctx context.Context, req *connect.Request[user.Empty],
) (*connect.Response[user.PaymentMethodList], error) {
	userData := ctx.Value("user").(auth.User)
	userIDNum, err := strconv.ParseInt(userData.ID, 10, 64)
	if err != nil {
		res := connect.NewResponse(&user.PaymentMethodList{})
		return res, err
	}

	paymentMethods, err := q.GetPaymentMethodsByUserID(ctx, userIDNum)
	if err != nil {
		res := connect.NewResponse(&user.PaymentMethodList{})
		return res, err
	}

	var methods []*user.PaymentMethodRes
	for _, method := range paymentMethods {
		methods = append(methods, &user.PaymentMethodRes{
			Id: strconv.FormatInt(method.ID, 10),
			MethodType: strconv.FormatInt(method.MethodTypeId, 10),
			Data: method.Data,
		})
	}

	res := connect.NewResponse(&user.PaymentMethodList{
		Methods: methods,
	})
	return res, nil
}


// rpc GetTransactionHistory(Empty) returns (TransactionList);
func (s *server) GetTransactionHistory(
	ctx context.Context, req *connect.Request[user.Empty],
) (*connect.Response[user.TransactionList], error) {
	userData := ctx.Value("user").(auth.User)
	userIDNum, err := strconv.ParseInt(userData.ID, 10, 64)
	if err != nil {
		res := connect.NewResponse(&user.TransactionList{})
		return res, err
	}

	t_history, err := q.GetTransactionHistoryByUserID(ctx, userIDNum)
	if err != nil {
		res := connect.NewResponse(&user.TransactionList{})
		return res, err
	}

	var history []*user.Transaction
	for _, transaction := range t_history {
		t_amount, err := transaction.Amount.Float64Value()
		if err != nil {
			res := connect.NewResponse(&user.TransactionList{})
			return res, err
		}
		t_datetime, err := transaction.CreatedAt.TimestampValue()
		if err != nil {
			res := connect.NewResponse(&user.TransactionList{})
			return res, err
		}
		history = append(history, &user.Transaction{
			Id: strconv.FormatInt(transaction.ID, 10),
			MethodId: strconv.FormatInt(transaction.PaymentMethodId, 10),
			Amount: t_amount.Float64,
			Datetime: timestamppb.New(t_datetime.Time),
		})
	}

	res := connect.NewResponse(&user.TransactionList{
		History: history,
	})
	return res, nil
}

func Register(r *chi.Mux) {
	interceptors := connect.WithInterceptors(auth.NewAuthInterceptor())
	path, handler := userconnect.NewUserServiceHandler(&server{}, interceptors)
	r.Mount(path, handler)
}
