// Package namedfirewallipvsixforwardfilterrule code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallipvsixforwardfilterrule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r firewallIPvsixForwardFilterRule) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Firewall  
⯯  
IPv6 firewall  
⯯  
IPv6 forward firewall  
⯯  
IPv6 firewall forward filter  
⯯  
**IPv6 Firewall forward filter rule number**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
