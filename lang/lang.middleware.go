package lang

import (
	"github.com/LordPax/godular/core"
	"github.com/gin-gonic/gin"
)

type ILangMiddleware interface {
	core.IProvider
	Lang(param string) gin.HandlerFunc
}

type LangMiddleware struct {
	*core.Provider
	langService ILangService
}

func NewLangMiddleware(module core.IModule) *LangMiddleware {
	return &LangMiddleware{
		Provider:    core.NewProvider("LangMiddleware"),
		langService: module.Get("LangService").(ILangService),
	}
}

func (lm *LangMiddleware) Lang(param string) gin.HandlerFunc {
	return func(c *gin.Context) {
		langParam := c.GetHeader(param)
		locale := lm.langService.GetLocale(langParam)
		c.Set("locale", locale)
		c.Next()
	}
}
