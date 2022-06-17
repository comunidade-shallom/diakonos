package fileutils

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"
)

func ChangeExtension(original, ext string) string {
	name := strings.TrimSuffix(original, filepath.Ext(original))

	return fmt.Sprintf("%s.%s", name, ext)
}

func ChangeLocation(original, targetDir, ext string) string {
	name := path.Base(original)

	if ext != "" {
		name = ChangeExtension(name, ext)
	}

	return path.Join(targetDir, name)
}
