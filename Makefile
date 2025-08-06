.PHONY: build clean

build:
	@echo "building"
	CGO_ENABLED=0 go build .
	@echo "build done"