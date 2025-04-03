package tools

import (
	"context"
	"fmt"

	"github.com/gridscale/gsclient-go/v3"
	"github.com/mark3labs/mcp-go/mcp"
)

func CreateStorage(gs *gsclient.Client) HandlerFactory {
	return func() (mcp.Tool, Handler) {
		opts := []mcp.ToolOption{
			mcp.WithDescription("Create a new storage which can be attached to a server."),
			mcp.WithString("name", mcp.Required()),
			mcp.WithNumber("capacity",
				// TODO: Restrict to integer values
				mcp.Required(),
				mcp.Min(10),
				mcp.Max(4000),
				mcp.Description("Capacity in GB"),
				mcp.DefaultNumber(10),
			),
			// mcp.WithObject("template",
			// 	mcp.Description("Optional storage template to apply to the provisioned storage"),
			// 	mcp.Properties(map[string]any{}),
			// ),
		}

		tool := mcp.NewTool("create_storage", opts...)

		handler := Handler(func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			gsRequest := gsclient.StorageCreateRequest{
				Name: request.Params.Arguments["name"].(string),
			}

			capacity, err := GetIntParam(request.Params.Arguments, "capacity")
			if err != nil {
				return nil, fmt.Errorf("capacity parameteris not valid: %w", err)
			}
			gsRequest.Capacity = capacity

			// TODO: Check if we can somehow report progress back to the client. The mcp doc says: Use progress reporting for long operations https://modelcontextprotocol.io/docs/concepts/tools
			gsResponse, err := gs.CreateStorage(ctx, gsRequest)
			if err != nil {
				return nil, fmt.Errorf("failed to get create storage: %w", err)
			}

			return mcp.NewToolResultText(fmt.Sprintf("Storage created with ID: %s", gsResponse.ObjectUUID)), nil
		})

		return tool, handler
	}
}

func GetIntParam(args map[string]any, key string) (int, error) {
	if value, ok := args[key]; ok {
		if intValue, ok := value.(int); ok {
			return intValue, nil
		} else if floatValue, ok := value.(float64); ok {
			// FIXME: This is a workaround for the fact that mcp does not support int parameters. This should be removed asap
			return int(floatValue), nil
		}
	}
	return 0, fmt.Errorf("parameter %s not found or not an integer", key)
}
