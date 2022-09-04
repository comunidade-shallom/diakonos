//nolint:ireturn
package sources

import (
	"crypto/rand"
	"image"
	"math/big"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
	"golang.org/x/image/font"
	"gopkg.in/fogleman/gg.v1"
)

type Sources struct {
	Footer string `fig:"footer" yaml:"footer" default:"sources/footer.png"`
	Fonts  string `fig:"fonts" yaml:"fonts" default:"sources/fonts"`
	Covers string `fig:"covers" yaml:"covers" default:"sources/covers"`
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

func (s Sources) OpenRandomFont(points float64) (font.Face, error) {
	src, err := s.RandomFont()
	if err != nil {
		return nil, err
	}

	return gg.LoadFontFace(src, points)
}

func (s Sources) OpenRandomCover() (image.Image, error) {
	src, err := s.RandomCover()
	if err != nil {
		return nil, err
	}

	return imaging.Open(src)
}

func randomFile(base string) (string, error) {
	list, err := os.ReadDir(base)
	if err != nil {
		return "", err
	}

	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(list))))
	if err != nil {
		return "", err
	}

	return filepath.Join(base, list[nBig.Int64()].Name()), nil
}

func listFiles(base string) ([]string, error) {
	list, err := os.ReadDir(base)
	if err != nil {
		return nil, err
	}

	res := make([]string, len(list))

	for i, de := range list {
		res[i] = filepath.Join(base, de.Name())
	}

	return res, nil
}
