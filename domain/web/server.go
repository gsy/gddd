package web

type HttpServer interface {
	// TODO: add options to run, like address
	// TODO: configuration, validators
	Run(ctx context.Context)
	Stop(ctx context.Context)
	GracefulStop(ctx context.Context)
}
