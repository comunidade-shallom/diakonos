package covers

import (
	"fmt"
	"strconv"
	"strings"
)

type Size struct {
	Width  int
	Height int
}

func ParseSizes(raw string) []Size {
	if raw == "" {
		return []Size{}
	}

	lines := strings.Split(raw, ",")

	list := make([]Size, len(lines))

	for index, line := range lines {
		//nolint:gomnd
		vals := strings.SplitN(line, "x", 2)

		var width, height int

		//nolint:gomnd
		if len(vals) == 2 {
			width, _ = strconv.Atoi(vals[0])
			height, _ = strconv.Atoi(vals[1])
		} else {
			width, _ = strconv.Atoi(vals[0])
			height = width
		}

		list[index] = Size{
			Width:  width,
			Height: height,
		}
	}

	return list
}

func (s Size) String() string {
	return fmt.Sprintf("%vx%v", s.Width, s.Height)
}
