package sources

import (
	"crypto/rand"
	"image"
	"math/big"
	"os"
	"path/filepath"
	"strings"

	"github.com/comunidade-shallom/diakonos/pkg/files"
	"github.com/disintegration/imaging"
	"github.com/golang/freetype/truetype"
	"github.com/pterm/pterm"
)

var cache = map[string][]string{}

type Sources struct {
	Footer string `fig:"footer" yaml:"footer" default:"sources/footer.png"`
	Fonts  string `fig:"fonts" yaml:"fonts" default:"sources/fonts"`
	Covers string `fig:"covers" yaml:"covers" default:"sources/covers"`
	Colors Colors `fig:"colors" yaml:"colors"`
}

func (s Sources) ListFonts() ([]string, error) {
	return listFiles(s.Fonts)
}

func (s Sources) ListCovers() ([]string, error) {
	return listFiles(s.Covers)
}

func (s Sources) GetFooter() (image.Image, error) {
	if filepath.IsAbs(s.Footer) {
		return imaging.Open(s.Footer)
	}

	return nil, nil
}

func (s Sources) RandomFont() (string, error) {
	return randomFile(s.Fonts)
}

func (s Sources) RandomCover() (string, error) {
	return randomFile(s.Covers)
}

func (s Sources) RandomColorPallet() ColorPallet {
	return s.Colors.RandomPallet()
}

func (s Sources) OpenRandomFont() (*truetype.Font, error) {
	src, err := s.RandomFont()
	if err != nil {
		return nil, err
	}

	fontBytes, err := os.ReadFile(src)
	if err != nil {
		return nil, err
	}

	font, err := truetype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}

	pterm.Debug.Printfln("Load font: %s", files.GetRelative(src))

	return font, nil
}

func (s Sources) OpenRandomCover() (image.Image, error) {
	src, err := s.RandomCover()
	if err != nil {
		return nil, err
	}

	pterm.Debug.Printfln("Load image: %s", files.GetRelative(src))

	return imaging.Open(src)
}

func randomFile(base string) (string, error) {
	list, err := listFiles(base)
	if err != nil {
		return "", err
	}

	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(list))))
	if err != nil {
		return "", err
	}

	return list[nBig.Int64()], nil
}

func listFiles(base string) ([]string, error) {
	if cached, has := cache[base]; has {
		return cached, nil
	}

	list, err := os.ReadDir(base)
	if err != nil {
		return nil, err
	}

	res := make([]string, 0)

	for _, de := range list {
		name := de.Name()

		// ignore dotfiles
		if strings.HasPrefix(name, ".") {
			continue
		}

		res = append(res, filepath.Join(base, name))
	}

	return res, nil
}
