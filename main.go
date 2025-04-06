package main

import (
	"flag"
	"fmt"

	"github.com/fabiante/gridscale-mcp/resources"
	"github.com/fabiante/gridscale-mcp/tools"
	"github.com/gridscale/gsclient-go/v3"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	userKey := flag.String("user-key", "", "Gridscale API user key")
	userToken := flag.String("user-token", "", "Gridscale API user token")
	flag.Parse()

	if *userKey == "" {
		fmt.Println("Error: --user-key flag is required")
		flag.Usage()
		return
	}

	if *userToken == "" {
		fmt.Println("Error: --user-token flag is required")
		flag.Usage()
		return
	}

	// Create a new MCP server
	s := server.NewMCPServer(
		"Gridscale API",
		"0.1.0",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
	)

	gs := newGSClient(*userKey, *userToken)

	addTool(s, tools.CreateIP(gs))
	addTool(s, tools.CreateStorage(gs))
	addResource(s, resources.GetStorageTemplate(gs)) // FIXME: Somehow the LLM does not respond with a list of templates if asked. Check why. Is that an issue of this code or the LLM?

	// Start the server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

func addTool(s *server.MCPServer, factory tools.HandlerFactory) {
	tool, handler := factory()
	s.AddTool(tool, handler)
}

func addResourceTemplate(s *server.MCPServer, factory resources.TemplateHandlerFactory) {
	tool, handler := factory()
	s.AddResourceTemplate(tool, handler)
}

func addResource(s *server.MCPServer, factory resources.HandlerFactory) {
	tool, handler := factory()
	s.AddResource(tool, handler)
}

func newGSClient(userKey, userToken string) *gsclient.Client {
	config := gsclient.DefaultConfiguration(userKey, userToken)
	return gsclient.NewClient(config)
}
