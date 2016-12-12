.PHONY: clean install test build all

clean:
	rm -f ./bin/secretstore

install:
	govendor install

test:
	govendor test +local

build:
	mkdir -p ./bin
	govendor build -o ./bin/secretstore .

all: build test