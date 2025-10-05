BIN := "./app"

build:
	go build -o $(BIN)

run: build
	$(BIN) -from