# gosearch developer tasks. Run `make help` to see all commands.

.PHONY: help build run test race cover fmt vet install demo clean

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2}'

build: ## Build the ./gosearch binary
	go build -o gosearch .

run: ## Build and run (use ARGS="..."), e.g. make run ARGS="-i func ."
	go run . $(ARGS)

test: ## Run all tests
	go test ./...

race: ## Run tests with the race detector
	go test -race ./...

cover: ## Run tests with coverage
	go test -cover ./...

fmt: ## Format all code
	go fmt ./...

vet: ## Run go vet
	go vet ./...

install: ## Install gosearch into your GOBIN
	go install .

demo: ## Run a quick self-demo (searches this project)
	./demo.sh

clean: ## Remove build artifacts
	rm -f gosearch cover.out
