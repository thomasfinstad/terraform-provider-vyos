// Package namedfirewallipvfourforwardfilterrule code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallipvfourforwardfilterrule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r firewallIPvfourForwardFilterRule) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Firewall  
⯯  
IPv4 firewall  
⯯  
IPv4 forward firewall  
⯯  
IPv4 firewall forward filter  
⯯  
**IPv4 Firewall forward filter rule number**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
