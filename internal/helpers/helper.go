package helpers

import (
	"strings"
)

func JoinURL(base, path string) string {
	sep := "/"
	base = strings.TrimSuffix(base, sep)
	path = strings.TrimPrefix(path, sep)

	return base + sep + path
}
