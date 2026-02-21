dev:
	go build

release:
	go build -ldflags='-s -w'
	upx --brute -9 calign