LDFLAGS=-s -w

all: clean build data

build:
	$(info ************ BUILDING EXECUTABLE FILE ************)
	go build -ldflags "$(LDFLAGS)" -o ./build/vidpovid-bot-go

data:
	$(info ************ BUILDING DATA FILES ************)
	# Config example
	cp ./.env.example ./build/.env.example

clean:
	$(info ************ CLEANING ************)
	rm -rf ./build

run:
	$(info ************ RUNNING ************)
	go run -ldflags "$(LDFLAGS)" main.go

.PHONY: all run
