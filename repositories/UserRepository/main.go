package UserRepository

import (
	"go-core/models/UserModel"

	"github.com/jinzhu/gorm"
)


func NewUserRepository(db *gorm.DB)*UsersRepository{
	return &UsersRepository{db:db}
}

type UsersRepository struct {
	db *gorm.DB
}

func(repo UsersRepository)GetAll()(data []UserModel.User,err error){
	loads := []string{"Region"}
	result := repo.db
	for _,s := range loads{
		result = result.Preload(s)
	}
	result.Find(&data)
	err = result.Error
	return
}