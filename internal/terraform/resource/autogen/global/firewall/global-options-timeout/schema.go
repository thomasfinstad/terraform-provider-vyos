// Package globalfirewallglobaloptionstimeout code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalfirewallglobaloptionstimeout

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r firewallGlobalOptionsTimeout) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

Firewall  
⯯  
Global Options  
⯯  
**Connection timeout options**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
