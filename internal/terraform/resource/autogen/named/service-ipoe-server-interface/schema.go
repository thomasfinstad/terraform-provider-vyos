// Package namedserviceipoeserverinterface code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedserviceipoeserverinterface

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r serviceIPoeServerInterface) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Internet Protocol over Ethernet (IPoE) Server

Interface to listen dhcp or unclassified packets

`,
		Attributes: r.model.ResourceAttributes(),
	}
}
