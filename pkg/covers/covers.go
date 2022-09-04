//nolint:varnamelen,gomnd
package covers

import (
	"image"
	"image/color"

	"github.com/comunidade-shallom/diakonos/pkg/sources"
	"github.com/disintegration/imaging"
	"gopkg.in/fogleman/gg.v1"
)

const boxMargin = 20.0

type Generator struct {
	Sources   sources.Sources
	FooterSRC string
	Size      int
	Text      string
}

func (g Generator) Generate() (image.Image, error) {
	dc := gg.NewContext(g.Size, g.Size)
	err := g.addBackgroundImage(dc)
	if err != nil {
		return nil, err
	}

	g.addBox(dc)

	err = g.addFooter(dc)
	if err != nil {
		return nil, err
	}

	err = g.addMainText(dc)
	if err != nil {
		return nil, err
	}

	return dc.Image(), nil
}

func (g Generator) getTextColor() color.Color {
	return color.White
}

func (g Generator) getTextShaddowColor() color.Color {
	return color.Black
}

func (g Generator) addBackgroundImage(dc *gg.Context) error {
	backgroundFilename, err := g.Sources.OpenRandomCover()
	if err != nil {
		return err
	}

	backgroundImage := imaging.Fill(
		backgroundFilename,
		dc.Width(),
		dc.Height(),
		imaging.Center,
		imaging.Lanczos,
	)

	dc.DrawImage(backgroundImage, 0, 0)

	return nil
}

func (g Generator) addBox(dc *gg.Context) {
	x := boxMargin
	y := boxMargin
	//nolint:gomnd
	w := float64(dc.Width()) - (2.0 * boxMargin)
	//nolint:gomnd
	h := float64(dc.Height()) - (2.0 * boxMargin)
	dc.SetColor(color.RGBA{0, 0, 0, 204})
	dc.DrawRectangle(x, y, w, h)
	dc.Fill()
}

func (g Generator) addFooterTXT(dc *gg.Context) error {
	//nolint:gomnd
	fontFace, err := g.Sources.OpenRandomFont(40)
	if err != nil {
		return err
	}

	dc.SetFontFace(fontFace)
	dc.SetColor(g.getTextColor())

	marginX := 50.0
	//nolint:gomnd
	marginY := -10.0
	textWidth, textHeight := dc.MeasureString(g.FooterSRC)

	x := float64(dc.Width()) - textWidth - marginX
	y := float64(dc.Height()) - textHeight - marginY

	dc.DrawString(g.FooterSRC, x, y)

	return nil
}

func (g Generator) addFooter(dc *gg.Context) error {
	footer, err := g.Sources.GetFooter()
	if err != nil {
		return err
	}

	if footer == nil {
		return g.addFooterTXT(dc)
	}

	width := dc.Width() / 2
	maxHeight := dc.Height() / 10

	footer = imaging.Fit(footer, width, maxHeight, imaging.Lanczos)
	size := footer.Bounds().Size()

	marginTop := dc.Height() - maxHeight
	marginLefth := (dc.Width() - size.X) / 2

	dc.DrawImage(footer, marginLefth, marginTop)

	return nil
}

func (g Generator) addMainText(dc *gg.Context) error {
	//nolint:
	fontFace, err := g.Sources.OpenRandomFont(130)
	if err != nil {
		return err
	}

	text := g.Text

	dc.SetFontFace(fontFace)

	//nolint:gomnd
	textRightMargin := boxMargin + 10.0
	//nolint:gomnd
	textTopMargin := boxMargin + 90.0
	lineSpacing := 1.5

	//nolint:gomnd
	maxWidth := float64(dc.Width()) - (textRightMargin * 2)

	x := textRightMargin
	y := textTopMargin

	dc.SetColor(g.getTextShaddowColor())
	//nolint:gomnd
	dc.DrawStringWrapped(text, x+2, y+3, 0, 0, maxWidth, lineSpacing, gg.AlignCenter)
	dc.SetColor(g.getTextColor())
	dc.DrawStringWrapped(text, x, y, 0, 0, maxWidth, lineSpacing, gg.AlignCenter)

	return nil
}
