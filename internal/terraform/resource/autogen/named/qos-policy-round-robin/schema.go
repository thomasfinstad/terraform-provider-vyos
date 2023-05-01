// Package namedqospolicyroundrobin code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedqospolicyroundrobin

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r qosPolicyRoundRobin) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Quality of Service (QoS)

Service Policy definitions

Deficit Round Robin Scheduler

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Policy name  |

`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
