all: build test
 
test:
	go test -cover

cover:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out

clean:
	go clean

format:
	gofmt -s -w .

push: 
	@echo "Checking if in main branch"
	@git branch --show-current | grep -q "main" || (echo "Not in main branch, exiting"; exit 1)

	git tag v$(VERSION)
	git push --tags
	git push
	GOPROXY=proxy.golang.org go list -m github.com/rodgco/bigbitvector@v$(VERSION)
	@echo "Published version $(VERSION)"

info:
	@echo Makefile-SemVer demonstration, use autocomplete to see available commands.
	@echo Current version: v$(VERSION)

# https://github.com/malcos/makefile-semver
include Makefile.semver
