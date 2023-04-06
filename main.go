package main

import (
	"github.com/flunka/greencode/endpoints"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", endpoints.Ping)
	r.POST("/transactions/report", endpoints.Report)
	r.Run()
}
