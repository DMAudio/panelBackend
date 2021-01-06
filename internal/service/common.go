package service

import (
	"context"
	"sync"
	"time"

	"github.com/DMAudio/panelBackend/package/log"
)

type Service interface {
	Type() string
	Serve(stopped func())
	Shutdown(done func())
}

var servicesMap map[string]Service

func Register(service Service) {
	if service == nil {
		return
	}
	if servicesMap == nil {
		servicesMap = map[string]Service{}
	}
	servicesMap[service.Type()] = service
}

func Serve() {
	if len(servicesMap) == 0 {
		return
	}
	var wait = sync.WaitGroup{}
	for name, service := range servicesMap {
		mapKey := name
		wait.Add(1)
		service.Serve(func() {
			wait.Done()
			log.Warnf("service `%s` stopped", mapKey)
		})
	}
	wait.Wait()
}

var ShutdownTimeout = time.Second * 3 // todo config

func Shutdown() {
	var wait = sync.WaitGroup{}
	wait.Add(len(servicesMap))
	for _, e := range servicesMap {
		var done func()
		var ctx context.Context
		if ShutdownTimeout > 0 {
			ctx, done = context.WithTimeout(context.Background(), ShutdownTimeout)
		} else {
			ctx, done = context.WithCancel(context.Background())
		}
		e.Shutdown(done)
		go func() {
			<-ctx.Done()
			wait.Done()
		}()
	}
	wait.Wait()
}
