package BaseRepository

import (
	"golang-core/models/PaginationModel"

	"github.com/jinzhu/gorm"
)

func NewBaseRepository()*BaseRepository{
	return &BaseRepository{}
}

type BaseRepository struct {
}

func (repo BaseRepository)Paginate(dbQuery *gorm.DB, pagination PaginationModel.Pagination)(result *gorm.DB){
	
	offset := pagination.Limit*(pagination.Page-1)

	result = dbQuery.Limit(pagination.Limit).Offset(offset).Order(pagination.Order)

	return 
}

func(repo BaseRepository)Peload(dbQuery *gorm.DB,loads []string)(result *gorm.DB){
	result = dbQuery
	for _,s := range loads{
		result = result.Preload(s)
	}
	return
}