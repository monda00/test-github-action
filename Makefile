SHELL := /bin/bash

PLIST=info.plist
EXEC_BIN=sample-alfred-workflow
DIST_FILE=sample.alfredworkflow
GO_SRCS=$(shell find -f . \( -name \*.go \))

all: $(DIST_FILE)

$(EXEC_BIN): $(GO_SRCS)
	go build -o $(EXEC_BIN)

$(DIST_FILE): $(EXEC_BIN) $(CREDITS) $(PLIST)
	zip -r $(DIST_FILE) $(PLIST) $(EXEC_BIN)
