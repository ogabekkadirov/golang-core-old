package router

import (
	"go-core/controllers/RegionsController"
	"go-core/controllers/TestController"
	"go-core/controllers/UsersController"
	"go-core/responseCodes"
	"go-core/utils/response"

	"github.com/gin-gonic/gin"
)


func Init() *gin.Engine {
	router := gin.Default()
	router.GET("ping", func(ctx *gin.Context) {
		response.SuccessResult(ctx, responseCodes.StatusOK, "pong")
	})
	v1 := router.Group("api/v1")
	{
		testRouter := v1.Group("/test")
		{
			testRouter.GET("get-test",TestController.Intex)
		}
		usersRouter := v1.Group("/users")
		{
			usersRouter.GET("/", UsersController.Index)
		}
		regionsRouter := v1.Group("/regions")
		{
			regionsRouter.GET("/", RegionsController.Index)
		}
	}

	return router
}