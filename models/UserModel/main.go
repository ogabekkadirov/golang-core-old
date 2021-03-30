package UserModel

import "go-core/models/RegionModel"

type User struct {
	ID       uint   `json:"id"`
	Login    string `json:"login"`
	Fullname string `json:"fullname"`
	RegionID uint   `json:"region_id"`
	Region   RegionModel.Region `gorm:"forignkey:RegionID"` 
}
