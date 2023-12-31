run: ## Run e-commerce poject on host machine
	go run cmd/main.go start

build: ## Create executable file on host machine
	go build cmd/main -o e-commerce

# clean: ## Clean database file for a fresh start
# 	rm test.db

# test: ## Run all unit tests in the project
# 	go test -v ./...

# test-cover: ## Run all unit tests in the project with test coverage
# 	go test -v ./... -covermode=count -coverprofile=coverage.out