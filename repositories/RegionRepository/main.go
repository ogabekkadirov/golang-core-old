package RegionRepository

import (
	"go-core/models/RegionModel"

	"github.com/jinzhu/gorm"
)


func NewUserRepository(db *gorm.DB)*RegionRepository{
	return &RegionRepository{db:db}
}

type RegionRepository struct {
	db *gorm.DB
}

func(repo RegionRepository) GetAll()(data []RegionModel.Region, err error){
	
	result := repo.db.Find(&data)
	err = result.Error
	
	return
}