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

// VpnIPsecProfileAuthentication describes the resource data model.
type VpnIPsecProfileAuthentication struct {
	// LeafNodes
	LeafVpnIPsecProfileAuthenticationMode            types.String `tfsdk:"mode"`
	LeafVpnIPsecProfileAuthenticationPreSharedSecret types.String `tfsdk:"pre_shared_secret"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VpnIPsecProfileAuthentication) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vpn", "ipsec", "profile", "authentication"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVpnIPsecProfileAuthenticationMode.IsNull() || o.LeafVpnIPsecProfileAuthenticationMode.IsUnknown()) {
		vyosData["mode"] = o.LeafVpnIPsecProfileAuthenticationMode.ValueString()
	}
	if !(o.LeafVpnIPsecProfileAuthenticationPreSharedSecret.IsNull() || o.LeafVpnIPsecProfileAuthenticationPreSharedSecret.IsUnknown()) {
		vyosData["pre-shared-secret"] = o.LeafVpnIPsecProfileAuthenticationPreSharedSecret.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VpnIPsecProfileAuthentication) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vpn", "ipsec", "profile", "authentication"}})

	// Leafs
	if value, ok := vyosData["mode"]; ok {
		o.LeafVpnIPsecProfileAuthenticationMode = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecProfileAuthenticationMode = basetypes.NewStringNull()
	}
	if value, ok := vyosData["pre-shared-secret"]; ok {
		o.LeafVpnIPsecProfileAuthenticationPreSharedSecret = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecProfileAuthenticationPreSharedSecret = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vpn", "ipsec", "profile", "authentication"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VpnIPsecProfileAuthentication) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"mode":              types.StringType,
		"pre_shared_secret": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnIPsecProfileAuthentication) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"mode": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Authentication mode

|  Format  |  Description  |
|----------|---------------|
|  pre-shared-secret  |  Use a pre-shared secret key  |

`,
		},

		"pre_shared_secret": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Pre-shared secret key

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Pre-shared secret key  |

`,
		},

		// TagNodes

		// Nodes

	}
}