package fileutils

import (
	"encoding/json"
	"time"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

type VideoFormat struct {
	FileName       string            `json:"filename"`
	BitRate        string            `json:"bit_rate"`
	FormatLongName string            `json:"format_long_name"`
	FormatName     string            `json:"format_name"`
	ProbeScore     int               `json:"probe_score"`
	Programs       int               `json:"nb_programs"`
	Size           string            `json:"size"`
	Duration       string            `json:"duration"`
	StartTime      string            `json:"start_time"`
	Streams        int               `json:"nb_streams"`
	Tags           map[string]string `json:"tags"`
}

type VideoStream struct {
	CodecName      string            `json:"codec_name"`
	CodecType      string            `json:"codec_type"`
	CodecLongName  string            `json:"codec_long_name"`
	CodecTag       string            `json:"codec_tag"`
	CodecTagString string            `json:"codec_tag_string"`
	Index          int               `json:"index"`
	Profile        string            `json:"profile"`
	Tags           map[string]string `json:"tags"`
}

type VideoInfo struct {
	Streams []VideoStream `json:"streams"`
	Format  VideoFormat   `json:"format"`
}

func GetVideoInfo(source string) (info VideoInfo, err error) {
	j, err := ffmpeg.ProbeWithTimeout(source, time.Second*5, ffmpeg.KwArgs{})
	if err != nil {
		return info, err
	}

	err = json.Unmarshal([]byte(j), &info)

	return info, err
}

func GetAudioStreams(source string) (list []VideoStream, err error) {
	info, err := GetVideoInfo(source)
	if err != nil {
		return list, err
	}

	for _, s := range info.Streams {
		if s.CodecType == "audio" {
			list = append(list, s)
		}
	}

	return list, err
}
