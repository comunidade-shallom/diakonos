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
		list[index] = ParseSize(line)
	}

	return list
}

func ParseSize(raw string) Size {
	//nolint:gomnd
	vals := strings.SplitN(raw, "x", 2)

	var width, height int

	//nolint:gomnd
	if len(vals) == 2 {
		width, _ = strconv.Atoi(vals[0])
		height, _ = strconv.Atoi(vals[1])
	} else {
		width, _ = strconv.Atoi(vals[0])
		height = width
	}

	return Size{
		Width:  width,
		Height: height,
	}
}

func BuildSizes(lines []string) []Size {
	res := make([]Size, len(lines))

	for index, line := range lines {
		res[index] = ParseSize(line)
	}

	return res
}

func (s Size) String() string {
	return fmt.Sprintf("%vx%v", s.Width, s.Height)
}
