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

// InterfacesWirelessIPvsixAddress describes the resource data model.
type InterfacesWirelessIPvsixAddress struct {
	// LeafNodes
	LeafInterfacesWirelessIPvsixAddressAutoconf           types.String `tfsdk:"autoconf"`
	LeafInterfacesWirelessIPvsixAddressEuisixfour         types.String `tfsdk:"eui64"`
	LeafInterfacesWirelessIPvsixAddressNoDefaultLinkLocal types.String `tfsdk:"no_default_link_local"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *InterfacesWirelessIPvsixAddress) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"interfaces", "wireless", "ipv6", "address"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafInterfacesWirelessIPvsixAddressAutoconf.IsNull() || o.LeafInterfacesWirelessIPvsixAddressAutoconf.IsUnknown()) {
		vyosData["autoconf"] = o.LeafInterfacesWirelessIPvsixAddressAutoconf.ValueString()
	}
	if !(o.LeafInterfacesWirelessIPvsixAddressEuisixfour.IsNull() || o.LeafInterfacesWirelessIPvsixAddressEuisixfour.IsUnknown()) {
		vyosData["eui64"] = o.LeafInterfacesWirelessIPvsixAddressEuisixfour.ValueString()
	}
	if !(o.LeafInterfacesWirelessIPvsixAddressNoDefaultLinkLocal.IsNull() || o.LeafInterfacesWirelessIPvsixAddressNoDefaultLinkLocal.IsUnknown()) {
		vyosData["no-default-link-local"] = o.LeafInterfacesWirelessIPvsixAddressNoDefaultLinkLocal.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *InterfacesWirelessIPvsixAddress) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"interfaces", "wireless", "ipv6", "address"}})

	// Leafs
	if value, ok := vyosData["autoconf"]; ok {
		o.LeafInterfacesWirelessIPvsixAddressAutoconf = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessIPvsixAddressAutoconf = basetypes.NewStringNull()
	}
	if value, ok := vyosData["eui64"]; ok {
		o.LeafInterfacesWirelessIPvsixAddressEuisixfour = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessIPvsixAddressEuisixfour = basetypes.NewStringNull()
	}
	if value, ok := vyosData["no-default-link-local"]; ok {
		o.LeafInterfacesWirelessIPvsixAddressNoDefaultLinkLocal = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessIPvsixAddressNoDefaultLinkLocal = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"interfaces", "wireless", "ipv6", "address"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o InterfacesWirelessIPvsixAddress) AttributeTypes() map[string]attr.Type {
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
func (o InterfacesWirelessIPvsixAddress) ResourceSchemaAttributes() map[string]schema.Attribute {
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
