package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

type HealthCheckHandler struct{}

func (HealthCheckHandler) GetRoutes() []Route {
	return []Route{
		{
			Path:   "/health",
			Method: http.MethodGet,
			Handler: func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"status": "ok",
				})
			},
		},
		{
			Path:   "/readiness",
			Method: http.MethodGet,
			Handler: func(c *gin.Context) {
				c.String(http.StatusOK, "ok")
			},
		},
	}
}
