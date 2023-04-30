// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ProtocolsOspfAreaVirtualLink describes the resource data model.
type ProtocolsOspfAreaVirtualLink struct {
	// LeafNodes
	ProtocolsOspfAreaVirtualLinkDeadInterval       customtypes.CustomStringValue `tfsdk:"dead_interval" json:"dead-interval,omitempty"`
	ProtocolsOspfAreaVirtualLinkHelloInterval      customtypes.CustomStringValue `tfsdk:"hello_interval" json:"hello-interval,omitempty"`
	ProtocolsOspfAreaVirtualLinkRetransmitInterval customtypes.CustomStringValue `tfsdk:"retransmit_interval" json:"retransmit-interval,omitempty"`
	ProtocolsOspfAreaVirtualLinkTransmitDelay      customtypes.CustomStringValue `tfsdk:"transmit_delay" json:"transmit-delay,omitempty"`

	// TagNodes

	// Nodes
	ProtocolsOspfAreaVirtualLinkAuthentication types.Object `tfsdk:"authentication" json:"authentication,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ProtocolsOspfAreaVirtualLink) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"dead_interval": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Interval after which a neighbor is declared dead

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Neighbor dead interval (seconds)  |
`,

			// Default:          stringdefault.StaticString(`40`),
			Computed: true,
		},

		"hello_interval": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Interval between hello packets

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Hello interval (seconds)  |
`,

			// Default:          stringdefault.StaticString(`10`),
			Computed: true,
		},

		"retransmit_interval": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Interval between retransmitting lost link state advertisements

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Retransmit interval (seconds)  |
`,

			// Default:          stringdefault.StaticString(`5`),
			Computed: true,
		},

		"transmit_delay": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Link state transmit delay

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Link state transmit delay (seconds)  |
`,

			// Default:          stringdefault.StaticString(`1`),
			Computed: true,
		},

		// TagNodes

		// Nodes

		"authentication": schema.SingleNestedAttribute{
			Attributes: ProtocolsOspfAreaVirtualLinkAuthentication{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Authentication

`,
		},
	}
}
