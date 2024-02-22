// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsBgpNeighborLocalAs describes the resource data model.
type ProtocolsBgpNeighborLocalAs struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.Number `tfsdk:"local_as_id" vyos:",self-id"`

	ParentIDProtocolsBgpNeighbor types.String `tfsdk:"neighbor" vyos:"neighbor,parent-id"`

	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeProtocolsBgpNeighborLocalAsNoPrepend *ProtocolsBgpNeighborLocalAsNoPrepend `tfsdk:"no_prepend" vyos:"no-prepend,omitempty"`
}

// SetID configures the resource ID
func (o *ProtocolsBgpNeighborLocalAs) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBgpNeighborLocalAs) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"protocols",

		"bgp",

		"neighbor",
		o.ParentIDProtocolsBgpNeighbor.ValueString(),

		"local-as",
		o.SelfIdentifier.ValueBigFloat().String(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborLocalAs) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"local_as_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Specify alternate ASN for this BGP process

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-4294967294  &emsp; |  Autonomous System Number (ASN)  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		"neighbor_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `BGP neighbor

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  BGP neighbor IP address  |
    |  ipv6  &emsp; |  BGP neighbor IPv6 address  |
    |  txt  &emsp; |  Interface name  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		// Nodes

		"no_prepend": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborLocalAsNoPrepend{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Disable prepending local-as from/to updates for eBGP peers

`,
		},
	}
}
