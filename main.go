package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Any("/auth", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")
		if username == "admin" && password == "cGFzc3dvcmQK" {
			context.JSON(http.StatusOK, nil)
		} else {
			context.AbortWithStatus(http.StatusUnauthorized)
		}
	})
	router.Run(":8080")
}
