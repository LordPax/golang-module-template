package lang

import (
	"github.com/LordPax/godular/core"
)

type ILangService interface {
	core.IProvider
	AddStrings(strings LangString, langs ...string)
	GetLocale(l string) ILocale
}

type LangService struct {
	*core.Provider
	locale map[string]ILocale
}

func NewLangService(module core.IModule) *LangService {
	return &LangService{
		Provider: core.NewProvider("LangService"),
		locale:   make(map[string]ILocale),
	}
}

func (ls *LangService) AddStrings(strings LangString, langs ...string) {
	for _, lang := range langs {
		if _, ok := ls.locale[lang]; !ok {
			ls.locale[lang] = NewLocalize(lang, strings)
			continue
		}
		ls.locale[lang].Append(strings)
	}
}

func (ls *LangService) GetLocale(l string) ILocale {
	locale, ok := ls.locale[l]
	if !ok {
		return ls.locale["en_US"]
	}

	return locale
}
