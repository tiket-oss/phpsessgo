PROJECT_NAME := $(shell basename "$(PWD)")
SAMPLE_BINARY := $(PROJECT_NAME)_sample

## build-sample: Build the binary for sample
build-sample:
	@echo "  >  Building binary..."
	@go build -o $(SAMPLE_BINARY) ./sample

.PHONY: build-sample