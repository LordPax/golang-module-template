package query

import (
	"net/http"

	"github.com/LordPax/godular/core"
	"github.com/gin-gonic/gin"
)

type IQueryService interface {
	core.IProvider
	QueryFilter() gin.HandlerFunc
}

type QueryService struct {
	*core.Provider
}

func NewQueryService(module core.IModule) *QueryService {
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
