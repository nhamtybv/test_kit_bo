cli:
	go run cmd/cli/main.go

run:
	go run cmd/web/main.go

build:
	GOOS=windows GOARCH=amd64 go build -o /mnt/d/temp/test_kit/cli_tool.exe ./cmd/cli
	GOOS=windows GOARCH=amd64 go build -o /mnt/d/temp/test_kit/web_tool.exe ./cmd/web
#	GOOS=linux GOARCH=amd64 go build -o /mnt/d/temp/test_kit/sv_amd64 ./cmd/web
clear:
	rm -rf bin