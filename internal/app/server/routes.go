package server

import (
	"net/http"

	"github.com/evertontomalok/go-rest-sample/internal/adapters/handlers"
	"github.com/gin-gonic/gin"
)

type Route struct {
	Path    string
	Method  string
	Handler gin.HandlerFunc
}

var healthCheck = []Route{
	{
		"/health",
		http.MethodGet,
		handlers.Health,
	},
	{
		"/readiness",
		http.MethodGet,
		handlers.Readiness,
	},
}

var routes = []Route{
	{
		"/person/:personId",
		http.MethodGet,
		handlers.GetPersonById,
	},
}
