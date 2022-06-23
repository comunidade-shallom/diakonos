package merge

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Config struct {
	OutputDir string `fig:"output_dir" yaml:"output_dir" default:"outputs/merges"`
}

type Params struct {
	Sources   []string
	OutputDir string
	Name      string
}

func (c Config) Apply(params Params) (Params, error) {
	if params.OutputDir == "" {
		params.OutputDir = c.OutputDir
	}

	pwd, _ := os.Getwd()

	for i, v := range params.Sources {
		if path.IsAbs(v) {
			continue
		}

		params.Sources[i] = path.Join(pwd, v)
	}

	return params, nil
}

func (p Params) Filename() string {
	name := p.Name

	if name == "" {
		name = p.HashName()
	}

	if !path.IsAbs(name) {
		name = path.Join(p.OutputDir, name)
	}

	return name
}

func (p Params) tempFile() (*os.File, error) {
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

	file.WriteString(list.String()) //nolint:errcheck

	return file, err
}

func (p Params) HashName() string {
	var buffer bytes.Buffer

	for _, v := range p.Sources {
		buffer.WriteString(v)
		buffer.WriteString("#")
	}

	sum := sha256.Sum256(buffer.Bytes())
	ext := filepath.Ext(p.Sources[0])

	return base64.RawURLEncoding.EncodeToString(sum[:]) + ext
}
