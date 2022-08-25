package audios

import (
	"context"
	"os"

	"github.com/dhowden/tag"
)

type AudioTags struct {
	Title   string
	Artist  string
	Album   string
	Year    int
	URL     string
	Comment string
}

func Info(_ context.Context, source string) (AudioTags, error) {
	file, err := os.Open(source)
	if err != nil {
		return AudioTags{}, err
	}

	defer file.Close()

	meta, err := tag.ReadFrom(file)
	if err != nil {
		return AudioTags{}, err
	}

	return AudioTags{
		Title:   meta.Title(),
		Artist:  meta.Artist(),
		Album:   meta.Album(),
		Year:    meta.Year(),
		Comment: meta.Comment(),
		URL:     getURL(meta),
	}, nil
}

func getURL(meta tag.Metadata) string {
	raw := meta.Raw()

	for _, k := range []string{"WXXX", "WXX"} {
		url, ok := raw[k]

		if ok {
			val, ok := url.(*tag.Comm)

			if ok {
				return val.Text
			}

			return ""
		}
	}

	return ""
}
