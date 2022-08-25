package audios

import (
	"context"
	"fmt"

	"github.com/bogem/id3v2/v2"
)

type AudioTags struct {
	Title   string
	Artist  string
	Album   string
	Year    string
	URL     string
	Comment string
}

func Info(_ context.Context, source string) (AudioTags, error) {
	tag, err := id3v2.Open(source, id3v2.Options{Parse: true})
	if err != nil {
		return AudioTags{}, err
	}

	defer tag.Close()

	frames := tag.AllFrames()

	for key, values := range frames {
		fmt.Println(". " + key)

		for _, frame := range values {
			ud, ok := frame.(id3v2.UserDefinedTextFrame)

			if ok {
				fmt.Println(".. " + ud.Description)
				fmt.Println(".. " + ud.Value)

				continue
			}

			tx, ok := frame.(id3v2.TextFrame)

			if ok {
				fmt.Println(".. " + tx.Text)

				continue
			}

			pic, ok := frame.(id3v2.PictureFrame)

			if ok {
				fmt.Println(".. " + pic.Description)
				fmt.Println(".. " + pic.MimeType)

				continue
			}

			cm, ok := frame.(id3v2.CommentFrame)

			if ok {
				fmt.Println(".. " + cm.Description)
				fmt.Println(".. " + cm.Text)

				continue
			}

			id, ok := frame.(id3v2.UFIDFrame)

			if ok {
				fmt.Println(".. " + id.OwnerIdentifier)
				fmt.Println(".. " + id.UniqueIdentifier())
				fmt.Println(".. " + string(id.Identifier))

				continue
			}

			cp, ok := frame.(id3v2.ChapterFrame)

			if ok {
				fmt.Println(".. " + cp.ElementID)
				fmt.Println(".. " + cp.Title.Text)
				fmt.Println(".. " + cp.Description.Text)
				fmt.Println(".. " + cp.UniqueIdentifier())

				continue
			}

			fmt.Println(".. Unknow")
		}

		fmt.Println("...")
	}

	return AudioTags{
		Title:  tag.Title(),
		Artist: tag.Artist(),
		Album:  tag.Album(),
		Year:   tag.Year(),
		// URL:     tag.GetTextFrame("WXXX").Text,
		Comment: getComment(tag),
	}, err
}

func getComment(tag *id3v2.Tag) string {
	comments := tag.GetFrames(tag.CommonID("Comments"))
	if len(comments) > 0 {
		cm, ok := comments[0].(id3v2.CommentFrame)

		if ok {
			return cm.Text
		}

		return ""
	}

	return ""
}
