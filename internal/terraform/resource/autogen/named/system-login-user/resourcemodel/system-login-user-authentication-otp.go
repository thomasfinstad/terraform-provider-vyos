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

// SystemLoginUserAuthenticationOtp describes the resource data model.
type SystemLoginUserAuthenticationOtp struct {
	// LeafNodes
	LeafSystemLoginUserAuthenticationOtpRateLimit  types.String `tfsdk:"rate_limit"`
	LeafSystemLoginUserAuthenticationOtpRateTime   types.String `tfsdk:"rate_time"`
	LeafSystemLoginUserAuthenticationOtpWindowSize types.String `tfsdk:"window_size"`
	LeafSystemLoginUserAuthenticationOtpKey        types.String `tfsdk:"key"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *SystemLoginUserAuthenticationOtp) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"system", "login", "user", "authentication", "otp"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafSystemLoginUserAuthenticationOtpRateLimit.IsNull() || o.LeafSystemLoginUserAuthenticationOtpRateLimit.IsUnknown()) {
		vyosData["rate-limit"] = o.LeafSystemLoginUserAuthenticationOtpRateLimit.ValueString()
	}
	if !(o.LeafSystemLoginUserAuthenticationOtpRateTime.IsNull() || o.LeafSystemLoginUserAuthenticationOtpRateTime.IsUnknown()) {
		vyosData["rate-time"] = o.LeafSystemLoginUserAuthenticationOtpRateTime.ValueString()
	}
	if !(o.LeafSystemLoginUserAuthenticationOtpWindowSize.IsNull() || o.LeafSystemLoginUserAuthenticationOtpWindowSize.IsUnknown()) {
		vyosData["window-size"] = o.LeafSystemLoginUserAuthenticationOtpWindowSize.ValueString()
	}
	if !(o.LeafSystemLoginUserAuthenticationOtpKey.IsNull() || o.LeafSystemLoginUserAuthenticationOtpKey.IsUnknown()) {
		vyosData["key"] = o.LeafSystemLoginUserAuthenticationOtpKey.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *SystemLoginUserAuthenticationOtp) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"system", "login", "user", "authentication", "otp"}})

	// Leafs
	if value, ok := vyosData["rate-limit"]; ok {
		o.LeafSystemLoginUserAuthenticationOtpRateLimit = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafSystemLoginUserAuthenticationOtpRateLimit = basetypes.NewStringNull()
	}
	if value, ok := vyosData["rate-time"]; ok {
		o.LeafSystemLoginUserAuthenticationOtpRateTime = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafSystemLoginUserAuthenticationOtpRateTime = basetypes.NewStringNull()
	}
	if value, ok := vyosData["window-size"]; ok {
		o.LeafSystemLoginUserAuthenticationOtpWindowSize = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafSystemLoginUserAuthenticationOtpWindowSize = basetypes.NewStringNull()
	}
	if value, ok := vyosData["key"]; ok {
		o.LeafSystemLoginUserAuthenticationOtpKey = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafSystemLoginUserAuthenticationOtpKey = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"system", "login", "user", "authentication", "otp"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o SystemLoginUserAuthenticationOtp) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"rate_limit":  types.StringType,
		"rate_time":   types.StringType,
		"window_size": types.StringType,
		"key":         types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemLoginUserAuthenticationOtp) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"rate_limit": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Limit number of logins (rate-limit) per rate-time

|  Format  |  Description  |
|----------|---------------|
|  u32:1-10  |  Number of attempts  |

`,

			// Default:          stringdefault.StaticString(`3`),
			Computed: true,
		},

		"rate_time": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Limit number of logins (rate-limit) per rate-time

|  Format  |  Description  |
|----------|---------------|
|  u32:15-600  |  Time interval  |

`,

			// Default:          stringdefault.StaticString(`30`),
			Computed: true,
		},

		"window_size": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set window of concurrently valid codes

|  Format  |  Description  |
|----------|---------------|
|  u32:1-21  |  Window size  |

`,

			// Default:          stringdefault.StaticString(`3`),
			Computed: true,
		},

		"key": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Key/secret the token algorithm (see RFC4226)

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Base32 encoded key/token  |

`,
		},

		// TagNodes

		// Nodes

	}
}