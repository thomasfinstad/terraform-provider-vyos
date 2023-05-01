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

// ProtocolsFailoverRoute describes the resource data model.
type ProtocolsFailoverRoute struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes

	// TagNodes
	TagProtocolsFailoverRouteNextHop types.Map `tfsdk:"next_hop"`

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsFailoverRoute) GetVyosPath() []string {
	return []string{
		"protocols",
		"failover",
		"route",
		o.ID.ValueString(),
	}
}

// TerraformToVyos converts terraform data to vyos data
func (o *ProtocolsFailoverRoute) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"protocols", "failover", "route"}})

	vyosData := make(map[string]interface{})

	// Leafs

	// Tags
	if !(o.TagProtocolsFailoverRouteNextHop.IsNull() || o.TagProtocolsFailoverRouteNextHop.IsUnknown()) {
		subModel := make(map[string]ProtocolsFailoverRouteNextHop)
		diags.Append(o.TagProtocolsFailoverRouteNextHop.ElementsAs(ctx, &subModel, false)...)

		subData := make(map[string]interface{})
		for k, v := range subModel {
			subData[k] = v.TerraformToVyos(ctx, diags)
		}
		vyosData["next-hop"] = subData
	}

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ProtocolsFailoverRoute) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"protocols", "failover", "route"}})

	// Leafs

	// Tags
	if value, ok := vyosData["next-hop"]; ok {
		data, d := types.MapValueFrom(ctx, types.ObjectType{AttrTypes: ProtocolsFailoverRouteNextHop{}.AttributeTypes()}, value.(map[string]interface{}))
		diags.Append(d...)
		o.TagProtocolsFailoverRouteNextHop = data
	} else {
		o.TagProtocolsFailoverRouteNextHop = basetypes.NewMapNull(types.ObjectType{})
	}

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"protocols", "failover", "route"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ProtocolsFailoverRoute) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs

		// Tags
		"next_hop": types.MapType{ElemType: types.ObjectType{AttrTypes: ProtocolsFailoverRouteNextHop{}.AttributeTypes()}},

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsFailoverRoute) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Failover IPv4 route

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  IPv4 failover route  |

`,
		},

		// LeafNodes

		// TagNodes

		"next_hop": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ProtocolsFailoverRouteNextHop{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Next-hop IPv4 router address

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Next-hop router address  |

`,
		},

		// Nodes

	}
}