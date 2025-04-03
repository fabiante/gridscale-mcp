install_location := "/usr/bin/gridscale-mcp"

default: build install

build:
    go build -o gridscale-mcp ./main.go

install:
    sudo mv gridscale-mcp {{install_location}}

uninstall:
    sudo rm -f {{install_location}}

# See: https://github.com/modelcontextprotocol/inspector
inspect:
    npx @modelcontextprotocol/inspector node build/index.js
