package tools

import (
	"context"
	"fmt"

	"github.com/gridscale/gsclient-go/v3"
	"github.com/mark3labs/mcp-go/mcp"
)

func CreateIP(gs *gsclient.Client) HandlerFactory {
	return func() (mcp.Tool, Handler) {
		opts := []mcp.ToolOption{
			mcp.WithDescription("Create a new public IP address (v4 or v6) which can be attached to a server on its public network interface."),
			mcp.WithString("name"),
			mcp.WithNumber("family", mcp.Required(), mcp.Enum("4", "6")), // TODO: Restrict to integer values
		}

		tool := mcp.NewTool("create_ip", opts...)

		handler := Handler(func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			gsRequest := gsclient.IPCreateRequest{}

			if ip, err := request.RequireInt("family"); err == nil {
				gsRequest.Family = gsclient.IPAddressType(ip)
			} else {
				return newInvalidParamErrorResult("family", err)
			}

			if name, err := request.RequireString("name"); err == nil {
				gsRequest.Name = name
			} else {
				return newInvalidParamErrorResult("name", err)
			}

			gsResponse, err := gs.CreateIP(ctx, gsRequest)
			if err != nil {
				return mcp.NewToolResultErrorFromErr("failed to create ip", err), nil
			}

			return mcp.NewToolResultText(fmt.Sprintf("IP %s created with ID: %s", gsResponse.IP, gsResponse.ObjectUUID)), nil
		})

		return tool, handler
	}
}
