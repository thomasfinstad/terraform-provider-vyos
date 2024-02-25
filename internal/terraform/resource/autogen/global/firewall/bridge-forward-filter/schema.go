// Package globalfirewallbridgeforwardfilter code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalfirewallbridgeforwardfilter

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r firewallBrIDgeForwardFilter) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

<div style="text-align: center">
Firewall

<br>
&darr;
<br>
Bridge firewall

<br>
&darr;
<br>
Bridge forward firewall

<br>
&darr;
<br>
<b>
Bridge firewall forward filter
</b>
</div>
`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}