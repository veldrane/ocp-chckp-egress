 # Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=chckpApi
PREFIX=/usr/local
CMD_DIR=cmd/checkpoint

all: build



build:
	$(GOBUILD) -o $(BINARY_NAME) $(CMD_DIR)/main.go  $(CMD_DIR)/http.go $(CMD_DIR)/background.go

goabuild: 
	goa gen checkpoint/design
	goa example checkpoint/design

goaclean:
	rm -rf ./gen

install:
	install $(BINARY_NAME) $(PREFIX)/bin/chckpApi

clean:
	rm -f $(BINARY_NAME)
