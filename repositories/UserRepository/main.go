package UserRepository

import (
	"errors"
	"fmt"
	"golang-core/models/UserModel"
	"golang-core/responseCodes"
	"golang-core/utils/gcjwt"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)


func NewUserRepository(db *gorm.DB)*UsersRepository{
	return &UsersRepository{db:db}
}

type UsersRepository struct {
	db *gorm.DB
}

func(repo UsersRepository)GetAll()(data []UserModel.User,err error){
	loads := []string{"Role"}
	result := repo.db
	for _,s := range loads{
		result = result.Preload(s)
	}
	result.Find(&data)
	err = result.Error
	return
}

func(repo UsersRepository)Login(credentials UserModel.Credentials)(data map[string]interface{}, err error){

	user := &UserModel.User{}

	repo.db.Where("username = ?", credentials.Username).First(&user)
	
	if user.ID == 0 {
		err = errors.New(string(responseCodes.StatusNotFound))
		return
	}

	if err = bcrypt.CompareHashAndPassword(user.Password, []byte(credentials.Password)); err != nil{
		return
	}
	fmt.Println(user)
	token,tokenErr := gcjwt.CreateToken(*user)

	if tokenErr != nil {
		err = tokenErr
		return 
	}

	data = make(map[string]interface{})
	data["access_token"]=token

	return 
}	