
template:
	templ generate --watch
build: tmp
	go build -o bin/main

run: build
	./bin/main