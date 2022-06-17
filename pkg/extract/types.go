package extract

type ExtractParams struct {
	Source    string
	OutputDir string
}

type ExtractedFile struct {
	ExtractParams
	Name string
}
