set-up: ## setup go mod
	git config --local commit.template .github/.commit_template
	go mod tidy

generate: ## generate and format code
	go generate
	find . -print | grep --regex '.*\.go' | xargs goimports -w -local "github.com/knakazawa9/package"

test: ## exec go test
	go test -v ./...

help: ## show help
	@echo "Usage: make [command] \n"
	@echo 'The commands are: '
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[39m%s: \033[33m%s\n", $$1, $$2}'
