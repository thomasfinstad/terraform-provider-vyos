// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// InterfacesBondingVifSIPvsixAddress describes the resource data model.
type InterfacesBondingVifSIPvsixAddress struct {
	// LeafNodes
	LeafInterfacesBondingVifSIPvsixAddressAutoconf           types.String `tfsdk:"autoconf"`
	LeafInterfacesBondingVifSIPvsixAddressEuisixfour         types.String `tfsdk:"eui64"`
	LeafInterfacesBondingVifSIPvsixAddressNoDefaultLinkLocal types.String `tfsdk:"no_default_link_local"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *InterfacesBondingVifSIPvsixAddress) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"interfaces", "bonding", "vif-s", "ipv6", "address"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafInterfacesBondingVifSIPvsixAddressAutoconf.IsNull() || o.LeafInterfacesBondingVifSIPvsixAddressAutoconf.IsUnknown()) {
		vyosData["autoconf"] = o.LeafInterfacesBondingVifSIPvsixAddressAutoconf.ValueString()
	}
	if !(o.LeafInterfacesBondingVifSIPvsixAddressEuisixfour.IsNull() || o.LeafInterfacesBondingVifSIPvsixAddressEuisixfour.IsUnknown()) {
		vyosData["eui64"] = o.LeafInterfacesBondingVifSIPvsixAddressEuisixfour.ValueString()
	}
	if !(o.LeafInterfacesBondingVifSIPvsixAddressNoDefaultLinkLocal.IsNull() || o.LeafInterfacesBondingVifSIPvsixAddressNoDefaultLinkLocal.IsUnknown()) {
		vyosData["no-default-link-local"] = o.LeafInterfacesBondingVifSIPvsixAddressNoDefaultLinkLocal.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *InterfacesBondingVifSIPvsixAddress) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"interfaces", "bonding", "vif-s", "ipv6", "address"}})

	// Leafs
	if value, ok := vyosData["autoconf"]; ok {
		o.LeafInterfacesBondingVifSIPvsixAddressAutoconf = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingVifSIPvsixAddressAutoconf = basetypes.NewStringNull()
	}
	if value, ok := vyosData["eui64"]; ok {
		o.LeafInterfacesBondingVifSIPvsixAddressEuisixfour = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingVifSIPvsixAddressEuisixfour = basetypes.NewStringNull()
	}
	if value, ok := vyosData["no-default-link-local"]; ok {
		o.LeafInterfacesBondingVifSIPvsixAddressNoDefaultLinkLocal = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingVifSIPvsixAddressNoDefaultLinkLocal = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"interfaces", "bonding", "vif-s", "ipv6", "address"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o InterfacesBondingVifSIPvsixAddress) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"autoconf":              types.StringType,
		"eui64":                 types.StringType,
		"no_default_link_local": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesBondingVifSIPvsixAddress) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"autoconf": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enable acquisition of IPv6 address using stateless autoconfig (SLAAC)

`,
		},

		"eui64": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Prefix for IPv6 address with MAC-based EUI-64

|  Format  |  Description  |
|----------|---------------|
|  <h:h:h:h:h:h:h:h/64>  |  IPv6 /64 network  |

`,
		},

		"no_default_link_local": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Remove the default link-local address from the interface

`,
		},

		// TagNodes

		// Nodes

	}
}