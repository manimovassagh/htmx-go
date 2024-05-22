build:
	go build -o bin/main

htmlgen:
	templ generate

run: build
	@templ generate
	./bin/main