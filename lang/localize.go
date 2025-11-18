package lang

type LangString map[string]string

type ILocale interface {
	Get(key string) string
	SetLang(lang string)
	Set(key string, value string)
	Append(strings LangString)
}

type Locale struct {
	lang    string
	strings LangString
}

func NewLocalize(lang string, strings LangString) *Locale {
	return &Locale{
		lang:    lang,
		strings: strings,
	}
}

func (l *Locale) SetLang(lang string) {
	l.lang = lang
}

func (l *Locale) Get(key string) string {
	return l.strings[key]
}

func (l *Locale) Set(key string, value string) {
	l.strings[key] = value
}

func (l *Locale) Append(strings LangString) {
	for k, v := range strings {
		l.Set(k, v)
	}
}
