# settings
IS_RELEASE ?= false

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

# tools
GO = go
WINRES = $(GO) run github.com/tc-hib/go-winres@latest
WINDEPLOYQT = windeployqt6
UPX = upx

# output
BINARY = ./bin/$(GOOS)-$(GOARCH)

EXE = $(BINARY)/gcharted$(shell go env GOEXE)

# flags
UPX_FLAGS = --best --lzma

GO_GCFLAGS =
GO_LDFLAGS =
GO_FLAGS = -v

ifeq ($(IS_RELEASE),true)
	GO_GCFLAGS += -dwarf=false
	GO_LDFLAGS += -s -w
	GO_FLAGS += -trimpath
	ifeq ($(GOOS),windows)
	GO_LDFLAGS += -H windowsgui
	endif
endif

GO_FLAGS += -gcflags="$(GO_GCFLAGS)" -ldflags="$(GO_LDFLAGS)" -buildvcs=true

.PHONY: all
all: build upx

.PHONY: run
run:
	$(GO) get
	CGO_ENABLED=1 GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) run $(GO_FLAGS) .

.PHONY: build
build: clean
	mkdir -p $(BINARY)

	$(GO) get
ifeq ($(GOOS),windows)
	$(WINRES) make
endif
	CGO_ENABLED=1 GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build $(GO_FLAGS) -o $(EXE) .
ifeq ($(GOOS),windows)
	$(WINDEPLOYQT) $(EXE) --compiler-runtime
	./scripts/ldd_deploy.sh -i $(EXE) -o $(BINARY)
endif

.PHONY: upx
upx: build
ifeq ($(IS_RELEASE),true)
	$(UPX) $(UPX_FLAGS) $(EXE)
endif

.PHONY: clean
clean:
	rm -rf $(BINARY)
ifeq ($(GOOS),windows)
	rm -f rsrc_windows_*.syso
endif
