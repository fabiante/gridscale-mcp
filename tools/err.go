package tools

import (
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

func newUnparsableIntErrorResult(param string, err error) (*mcp.CallToolResult, error) {
	return mcp.NewToolResultErrorFromErr(fmt.Sprintf("invalid %q given, it must be parsable as integer", param), err), nil
}
