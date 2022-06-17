# diakonos

Tools to speed media content development

## Requeriments

- [ffmpeg](https://ffmpeg.org/download.html)

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

- [Go ~> 1.18](https://go.dev/dl/)
- [Task v3](https://taskfile.dev/)
- [ffmpeg v5](https://ffmpeg.org/)
