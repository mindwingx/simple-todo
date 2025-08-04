package locale

import (
	"encoding/json"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	src "microservice"
	"microservice/config"
	"microservice/internal/adapter/registry"
)

type locale struct {
	service config.Service
	config  config.Locale
	bundle  *i18n.Bundle
}

func New(registry registry.IRegistry) ILocale {
	lang := new(locale)
	registry.Parse(&lang.service)
	registry.Parse(&lang.config)
	lang.bundle = i18n.NewBundle(language.English)

	return lang
}

func (l *locale) Init() {
	var path string

	if l.service.Debug == false {
		path = "%s/locale/%s"
	} else {
		path = "%s/internal/adapter/locale/translation/%s"
	}

	l.bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	l.bundle.MustLoadMessageFile(fmt.Sprintf(path, src.Root(), "en-US.json"))
}

func (l *locale) Get(key string) string {
	localizer := i18n.NewLocalizer(l.bundle, l.config.Lang)

	localizedMessage, _ := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: key, // other fields are available with the i18n.Message struct
		},
	})

	return localizedMessage
}

func (l *locale) Plural(key string, params map[string]string) string {
	localizer := i18n.NewLocalizer(l.bundle, l.config.Lang)
	data := make(map[string]string)

	for localizerKey, localizerValue := range params {
		data[localizerKey] = localizerValue
	}

	formattedLocalizer := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: key,
		},
		TemplateData: data,
	})

	return formattedLocalizer
}
