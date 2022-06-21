package merge

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
)

func buildHash(sources []string) string {
	var buffer bytes.Buffer

	for _, v := range sources {
		buffer.WriteString(v)
		buffer.WriteString("#")
	}

	sum := sha256.Sum256([]byte(buffer.String()))

	return base64.RawURLEncoding.EncodeToString(sum[:])
}

func buildHashName(sources []string) string {
	return buildHash(sources) + ".mp4"
}
