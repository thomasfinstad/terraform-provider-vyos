// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &QosPolicyCakeFlowIsolation{}

// QosPolicyCakeFlowIsolation describes the resource data model.
type QosPolicyCakeFlowIsolation struct {
	// LeafNodes
	LeafQosPolicyCakeFlowIsolationBlind         types.Bool `tfsdk:"blind" vyos:"blind,omitempty"`
	LeafQosPolicyCakeFlowIsolationSrcHost       types.Bool `tfsdk:"src_host" vyos:"src-host,omitempty"`
	LeafQosPolicyCakeFlowIsolationDstHost       types.Bool `tfsdk:"dst_host" vyos:"dst-host,omitempty"`
	LeafQosPolicyCakeFlowIsolationHost          types.Bool `tfsdk:"host" vyos:"host,omitempty"`
	LeafQosPolicyCakeFlowIsolationFlow          types.Bool `tfsdk:"flow" vyos:"flow,omitempty"`
	LeafQosPolicyCakeFlowIsolationDualSrcHost   types.Bool `tfsdk:"dual_src_host" vyos:"dual-src-host,omitempty"`
	LeafQosPolicyCakeFlowIsolationDualDstHost   types.Bool `tfsdk:"dual_dst_host" vyos:"dual-dst-host,omitempty"`
	LeafQosPolicyCakeFlowIsolationTrIPleIsolate types.Bool `tfsdk:"triple_isolate" vyos:"triple-isolate,omitempty"`
	LeafQosPolicyCakeFlowIsolationNat           types.Bool `tfsdk:"nat" vyos:"nat,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyCakeFlowIsolation) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"blind": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disables flow isolation, all traffic passes through a single queue

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"src_host": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Flows are defined only by source address

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"dst_host": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Flows are defined only by destination address

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"host": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Flows are defined by source-destination host pairs

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"flow": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Flows are defined by the entire 5-tuple

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"dual_src_host": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Flows are defined by the 5-tuple, fairness is applied first over source addresses, then over individual flows

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"dual_dst_host": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Flows are defined by the 5-tuple, fairness is applied first over destination addresses, then over individual flows

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"triple_isolate": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Flows are defined by the 5-tuple, fairness is applied over source and destination addresses and also over individual flows (default)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"nat": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Perform NAT lookup before applying flow-isolation rules

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
