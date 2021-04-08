package pagination

import (
	"golang-core/models/PaginationModel"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetPagination(ctx *gin.Context) PaginationModel.Pagination {

	limit := 100
	page := 1
	order := "created_at asc"

	query := ctx.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]

		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			order = queryValue
			break
		}
	}
	return PaginationModel.Pagination{
		Limit: limit,
		Page:  page,
		Order: order,
	}
}
