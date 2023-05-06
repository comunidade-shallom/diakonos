# diakonos

Tools to speed media content development

## Requeriments

- [ffmpeg](https://ffmpeg.org/download.html)

## Usage

```
diakonos --help
```

### Config

All configs are load from `diakonos.yml`.

This file will be created if not exist.

```yaml
# Base output dir, will be used as bed of other output dirs
base_output_dir: ./outputs

# default download options
download:
  output_dir: ./downloads
  quality: hd1080
  mime_type: mp4

# cut options
cut:
  output_dir: ./cuts

# audio options
audio:
  output_dir: ./audios

# merge options
merge:
  output_dir: ./merges

# source options (cover generation)
sources:
  # footer image to be used in covers
  footer: ./sources/footer.png
  fonts: ./sources/fonts
  covers: ./sources/covers
  colors:
    - "#000000"
    - "#976f4e"
    - "#4e7197"
    - "#374f6a"
    - "#978a4e"
    - "#6a6137"
    - "#24180f"
    - "#0f1c24"
    - "#0a1419"
    - "#24200f"
    - "#19160a"
```

### Commands

#### `diakonos config`

Display loaded config

```sh
diakonos config
```

#### `diakonos cover`

Generate covers

```sh
diakonos cover "Text to be used in cover"
```

```sh
diakonos cover --font-size 170 "Custom font size"
diakonos cover --sizes 1280x720,1080x1080 "Multiple sizes"
diakonos cover --width 1280 --height 720 "Custom cover size"
diakonos cover --sizes 1280x720,1080x1080 --times 10 "Generate 20 covers"
```

#### `diakonos audio normalize`

Normalize the audio of a MP3 file.

```sh
diakonos audio normalize /path/to/audio/file.mp3
```

#### `diakonos audio info`

Display infos from a MP3 file.

```sh
diakonos audio info /path/to/audio/file.mp3
```

#### `diakonos video cut`

Cut video using a time period.

```sh
diakonos video cut --start 12m --finish 73m /path/to/video/file.mp4
```

#### `diakonos video extract`

Extract audio from a video file

```sh
diakonos video extract /path/to/video/file.mp4
```

#### `diakonos youtube download`

Download a youtube video from a URL

```sh
diakonos youtube download --quality 1080p https://www.youtube.com/watch?v=gs2dr7jzX-M
```

#### `diakonos youtube cut`

Download a youtube video, cut and extract the audio.

```sh
diakonos youtube cut --extract-audio --quality 1080p --start 12m --finish 73m  https://www.youtube.com/watch?v=gs2dr7jzX-M
```

#### `diakonos pipeline`

Dynamic process a list of commands in a single pipeline process

```sh
diakonos pipeline pipeline.file.yml
```

##### Pipeline YALM

```yaml
# Data will be used as template variable inside other parts of YAML file.
data:
  author: Paul
  date: 28/08/2022
  name: In the beginning was the Word, and the Word was with God, and the Word was God.
  source: https://www.youtube.com/watch?v=TG1Xrdvf5Qw

# name will be used in output folder
name: "{{ .Data.name }}"

# list of acions
actions:
  # unique identifier of the action
  - id: raw
    # type of action
    type: youtube-download
    # source, it will be used in the action
    source: "{{ .Data.source }}"

  - id: cut
    type: video-cut
    # sources can make reference to the output of other actions
    source:
      action: raw
    # custom parameters for this action
    params:
      start: 13m
      finish: 1h

  - id: audio
    type: video-extract-audio
    source:
      action: video-cut

  - id: audio-normalized
    type: audio-normalize
    source:
      action: audio

  - id: audio-add-tags
    type: audio-define-tags
    source:
      action: audio-normalized
    params:
      title: "{{ .Data.name }} - {{ .Data.author }} - {{ .Data.date }}"
      artist: Comunidade Bastista Shallom em Meriti
      album: Comunidade Shallom 2022
      url: https://comunidadeshallom.com.br
      year: 2022
      comment: |
        Whatch the full version in {{ .Data.source }}

  - id: cover-audio
    type: cover-generate
    params:
      text: "{{ .Data.name }}"
      times: 20
      font-size: 60
      sizes:
        - 1080x1080
        - 1280x720
```

## Development

```sh
task setup # install dependencies
task build # build cli
```

```sh
# run from source
task run:cli -- youtube download 'https://www.youtube.com/watch?v=8yAbX8W3Caw'

# build and run
task build
./bin/diakonos-cli-linux-amd64 config
```

### Requeriments

- [Go ~> 1.20](https://go.dev/dl/)
- [Task v3](https://taskfile.dev/)
- [ffmpeg v5](https://ffmpeg.org/)
