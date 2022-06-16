package cut

import "time"

type CutParams struct {
	Source    string
	OutputDir string
	Start     time.Duration
	Finish    time.Duration
}

type CroppedFile struct {
	CutParams
	Name string
}
