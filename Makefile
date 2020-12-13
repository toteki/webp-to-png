all: strip build run

build:
	go build main.go

run:
	./main

strip:
	rm -rf .git
	rm .gitignore
	rm LICENSE
	rm README.md
