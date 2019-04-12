PROJECT_NAME := $(shell basename "$(PWD)")
SAMPLE_BINARY := $(PROJECT_NAME)_sample
MOCK_TARGET := mock

files = session_handler session_id_creator session_encoder

sample: build-sample run-sample

build-sample:
	@echo "  >  Building sample..."
	@go build -o $(SAMPLE_BINARY) ./sample
	
run-sample: 
	@echo "  >  Run sample..."
	@./$(SAMPLE_BINARY)
		
mock:
	@echo "  >  Generate mock class..."
	@go get github.com/golang/mock/gomock
	@go install github.com/golang/mock/mockgen
	@$(foreach file,$(files), $(GOPATH)/bin/mockgen -source=$(file).go -destination=$(MOCK_TARGET)/$(file).go -package=$(MOCK_TARGET);)

.PHONY: sample build-sample run-sample  mock