help: Makefile
	@printf "${BLUE}Choose a command run:${NC}\n"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/    /'


## make build: Build etherman
build:
	@mkdir -p build
	@go build -o build/etherman main.go
	@printf "Build etherman successfully\n"

## make install: Install etherman
install:
	@go install
	@printf "Install etherman successfully\n"

lint:
	golangci-lint run

.PHONY: build lint install