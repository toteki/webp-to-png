all: build strip

build:
	go build main.go

strip:
	rm -rf .git
	rm .gitignore
	rm LICENSE
	rm README.md
	rm main.go
	rm Makefile
	mkdir input
	mkdir output
