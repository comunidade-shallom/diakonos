//nolint:varnamelen,gomnd
package covers

import (
	"image"
	"image/color"

	"github.com/disintegration/imaging"
	"github.com/muesli/gamut"
	"golang.org/x/image/font"
	"gopkg.in/fogleman/gg.v1"
)

const boxMargin = 20.0

type Builder struct {
	Text       string
	Size       int
	Color      color.Color
	Font       font.Face
	Background image.Image
	Footer     image.Image
}

func (g Builder) Build() image.Image {
	dc := gg.NewContext(g.Size, g.Size)

	g.addBackground(dc)
	g.addBox(dc)
	g.addMainText(dc)
	g.addFooter(dc)

	return dc.Image()
}

func (g Builder) TextColor() color.Color {
	return gamut.Contrast(g.Color)
}

func (g Builder) TextShaddowColor() color.Color {
	return gamut.Complementary(g.Color)
}

func (g Builder) addBackground(dc *gg.Context) {
	if g.Background == nil {
		return
	}

	backgroundImage := MaybeApplyFilters(imaging.Fill(
		g.Background,
		dc.Width(),
		dc.Height(),
		imaging.Center,
		imaging.Lanczos,
	))

	dc.DrawImage(backgroundImage, 0, 0)
}

func (g Builder) addBox(dc *gg.Context) {
	R, G, B, _ := color.RGBAModel.Convert(g.Color).RGBA()

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
}

func (g Builder) addFooter(dc *gg.Context) {
	width := dc.Width() / 2
	maxHeight := dc.Height() / 10

	footer := imaging.Fit(g.Footer, width, maxHeight, imaging.Lanczos)
	size := footer.Bounds().Size()

	marginTop := dc.Height() - maxHeight
	marginLefth := (dc.Width() - size.X) / 2

	dc.DrawImage(footer, marginLefth, marginTop)
}

func (g Builder) addMainText(dc *gg.Context) {
	text := g.Text

	dc.SetFontFace(g.Font)

	//nolint:gomnd
	textRightMargin := boxMargin + 10.0
	//nolint:gomnd
	textTopMargin := boxMargin + 90.0
	lineSpacing := 1.5

	//nolint:gomnd
	maxWidth := float64(dc.Width()) - (textRightMargin * 2)

	x := textRightMargin
	y := textTopMargin

	dc.SetColor(g.TextShaddowColor())
	//nolint:gomnd
	dc.DrawStringWrapped(text, x+2, y+3, 0, 0, maxWidth, lineSpacing, gg.AlignCenter)
	dc.SetColor(g.TextColor())
	dc.DrawStringWrapped(text, x, y, 0, 0, maxWidth, lineSpacing, gg.AlignCenter)
}
