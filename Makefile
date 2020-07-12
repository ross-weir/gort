
build:
	go build ./cmd/gort
.PHONY: build

build_race:
	go build -race ./cmd/gort
.PHONY: build_race

test:
	go test -v ./...
.PHONY: test

release:
	git tag -a $(ver) -m "$(msg)"
	git push origin $(ver)
.PHONY: release