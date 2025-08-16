package server

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/evertontomalok/go-rest-sample/internal/adapters/handlers"
	config "github.com/evertontomalok/go-rest-sample/internal/app"
	"github.com/evertontomalok/go-rest-sample/internal/app/domain/entities"
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
	injectRoutes(router, repo)

	return router
}

func injectRoutes(router *gin.Engine, repo ports.Repository) {
	var healthCheck = []entities.Route{
		{
			Path:    "/health",
			Method:  http.MethodGet,
			Handler: handlers.Health,
		},
		{
			Path:    "/readiness",
			Method:  http.MethodGet,
			Handler: handlers.Readiness,
		},
	}
	for _, route := range healthCheck {
		router.Handle(route.Method, route.Path, route.Handler)
	}

	apiGroup := router.Group("/api")

	personHandlers := handlers.NewPersonHandler(repo)
	for _, route := range personHandlers.GetPersonRoutes() {
		apiGroup.Handle(route.Method, route.Path, route.Handler)
	}
}
