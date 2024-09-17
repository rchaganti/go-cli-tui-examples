package helper

import (
	"path"
	"strings"
)

func ParseConfigPath(configPath string) (string, string, string) {
	dir, file := path.Split(configPath)
	ext := strings.Trim(path.Ext(configPath), ".")
	return dir, file, ext
}
