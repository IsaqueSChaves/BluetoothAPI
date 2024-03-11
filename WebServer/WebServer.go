package WebServer

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaqueschaves/BluetoothAPI/BluetoothManager"
)


func Start() {
	fmt.Println("Starting Web Server...")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
	  c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	  })
	})
	r.POST("/disconnect/:deviceAddress", disconnectHandler)
	r.POST("/connect/:deviceAddress", connectHandler)
	r.Run(":8085")
}

func disconnectHandler(c *gin.Context) {
	deviceAddress := c.Param("deviceAddress")
	err := BluetoothManager.Disconnect(deviceAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
            "message": err.Error(),
        })
	}
	c.JSON(http.StatusOK, gin.H{
        "message": "Disconnected",
    })
}

func connectHandler(c *gin.Context) {
	deviceAddress := c.Param("deviceAddress")
	err := BluetoothManager.Connect(deviceAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
            "message": err.Error(),
        })
	}
	c.JSON(http.StatusOK, gin.H{
        "message": "Connected",
    })
}