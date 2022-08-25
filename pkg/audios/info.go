package audios

import (
	"context"
	"os"
	"strconv"

	"github.com/bogem/id3v2/v2"
	"github.com/comunidade-shallom/diakonos/pkg/support/collection"
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

func WriteTags(_ context.Context, tags AudioTags, source string) error {
	file, err := id3v2.Open(source, id3v2.Options{Parse: false})
	if err != nil {
		return err
	}

	defer file.Close()

	file.SetTitle(tags.Title)
	file.SetArtist(tags.Artist)
	file.SetAlbum(tags.Album)

	if tags.Year > 0 {
		file.SetYear(strconv.Itoa(tags.Year))
	}

	comment := id3v2.CommentFrame{
		Encoding: id3v2.EncodingUTF8,
		Language: "por",
		Text:     tags.Comment,
	}

	file.AddCommentFrame(comment)

	url := id3v2.UserDefinedTextFrame{
		Encoding: id3v2.EncodingUTF8,
		Value:    tags.URL,
	}

	file.AddFrame("WXXX", url)

	return file.Save()
}

func (t AudioTags) FromRaw(raw collection.Params) AudioTags {
	t.Title = raw.String("title")
	t.Artist = raw.String("artist")
	t.Album = raw.String("album")
	t.URL = raw.String("url")
	t.Comment = raw.String("comment")
	t.Year = raw.Int("year")

	return t
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
