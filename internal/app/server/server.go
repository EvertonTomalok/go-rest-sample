package server

import (
	"context"
	"log"
	"net"
	"net/http"

	config "github.com/evertontomalok/go-rest-sample/internal/app"

	"github.com/evertontomalok/go-rest-sample/pkg/utils"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	injectRoutes(router)

	return router
}

func RunServer(ctx context.Context, config config.Config) {
	done := utils.MakeDoneSignal()

	server := &http.Server{
		Addr:    net.JoinHostPort(config.App.Host, config.App.Port),
		Handler: Router(),
	}

	go func() {
		log.Printf("Server started at %s:%s", config.App.Host, config.App.Port)

		if err := server.ListenAndServe(); err != nil {
			log.Panicf("Error trying to start server. %+v", err)
		}
	}()

	<-done
	log.Println("Stopping server...")
}

func injectRoutes(router *gin.Engine) {
	for _, route := range healthCheck {
		router.Handle(route.Method, route.Path, route.Handler)
	}

	apiGroup := router.Group("/api")
	for _, route := range routes {
		apiGroup.Handle(route.Method, route.Path, route.Handler)
	}
}
