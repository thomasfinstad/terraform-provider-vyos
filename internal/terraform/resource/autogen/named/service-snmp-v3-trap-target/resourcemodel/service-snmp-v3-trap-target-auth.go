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

// ServiceSnmpVthreeTrapTargetAuth describes the resource data model.
type ServiceSnmpVthreeTrapTargetAuth struct {
	// LeafNodes
	LeafServiceSnmpVthreeTrapTargetAuthEncryptedPassword types.String `tfsdk:"encrypted_password"`
	LeafServiceSnmpVthreeTrapTargetAuthPlaintextPassword types.String `tfsdk:"plaintext_password"`
	LeafServiceSnmpVthreeTrapTargetAuthType              types.String `tfsdk:"type"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *ServiceSnmpVthreeTrapTargetAuth) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"service", "snmp", "v3", "trap-target", "auth"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafServiceSnmpVthreeTrapTargetAuthEncryptedPassword.IsNull() || o.LeafServiceSnmpVthreeTrapTargetAuthEncryptedPassword.IsUnknown()) {
		vyosData["encrypted-password"] = o.LeafServiceSnmpVthreeTrapTargetAuthEncryptedPassword.ValueString()
	}
	if !(o.LeafServiceSnmpVthreeTrapTargetAuthPlaintextPassword.IsNull() || o.LeafServiceSnmpVthreeTrapTargetAuthPlaintextPassword.IsUnknown()) {
		vyosData["plaintext-password"] = o.LeafServiceSnmpVthreeTrapTargetAuthPlaintextPassword.ValueString()
	}
	if !(o.LeafServiceSnmpVthreeTrapTargetAuthType.IsNull() || o.LeafServiceSnmpVthreeTrapTargetAuthType.IsUnknown()) {
		vyosData["type"] = o.LeafServiceSnmpVthreeTrapTargetAuthType.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ServiceSnmpVthreeTrapTargetAuth) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"service", "snmp", "v3", "trap-target", "auth"}})

	// Leafs
	if value, ok := vyosData["encrypted-password"]; ok {
		o.LeafServiceSnmpVthreeTrapTargetAuthEncryptedPassword = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceSnmpVthreeTrapTargetAuthEncryptedPassword = basetypes.NewStringNull()
	}
	if value, ok := vyosData["plaintext-password"]; ok {
		o.LeafServiceSnmpVthreeTrapTargetAuthPlaintextPassword = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceSnmpVthreeTrapTargetAuthPlaintextPassword = basetypes.NewStringNull()
	}
	if value, ok := vyosData["type"]; ok {
		o.LeafServiceSnmpVthreeTrapTargetAuthType = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceSnmpVthreeTrapTargetAuthType = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"service", "snmp", "v3", "trap-target", "auth"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ServiceSnmpVthreeTrapTargetAuth) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"encrypted_password": types.StringType,
		"plaintext_password": types.StringType,
		"type":               types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceSnmpVthreeTrapTargetAuth) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"encrypted_password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Defines the encrypted key for authentication

`,
		},

		"plaintext_password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Defines the clear text key for authentication

`,
		},

		"type": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Define used protocol

|  Format  |  Description  |
|----------|---------------|
|  md5  |  Message Digest 5  |
|  sha  |  Secure Hash Algorithm  |

`,

			// Default:          stringdefault.StaticString(`md5`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}