APP := sorter
BIN_DIR := bin
CMD_DIR := ./cmd/$(APP)
GOCACHE := $(CURDIR)/.gocache

.PHONY: build test run clean

build:
	mkdir -p $(BIN_DIR)
	GOCACHE=$(GOCACHE) go build -o $(BIN_DIR)/$(APP) $(CMD_DIR)

test:
	GOCACHE=$(GOCACHE) go test ./...

run:
	GOCACHE=$(GOCACHE) go run $(CMD_DIR)

clean:
	rm -rf $(BIN_DIR) .gocache
