# Makefile

# Define the services
SERVICES := api
# Define the output directory for the built binaries
BIN_DIR := bin

# Default target
all: build

# Build each service
build: $(SERVICES)

$(SERVICES):
	@echo "Building $@..."
	@go build -o $(BIN_DIR)/$@ ./cmd/$@

# Run each service in the background and save their PIDs
.PHONY: run
run: $(SERVICES)
	@mkdir -p pids
	@for service in $(SERVICES); do \
		echo "Running $$service..."; \
		./$(BIN_DIR)/$$service & echo $$! > pids/$$service.pid; \
	done

# Run individual services
.PHONY: $(addprefix run-,$(SERVICES))
$(addprefix run-,$(SERVICES)): run-%:
	@echo "Running $*..."
	@./$(BIN_DIR)/$* & echo $$! > pids/$*.pid

# Stop all services
.PHONY: stop
stop:
	@for service in $(SERVICES); do \
		echo "Stopping $$service..."; \
		if [ -f pids/$$service.pid ]; then kill `cat pids/$$service.pid` && rm pids/$$service.pid; fi; \
	done

# Stop individual services
.PHONY: $(addprefix stop-,$(SERVICES))
$(addprefix stop-,$(SERVICES)): stop-%:
	@echo "Stopping $*..."
	@if [ -f pids/$*.pid ]; then kill `cat pids/$*.pid` && rm pids/$*.pid; fi

# Clean up built binaries and PID files
.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm -rf $(BIN_DIR) pids

# Test target (if you have tests)
.PHONY: test
test:
	@echo "Running tests..."
	@go test ./...

# Ensure output directory exists
$(BIN_DIR):
	@mkdir -p $(BIN_DIR)

# Build and run all services
.PHONY: all
all: build run
