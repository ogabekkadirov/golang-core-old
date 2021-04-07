package UserService

import (
	"golang-core/models/UserModel"
	"golang-core/store"
)


func Index() (result []UserModel.User,err error){
	store, storeErr := store.GetStore()
	if storeErr != nil {
		err = storeErr
		return
	}
	result, err = store.Users.GetAll()
	
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