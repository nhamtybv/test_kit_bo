cli:
	go run cmd/cli/main.go

run:
	go run cmd/web/main.go

build:
	GOOS=windows GOARCH=386 go build -o /mnt/d/temp/test_kit/test_tool.exe ./cmd/cli
	GOOS=windows GOARCH=386 go build -o /mnt/d/temp/test_kit/test_tool_web.exe ./cmd/web
	GOOS=linux GOARCH=amd64 go build -o /mnt/d/temp/test_kit/sv_amd64 ./cmd/web
clear:
	rm -rf bin