package logUser

import (
	"golang-api/core"
	"golang-api/log"
	"golang-api/query"
	"golang-api/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LogUserController struct {
	*core.Provider
	logService     *log.LogService
	logMiddleware  *log.LogMiddleware
	userMiddleware *user.UserMiddleware
	queryService   *query.QueryService
}

func NewLogUserController(module *LogUserModule) *LogUserController {
	return &LogUserController{
		Provider:       core.NewProvider("LogUserController"),
		logService:     module.Get("LogService").(*log.LogService),
		logMiddleware:  module.Get("LogMiddleware").(*log.LogMiddleware),
		userMiddleware: module.Get("UserMiddleware").(*user.UserMiddleware),
		queryService:   module.Get("QueryService").(*query.QueryService),
	}
}

func (lc *LogUserController) RegisterRoutes(rg *gin.RouterGroup) {
	logs := rg.Group("/logs")
	logs.GET("/",
		lc.userMiddleware.IsLoggedIn(true),
		lc.userMiddleware.IsAdmin(),
		lc.queryService.QueryFilter(),
		lc.FindAll,
	)
	logs.GET("/:log",
		lc.userMiddleware.IsLoggedIn(true),
		lc.userMiddleware.IsAdmin(),
		lc.logMiddleware.FindOne("log"),
		lc.FindOne,
	)
}

func (lc *LogUserController) FindAll(c *gin.Context) {
	q, _ := c.MustGet("query").(query.QueryFilter)

	logs, err := lc.logService.FindAll(q)
	if err != nil {
		lc.logService.Errorf([]string{"LogUserController", "FindAll"}, "%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, logs)
}

func (lc *LogUserController) FindOne(c *gin.Context) {
	log, _ := c.MustGet("log").(*log.Log)
	c.JSON(http.StatusOK, log)
}
