INSTALL_LOCATION := /usr/bin/gridscale-mcp

.PHONY: default build install uninstall inspect

default: build

build:
	go build -o gridscale-mcp ./main.go

install:
	sudo mv gridscale-mcp $(INSTALL_LOCATION)

uninstall:
	sudo rm -f $(INSTALL_LOCATION)

# See: https://github.com/modelcontextprotocol/inspector
inspect:
	npx @modelcontextprotocol/inspector node build/index.js
