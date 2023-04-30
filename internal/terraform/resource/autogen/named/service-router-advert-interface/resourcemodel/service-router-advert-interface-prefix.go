// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ServiceRouterAdvertInterfacePrefix describes the resource data model.
type ServiceRouterAdvertInterfacePrefix struct {
	// LeafNodes
	ServiceRouterAdvertInterfacePrefixNoAutonomousFlag  customtypes.CustomStringValue `tfsdk:"no_autonomous_flag" json:"no-autonomous-flag,omitempty"`
	ServiceRouterAdvertInterfacePrefixNoOnLinkFlag      customtypes.CustomStringValue `tfsdk:"no_on_link_flag" json:"no-on-link-flag,omitempty"`
	ServiceRouterAdvertInterfacePrefixDeprecatePrefix   customtypes.CustomStringValue `tfsdk:"deprecate_prefix" json:"deprecate-prefix,omitempty"`
	ServiceRouterAdvertInterfacePrefixDecrementLifetime customtypes.CustomStringValue `tfsdk:"decrement_lifetime" json:"decrement-lifetime,omitempty"`
	ServiceRouterAdvertInterfacePrefixPreferredLifetime customtypes.CustomStringValue `tfsdk:"preferred_lifetime" json:"preferred-lifetime,omitempty"`
	ServiceRouterAdvertInterfacePrefixValIDLifetime     customtypes.CustomStringValue `tfsdk:"valid_lifetime" json:"valid-lifetime,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ServiceRouterAdvertInterfacePrefix) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"no_autonomous_flag": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Prefix can not be used for stateless address auto-configuration

`,
		},

		"no_on_link_flag": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Prefix can not be used for on-link determination

`,
		},

		"deprecate_prefix": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Upon shutdown, this option will deprecate the prefix by announcing it in the shutdown RA

`,
		},

		"decrement_lifetime": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Lifetime is decremented by the number of seconds since the last RA - use in conjunction with a DHCPv6-PD prefix

`,
		},

		"preferred_lifetime": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Time in seconds that the prefix will remain preferred

|  Format  |  Description  |
|----------|---------------|
|  u32  |  Time in seconds that the prefix will remain preferred  |
|  infinity  |  Prefix will remain preferred forever  |
`,

			// Default:          stringdefault.StaticString(`14400`),
			Computed: true,
		},

		"valid_lifetime": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Time in seconds that the prefix will remain valid

|  Format  |  Description  |
|----------|---------------|
|  u32:1-4294967295  |  Time in seconds that the prefix will remain valid  |
|  infinity  |  Prefix will remain preferred forever  |
`,

			// Default:          stringdefault.StaticString(`2592000`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
