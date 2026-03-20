BINARY := ghpp
MODULE := github.com/douhashi/gh-project-promoter
GOFLAGS := -trimpath
LDFLAGS := -s -w

.PHONY: build test lint clean install

build:
	CGO_ENABLED=0 go build $(GOFLAGS) -ldflags '$(LDFLAGS)' -o $(BINARY) .

test:
	go test ./...

lint:
	go vet ./...

clean:
	rm -f $(BINARY)

install: build
	mv $(BINARY) $(GOPATH)/bin/$(BINARY)
