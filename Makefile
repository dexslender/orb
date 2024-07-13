OUT := "bin/orb.bot"
VERSION := $(shell git rev-parse --short HEAD)
FLAGS := -ldflags="-X 'main.version=${VERSION}'-w -s"

all: build
buildx: build exec

build:
	go build -o $(OUT) .
exec:
	@echo Executing binary:
	./$(OUT)
buildf:
	@echo Building with flags
	go build $(FLAGS) -o $(OUT) .
run:
	go run .
clean:
	rm $(OUT)

test:
	go test -v .
