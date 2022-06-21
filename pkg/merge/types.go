package merge

import (
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/comunidade-shallom/diakonos/pkg/fileutils"
)

type MergeParams struct {
	Sources   []string
	OutputDir string
	Name      string
}

type MergedFile struct {
	MergeParams
	Name string
}

func (p MergeParams) filename() string {
	name := p.Name

	if name == "" {
		name = buildHashName(p.Sources)
	}

	if !path.IsAbs(name) {
		name = path.Join(p.OutputDir, name)
	}

	return name
}

func (p MergeParams) tempFile() (*os.File, error) {
	file, err := ioutil.TempFile(p.OutputDir, "videos-to-merge.*.txt")

	if err != nil {
		return nil, err
	}

	var list strings.Builder

	for _, v := range p.Sources {
		// file 'movie.mp4'
		list.WriteString("file '")
		list.WriteString(v)
		list.WriteString("' \n")
	}

	file.WriteString(list.String())

	return file, err
}

func (m MergedFile) fileExists() bool {
	return fileutils.FileExists(m.Name)
}
