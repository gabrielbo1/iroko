package msg

import (
	"github.com/BurntSushi/toml"
	"github.com/gabrielbo1/iroko/domain"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle

func getBundle() *i18n.Bundle {
	if bundle == nil {
		bundle = i18n.NewBundle(language.Portuguese)
		bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
		bundle.MustLoadMessageFile("active.pt.toml")
		bundle.MustLoadMessageFile("active.en.toml")
	}
	return bundle
}

//Message - Set internationalized message in error.
func Message(err *domain.Err, lang, accept string) *domain.Err {
	loc := i18n.NewLocalizer(getBundle(), lang, accept)
	msg := loc.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: err.GetCode(),
		},
		TemplateData: err.GetMsgParam(),
	})
	return domain.NewErr().
		WithErr(err).
		WithMessage(msg)
}
