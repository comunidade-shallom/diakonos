//nolint:varnamelen,gomnd
package covers

import (
	"image"
	"image/color"

	"github.com/disintegration/imaging"
	"github.com/golang/freetype/truetype"
	"github.com/muesli/gamut"
	"github.com/pterm/pterm"
	"golang.org/x/image/font"
	"gopkg.in/fogleman/gg.v1"
)

const (
	boxMargin   = 15.0
	minFontSize = 25.0
)

type Builder struct {
	Text       string
	Width      int
	Height     int
	Color      color.Color
	Font       *truetype.Font
	FontSize   float64
	Background image.Image
	Footer     image.Image
	Filters    []Filter
}

func (g Builder) WithSize(size Size) Builder {
	g.Width = size.Width
	g.Height = size.Height

	return g
}

func (g Builder) Build() image.Image {
	dc := gg.NewContext(g.Width, g.Height)

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

	backgroundImage := ApplyFilters(imaging.Fill(
		g.Background,
		dc.Width(),
		dc.Height(),
		imaging.Center,
		imaging.Lanczos,
	), g.Filters)

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
	const lineSpacing = 1.5

	text := g.Text
	W := dc.Width()
	H := dc.Height()
	P := boxMargin * 1.2

	yPad := P

	//nolint:gomnd
	maxWidth := float64(W) - (P * 2)
	//nolint:gomnd
	maxHeight := (float64(H) - (P * 2)) * 0.9

	fontSize := g.FontSize

	var font font.Face

	updateFont := func() {
		font = truetype.NewFace(g.Font, &truetype.Options{
			Size: fontSize,
		})

		dc.SetFontFace(font)
	}

	updateFont()

	for {
		if fontSize < minFontSize {
			break
		}

		updateFont()

		lines := dc.WordWrap(text, maxWidth)
		linesCount := len(lines)
		mls := ""

		for index, line := range lines {
			mls += line
			// last line
			if index != linesCount-1 {
				mls += "\n"
			}
		}

		_, stringHeight := dc.MeasureMultilineString(mls, lineSpacing)

		verticalSpace := maxHeight - stringHeight

		if (verticalSpace / 2) > P {
			yPad = verticalSpace / 2
		}

		if stringHeight < (maxHeight - (2 * P)) {
			break
		}

		fontSize -= (fontSize * 0.1)
	}

	pterm.Debug.Printfln("font size: %v", fontSize)

	dc.SetColor(g.TextShaddowColor())
	//nolint:gomnd
	dc.DrawStringWrapped(text, P+1, yPad+1, 0, 0, maxWidth, lineSpacing, gg.AlignCenter)
	dc.SetColor(g.TextColor())
	dc.DrawStringWrapped(text, P, yPad, 0, 0, maxWidth, lineSpacing, gg.AlignCenter)
}
