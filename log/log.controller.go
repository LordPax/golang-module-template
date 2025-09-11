package log

import (
	"golang-api/core"

	"github.com/gin-gonic/gin"
)

type LogController struct {
	*core.Provider
	logService *LogService
}

func NewLogController(module *LogModule) *LogController {
	return &LogController{
		Provider:   core.NewProvider("LogController"),
		logService: module.Get("LogService").(*LogService),
	}
}

func (uc *LogController) RegisterRoutes(rg *gin.RouterGroup) {
	// logs := rg.Group("/logs")
}
