// Package namedpolicylocalrouterule code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedpolicylocalrouterule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r policyLocalRouteRule) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `IPv4 policy route of local traffic

Policy local-route rule set number

|  Format  |  Description  |
|----------|---------------|
|  u32:1-32765  |  Local-route rule number (1-32765)  |

`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
