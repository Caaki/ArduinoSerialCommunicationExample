package handlers

import (
	"SerialArduinoCommunication/commandHelper"
	"SerialArduinoCommunication/configuration"
	"github.com/gin-gonic/gin"
	"net/http"
)

var commandStringTemp string = "'9138, 4566, 490, 654, 434, 710, 434, 710, 434, 682, 462, 706, 486, 630, 542, 1758, 514, 602, 542, 1758, 490, 1786, 486, 1786, 486, 1786, 490, 1786, 486, 1786, 486, 654, 490, 1786, 486, 1786, 434, 1838, 434, 710, 434, 1838, 514, 630, 434, 710, 434, 710, 486, 654, 434, 710, 434, 710, 434, 1838, 434, 710, 434, 1838, 434, 1838, 434, 1838, 438, 1838, 434, 1000'"

func RemoteHandler(e *gin.Engine) {

	v1 := e.Group("remote")
	{
		v1.GET("", sendCommandToRemote)
	}
}

func sendCommandToRemote(c *gin.Context) {

	if configuration.RemoteInUse {
		c.JSON(http.StatusConflict, StatusConflict(map[string]interface{}{
			"error": "Remote is already in use",
		}))
		return
	}
	configuration.RemoteInUse = true
	response, err := commandHelper.SendCommand("python3 serialCommunication.py -a send -v " + commandStringTemp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, StatusInternalServerError(map[string]interface{}{
			"error": err.Error(),
		}))
		return
	}

	c.JSON(http.StatusOK, StatusOK(map[string]interface{}{
		"remoteData": response,
	}))
	configuration.RemoteInUse = false
}
