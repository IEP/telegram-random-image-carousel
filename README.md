# Telegram Random Image Carousel

## Compile

`CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -v -trimpath .`

## Usage

Create `config.yaml` and `data.json` (see example files) and finally run the bot by using `./telegram-random-image-carousel`.

Get the random image by triggering it using `/bucketName` command.