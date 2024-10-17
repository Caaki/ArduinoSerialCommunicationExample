package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	"net/http"
)

func HealthCheckHandler(e *gin.Engine) {

	v1 := e.Group("health")
	{
		v1.GET("", healthCheck)
	}
}
func healthCheck(c *gin.Context) {
	userAgentString := c.Request.UserAgent()
	ua := user_agent.New(userAgentString)
	browser, _ := ua.Browser()

	fmt.Println(browser, ua.Model(), ua.OS())

	c.JSON(http.StatusOK, StatusOK(map[string]interface{}{
		"status": "Server is up",
	}))
}
