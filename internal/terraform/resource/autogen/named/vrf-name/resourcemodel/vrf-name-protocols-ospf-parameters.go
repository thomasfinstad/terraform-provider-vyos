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

// VrfNameProtocolsOspfParameters describes the resource data model.
type VrfNameProtocolsOspfParameters struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfParametersAbrType                           types.String `tfsdk:"abr_type"`
	LeafVrfNameProtocolsOspfParametersOpaqueLsa                         types.String `tfsdk:"opaque_lsa"`
	LeafVrfNameProtocolsOspfParametersRfconefiveeightthreeCompatibility types.String `tfsdk:"rfc1583_compatibility"`
	LeafVrfNameProtocolsOspfParametersRouterID                          types.String `tfsdk:"router_id"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsOspfParameters) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "ospf", "parameters"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsOspfParametersAbrType.IsNull() || o.LeafVrfNameProtocolsOspfParametersAbrType.IsUnknown()) {
		vyosData["abr-type"] = o.LeafVrfNameProtocolsOspfParametersAbrType.ValueString()
	}
	if !(o.LeafVrfNameProtocolsOspfParametersOpaqueLsa.IsNull() || o.LeafVrfNameProtocolsOspfParametersOpaqueLsa.IsUnknown()) {
		vyosData["opaque-lsa"] = o.LeafVrfNameProtocolsOspfParametersOpaqueLsa.ValueString()
	}
	if !(o.LeafVrfNameProtocolsOspfParametersRfconefiveeightthreeCompatibility.IsNull() || o.LeafVrfNameProtocolsOspfParametersRfconefiveeightthreeCompatibility.IsUnknown()) {
		vyosData["rfc1583-compatibility"] = o.LeafVrfNameProtocolsOspfParametersRfconefiveeightthreeCompatibility.ValueString()
	}
	if !(o.LeafVrfNameProtocolsOspfParametersRouterID.IsNull() || o.LeafVrfNameProtocolsOspfParametersRouterID.IsUnknown()) {
		vyosData["router-id"] = o.LeafVrfNameProtocolsOspfParametersRouterID.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsOspfParameters) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "ospf", "parameters"}})

	// Leafs
	if value, ok := vyosData["abr-type"]; ok {
		o.LeafVrfNameProtocolsOspfParametersAbrType = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsOspfParametersAbrType = basetypes.NewStringNull()
	}
	if value, ok := vyosData["opaque-lsa"]; ok {
		o.LeafVrfNameProtocolsOspfParametersOpaqueLsa = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsOspfParametersOpaqueLsa = basetypes.NewStringNull()
	}
	if value, ok := vyosData["rfc1583-compatibility"]; ok {
		o.LeafVrfNameProtocolsOspfParametersRfconefiveeightthreeCompatibility = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsOspfParametersRfconefiveeightthreeCompatibility = basetypes.NewStringNull()
	}
	if value, ok := vyosData["router-id"]; ok {
		o.LeafVrfNameProtocolsOspfParametersRouterID = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsOspfParametersRouterID = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "ospf", "parameters"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsOspfParameters) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"abr_type":              types.StringType,
		"opaque_lsa":            types.StringType,
		"rfc1583_compatibility": types.StringType,
		"router_id":             types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfParameters) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"abr_type": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `OSPF ABR type

|  Format  |  Description  |
|----------|---------------|
|  cisco  |  Cisco ABR type  |
|  ibm  |  IBM ABR type  |
|  shortcut  |  Shortcut ABR type  |
|  standard  |  Standard ABR type  |

`,

			// Default:          stringdefault.StaticString(`cisco`),
			Computed: true,
		},

		"opaque_lsa": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enable the Opaque-LSA capability (rfc2370)

`,
		},

		"rfc1583_compatibility": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enable RFC1583 criteria for handling AS external routes

`,
		},

		"router_id": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Override default router identifier

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Router-ID in IP address format  |

`,
		},

		// TagNodes

		// Nodes

	}
}