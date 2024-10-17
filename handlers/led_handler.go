package handlers

import (
	"github.com/gin-gonic/gin"
)

func LedHandler(e *gin.Engine) {
	v1 := e.Group("led")
	{
		v1.GET("/on", sendData)
		v1.GET("/off", turnOffLed)
	}
}

func sendData(c *gin.Context) {

	//err := configuration.SendCommand("on")
	//if err != nil {
	//	return
	//}
	//
	//response, err := configuration.ReadMessage()
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, StatusBadRequest(map[string]interface{}{
	//		"error":    err.Error(),
	//		"response": response,
	//	}))
	//	return
	//}
	//
	//c.JSON(http.StatusOK, StatusOK(map[string]interface{}{
	//	"response": response,
	//}))
}

func turnOffLed(c *gin.Context) {

	//err := configuration.SendCommand("off")
	//if err != nil {
	//	return
	//}
	//
	//response, err := configuration.ReadMessage()
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, StatusBadRequest(map[string]interface{}{
	//		"error":    err.Error(),
	//		"response": response,
	//	}))
	//	return
	//}
	//
	//c.JSON(http.StatusOK, StatusOK(map[string]interface{}{
	//	"response": response,
	//}))
}
