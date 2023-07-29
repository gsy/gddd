package rpc

type RPCServer interface {
	Start(ctx context.Context) (err error)
	Stop(ctx context.Context) (err error)
	GracefulStop(ctx context.Context) (err error)
}
