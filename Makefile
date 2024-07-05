all: build test
 
test:
	go test -cover

cover:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out

clean:
	go clean

publish: 
	git push --tags
	git push
	GOPROXY=proxy.golang.org go list -m github.com/rodgco/bigbitvector@$(VERSION)
	@echo "Published version $(VERSION)"

info:
	@echo Makefile-SemVer demonstration, use autocomplete to see available commands.

# https://github.com/malcos/makefile-semver
include Makefile.semver
