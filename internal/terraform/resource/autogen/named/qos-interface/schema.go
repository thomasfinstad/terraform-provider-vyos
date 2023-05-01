// Package namedqosinterface code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedqosinterface

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r qosInterface) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Quality of Service (QoS)

Interface to apply QoS policy

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Interface name  |

`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
