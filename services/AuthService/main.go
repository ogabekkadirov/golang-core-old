package AuthService

import (
	"golang-core/models/UserModel"
	"golang-core/store"
)

func Login(credentials UserModel.Credentials) (result map[string]interface{}, err error) {

	Store, storeErr := store.GetStore()

	if storeErr != nil {
		err = storeErr
		return
	}

	result, err = Store.Auth.Login(credentials)

	return

}
