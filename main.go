package main

import (
	"banboom/mysql"
	"banboom/service"
	"github.com/gin-gonic/gin"
)
import "net/http"
import "banboom/morestrings"

func main() {
	mysql.Init()
	mysql.SelectAll()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": morestrings.ReverseRunes("!oG ,olleH"),
		})
	})

	r.GET("/dict", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": service.TranslateText("yydsa"),
		})
	})


	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	mysql.Close()
}
