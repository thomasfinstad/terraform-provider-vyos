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

// VrfNameIPvsix describes the resource data model.
type VrfNameIPvsix struct {
	// LeafNodes
	LeafVrfNameIPvsixDisableForwarding types.String `tfsdk:"disable_forwarding"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameIPvsix) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "ipv6"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameIPvsixDisableForwarding.IsNull() || o.LeafVrfNameIPvsixDisableForwarding.IsUnknown()) {
		vyosData["disable-forwarding"] = o.LeafVrfNameIPvsixDisableForwarding.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameIPvsix) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "ipv6"}})

	// Leafs
	if value, ok := vyosData["disable-forwarding"]; ok {
		o.LeafVrfNameIPvsixDisableForwarding = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameIPvsixDisableForwarding = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "ipv6"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameIPvsix) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"disable_forwarding": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameIPvsix) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"disable_forwarding": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable IP forwarding on this interface

`,
		},

		// TagNodes

		// Nodes

	}
}
