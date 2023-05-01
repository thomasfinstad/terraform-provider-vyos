// Package namedservicesnmpvthreeview code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicesnmpvthreeview

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r serviceSnmpVthreeView) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Simple Network Management Protocol (SNMP)

Simple Network Management Protocol (SNMP) v3

Specifies the view with name viewname

`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
