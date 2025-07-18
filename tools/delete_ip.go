package tools

import (
	"context"

	"github.com/gridscale/gsclient-go/v3"
	"github.com/mark3labs/mcp-go/mcp"
)

func DeleteIP(gs *gsclient.Client) HandlerFactory {
	return func() (mcp.Tool, Handler) {
		opts := []mcp.ToolOption{
			mcp.WithDescription("Delete a public IP address. This is only possible if it is currently not attached to a server."),
			mcp.WithString("uuid", mcp.Description("UUID of the IP address to delete")),
		}

		tool := mcp.NewTool("delete_ip", opts...)

		handler := Handler(func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var uuid string
			if u, err := request.RequireString("uuid"); err == nil {
				uuid = u
			} else {
				return newInvalidParamErrorResult("uuid", err)
			}

			err := gs.DeleteIP(ctx, uuid)
			if err != nil {
				return mcp.NewToolResultErrorFromErr("failed to delete IP", err), nil
			}

			return mcp.NewToolResultText("IP has been deleted"), nil
		})

		return tool, handler
	}
}
