package lang

import (
	"github.com/LordPax/godular/core"
)

type ILangService interface {
	core.IProvider
	AddStrings(strings LangString, lang ...string)
	GetLocale(lang string) ILocale
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

func (l *LangService) AddStrings(strings LangString, lang ...string) {
	for _, la := range lang {
		if _, ok := l.locale[la]; !ok {
			l.locale[la] = NewLocalize(la, strings)
			continue
		}
		l.locale[la].Append(strings)
	}
}

func (l *LangService) GetLocale(lang string) ILocale {
	locale, ok := l.locale[lang]
	if !ok {
		return l.locale["en"]
	}

	return locale
}
