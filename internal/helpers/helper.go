package helpers

import (
	"strings"

	"beseller/internal/config"
)

func JoinURL(base, path string) string {
	sep := "/"
	base = strings.TrimSuffix(base, sep)
	path = strings.TrimPrefix(path, sep)

	return base + sep + path
}

func NewTestConfig() *config.Config {
	return &config.Config{
		App: config.App{
			AppURL:   "http://app.com/",
			ImageURL: "pics/items",
		},
	}
}
