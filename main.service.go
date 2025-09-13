package main

import (
	"golang-api/core"
	"golang-api/database"
	"golang-api/gin"
)

type MainService struct {
	*core.Provider
	databaseService *database.DatabaseService
	ginService      *gin.GinService
}

func NewMainService(module *MainModule) *MainService {
	return &MainService{
		Provider:        core.NewProvider("MainService"),
		databaseService: module.Get("DatabaseService").(*database.DatabaseService),
		ginService:      module.Get("GinService").(*gin.GinService),
	}
}

func (ms *MainService) Start() {
	defer ms.databaseService.Close()
	ms.ginService.Run()
}
