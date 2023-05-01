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

// NatDestinationRuleTranSLAtionOptions describes the resource data model.
type NatDestinationRuleTranSLAtionOptions struct {
	// LeafNodes
	LeafNatDestinationRuleTranSLAtionOptionsAddressMapping types.String `tfsdk:"address_mapping"`
	LeafNatDestinationRuleTranSLAtionOptionsPortMapping    types.String `tfsdk:"port_mapping"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *NatDestinationRuleTranSLAtionOptions) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"nat", "destination", "rule", "translation", "options"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafNatDestinationRuleTranSLAtionOptionsAddressMapping.IsNull() || o.LeafNatDestinationRuleTranSLAtionOptionsAddressMapping.IsUnknown()) {
		vyosData["address-mapping"] = o.LeafNatDestinationRuleTranSLAtionOptionsAddressMapping.ValueString()
	}
	if !(o.LeafNatDestinationRuleTranSLAtionOptionsPortMapping.IsNull() || o.LeafNatDestinationRuleTranSLAtionOptionsPortMapping.IsUnknown()) {
		vyosData["port-mapping"] = o.LeafNatDestinationRuleTranSLAtionOptionsPortMapping.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *NatDestinationRuleTranSLAtionOptions) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"nat", "destination", "rule", "translation", "options"}})

	// Leafs
	if value, ok := vyosData["address-mapping"]; ok {
		o.LeafNatDestinationRuleTranSLAtionOptionsAddressMapping = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafNatDestinationRuleTranSLAtionOptionsAddressMapping = basetypes.NewStringNull()
	}
	if value, ok := vyosData["port-mapping"]; ok {
		o.LeafNatDestinationRuleTranSLAtionOptionsPortMapping = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafNatDestinationRuleTranSLAtionOptionsPortMapping = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"nat", "destination", "rule", "translation", "options"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o NatDestinationRuleTranSLAtionOptions) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"address_mapping": types.StringType,
		"port_mapping":    types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o NatDestinationRuleTranSLAtionOptions) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address_mapping": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Address mapping options

|  Format  |  Description  |
|----------|---------------|
|  persistent  |  Gives a client the same source or destination-address for each connection  |
|  random  |  Random source or destination address allocation for each connection  |

`,

			// Default:          stringdefault.StaticString(`random`),
			Computed: true,
		},

		"port_mapping": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Port mapping options

|  Format  |  Description  |
|----------|---------------|
|  random  |  Randomize source port mapping  |
|  fully-random  |  Full port randomization  |
|  none  |  Do not apply port randomization  |

`,

			// Default:          stringdefault.StaticString(`none`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
