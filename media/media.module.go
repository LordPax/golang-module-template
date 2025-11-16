package media

import (
	"golang-api/database"
	"golang-api/dotenv"
	"golang-api/log"

	"github.com/LordPax/godular/core"
)

var module *MediaModule

type MediaModule struct {
	*core.Module
}

func NewMediaModule() *MediaModule {
	module := &MediaModule{
		Module: core.NewModule("MediaModule"),
	}

	module.AddModule(dotenv.Module())
	module.AddModule(database.Module())
	module.AddModule(log.Module())
	module.AddProvider(NewMediaModel(module))
	module.AddProvider(NewOpenstackService(module))
	module.AddProvider(NewMediaService(module))
	module.AddProvider(NewMediaMiddleware(module))

	return module
}

func Module() *MediaModule {
	if module == nil {
		module = NewMediaModule()
	}
	return module
}
