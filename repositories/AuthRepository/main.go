package AuthRepository

import (
	"errors"
	"golang-core/models/UserModel"
	"golang-core/repositories/BaseRepository"
	"golang-core/responseCodes"
	"golang-core/utils/gcjwt"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db:             db,
		BaseRepository: BaseRepository.NewBaseRepository(),
	}
}

type AuthRepository struct {
	db             *gorm.DB
	BaseRepository *BaseRepository.BaseRepository
}

func (repo AuthRepository) Login(credentials UserModel.Credentials) (result map[string]interface{}, err error) {

	user := &UserModel.User{}

	repo.db.Where("username = ?", credentials.Username).First(&user)

	if user.ID == 0 {
		err = errors.New(string(responseCodes.StatusNotFound))
		return
	}

	if err = bcrypt.CompareHashAndPassword(user.Password, []byte(credentials.Password)); err != nil {
		return
	}
	token, tokenErr := gcjwt.CreateToken(*user)

	if tokenErr != nil {
		err = tokenErr
		return
	}

	result = map[string]interface{}{
		"access_token": token,
	}
	return
}
