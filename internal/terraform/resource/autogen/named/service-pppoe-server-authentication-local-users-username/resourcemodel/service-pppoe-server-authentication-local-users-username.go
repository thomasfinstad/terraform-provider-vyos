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

// ServicePppoeServerAuthenticationLocalUsersUsername describes the resource data model.
type ServicePppoeServerAuthenticationLocalUsersUsername struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafServicePppoeServerAuthenticationLocalUsersUsernameDisable  types.String `tfsdk:"disable"`
	LeafServicePppoeServerAuthenticationLocalUsersUsernamePassword types.String `tfsdk:"password"`
	LeafServicePppoeServerAuthenticationLocalUsersUsernameStaticIP types.String `tfsdk:"static_ip"`

	// TagNodes

	// Nodes
	NodeServicePppoeServerAuthenticationLocalUsersUsernameRateLimit types.Object `tfsdk:"rate_limit"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServicePppoeServerAuthenticationLocalUsersUsername) GetVyosPath() []string {
	return []string{
		"service",
		"pppoe-server",
		"authentication",
		"local-users",
		"username",
		o.ID.ValueString(),
	}
}

// TerraformToVyos converts terraform data to vyos data
func (o *ServicePppoeServerAuthenticationLocalUsersUsername) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"service", "pppoe-server", "authentication", "local-users", "username"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafServicePppoeServerAuthenticationLocalUsersUsernameDisable.IsNull() || o.LeafServicePppoeServerAuthenticationLocalUsersUsernameDisable.IsUnknown()) {
		vyosData["disable"] = o.LeafServicePppoeServerAuthenticationLocalUsersUsernameDisable.ValueString()
	}
	if !(o.LeafServicePppoeServerAuthenticationLocalUsersUsernamePassword.IsNull() || o.LeafServicePppoeServerAuthenticationLocalUsersUsernamePassword.IsUnknown()) {
		vyosData["password"] = o.LeafServicePppoeServerAuthenticationLocalUsersUsernamePassword.ValueString()
	}
	if !(o.LeafServicePppoeServerAuthenticationLocalUsersUsernameStaticIP.IsNull() || o.LeafServicePppoeServerAuthenticationLocalUsersUsernameStaticIP.IsUnknown()) {
		vyosData["static-ip"] = o.LeafServicePppoeServerAuthenticationLocalUsersUsernameStaticIP.ValueString()
	}

	// Tags

	// Nodes
	if !(o.NodeServicePppoeServerAuthenticationLocalUsersUsernameRateLimit.IsNull() || o.NodeServicePppoeServerAuthenticationLocalUsersUsernameRateLimit.IsUnknown()) {
		var subModel ServicePppoeServerAuthenticationLocalUsersUsernameRateLimit
		diags.Append(o.NodeServicePppoeServerAuthenticationLocalUsersUsernameRateLimit.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["rate-limit"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ServicePppoeServerAuthenticationLocalUsersUsername) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"service", "pppoe-server", "authentication", "local-users", "username"}})

	// Leafs
	if value, ok := vyosData["disable"]; ok {
		o.LeafServicePppoeServerAuthenticationLocalUsersUsernameDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServicePppoeServerAuthenticationLocalUsersUsernameDisable = basetypes.NewStringNull()
	}
	if value, ok := vyosData["password"]; ok {
		o.LeafServicePppoeServerAuthenticationLocalUsersUsernamePassword = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServicePppoeServerAuthenticationLocalUsersUsernamePassword = basetypes.NewStringNull()
	}
	if value, ok := vyosData["static-ip"]; ok {
		o.LeafServicePppoeServerAuthenticationLocalUsersUsernameStaticIP = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServicePppoeServerAuthenticationLocalUsersUsernameStaticIP = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := vyosData["rate-limit"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, ServicePppoeServerAuthenticationLocalUsersUsernameRateLimit{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeServicePppoeServerAuthenticationLocalUsersUsernameRateLimit = data

	} else {
		o.NodeServicePppoeServerAuthenticationLocalUsersUsernameRateLimit = basetypes.NewObjectNull(ServicePppoeServerAuthenticationLocalUsersUsernameRateLimit{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"service", "pppoe-server", "authentication", "local-users", "username"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ServicePppoeServerAuthenticationLocalUsersUsername) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"disable":   types.StringType,
		"password":  types.StringType,
		"static_ip": types.StringType,

		// Tags

		// Nodes
		"rate_limit": types.ObjectType{AttrTypes: ServicePppoeServerAuthenticationLocalUsersUsernameRateLimit{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServicePppoeServerAuthenticationLocalUsersUsername) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `User name for authentication

`,
		},

		// LeafNodes

		"disable": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
		},

		"password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Password for authentication

`,
		},

		"static_ip": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Static client IP address

`,

			// Default:          stringdefault.StaticString(`&`),
			Computed: true,
		},

		// TagNodes

		// Nodes

		"rate_limit": schema.SingleNestedAttribute{
			Attributes: ServicePppoeServerAuthenticationLocalUsersUsernameRateLimit{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Upload/Download speed limits

`,
		},
	}
}