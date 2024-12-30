package i18n

import (
	"embed"
	"io/fs"
	"log/slog"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

//go:embed *.toml
var fsys embed.FS

var localizer *i18n.Localizer

func Init(loc string) error {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	files, err := fs.Glob(fsys, "*.toml")
	if err != nil {
		return err
	}

	for _, file := range files {
		slog.Info("loading locale", "file", file)
		if _, err := bundle.LoadMessageFileFS(fsys, file); err != nil {
			return err
		}
	}

	slog.Info("using locale", "loc", loc)
	localizer = i18n.NewLocalizer(bundle, loc, "en")

	return nil
}

func L(msgid string) string {
	return localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: msgid})
}

func LT(msgid string, tmplData map[string]interface{}) string {
	return localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: msgid, TemplateData: tmplData})
}

func LC(cfg *i18n.LocalizeConfig) string {
	return localizer.MustLocalize(cfg)
}
