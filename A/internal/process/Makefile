SERVICE_NAME = process

env:
	export $(grep -v '^#' ~/go/src/snappfood/A/pkg/.env | xargs -d '\n')

build:
	go build -race -o ./bin/${SERVICE_NAME} ~/go/src/snappfood/A/cmd/process/...

run:
	go run -race ~/go/src/snappfood/A/cmd/process/main.go