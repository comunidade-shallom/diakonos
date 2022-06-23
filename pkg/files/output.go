package files

import (
	"path"
)

type Output struct {
	Filename string
}

func (o Output) Exists() bool {
	return FileExists(o.Filename)
}

func (o Output) NameRelative() string {
	return GetRelative(o.Filename)
}

func (o Output) Name() string {
	return path.Base(o.Filename)
}
