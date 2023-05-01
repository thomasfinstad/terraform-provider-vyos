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

// InterfacesWwanAuthentication describes the resource data model.
type InterfacesWwanAuthentication struct {
	// LeafNodes
	LeafInterfacesWwanAuthenticationUsername types.String `tfsdk:"username"`
	LeafInterfacesWwanAuthenticationPassword types.String `tfsdk:"password"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *InterfacesWwanAuthentication) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"interfaces", "wwan", "authentication"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafInterfacesWwanAuthenticationUsername.IsNull() || o.LeafInterfacesWwanAuthenticationUsername.IsUnknown()) {
		vyosData["username"] = o.LeafInterfacesWwanAuthenticationUsername.ValueString()
	}
	if !(o.LeafInterfacesWwanAuthenticationPassword.IsNull() || o.LeafInterfacesWwanAuthenticationPassword.IsUnknown()) {
		vyosData["password"] = o.LeafInterfacesWwanAuthenticationPassword.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *InterfacesWwanAuthentication) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"interfaces", "wwan", "authentication"}})

	// Leafs
	if value, ok := vyosData["username"]; ok {
		o.LeafInterfacesWwanAuthenticationUsername = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWwanAuthenticationUsername = basetypes.NewStringNull()
	}
	if value, ok := vyosData["password"]; ok {
		o.LeafInterfacesWwanAuthenticationPassword = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWwanAuthenticationPassword = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"interfaces", "wwan", "authentication"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o InterfacesWwanAuthentication) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"username": types.StringType,
		"password": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWwanAuthentication) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"username": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Username used for authentication

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Username  |

`,
		},

		"password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Password used for authentication

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Password  |

`,
		},

		// TagNodes

		// Nodes

	}
}
