// Package namedfirewallipvsixoutputrawrule code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallipvsixoutputrawrule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r firewallIPvsixOutputRawRule) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Firewall  
⯯  
IPv6 firewall  
⯯  
IPv6 output firewall  
⯯  
IPv6 firewall output raw  
⯯  
**IPv6 Firewall output raw rule number**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
