package covers

import (
	"github.com/comunidade-shallom/diakonos/pkg/sources"
)

type GeneratorSource struct {
	Sources   sources.Sources
	Width     int
	Height    int
	FontSize  float64
	FooterSRC string
	Text      string
}

func (g GeneratorSource) Random() (Builder, error) {
	fontSize := g.FontSize
	if fontSize == 0 {
		fontSize = 130
	}

	width := g.Width
	if width == 0 {
		width = 1080
	}

	height := g.Height
	if height == 0 {
		height = width
	}

	background, err := g.Sources.OpenRandomCover()
	if err != nil {
		return Builder{}, err
	}

	footer, err := g.Sources.GetFooter()
	if err != nil {
		return Builder{}, err
	}

	font, err := g.Sources.OpenRandomFont()
	if err != nil {
		return Builder{}, err
	}

	return Builder{
		Text:        g.Text,
		ColorPallet: g.Sources.RandomColorPallet(),
		Font:        font,
		FontSize:    fontSize,
		Height:      height,
		Width:       width,
		Background:  background,
		Footer:      footer,
		Filters:     RandomFilters(),
	}, nil
}
