run:
	go run .

build.mac:
	GOOS=darwin GOARCH=arm64 go build -o mac_db

build.windows:
	GOOS=windows GOARCH=amd64 go build -o windows_db
