package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", Ping) 
	r.POST("/transactions/report", Report)
	r.Run()
}
