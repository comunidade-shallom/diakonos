package sources

import (
	"image/color"
	"strings"

	"github.com/comunidade-shallom/diakonos/pkg/support"
	"github.com/muesli/gamut"
)

// https://colors.muz.li/palette/976f4e/4e7197/374f6a/978a4e/6a6137
// https://colors.muz.li/palette/24180f/0f1c24/0a1419/24200f/19160a

const defaultColors = "#000000,#976f4e,#4e7197,#374f6a,#978a4e,#6a6137,#24180f,#0f1c24,#0a1419,#24200f,#19160a"

type (
	// Colors collection.
	Colors []string
	// Color Pallet.
	ColorPallet struct {
		// Base color, to be used as background
		Base color.Color
		// Text color, contract from base
		Text color.Color
		// Shadow color, complementary from base
		Shadow color.Color
	}
)

// Random color.
func (cl Colors) Random() color.Color {
	if len(cl) == 0 {
		return color.Black
	}

	hex := cl[support.RandInt(int64(len(cl)))]

	return gamut.Hex(hex)
}

// RandomPallet.
func (cl Colors) RandomPallet() ColorPallet {
	base := cl.Random()

	return ColorPallet{
		Base:   base,
		Text:   gamut.Contrast(base),
		Shadow: gamut.Complementary(base),
	}
}

func DefaultColors() []string {
	return strings.Split(defaultColors, ",")
}
