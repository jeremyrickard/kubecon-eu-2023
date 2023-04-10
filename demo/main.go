package main

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jeremyrickard/scaling-spork/api"
	"github.com/mitchellh/mapstructure"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})
	r.POST("/send", func(c *gin.Context) {
		jsonData, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "couldn't read message")
		}
		var message Message
		err = mapstructure.Decode(jsonData, &message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "bad message")
		}
		client := api.Client{}
		if err := client.DoSomething(); err != nil {
			c.JSON(http.StatusInternalServerError, "my client was broken")
		}
		c.JSON(http.StatusOK, gin.H{
			"sender": message.Sender,
			"text":   message.Text,
		})
	})
	r.Run()
}
