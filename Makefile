GO=go

BINARY=./bin

FLAGS=-v

GOOS=windows
GOARCH=amd64

all: build

build:
	mkdir $(BINARY)
	mkdir $(BINARY)/$(GOOS)-$(GOARCH)
	$(GO) mod tidy
	$(GO) get
	GO111MODULE=on GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build $(FLAGS) -o $(BINARY)/$(GOOS)-$(GOARCH)

clean:
	rm -rf $(BINARY)