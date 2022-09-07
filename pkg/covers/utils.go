package covers

import (
	"crypto/sha256"
	"encoding/base64"

	"github.com/gosimple/slug"
)

const folderMaxLength = 20

func BuildFolderName(text string) string {
	folder := slug.Make(text)
	if len(folder) > folderMaxLength {
		hasher := sha256.New()
		hasher.Write([]byte(text))

		folder = base64.StdEncoding.EncodeToString(hasher.Sum(nil))[:5] + "-" + folder[:folderMaxLength]
	}

	return folder
}
