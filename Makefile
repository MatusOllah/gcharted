GO=go
FYNE=fyne

TARGET=windows
BINARY=./bin/$(TARGET)

all: build

build: clean
build:
	mkdir $(BINARY)
	$(GO) mod tidy
	$(GO) get
	$(FYNE) build --target $(TARGET) -o $(BINARY)

clean:
	rm -rf $(BINARY)