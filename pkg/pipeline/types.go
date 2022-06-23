package pipeline

type Action string
type Source struct {
	Value string
}

type Output struct {
	Filename string
}

const (
	Download     Action = "download"
	CutVideo     Action = "cut-video"
	ExtractAudio Action = "extract-audio"
)
