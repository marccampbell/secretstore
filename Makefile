.PHONY: clean install test build all

clean:
	rm -f ./bin/kube-vault

install:
	govendor install

test:
	govendor test +local

build:
	mkdir -p ./bin
	govendor build -o ./bin/kube-vault .

all: build test