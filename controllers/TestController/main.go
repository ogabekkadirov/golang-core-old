package TestController

import (
	"golang-core/responseCodes"
	"golang-core/utils/response"

	"github.com/gin-gonic/gin"
)

func Intex(ctx *gin.Context) {
	a := "test"
	response.SuccessResult(ctx, responseCodes.StatusOK, a)
	return
}
