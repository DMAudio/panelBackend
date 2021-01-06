package api_http

import (
	"context"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"

	"github.com/DMAudio/panelBackend/internal/service"
	"github.com/DMAudio/panelBackend/package/log"
)

const ServiceType = "api_http"

var serviceInitialize sync.Once
var serviceInstance *ServiceST

func GetService() *ServiceST {
	if serviceInstance == nil{
		serviceInitialize.Do(func() {
			serviceInstance = &ServiceST{}
		})
	}
	return serviceInstance
}

type ServiceST struct {
	server *http.Server
	engine *gin.Engine
}

func (s ServiceST) Type() string {
	return ServiceType
}

func (s ServiceST) Serve(stopped func()) {
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

func (s ServiceST) Shutdown(done func()) {
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
	service.Register(GetService())
}
