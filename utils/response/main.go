package response

import (
	"golang-core/responseCodes"

	"github.com/gin-gonic/gin"
)


func SuccessResult(ctx *gin.Context, httpStatus int, result interface{})(mapResult map[string]interface{}){

	mapResult = make(map[string]interface{})
	mapResult["success"] = true
	mapResult["message"] = "ok"
	mapResult["result"] = result
	mapResult["status"] = httpStatus
	
	ctx.JSON(httpStatus, mapResult)
	return
}

func ErrorResult(ctx *gin.Context, httpStatus int)(mapResult map[string]interface{}){

	mapResult = make(map[string]interface{})
	mapResult["success"] = false
	mapResult["message"] = responseCodes.StatusText(httpStatus)
	mapResult["status"] = httpStatus

	ctx.AbortWithStatusJSON(httpStatus, mapResult)
	return
}