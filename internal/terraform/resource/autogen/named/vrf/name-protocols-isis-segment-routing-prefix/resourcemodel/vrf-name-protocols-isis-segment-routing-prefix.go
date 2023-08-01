// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsIsisSegmentRoutingPrefix describes the resource data model.
type VrfNameProtocolsIsisSegmentRoutingPrefix struct {
	SelfIdentifier types.String `tfsdk:"prefix_id" vyos:",self-id"`

	ParentIDVrfName types.String `tfsdk:"name" vyos:"name,parent-id"`

	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsIsisSegmentRoutingPrefixAbsolute *VrfNameProtocolsIsisSegmentRoutingPrefixAbsolute `tfsdk:"absolute" vyos:"absolute,omitempty"`
	NodeVrfNameProtocolsIsisSegmentRoutingPrefixIndex    *VrfNameProtocolsIsisSegmentRoutingPrefixIndex    `tfsdk:"index" vyos:"index,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VrfNameProtocolsIsisSegmentRoutingPrefix) GetVyosPath() []string {
	return []string{
		"vrf",

		"name",
		o.ParentIDVrfName.ValueString(),

		"protocols",

		"isis",

		"segment-routing",

		"prefix",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisSegmentRoutingPrefix) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, an amalgamation of the `prefix_id` and the parents `*_id` fields seperated by dunder `__` starting with top level ancestor.",
		},
		"prefix_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Static IPv4/IPv6 prefix segment/label mapping

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  IPv4 prefix segment  |
    |  ipv6net  &emsp; |  IPv6 prefix segment  |

`,
		},

		"name_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Virtual Routing and Forwarding instance

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  VRF instance name  |

`,
		},

		// LeafNodes

		// Nodes

		"absolute": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisSegmentRoutingPrefixAbsolute{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Specify the absolute value of prefix segment/label ID

`,
		},

		"index": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisSegmentRoutingPrefixIndex{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Specify the index value of prefix segment/label ID

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsIsisSegmentRoutingPrefix) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsIsisSegmentRoutingPrefix) UnmarshalJSON(_ []byte) error {
	return nil
}
