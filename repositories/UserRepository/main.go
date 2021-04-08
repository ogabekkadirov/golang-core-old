package UserRepository

import (
	"errors"
	"golang-core/models/PaginationModel"
	"golang-core/models/UserModel"
	"golang-core/repositories/BaseRepository"
	"golang-core/responseCodes"

	"github.com/jinzhu/gorm"
)

func NewUserRepository(db *gorm.DB) *UsersRepository {
	return &UsersRepository{db: db, BaseRepository: BaseRepository.NewBaseRepository()}
}

type UsersRepository struct {
	db             *gorm.DB
	BaseRepository *BaseRepository.BaseRepository
}

func (repo UsersRepository) GetAll(pagination PaginationModel.Pagination) (result map[string]interface{}, err error) {

	dbQuery := repo.db.Table("users")

	loads := []string{"Role"}

	var total int64

	dbQuery.Count(&total)

	resultBody := []UserModel.User{}

	dbQuery = repo.BaseRepository.Peload(dbQuery, loads)

	repo.BaseRepository.Paginate(dbQuery, pagination).Find(&resultBody)

	result = map[string]interface{}{
		"data":       resultBody,
		"pagination": pagination,
		"total":      total,
	}

	return
}
func (repo UsersRepository) Get(id uint) (result UserModel.User, err error) {

	loads := []string{"Role"}

	dbQuery := repo.db.Where("id = ?", id)

	dbQuery = repo.BaseRepository.Peload(dbQuery, loads)

	dbQuery.First(&result)

	if result.ID == 0 {
		err = errors.New(string(responseCodes.StatusNotFound))
		return
	}
	return
}
