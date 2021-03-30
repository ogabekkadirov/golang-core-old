package TestController

import (
	"ftp/utils/response"
	"golang-core/responseCodes"

	"github.com/gin-gonic/gin"
)


func Intex(ctx *gin.Context) {
	a := "test"
	response.SuccessResult(ctx, responseCodes.StatusOK, a)
	return 
}