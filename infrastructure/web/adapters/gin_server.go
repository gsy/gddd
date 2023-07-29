package adapters

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type ginServer struct {
	server      *http.Server
	quitChannel chan os.Signal
}

func NewGinServer() *ginServer {
	return &ginServer{
		server: &http.Server{
			Addr:   ":8080",
			engine: gin.Default(),
		},
		quitChannel: sigint,
	}
}

func (s *ginServer) Run(ctx context.Context) {
	if err = s.server.ListenAndServe(); err != nil {
		log.Fatalf("HTTP server run: %v", err)
	}
}

func (s *ginServer) Stop(ctx context.Context) {
	if err := s.server.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP server Shutdown: %v", err)
	}
}

func (s *ginServer) GracefulStop(ctx context.Context) (err error) {
	return s.Stop(ctx)
}
