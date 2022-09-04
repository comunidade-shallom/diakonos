//nolint:gomnd
package covers

import (
	"crypto/rand"
	"image"
	"math/big"

	"github.com/disintegration/imaging"
	"github.com/pterm/pterm"
)

type Filter func(source image.Image) image.Image

var filters map[string]Filter = map[string]Filter{
	"Blur":        ApplyBlur,
	"AdjustGamma": ApplyAdjustGamma,
	"Sharpen":     ApplySharpen,
}

func ApplySharpen(source image.Image) image.Image {
	sigma := float64(Intn(1<<3))/(1<<3) + 1

	return imaging.Sharpen(source, sigma)
}

func ApplyAdjustGamma(source image.Image) image.Image {
	sigma := float64(Intn(1<<3))/(1<<3) + 1

	return imaging.AdjustGamma(source, sigma)
}

func ApplyBlur(source image.Image) image.Image {
	sigma := float64(Intn(1<<3))/(1<<3) + 1

	return imaging.Blur(source, sigma)
}

func MaybeApplyFilters(source image.Image) image.Image {
	for name, filter := range filters {
		if (Intn(50) % 2) == 0 {
			pterm.Debug.Printfln("Applying %s", name)

			source = filter(source)
		}
	}

	return source
}

func Intn(max int64) int64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		panic(err)
	}

	return nBig.Int64()
}
