package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// NoRouteHandler - handle 404 route
func NoRouteHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, StatusNotFound(map[string]interface{}{
		"0": "error_api_not_found",
		"1": c.Request.RequestURI,
	}))
}
