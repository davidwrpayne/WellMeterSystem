.PHONY: all

all: compile

compile:
	go build -o well-meter-system cmd/main.go