.PHONY: version test vet lint release

version:
	@cat VERSION

test:
	go test -race ./...

vet:
	go vet ./...

lint: vet test

release:
	./scripts/version.sh
