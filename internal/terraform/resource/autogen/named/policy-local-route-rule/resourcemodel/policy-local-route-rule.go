// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// PolicyLocalRouteRule describes the resource data model.
type PolicyLocalRouteRule struct {
	// LeafNodes
	PolicyLocalRouteRuleFwmark           customtypes.CustomStringValue `tfsdk:"fwmark" json:"fwmark,omitempty"`
	PolicyLocalRouteRuleSource           customtypes.CustomStringValue `tfsdk:"source" json:"source,omitempty"`
	PolicyLocalRouteRuleDestination      customtypes.CustomStringValue `tfsdk:"destination" json:"destination,omitempty"`
	PolicyLocalRouteRuleInboundInterface customtypes.CustomStringValue `tfsdk:"inbound_interface" json:"inbound-interface,omitempty"`

	// TagNodes

	// Nodes
	PolicyLocalRouteRuleSet types.Object `tfsdk:"set" json:"set,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o PolicyLocalRouteRule) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"fwmark": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Match fwmark value

|  Format  |  Description  |
|----------|---------------|
|  u32:1-2147483647  |  Address to match against  |
`,
		},

		"source": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Source address or prefix

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Address to match against  |
|  ipv4net  |  Prefix to match against  |
`,
		},

		"destination": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Destination address or prefix

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Address to match against  |
|  ipv4net  |  Prefix to match against  |
`,
		},

		"inbound_interface": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Inbound Interface

`,
		},

		// TagNodes

		// Nodes

		"set": schema.SingleNestedAttribute{
			Attributes: PolicyLocalRouteRuleSet{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Packet modifications

`,
		},
	}
}
