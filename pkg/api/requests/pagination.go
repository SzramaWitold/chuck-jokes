package requests

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type PaginationRequest struct {
	Page    int
	PerPage int
}

func NewPaginationRequest(c *gin.Context) PaginationRequest {
	page, perPage := getPaginationSetup(c)
	return PaginationRequest{
		Page:    page,
		PerPage: perPage,
	}
}

func getPaginationSetup(c *gin.Context) (int, int) {
	query := c.Request.URL.Query()
	page, err := strconv.Atoi(query.Get("page"))

	if err != nil {
		log.Println("Wrong type provide as a page parameter")
	}

	perPage, err := strconv.Atoi(query.Get("per_page"))

	if err != nil {
		log.Println("Wrong type provide as a per_page parameter")
	}
	return page, perPage
}
