package UsersController

import (
	"golang-core/models/UserModel"
	"golang-core/responseCodes"
	"golang-core/services/UserService"
	"golang-core/utils/response"

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

func Login(ctx *gin.Context){
	
	credentials := UserModel.Credentials{}

	if err := ctx.ShouldBindJSON(&credentials); err != nil{
		response.ErrorResult(ctx, responseCodes.StatusBadRequest)
		return
	}

	result, err := UserService.Login(credentials)

	if err != nil{
		response.ErrorResult(ctx, responseCodes.StatusUnauthorized)
		return
	}

	response.SuccessResult(ctx, responseCodes.StatusOK, result)
}