package tools

import (
	"context"
	"fmt"

	"github.com/fabiante/gridscale-mcp/util"
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

			capacity, err := util.GetIntParam(request.Params.Arguments, "capacity")
			if err != nil {
				return newUnparsableIntErrorResult("capacity", err)
			}
			gsRequest.Capacity = capacity

			// TODO: Check if we can somehow report progress back to the client. The mcp doc says: Use progress reporting for long operations https://modelcontextprotocol.io/docs/concepts/tools
			gsResponse, err := gs.CreateStorage(ctx, gsRequest)
			if err != nil {
				return mcp.NewToolResultErrorFromErr("failed to create storage", err), nil
			}

			return mcp.NewToolResultText(fmt.Sprintf("Storage created with ID: %s", gsResponse.ObjectUUID)), nil
		})

		return tool, handler
	}
}
