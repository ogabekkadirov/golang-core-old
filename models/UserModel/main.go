package UserModel

import (
	"golang-core/models/RoleModel"
)


type User struct {
	ID       uint   `json:"id"`
	Username string `json:"login"`
	Fullname string `json:"fullname"`
	Password []byte `json:"-"`
	RoleID 	uint   `json:"role_id"`
	Role   RoleModel.Role `gorm:"forignkey:RoleID"` 
}

type UsersList struct{
	Users []User `json:"users"`
}
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

