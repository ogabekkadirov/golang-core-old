package UsersController

import (
	"go-core/responseCodes"
	"go-core/services/UserService"
	"go-core/utils/response"

	"github.com/gin-gonic/gin"
)


func Index(ctx *gin.Context){
	result, err := UserService.Index()
	if err != nil {
		response.ErrorResult(ctx, responseCodes.StatusBadRequest)
		return
	}
	response.SuccessResult(ctx, responseCodes.StatusOK, result)
	return
}