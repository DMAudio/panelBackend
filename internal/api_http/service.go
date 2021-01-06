package api_http

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/DMAudio/panelBackend/internal/service"
	"github.com/DMAudio/panelBackend/package/log"
)

type apiService struct {
	server *http.Server
	engine *gin.Engine
}

func (s apiService) Type() string {
	return "api_http"
}

func (s apiService) Serve(stopped func()) {
	gin.SetMode(gin.DebugMode)
	engine := gin.Default()
	engine.Use(func(c *gin.Context) {
		c.Next()
	})
	engine.Use(gin.Recovery())
	s.engine = engine
	go func() {
		s.server = &http.Server{Handler: s.engine}
		_ = s.server.ListenAndServe() // todo config
		stopped()
	}()
}

func (s apiService) Shutdown(done func()) {
	if s.server == nil {
		done()
		return
	}
	s.server.RegisterOnShutdown(done)
	err := s.server.Shutdown(context.Background())
	if err != nil {
		log.Error(err)
	}
}

func Load() {
	service.Register(&apiService{})
}
