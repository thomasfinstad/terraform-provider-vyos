// Package namedpolicyrouterule code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedpolicyrouterule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r policyRouteRule) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Policy route rule set name for IPv4

Policy rule number

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-999999  |  Number of policy rule  |

`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
