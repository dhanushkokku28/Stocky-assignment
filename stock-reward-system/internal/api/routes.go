package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/rewards", CreateReward)
	router.GET("/stocks/today", GetTodaysStocks)
	router.GET("/stocks/historical", GetHistoricalINRValues)
	router.GET("/user/stats", GetUserStats)
}