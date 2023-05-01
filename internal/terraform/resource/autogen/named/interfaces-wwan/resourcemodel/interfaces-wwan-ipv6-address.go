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

// InterfacesWwanIPvsixAddress describes the resource data model.
type InterfacesWwanIPvsixAddress struct {
	// LeafNodes
	LeafInterfacesWwanIPvsixAddressAutoconf           types.String `tfsdk:"autoconf"`
	LeafInterfacesWwanIPvsixAddressEuisixfour         types.String `tfsdk:"eui64"`
	LeafInterfacesWwanIPvsixAddressNoDefaultLinkLocal types.String `tfsdk:"no_default_link_local"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *InterfacesWwanIPvsixAddress) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"interfaces", "wwan", "ipv6", "address"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafInterfacesWwanIPvsixAddressAutoconf.IsNull() || o.LeafInterfacesWwanIPvsixAddressAutoconf.IsUnknown()) {
		vyosData["autoconf"] = o.LeafInterfacesWwanIPvsixAddressAutoconf.ValueString()
	}
	if !(o.LeafInterfacesWwanIPvsixAddressEuisixfour.IsNull() || o.LeafInterfacesWwanIPvsixAddressEuisixfour.IsUnknown()) {
		vyosData["eui64"] = o.LeafInterfacesWwanIPvsixAddressEuisixfour.ValueString()
	}
	if !(o.LeafInterfacesWwanIPvsixAddressNoDefaultLinkLocal.IsNull() || o.LeafInterfacesWwanIPvsixAddressNoDefaultLinkLocal.IsUnknown()) {
		vyosData["no-default-link-local"] = o.LeafInterfacesWwanIPvsixAddressNoDefaultLinkLocal.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *InterfacesWwanIPvsixAddress) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"interfaces", "wwan", "ipv6", "address"}})

	// Leafs
	if value, ok := vyosData["autoconf"]; ok {
		o.LeafInterfacesWwanIPvsixAddressAutoconf = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWwanIPvsixAddressAutoconf = basetypes.NewStringNull()
	}
	if value, ok := vyosData["eui64"]; ok {
		o.LeafInterfacesWwanIPvsixAddressEuisixfour = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWwanIPvsixAddressEuisixfour = basetypes.NewStringNull()
	}
	if value, ok := vyosData["no-default-link-local"]; ok {
		o.LeafInterfacesWwanIPvsixAddressNoDefaultLinkLocal = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWwanIPvsixAddressNoDefaultLinkLocal = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"interfaces", "wwan", "ipv6", "address"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o InterfacesWwanIPvsixAddress) AttributeTypes() map[string]attr.Type {
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
func (o InterfacesWwanIPvsixAddress) ResourceSchemaAttributes() map[string]schema.Attribute {
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