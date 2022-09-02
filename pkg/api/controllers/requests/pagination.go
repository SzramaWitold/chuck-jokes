package requests

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FindCollection struct {
	Page    int
	PerPage int
}

func (r *RequestValidator) NewFindCollection(c *gin.Context) FindCollection {
	page, perPage := getPaginationSetup(c)

	return FindCollection{
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
