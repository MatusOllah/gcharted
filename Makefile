# settings
IS_RELEASE ?= false

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

# tools
GO = go
WINRES = $(GO) run github.com/tc-hib/go-winres@latest
UPX = upx

# output
BINARY = ./bin/$(GOOS)-$(GOARCH)

EXE = $(BINARY)/gcharted$(shell go env GOEXE)

# flags
CGO_LDFLAGS = -lmsvcrt
GO_GCFLAGS =
GO_LDFLAGS =
GO_FLAGS = -v

UPX_FLAGS = -f --best --lzma

ifeq ($(IS_RELEASE),true)
	GO_GCFLAGS += -dwarf=false
	GO_LDFLAGS += -s -w
	GO_FLAGS += -trimpath
	ifeq ($(GOOS),windows)
	GO_LDFLAGS += -H windowsgui -extldflags=-static
	endif
endif

GO_FLAGS += -gcflags="$(GO_GCFLAGS)" -ldflags="$(GO_LDFLAGS)" -buildvcs=true -buildmode=pie

.PHONY: all
all: clean build

.PHONY: run
run:
	$(GO) get
	CGO_ENABLED=1 CGO_LDFLAGS=$(CGO_LDFLAGS) GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) run $(GO_FLAGS) .

.PHONY: run-debug
run-debug:
	$(GO) get
	CGO_ENABLED=1 CGO_LDFLAGS=$(CGO_LDFLAGS) GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) run $(GO_FLAGS) . -log-level=debug

.PHONY: build
build: $(BINARY)

$(BINARY):
	mkdir -p $(BINARY)

	$(GO) get
ifeq ($(GOOS),windows)
	$(WINRES) make
endif
	CGO_ENABLED=1 CGO_LDFLAGS=$(CGO_LDFLAGS) GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build $(GO_FLAGS) -o $(EXE) .

ifeq ($(IS_RELEASE),true)
	strip $(EXE)
	$(UPX) $(UPX_FLAGS) $(EXE)
endif

.PHONY: clean
clean:
	rm -rf $(BINARY)
ifeq ($(GOOS),windows)
	rm -f ./rsrc_windows_*.syso
endif

.PHONY: clean-all
clean-all:
	rm -rf ./bin/
ifeq ($(GOOS),windows)
	rm -f ./rsrc_windows_*.syso
endif
