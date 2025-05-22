package tools

import (
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

func newInvalidParamErrorResult(param string, err error) (*mcp.CallToolResult, error) {
	return mcp.NewToolResultErrorFromErr(fmt.Sprintf("invalid %q given", param), err), nil
}
