package v1

import (
	"context"

	svcmodel "go_school/3_http_server/service/model"
)

type Service interface {
	CreateUser(ctx context.Context, user svcmodel.User) error
	ListUsers(ctx context.Context) []svcmodel.User
	GetUser(ctx context.Context, email string) (svcmodel.User, error)
	UpdateUser(ctx context.Context, email string, user svcmodel.User) (svcmodel.User, error)
	DeleteUser(ctx context.Context, email string) error
}
