// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsBgpAddressFamilyIPvfourUnicastNetwork describes the resource data model.
type ProtocolsBgpAddressFamilyIPvfourUnicastNetwork struct {
	SelfIdentifier types.String `tfsdk:"network_id" vyos:",self-id"`

	// LeafNodes
	LeafProtocolsBgpAddressFamilyIPvfourUnicastNetworkBackdoor types.Bool   `tfsdk:"backdoor" vyos:"backdoor,omitempty"`
	LeafProtocolsBgpAddressFamilyIPvfourUnicastNetworkRouteMap types.String `tfsdk:"route_map" vyos:"route-map,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBgpAddressFamilyIPvfourUnicastNetwork) GetVyosPath() []string {
	return []string{
		"protocols",

		"bgp",

		"address-family",

		"ipv4-unicast",

		"network",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpAddressFamilyIPvfourUnicastNetwork) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, an amalgamation of the `network_id` and the parents `*_id` fields seperated by dunder `__` starting with top level ancestor.",
		},
		"network_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `BGP network

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  BGP network  |

`,
		},

		// LeafNodes

		"backdoor": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Network as a backdoor route

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"route_map": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specify route-map name to use

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Route map name  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBgpAddressFamilyIPvfourUnicastNetwork) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsBgpAddressFamilyIPvfourUnicastNetwork) UnmarshalJSON(_ []byte) error {
	return nil
}
