package alert

import "context"

const (
	AppName = "alert"
)

type Service interface {
	RpcServer
	CreateTable(ctx context.Context) error
}
