//nolint:gomnd
package covers

import (
	"image"

	"github.com/comunidade-shallom/diakonos/pkg/support"
	"github.com/disintegration/imaging"
	"github.com/pterm/pterm"
)

type Filter func(source image.Image) image.Image

var filters = map[string]Filter{
	"Blur":        ApplyBlur,
	"AdjustGamma": ApplyAdjustGamma,
	"Sharpen":     ApplySharpen,
}

func ApplySharpen(source image.Image) image.Image {
	sigma := float64(support.RandInt(1<<3))/(1<<3) + 1

	return imaging.Sharpen(source, sigma)
}

func ApplyAdjustGamma(source image.Image) image.Image {
	sigma := float64(support.RandInt(1<<3))/(1<<3) + 1

	return imaging.AdjustGamma(source, sigma)
}

func ApplyBlur(source image.Image) image.Image {
	sigma := float64(support.RandInt(1<<3))/(1<<3) + 1

	return imaging.Blur(source, sigma)
}

func MaybeApplyFilters(source image.Image) image.Image {
	for name, filter := range filters {
		if (support.RandInt(50) % 2) == 0 {
			pterm.Debug.Printfln("Applying %s", name)

			source = filter(source)
		}
	}

	return source
}

func ApplyFilters(source image.Image, list []Filter) image.Image {
	for _, filter := range list {
		source = filter(source)
	}

	return source
}

func RandomFilters() []Filter {
	res := []Filter{}

	for name, filter := range filters {
		if (support.RandInt(50) % 2) == 0 {
			pterm.Debug.Printfln("Selecting %s", name)

			res = append(res, filter)
		}
	}

	return res
}
