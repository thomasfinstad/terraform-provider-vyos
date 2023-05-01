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

// VrfNameProtocolsIsisInterfacePassword describes the resource data model.
type VrfNameProtocolsIsisInterfacePassword struct {
	// LeafNodes
	LeafVrfNameProtocolsIsisInterfacePasswordPlaintextPassword types.String `tfsdk:"plaintext_password"`
	LeafVrfNameProtocolsIsisInterfacePasswordMdfive            types.String `tfsdk:"md5"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsIsisInterfacePassword) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "isis", "interface", "password"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsIsisInterfacePasswordPlaintextPassword.IsNull() || o.LeafVrfNameProtocolsIsisInterfacePasswordPlaintextPassword.IsUnknown()) {
		vyosData["plaintext-password"] = o.LeafVrfNameProtocolsIsisInterfacePasswordPlaintextPassword.ValueString()
	}
	if !(o.LeafVrfNameProtocolsIsisInterfacePasswordMdfive.IsNull() || o.LeafVrfNameProtocolsIsisInterfacePasswordMdfive.IsUnknown()) {
		vyosData["md5"] = o.LeafVrfNameProtocolsIsisInterfacePasswordMdfive.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsIsisInterfacePassword) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "isis", "interface", "password"}})

	// Leafs
	if value, ok := vyosData["plaintext-password"]; ok {
		o.LeafVrfNameProtocolsIsisInterfacePasswordPlaintextPassword = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsIsisInterfacePasswordPlaintextPassword = basetypes.NewStringNull()
	}
	if value, ok := vyosData["md5"]; ok {
		o.LeafVrfNameProtocolsIsisInterfacePasswordMdfive = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsIsisInterfacePasswordMdfive = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "isis", "interface", "password"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsIsisInterfacePassword) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"plaintext_password": types.StringType,
		"md5":                types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisInterfacePassword) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"plaintext_password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Plain-text authentication type

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Circuit password  |

`,
		},

		"md5": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `MD5 authentication type

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Level-wide password  |

`,
		},

		// TagNodes

		// Nodes

	}
}