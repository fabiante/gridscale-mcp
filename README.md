# Gridscale MCP

This is a [MCP Server](https://modelcontextprotocol.io/docs/concepts/architecture) which allows AI modells
to interact with the [Gridscale API](https://my.gridscale.io/product-documentation/cloud-computing/apis/public-api/introduction/)

You can use this with tools such as [5ire](https://5ire.app/) to let your LLM provision infrastructre as you wish.

> [!CAUTION]
> If you use this program, you are responsible for the consequences. You should only use this in an empty Gridscale project
> where the AI model will not cause any harm like deleting your production database.

## Features (Tools and Resources)

With this MCP server you can use the following tools and resources:

| Type | Name                  | Note                             | Supported | Support Wanted |
|------|-----------------------|----------------------------------|-----------|----------------|
| Tool | create_ip             |                                  | Yes       |                |
| Tool | delete_ip             |                                  | Yes       |                |
| Tool | create_storage        | Not all parameters are supported | Partially |                |
| Tool | create_server         |                                  | No        | Yes            |
| Tool | create_network        |                                  | No        | Yes            |
| Res  | get_storage_templates |                                  | Yes       |                |

If you want to contribute, try adding one of the above missing tools or resources. 

## Building + Installation

To build this program you need a [Go](https://go.dev/) environment.

The recommended way is to use `make`:

```shell
make build
```

This will have created a binary called `gridscale-mcp` in the project directory.

## Install

It simplifies things if you install the binary to a location in your `$PATH`, for example `/usr/bin` on Linux.

You can use `make` for that:

```shell
# Install
make install

# Uninstall
make uninstall
```

## Usage

> [!CAUTION]
> If you use this program, you are responsible for the consequences. You should only use this in an empty Gridscale project
> where the AI model will not cause any harm like deleting your production database.

### Prerequisites

You need:

- To have a build of `gridscale-mcp`. Follow the steps in the [Building](#building) section to first build the program.
- Public API credentials for the [Gridscale API](https://my.gridscale.io/APIDoc/)

### Usage in an MCP client

To run this server you need two arguments:

```shell
gridscale-mcp --user-key <your-gs-user-key> --user-token <your-gs-user-token>
```

I recommend to use [5ire](https://github.com/nanbingxyz/5ire) as desktop client to interact with an LLM
and this MCP server. You are responsible for setting up an LLM (for example Gemini).

In 5ire the MCP server is setup as tool like so:

```json
{
  "name": "Gridscale",
  "key": "gridscale",
  "command": "gridscale-mcp",
  "args": [
    "--user-key",
    "...",
    "--user-token",
    "..."
  ]
}
```

## Debugging

You can directly communicate with you MCP server as with any other API server.

The [MCP Inspector](https://github.com/modelcontextprotocol/inspector) is a great tool for this.

If you have Node.js installed, simply run:

```shell
make inspect
```
