package lang_test

import (
	"golang-api/lang"

	"github.com/LordPax/godular/core"
)

type LangServiceMock struct {
	*core.Provider
	*core.Mockable
}

func NewLangService(module core.IModule) *LangServiceMock {
	return &LangServiceMock{
		Provider: core.NewProvider("LangService"),
		Mockable: core.NewMockable(),
	}
}

func (ls *LangServiceMock) AddStrings(strings lang.LangString, langs ...string) {
	ls.MethodCalled("AddStrings", strings, langs)
	ls.CallFunc("AddStrings")
	return
}

func (ls *LangServiceMock) GetLocale(l string) lang.ILocale {
	ls.MethodCalled("GetLocale", l)
	return ls.CallFunc("GetLocale").(lang.ILocale)
}
