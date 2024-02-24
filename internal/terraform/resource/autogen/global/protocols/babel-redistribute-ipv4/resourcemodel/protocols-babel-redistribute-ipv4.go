// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsBabelRedistributeIPvfour describes the resource data model.
type ProtocolsBabelRedistributeIPvfour struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafProtocolsBabelRedistributeIPvfourBgp       types.Bool `tfsdk:"bgp" vyos:"bgp,omitempty"`
	LeafProtocolsBabelRedistributeIPvfourConnected types.Bool `tfsdk:"connected" vyos:"connected,omitempty"`
	LeafProtocolsBabelRedistributeIPvfourEigrp     types.Bool `tfsdk:"eigrp" vyos:"eigrp,omitempty"`
	LeafProtocolsBabelRedistributeIPvfourIsis      types.Bool `tfsdk:"isis" vyos:"isis,omitempty"`
	LeafProtocolsBabelRedistributeIPvfourKernel    types.Bool `tfsdk:"kernel" vyos:"kernel,omitempty"`
	LeafProtocolsBabelRedistributeIPvfourNhrp      types.Bool `tfsdk:"nhrp" vyos:"nhrp,omitempty"`
	LeafProtocolsBabelRedistributeIPvfourOspf      types.Bool `tfsdk:"ospf" vyos:"ospf,omitempty"`
	LeafProtocolsBabelRedistributeIPvfourRIP       types.Bool `tfsdk:"rip" vyos:"rip,omitempty"`
	LeafProtocolsBabelRedistributeIPvfourStatic    types.Bool `tfsdk:"static" vyos:"static,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *ProtocolsBabelRedistributeIPvfour) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBabelRedistributeIPvfour) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"protocols",

		"babel",

		"redistribute",

		"ipv4",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBabelRedistributeIPvfour) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"bgp": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Redistribute BGP routes

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"connected": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Redistribute connected routes

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"eigrp": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Redistribute EIGRP routes

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"isis": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Redistribute IS-IS routes

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"kernel": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Redistribute kernel routes

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"nhrp": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Redistribute NHRP routes

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"ospf": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Redistribute OSPF routes

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"rip": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Redistribute RIP routes

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"static": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Redistribute static routes

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},
	}
}
