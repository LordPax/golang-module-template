package query

import (
	"golang-api/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IQueryService interface {
	core.IProvider
	QueryFilter() gin.HandlerFunc
}

type QueryService struct {
	*core.Provider
}

func NewQueryService(module *QueryModule) *QueryService {
	return &QueryService{
		Provider: core.NewProvider("QueryService"),
	}
}

func (qf *QueryService) QueryFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.Request.URL.Query()

		queryFilter, err := NewQueryFilter(query)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing query"})
			c.Abort()
			return
		}

		c.Set("query", queryFilter)
		c.Next()
	}
}
