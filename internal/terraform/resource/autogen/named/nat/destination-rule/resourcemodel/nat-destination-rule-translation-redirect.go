// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &NatDestinationRuleTranSLAtionRedirect{}

// NatDestinationRuleTranSLAtionRedirect describes the resource data model.
type NatDestinationRuleTranSLAtionRedirect struct {
	// LeafNodes
	LeafNatDestinationRuleTranSLAtionRedirectPort types.String `tfsdk:"port" vyos:"port,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o NatDestinationRuleTranSLAtionRedirect) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"port": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Port number

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Numeric IP port  |
    |  range  &emsp; |  Numbered port range (e.g., 1001-1005)  |

`,
		},

		// Nodes

	}
}
