package router

import (
	"golang-core/controllers/RegionsController"
	"golang-core/controllers/TestController"
	"golang-core/controllers/UsersController"
	"golang-core/responseCodes"
	"golang-core/utils/response"

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
			testRouter.GET("get-test", TestController.Intex)
		}
		usersRouter := v1.Group("/users")
		{
			usersRouter.GET("/", UsersController.Index)
		}
		authRouter := v1.Group("auth")
		{
			authRouter.POST("/login", UsersController.Login)
		}
		regionsRouter := v1.Group("/regions")
		{
			regionsRouter.GET("/", RegionsController.Index)
		}
	}

	return router
}
