package store

import (
	"errors"
	"golang-core/database"
	"golang-core/repositories/AuthRepository"
	"golang-core/repositories/RegionRepository"
	"golang-core/repositories/UserRepository"

	"github.com/jinzhu/gorm"
)

func NewStore(db *gorm.DB) *Store {
	return &Store{
		Users:   UserRepository.NewUserRepository(db),
		Regions: RegionRepository.NewUserRepository(db),
		Auth:    AuthRepository.NewAuthRepository(db),
	}
}

type Store struct {
	Users   *UserRepository.UsersRepository
	Regions *RegionRepository.RegionRepository
	Auth    *AuthRepository.AuthRepository
}

func GetStore() (dbStore *Store, err error) {
	db := database.GetDB()

	if db == nil {
		err = errors.New("db is nil in store")
		return
	}
	dbStore = NewStore(db)
	return
}
