// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourLabeledUnicastAllowasIn{}

// VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourLabeledUnicastAllowasIn describes the resource data model.
type VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourLabeledUnicastAllowasIn struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourLabeledUnicastAllowasInNumber types.Number `tfsdk:"number" vyos:"number,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourLabeledUnicastAllowasIn) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"number": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Number of occurrences of AS number

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-10  &emsp; |  Number of times AS is allowed in path  |

`,
		},

		// Nodes

	}
}
