package RegionService

import (
	"golang-core/models/RegionModel"
	"golang-core/store"
)


func Index() (data []RegionModel.Region, err error){

	store,storeErr := store.GetStore()
	if storeErr != nil {
		err = storeErr
		return
	}
	data,err = store.Regions.GetAll()
	
	return
}