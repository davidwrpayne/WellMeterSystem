.PHONY: all

all: compile

compile:
	go build -o wms cmd/main.go