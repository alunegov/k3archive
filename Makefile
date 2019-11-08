build-rpi:
	SET GOOS=linux& SET GOARCH=arm& SET GOARM=6& go build

build-win-32:
	SET GOOS=windows& SET GOARCH=386& go build

build-win-64:
	SET GOOS=windows& SET GOARCH=amd64& go build
