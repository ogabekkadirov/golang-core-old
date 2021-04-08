package UserService

import (
	"golang-core/models/PaginationModel"
	"golang-core/models/UserModel"
	"golang-core/store"
)


func Index(pagination PaginationModel.Pagination) (result map[string]interface{},err error){
	
	store, storeErr := store.GetStore()
	
	if storeErr != nil {
		err = storeErr
		return
	}
	
	result, err = store.Users.GetAll(pagination)
	
	return
}
func Get(id uint)(result UserModel.User, err error){
	
	store, storeErr := store.GetStore()
	
	if storeErr != nil {
		err = storeErr
		return 
	}
	
	result, err = store.Users.Get(id)

	return
}

func Login(credentials UserModel.Credentials)(result map[string]interface{}, err error){
	
	store, storeErr := store.GetStore()

	if storeErr != nil {
		err = storeErr
		return
	}
	
	result,err = store.Users.Login(credentials)

	return
}