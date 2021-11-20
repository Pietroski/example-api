# Makefile

mock-generate:
	go get -d github.com/golang/mock/mockgen
	go generate ./...

build:
	./scripts/build/build-local.sh

clean-build:
	rm -rf build
