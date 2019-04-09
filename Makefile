PROJECT_NAME := $(shell basename "$(PWD)")
SAMPLE_BINARY := $(PROJECT_NAME)_sample

sample: build-sample run-sample

build-sample:
	@echo "  >  Building binary..."
	@go build -o $(SAMPLE_BINARY) ./sample
	
run-sample: 
	@./$(SAMPLE_BINARY)

.PHONY: sample build-sample run-sample