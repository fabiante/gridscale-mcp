package resources

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
)

// TODO: Think about contributing this type to the lib
type Handler = func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error)

type TemplateHandlerFactory = func() (mcp.ResourceTemplate, Handler)

type HandlerFactory = func() (mcp.Resource, Handler)
