clean:
	rm -f springman
.PHONY: clean
build: clean
	go build -o springman

run: build
	./springman
.PHONY: build run
