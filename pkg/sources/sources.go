
package sources

import (
	"crypto/rand"
	"image"
	"image/color"
	"math/big"
	"os"
	"path/filepath"
	"strings"

	"github.com/comunidade-shallom/diakonos/pkg/files"
	"github.com/disintegration/imaging"
	"github.com/golang/freetype/truetype"
	"github.com/muesli/gamut"
	"github.com/pterm/pterm"
)

var cache = map[string][]string{}

type Sources struct {
	Footer string `fig:"footer" yaml:"footer" default:"sources/footer.png"`
	Fonts  string `fig:"fonts" yaml:"fonts" default:"sources/fonts"`
	Covers string `fig:"covers" yaml:"covers" default:"sources/covers"`
	// https://colors.muz.li/palette/976f4e/4e7197/374f6a/978a4e/6a6137
	// https://colors.muz.li/palette/24180f/0f1c24/0a1419/24200f/19160a
	//nolint:lll
	Colors []string `fig:"colors" yaml:"colors" default:"[#000000,#976f4e,#4e7197,#374f6a,#978a4e,#6a6137,#24180f,#0f1c24,#0a1419,#24200f,#19160a]"`
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

func (s Sources) RandomColor() (color.Color, error) {
	if len(s.Colors) == 0 {
		return color.Black, nil
	}

	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(s.Colors))))
	if err != nil {
		return nil, err
	}

	hex := s.Colors[nBig.Int64()]

	color := gamut.Hex(hex)

	pterm.Debug.Printfln("Color: %s", hex)

	return color, nil
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
