package store

import (
	"errors"
	"go-core/database"
	"go-core/repositories/RegionRepository"
	"go-core/repositories/UserRepository"

	"github.com/jinzhu/gorm"
)


func NewStore(db *gorm.DB)*Store{
	return &Store{
		Users: UserRepository.NewUserRepository(db),
		Regions: RegionRepository.NewUserRepository(db),
	}
}

type Store struct {
	Users *UserRepository.UsersRepository
	Regions *RegionRepository.RegionRepository
}

func GetStore()(dbStore *Store, err error){
	db := database.GetDB()

	if db == nil{
		err = errors.New("db is nil in store")
		return
	}
	dbStore = NewStore(db)
	return
}