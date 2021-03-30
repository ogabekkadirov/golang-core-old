package TestController

import (
	"ftp/utils/response"
	"go-core/responseCodes"

	"github.com/gin-gonic/gin"
)


func Intex(ctx *gin.Context) {
	a := "test"
	response.SuccessResult(ctx, responseCodes.StatusOK, a)
	return 
}