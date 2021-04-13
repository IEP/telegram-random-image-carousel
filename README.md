# Telegram Random Image Carousel

## Compile

`CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -v -trimpath .`

## Usage

Create `config.yaml` and `data.json` (see example files) and run with `./telegram-random-image-carousel`