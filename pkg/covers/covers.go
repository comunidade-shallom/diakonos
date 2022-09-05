//nolint:varnamelen,gomnd
package covers

import (
	"image"
	"image/color"

	"github.com/comunidade-shallom/diakonos/pkg/sources"
	"github.com/disintegration/imaging"
	"github.com/muesli/gamut"
	"gopkg.in/fogleman/gg.v1"
)

const boxMargin = 20.0

type Generator struct {
	Sources   sources.Sources
	Size      int
	FontSize  float64
	FooterSRC string
	Text      string
}

func (g Generator) Generate() (image.Image, error) {
	dc := gg.NewContext(g.Size, g.Size)
	if err := g.addBackgroundImage(dc); err != nil {
		return nil, err
	}

	boxColor, err := g.addBox(dc)
	if err != nil {
		return nil, err
	}

	err = g.addFooter(dc, boxColor)
	if err != nil {
		return nil, err
	}

	err = g.addMainText(dc, boxColor)
	if err != nil {
		return nil, err
	}

	return dc.Image(), nil
}

func (g Generator) getTextColor(boxColor color.Color) color.Color {
	return gamut.Contrast(boxColor)
}

func (g Generator) getTextShaddowColor(boxColor color.Color) color.Color {
	return gamut.Complementary(boxColor)
}

func (g Generator) addBackgroundImage(dc *gg.Context) error {
	backgroundFilename, err := g.Sources.OpenRandomCover()
	if err != nil {
		return err
	}

	backgroundImage := MaybeApplyFilters(imaging.Fill(
		backgroundFilename,
		dc.Width(),
		dc.Height(),
		imaging.Center,
		imaging.Lanczos,
	))

	dc.DrawImage(backgroundImage, 0, 0)

	return nil
}

func (g Generator) addBox(dc *gg.Context) (color.Color, error) {
	bgColor, err := g.Sources.RandomColor()
	if err != nil {
		return bgColor, err
	}

	bgColor = color.RGBAModel.Convert(bgColor)

	R, G, B, _ := bgColor.RGBA()

	boxColor := color.RGBA{
		R: uint8(R >> 8),
		G: uint8(G >> 8),
		B: uint8(B >> 8),
		A: 204,
	}

	x := boxMargin
	y := boxMargin
	//nolint:gomnd
	w := dc.Width() - (2.0 * boxMargin)
	//nolint:gomnd
	h := dc.Height() - (2.0 * boxMargin)

	box := gg.NewContext(w, h)
	box.SetColor(boxColor)
	box.DrawRectangle(0, 0, float64(w), float64(h))
	box.Fill()

	dc.DrawImage(box.Image(), int(x), int(y))

	return boxColor, nil
}

func (g Generator) addFooterTXT(dc *gg.Context, boxColor color.Color) error {
	//nolint:gomnd
	fontFace, err := g.Sources.OpenRandomFont(40)
	if err != nil {
		return err
	}

	dc.SetFontFace(fontFace)
	dc.SetColor(g.getTextColor(boxColor))

	marginX := 50.0
	//nolint:gomnd
	marginY := -10.0
	textWidth, textHeight := dc.MeasureString(g.FooterSRC)

	x := float64(dc.Width()) - textWidth - marginX
	y := float64(dc.Height()) - textHeight - marginY

	dc.DrawString(g.FooterSRC, x, y)

	return nil
}

func (g Generator) addFooter(dc *gg.Context, boxColor color.Color) error {
	footer, err := g.Sources.GetFooter()
	if err != nil {
		return err
	}

	if footer == nil {
		return g.addFooterTXT(dc, boxColor)
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

func (g Generator) addMainText(dc *gg.Context, boxColor color.Color) error {
	fontSize := g.FontSize
	if fontSize == 0 {
		fontSize = 130
	}

	fontFace, err := g.Sources.OpenRandomFont(fontSize)
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

	dc.SetColor(g.getTextShaddowColor(boxColor))
	//nolint:gomnd
	dc.DrawStringWrapped(text, x+2, y+3, 0, 0, maxWidth, lineSpacing, gg.AlignCenter)
	dc.SetColor(g.getTextColor(boxColor))
	dc.DrawStringWrapped(text, x, y, 0, 0, maxWidth, lineSpacing, gg.AlignCenter)

	return nil
}
