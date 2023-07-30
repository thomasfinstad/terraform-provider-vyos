// Package namedpolicyaccesslistsix code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedpolicyaccesslistsix

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r policyAccessListsix) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Routing policy

IPv6 access-list filter

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  Name of IPv6 access-list  |

`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
