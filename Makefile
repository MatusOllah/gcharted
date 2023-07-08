GO=go
FYNE=fyne

TARGET=windows
BINARY=./bin

all: build

build: clean
build:
	mkdir -p $(BINARY)/$(TARGET)
	$(GO) mod tidy
	$(GO) get
	$(FYNE) build --target $(TARGET) -o $(BINARY)/$(TARGET)

clean:
	rm -rf $(BINARY)/$(TARGET)