SHELL = /bin/bash

# Package and target names
APP_NAME = $(shell pwd | sed 's:.*/::')
TARGET = $(APP_NAME)

# Cononical version number provided to the target
VERSION := 0.0.1

# Linker flags to provide versioning to the target
LDFLAGS=-ldflags "-X=main.Version=$(VERSION)"

# Go paths and programs
BIN    = $(GOPATH)/bin
GOLINT = $(BIN)/golint
BINDATA = $(BIN)/go-bindata
GOBUILD = go build $(LDFLAGS) -o $(TARGET) ./server
GOTEST = go test -v -cover ./server

# Bindata assets
ASSETS= server/assets.go
BINDATA_FLAGS = -pkg=main -prefix=static -ignore=/*/.DS_*

# Node commands
NPM_BUILD = npm run build
NPM_TEST = npm run test
NPM_DEV = npm run start
BUNDLE = static/js/bundle.js
STYLE = static/css/style.css

build: clean assets
	$(GOTEST)
	$(GOBUILD)
	@echo "Done"
	
clean:
	@echo "Cleaning..."
	-$(shell rm -f $(BUNDLE))
	-$(shell rm -f $(STYLE))
	-$(shell rm -f $(ASSETS))
	-$(shell rm -f $(TARGET))

assets:
	@echo "Processing webpack and bindata"
	@$(NPM_BUILD)
	@$(BINDATA) $(BINDATA_FLAGS) -o $(ASSETS) static/...

backend: clean assets
	@echo "Building and unning the Go server..."
	$(GOTEST)
	$(GOBUILD)
	./$(TARGET)

frontend: clean
	@echo "Building and running the Node dev server..."
	$(NPM_DEV)
