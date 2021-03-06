PROJECT_NAME := $(shell basename "$(PWD)")
SAMPLE_BINARY := $(PROJECT_NAME)_sample
MOCK_TARGET := mock

files = session_handler session_id_creator session_encoder
		
mock:
	@echo "  >  Generate mock class..."
	@go get github.com/golang/mock/gomock
	@go install github.com/golang/mock/mockgen
	@$(foreach file,$(files), $(GOPATH)/bin/mockgen -source=$(file).go -destination=$(MOCK_TARGET)/$(file).go -package=$(MOCK_TARGET);)
	@$(GOPATH)/bin/mockgen -source=session_manager.go -destination=mock_session_manager.go -package=phpsessgo;

standard-http-example:
	@echo "  >  Building sample..."
	@go build -o $(SAMPLE_BINARY) ./examples/standard-http-example
	@echo "  >  Run sample..."
	@./$(SAMPLE_BINARY)

echo-middleware-example:
	@echo "  >  Building sample..."
	@go build -o $(SAMPLE_BINARY) ./examples/echo-middleware-example
	@echo "  >  Run sample..."
	@./$(SAMPLE_BINARY)

.PHONY: mock standard-http-example echo-middleware-example