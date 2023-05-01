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

// InterfacesWirelessCapabilitiesVhtCenterChannelFreq describes the resource data model.
type InterfacesWirelessCapabilitiesVhtCenterChannelFreq struct {
	// LeafNodes
	LeafInterfacesWirelessCapabilitiesVhtCenterChannelFreqFreqOne types.String `tfsdk:"freq_1"`
	LeafInterfacesWirelessCapabilitiesVhtCenterChannelFreqFreqTwo types.String `tfsdk:"freq_2"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *InterfacesWirelessCapabilitiesVhtCenterChannelFreq) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"interfaces", "wireless", "capabilities", "vht", "center-channel-freq"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafInterfacesWirelessCapabilitiesVhtCenterChannelFreqFreqOne.IsNull() || o.LeafInterfacesWirelessCapabilitiesVhtCenterChannelFreqFreqOne.IsUnknown()) {
		vyosData["freq-1"] = o.LeafInterfacesWirelessCapabilitiesVhtCenterChannelFreqFreqOne.ValueString()
	}
	if !(o.LeafInterfacesWirelessCapabilitiesVhtCenterChannelFreqFreqTwo.IsNull() || o.LeafInterfacesWirelessCapabilitiesVhtCenterChannelFreqFreqTwo.IsUnknown()) {
		vyosData["freq-2"] = o.LeafInterfacesWirelessCapabilitiesVhtCenterChannelFreqFreqTwo.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *InterfacesWirelessCapabilitiesVhtCenterChannelFreq) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"interfaces", "wireless", "capabilities", "vht", "center-channel-freq"}})

	// Leafs
	if value, ok := vyosData["freq-1"]; ok {
		o.LeafInterfacesWirelessCapabilitiesVhtCenterChannelFreqFreqOne = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessCapabilitiesVhtCenterChannelFreqFreqOne = basetypes.NewStringNull()
	}
	if value, ok := vyosData["freq-2"]; ok {
		o.LeafInterfacesWirelessCapabilitiesVhtCenterChannelFreqFreqTwo = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessCapabilitiesVhtCenterChannelFreqFreqTwo = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"interfaces", "wireless", "capabilities", "vht", "center-channel-freq"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o InterfacesWirelessCapabilitiesVhtCenterChannelFreq) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"freq_1": types.StringType,
		"freq_2": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWirelessCapabilitiesVhtCenterChannelFreq) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"freq_1": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VHT operating channel center frequency - center freq 1 (for use with 80, 80+80 and 160 modes)

|  Format  |  Description  |
|----------|---------------|
|  u32:34-173  |  5Ghz (802.11 a/h/j/n/ac) center channel index (use 42 for primary 80MHz channel 36)  |

`,
		},

		"freq_2": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VHT operating channel center frequency - center freq 2 (for use with the 80+80 mode)

|  Format  |  Description  |
|----------|---------------|
|  u32:34-173  |  5Ghz (802.11 a/h/j/n/ac) center channel index (use 58 for primary 80MHz channel 52)  |

`,
		},

		// TagNodes

		// Nodes

	}
}
