package logUser

import (
	"fmt"
	"golang-api/core"
	ginM "golang-api/gin"
	"golang-api/log"
	"golang-api/query"
	"golang-api/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ILogUserController interface {
	core.IProvider
	RegisterRoutes()
	FindAll(c *gin.Context)
	FindOne(c *gin.Context)
}

type LogUserController struct {
	*core.Provider
	logService     log.ILogService
	logMiddleware  log.ILogMiddleware
	userMiddleware user.IUserMiddleware
	queryService   query.IQueryService
	ginService     ginM.IGinService
}

func NewLogUserController(module *LogUserModule) *LogUserController {
	return &LogUserController{
		Provider:       core.NewProvider("LogUserController"),
		logService:     module.Get("LogService").(log.ILogService),
		logMiddleware:  module.Get("LogMiddleware").(log.ILogMiddleware),
		userMiddleware: module.Get("UserMiddleware").(user.IUserMiddleware),
		queryService:   module.Get("QueryService").(query.IQueryService),
		ginService:     module.Get("GinService").(ginM.IGinService),
	}
}

func (lc *LogUserController) OnInit() error {
	lc.RegisterRoutes()
	return nil
}

func (lc *LogUserController) RegisterRoutes() {
	fmt.Println("Registering User routes")
	logs := lc.ginService.GetGroup().Group("/logs")
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

// FindAll godoc
//
// @Summary      Get all logs
// @Description  get all logs
// @Tags         logs
// @Produce      json
// @Param        query  query     string  false  "Query filter"
// @Success      200  {array}   log.Log
// @Failure      500  {object}  utils.HttpError
// @Router       /api/logs [get]
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

// FindOne godoc
//
// @Summary      Get a log by ID
// @Description  get log by ID
// @Tags         logs
// @Produce      json
// @Param        log   path      string  true  "Log ID"
// @Success      200  {object}  log.Log
// @Failure      404  {object}  utils.HttpError
// @Router       /api/logs/{log} [get]
func (lc *LogUserController) FindOne(c *gin.Context) {
	log, _ := c.MustGet("log").(*log.Log)
	c.JSON(http.StatusOK, log)
}
