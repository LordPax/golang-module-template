package log

import (
	"golang-api/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ILogMiddleware interface {
	core.IProvider
	FindOne(name string) gin.HandlerFunc
}

type LogMiddleware struct {
	*core.Provider
	logService ILogService
}

func NewLogMiddleware(module *LogModule) *LogMiddleware {
	return &LogMiddleware{
		Provider:   core.NewProvider("LogMiddleware"),
		logService: module.Get("LogService").(ILogService),
	}
}

func (um *LogMiddleware) FindOne(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param(name)
		log, err := um.logService.FindByID(id)
		tags := []string{"LogMiddleware", "FindOne"}
		if err != nil {
			um.logService.Errorf(tags, "Log %s not found: %v", id, err)
			c.JSON(http.StatusNotFound, gin.H{"error": "Log not found"})
			c.Abort()
			return
		}

		c.Set("log", log)
		c.Next()
	}
}
