// Package namedpolicyaccesslist code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedpolicyaccesslist

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r policyAccessList) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Routing policy

IP access-list filter

|  Format  |  Description  |
|----------|---------------|
|  u32:1-99  |  IP standard access list  |
|  u32:100-199  |  IP extended access list  |
|  u32:1300-1999  |  IP standard access list (expanded range)  |
|  u32:2000-2699  |  IP extended access list (expanded range)  |

`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
