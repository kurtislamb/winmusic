BINARY=winmusic

build-install: build install

build:
	go build -o ./output/winmusic ./winmusic/cmd/.

install:
	sudo mv ./output/winmusic /usr/bin/winmusic