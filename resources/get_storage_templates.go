package resources

import (
	"context"
	"encoding/json"

	"github.com/gridscale/gsclient-go/v3"
	"github.com/mark3labs/mcp-go/mcp"
)

func GetStorageTemplate(gs *gsclient.Client) HandlerFactory {
	return func() (mcp.Resource, Handler) {
		opts := []mcp.ResourceOption{
			mcp.WithResourceDescription("Get storage templates. These can be used when creating storages to initialize it with an operating system. When using a template, you will likely have to supply a root password or SSH public key for authentication after the VM has booted."),
			mcp.WithMIMEType("application/json"),
		}

		res := mcp.NewResource("gs://storage_templates", "Storage templates", opts...)

		handler := Handler(func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			templates, err := gs.GetTemplateList(ctx)
			if err != nil {
				return []mcp.ResourceContents{}, mcp.NewToolResultErrorFromErr("failed to get storage templates", err)
			}
			if len(templates) == 0 {
				return []mcp.ResourceContents{}, mcp.NewToolResultError("no storage templates found", nil)
			}

			templatesJSON, err := json.Marshal(templates)
			if err != nil {
				return []mcp.ResourceContents{}, mcp.NewToolResultErrorFromErr("failed to marshal templates to JSON", err)
			}

			return []mcp.ResourceContents{
				mcp.TextResourceContents{
					URI:      request.Params.URI,
					MIMEType: "application/json",
					Text:     string(templatesJSON),
				},
			}, nil
		})

		return res, handler
	}
}
