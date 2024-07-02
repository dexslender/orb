OUT = "bin/orb"

all: build
build:
	go build -o $(OUT) .
exec:
	@echo Executing binary:
	./$(OUT)
run:
	go run .
clean:
	rm $(OUT)

test:
	go test -v .
