package server

import (
	"context"
	"log"
	"net"
	"net/http"

	config "github.com/evertontomalok/go-rest-sample/internal/app"
	"github.com/evertontomalok/go-rest-sample/internal/app/server/handlers"
	"github.com/evertontomalok/go-rest-sample/internal/ports"

	"github.com/evertontomalok/go-rest-sample/pkg/utils"
	"github.com/gin-gonic/gin"
)

func RunServer(ctx context.Context, config config.Config, repository ports.Repository) {
	done := utils.MakeDoneSignal()

	server := &http.Server{
		Addr:    net.JoinHostPort(config.App.Host, config.App.Port),
		Handler: Router(repository),
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

func Router(repo ports.Repository) *gin.Engine {
	router := gin.Default()

	healtCheckRoutes := handlers.HealthCheckRoutes{}
	for _, route := range healtCheckRoutes.GetRoutes() {
		router.Handle(route.Method, route.Path, route.Handler)
	}

	apiGroup := router.Group("/api")

	personHandlers := handlers.NewPersonHandler(repo)
	for _, route := range personHandlers.GetRoutes() {
		apiGroup.Handle(route.Method, route.Path, route.Handler)
	}

	return router
}
