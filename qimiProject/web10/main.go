package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	type msg struct {
		Name    string `json:"name"`
		Message string `json:"message"`
		Age     int    `json:"age"`
	}
	r.GET("json", func(c *gin.Context) {
		data := msg{"柯子", "名字", 22}
		c.JSON(http.StatusOK, data)
	})
	r.Run(":9090")
}
