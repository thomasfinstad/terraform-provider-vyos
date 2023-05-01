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

// FirewallNameRuleDestinationGeoIP describes the resource data model.
type FirewallNameRuleDestinationGeoIP struct {
	// LeafNodes
	LeafFirewallNameRuleDestinationGeoIPCountryCode  types.String `tfsdk:"country_code"`
	LeafFirewallNameRuleDestinationGeoIPInverseMatch types.String `tfsdk:"inverse_match"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *FirewallNameRuleDestinationGeoIP) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"firewall", "name", "rule", "destination", "geoip"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafFirewallNameRuleDestinationGeoIPCountryCode.IsNull() || o.LeafFirewallNameRuleDestinationGeoIPCountryCode.IsUnknown()) {
		vyosData["country-code"] = o.LeafFirewallNameRuleDestinationGeoIPCountryCode.ValueString()
	}
	if !(o.LeafFirewallNameRuleDestinationGeoIPInverseMatch.IsNull() || o.LeafFirewallNameRuleDestinationGeoIPInverseMatch.IsUnknown()) {
		vyosData["inverse-match"] = o.LeafFirewallNameRuleDestinationGeoIPInverseMatch.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *FirewallNameRuleDestinationGeoIP) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"firewall", "name", "rule", "destination", "geoip"}})

	// Leafs
	if value, ok := vyosData["country-code"]; ok {
		o.LeafFirewallNameRuleDestinationGeoIPCountryCode = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleDestinationGeoIPCountryCode = basetypes.NewStringNull()
	}
	if value, ok := vyosData["inverse-match"]; ok {
		o.LeafFirewallNameRuleDestinationGeoIPInverseMatch = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleDestinationGeoIPInverseMatch = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"firewall", "name", "rule", "destination", "geoip"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o FirewallNameRuleDestinationGeoIP) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"country_code":  types.StringType,
		"inverse_match": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallNameRuleDestinationGeoIP) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"country_code": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `GeoIP country code

|  Format  |  Description  |
|----------|---------------|
|  <country>  |  Country code (2 characters)  |

`,
		},

		"inverse_match": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Inverse match of country-codes

`,
		},

		// TagNodes

		// Nodes

	}
}
