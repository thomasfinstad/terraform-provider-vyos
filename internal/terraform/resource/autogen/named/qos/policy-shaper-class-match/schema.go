// Package namedqospolicyshaperclassmatch code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedqospolicyshaperclassmatch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r qosPolicyShaperClassMatch) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Quality of Service (QoS)  
⯯  
Service Policy definitions  
⯯  
Traffic shaping based policy (Hierarchy Token Bucket)  
⯯  
Class ID  
⯯  
**Class matching rule name**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
