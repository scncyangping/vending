package service

import "context"

type AuthSrv interface {
	CreateJwtToken(ctx context.Context)
}
