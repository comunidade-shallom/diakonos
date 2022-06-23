package files

import (
	"path"

	"github.com/comunidade-shallom/diakonos/pkg/fileutils"
)

type Output struct {
	Filename string
}

func (o Output) Exists() bool {
	return fileutils.FileExists(o.Filename)
}

func (o Output) NameRelative() string {
	return fileutils.GetRelative(o.Filename)
}

func (o Output) Name() string {
	return path.Base(o.Filename)
}
