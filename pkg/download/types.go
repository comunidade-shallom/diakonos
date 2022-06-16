package download

type DownloadParams struct {
	From      string
	OutputDir string
	Quality   string
	MimeType  string
}

type DownloadedFile struct {
	DownloadParams
	Name string
}
