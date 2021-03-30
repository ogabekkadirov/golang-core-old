package UserService

import (
	"go-core/models/UserModel"
	"go-core/store"
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