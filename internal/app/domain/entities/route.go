package entities

import "github.com/gin-gonic/gin"

type Route struct {
	Path    string
	Method  string
	Handler gin.HandlerFunc
}
