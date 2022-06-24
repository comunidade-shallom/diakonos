package pipeline

type (
	Action string
	Source struct {
		Value string
	}
)

const (
	Download     Action = "download"
	CutVideo     Action = "cut-video"
	ExtractAudio Action = "extract-audio"
)
