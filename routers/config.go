package routers

import (
	"github.com/chaosyeo/gin-starter/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterLotteriesRouters(router *gin.Engine) (*gin.Engine, gin.Error) {
	group := router.Group("/v1.0")
	group.POST("/lotteries", handlers.LotteriesHandler{}.Create)
	group.POST("/lotteries/remove", handlers.LotteriesHandler{}.Remove)
	group.POST("/lotteries/update", handlers.LotteriesHandler{}.Update)
	group.GET("/lotteries", handlers.LotteriesHandler{}.Collect)
	group.GET("/lotteries/fetch", handlers.LotteriesHandler{}.Fetch)
}