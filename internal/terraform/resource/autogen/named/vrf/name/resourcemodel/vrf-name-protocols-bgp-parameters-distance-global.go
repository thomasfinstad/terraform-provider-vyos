// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpParametersDistanceGlobal{}

// VrfNameProtocolsBgpParametersDistanceGlobal describes the resource data model.
type VrfNameProtocolsBgpParametersDistanceGlobal struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpParametersDistanceGlobalExternal types.Number `tfsdk:"external" vyos:"external,omitempty"`
	LeafVrfNameProtocolsBgpParametersDistanceGlobalInternal types.Number `tfsdk:"internal" vyos:"internal,omitempty"`
	LeafVrfNameProtocolsBgpParametersDistanceGlobalLocal    types.Number `tfsdk:"local" vyos:"local,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpParametersDistanceGlobal) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"external": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Administrative distance for external BGP routes

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-255  &emsp; |  Administrative distance for external BGP routes  |

`,
		},

		"internal": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Administrative distance for internal BGP routes

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-255  &emsp; |  Administrative distance for internal BGP routes  |

`,
		},

		"local": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Administrative distance for local BGP routes

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-255  &emsp; |  Administrative distance for internal BGP routes  |

`,
		},

		// Nodes

	}
}
