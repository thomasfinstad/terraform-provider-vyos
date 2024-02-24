// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VpnSstpAuthenticationRadiusRateLimit describes the resource data model.
type VpnSstpAuthenticationRadiusRateLimit struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafVpnSstpAuthenticationRadiusRateLimitAttribute  types.String `tfsdk:"attribute" vyos:"attribute,omitempty"`
	LeafVpnSstpAuthenticationRadiusRateLimitVendor     types.String `tfsdk:"vendor" vyos:"vendor,omitempty"`
	LeafVpnSstpAuthenticationRadiusRateLimitEnable     types.Bool   `tfsdk:"enable" vyos:"enable,omitempty"`
	LeafVpnSstpAuthenticationRadiusRateLimitMultIPlier types.String `tfsdk:"multiplier" vyos:"multiplier,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *VpnSstpAuthenticationRadiusRateLimit) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnSstpAuthenticationRadiusRateLimit) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"vpn",

		"sstp",

		"authentication",

		"radius",

		"rate-limit",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnSstpAuthenticationRadiusRateLimit) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"attribute": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `RADIUS attribute that contains rate information

`,

			// Default:          stringdefault.StaticString(`Filter-Id`),
			Computed: true,
		},

		"vendor": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Vendor dictionary

`,
		},

		"enable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable bandwidth shaping via RADIUS

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"multiplier": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Shaper multiplier

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  <0.001-1000>  &emsp; |  Shaper multiplier  |

`,

			// Default:          stringdefault.StaticString(`1`),
			Computed: true,
		},
	}
}
