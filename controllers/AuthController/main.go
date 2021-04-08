package AuthController

import (
	"golang-core/models/UserModel"
	"golang-core/responseCodes"
	"golang-core/services/AuthService"
	"golang-core/utils/response"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {

	credentials := UserModel.Credentials{}

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		response.ErrorResult(ctx, responseCodes.StatusBadRequest)
		return
	}

	result, err := AuthService.Login(credentials)

	if err != nil {
		response.ErrorResult(ctx, responseCodes.StatusUnauthorized)
		return
	}

	response.SuccessResult(ctx, responseCodes.StatusOK, result)
	return
}
