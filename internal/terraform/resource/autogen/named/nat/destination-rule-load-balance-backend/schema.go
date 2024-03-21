// Package namednatdestinationruleloadbalancebackend code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namednatdestinationruleloadbalancebackend

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r natDestinationRuleLoadBalanceBackend) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Network Address Translation (NAT) parameters  
⯯  
Destination NAT settings  
⯯  
Rule number for NAT  
⯯  
Apply NAT load balance  
⯯  
**Translated IP address**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
