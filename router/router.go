package router

import (
	"github.com/flunka/greencode/endpoints"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", endpoints.Ping)
	r.POST(endpoints.TransactionsEndpoint, endpoints.Report)
	r.POST(endpoints.ATMEndpoint, endpoints.Order)
	return r
}
