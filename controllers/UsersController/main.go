package UsersController

import (
	"golang-core/models/UserModel"
	"golang-core/responseCodes"
	"golang-core/services/UserService"
	"golang-core/utils/pagination"
	"golang-core/utils/response"
	"strconv"

	"github.com/gin-gonic/gin"
)


func Index(ctx *gin.Context){

	pagination := pagination.SetPagination(ctx)
	result, err := UserService.Index(pagination)
	if err != nil {
		response.ErrorResult(ctx, responseCodes.StatusBadRequest)
		return
	}
	response.SuccessResult(ctx, responseCodes.StatusOK, result)
	return
}

func Show(ctx *gin.Context){
	idP := ctx.Param("id")
	id,errConv := strconv.ParseUint(idP,10,32)
	if errConv != nil {
		response.ErrorResult(ctx, responseCodes.StatusBadRequest)
		return
	} 
	result,err := UserService.Get(uint(id))
	if err != nil {
		response.ErrorResult(ctx, responseCodes.StatusNotFound)
		return
	}
	response.SuccessResult(ctx, responseCodes.StatusOK, result)
		
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