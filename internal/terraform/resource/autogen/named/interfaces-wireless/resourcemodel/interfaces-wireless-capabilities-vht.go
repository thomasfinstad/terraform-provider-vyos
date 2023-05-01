// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesWirelessCapabilitiesVht describes the resource data model.
type InterfacesWirelessCapabilitiesVht struct {
	// LeafNodes
	LeafInterfacesWirelessCapabilitiesVhtAntennaCount        types.String `tfsdk:"antenna_count" json:"antenna-count,omitempty"`
	LeafInterfacesWirelessCapabilitiesVhtAntennaPatternFixed types.String `tfsdk:"antenna_pattern_fixed" json:"antenna-pattern-fixed,omitempty"`
	LeafInterfacesWirelessCapabilitiesVhtBeamform            types.String `tfsdk:"beamform" json:"beamform,omitempty"`
	LeafInterfacesWirelessCapabilitiesVhtChannelSetWIDth     types.String `tfsdk:"channel_set_width" json:"channel-set-width,omitempty"`
	LeafInterfacesWirelessCapabilitiesVhtLdpc                types.String `tfsdk:"ldpc" json:"ldpc,omitempty"`
	LeafInterfacesWirelessCapabilitiesVhtLinkAdaptation      types.String `tfsdk:"link_adaptation" json:"link-adaptation,omitempty"`
	LeafInterfacesWirelessCapabilitiesVhtMaxMpduExp          types.String `tfsdk:"max_mpdu_exp" json:"max-mpdu-exp,omitempty"`
	LeafInterfacesWirelessCapabilitiesVhtMaxMpdu             types.String `tfsdk:"max_mpdu" json:"max-mpdu,omitempty"`
	LeafInterfacesWirelessCapabilitiesVhtShortGi             types.String `tfsdk:"short_gi" json:"short-gi,omitempty"`
	LeafInterfacesWirelessCapabilitiesVhtTxPowersave         types.String `tfsdk:"tx_powersave" json:"tx-powersave,omitempty"`
	LeafInterfacesWirelessCapabilitiesVhtVhtCf               types.String `tfsdk:"vht_cf" json:"vht-cf,omitempty"`

	// TagNodes

	// Nodes
	NodeInterfacesWirelessCapabilitiesVhtCenterChannelFreq *InterfacesWirelessCapabilitiesVhtCenterChannelFreq `tfsdk:"center_channel_freq" json:"center-channel-freq,omitempty"`
	NodeInterfacesWirelessCapabilitiesVhtStbc              *InterfacesWirelessCapabilitiesVhtStbc              `tfsdk:"stbc" json:"stbc,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWirelessCapabilitiesVht) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"antenna_count": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Number of antennas on this card

|  Format  |  Description  |
|----------|---------------|
|  u32:1-8  |  Number of antennas for this card  |

`,
		},

		"antenna_pattern_fixed": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set if antenna pattern does not change during the lifetime of an association

`,
		},

		"beamform": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Beamforming capabilities

|  Format  |  Description  |
|----------|---------------|
|  single-user-beamformer  |  Support for operation as single user beamformer  |
|  single-user-beamformee  |  Support for operation as single user beamformee  |
|  multi-user-beamformer  |  Support for operation as multi user beamformer  |
|  multi-user-beamformee  |  Support for operation as multi user beamformee  |

`,
		},

		"channel_set_width": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VHT operating Channel width

|  Format  |  Description  |
|----------|---------------|
|  0  |  20 or 40 MHz channel width  |
|  1  |  80 MHz channel width  |
|  2  |  160 MHz channel width  |
|  3  |  80+80 MHz channel width  |

`,
		},

		"ldpc": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enable LDPC (Low Density Parity Check) coding capability

`,
		},

		"link_adaptation": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VHT link adaptation capabilities

|  Format  |  Description  |
|----------|---------------|
|  unsolicited  |  Station provides only unsolicited VHT MFB  |
|  both  |  Station can provide VHT MFB in response to VHT MRQ and unsolicited VHT MFB  |

`,
		},

		"max_mpdu_exp": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set the maximum length of A-MPDU pre-EOF padding that the station can receive

|  Format  |  Description  |
|----------|---------------|
|  u32:0-7  |  Maximum length of A-MPDU pre-EOF padding = 2 pow(13 + x) -1 octets  |

`,
		},

		"max_mpdu": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Increase Maximum MPDU length to 7991 or 11454 octets (otherwise: 3895 octets)

|  Format  |  Description  |
|----------|---------------|
|  7991  |  ncrease Maximum MPDU length to 7991 octets  |
|  11454  |  ncrease Maximum MPDU length to 11454 octets  |

`,
		},

		"short_gi": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Short GI capabilities

|  Format  |  Description  |
|----------|---------------|
|  80  |  Short GI for 80 MHz  |
|  160  |  Short GI for 160 MHz  |

`,
		},

		"tx_powersave": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enable VHT TXOP Power Save Mode

`,
		},

		"vht_cf": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Station supports receiving VHT variant HT Control field

`,
		},

		// TagNodes

		// Nodes

		"center_channel_freq": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessCapabilitiesVhtCenterChannelFreq{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `VHT operating channel center frequency

`,
		},

		"stbc": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessCapabilitiesVhtStbc{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Support for sending and receiving PPDU using STBC (Space Time Block Coding)

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesWirelessCapabilitiesVht) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesWirelessCapabilitiesVhtAntennaCount.IsNull() && !o.LeafInterfacesWirelessCapabilitiesVhtAntennaCount.IsUnknown() {
		jsonData["antenna-count"] = o.LeafInterfacesWirelessCapabilitiesVhtAntennaCount.ValueString()
	}

	if !o.LeafInterfacesWirelessCapabilitiesVhtAntennaPatternFixed.IsNull() && !o.LeafInterfacesWirelessCapabilitiesVhtAntennaPatternFixed.IsUnknown() {
		jsonData["antenna-pattern-fixed"] = o.LeafInterfacesWirelessCapabilitiesVhtAntennaPatternFixed.ValueString()
	}

	if !o.LeafInterfacesWirelessCapabilitiesVhtBeamform.IsNull() && !o.LeafInterfacesWirelessCapabilitiesVhtBeamform.IsUnknown() {
		jsonData["beamform"] = o.LeafInterfacesWirelessCapabilitiesVhtBeamform.ValueString()
	}

	if !o.LeafInterfacesWirelessCapabilitiesVhtChannelSetWIDth.IsNull() && !o.LeafInterfacesWirelessCapabilitiesVhtChannelSetWIDth.IsUnknown() {
		jsonData["channel-set-width"] = o.LeafInterfacesWirelessCapabilitiesVhtChannelSetWIDth.ValueString()
	}

	if !o.LeafInterfacesWirelessCapabilitiesVhtLdpc.IsNull() && !o.LeafInterfacesWirelessCapabilitiesVhtLdpc.IsUnknown() {
		jsonData["ldpc"] = o.LeafInterfacesWirelessCapabilitiesVhtLdpc.ValueString()
	}

	if !o.LeafInterfacesWirelessCapabilitiesVhtLinkAdaptation.IsNull() && !o.LeafInterfacesWirelessCapabilitiesVhtLinkAdaptation.IsUnknown() {
		jsonData["link-adaptation"] = o.LeafInterfacesWirelessCapabilitiesVhtLinkAdaptation.ValueString()
	}

	if !o.LeafInterfacesWirelessCapabilitiesVhtMaxMpduExp.IsNull() && !o.LeafInterfacesWirelessCapabilitiesVhtMaxMpduExp.IsUnknown() {
		jsonData["max-mpdu-exp"] = o.LeafInterfacesWirelessCapabilitiesVhtMaxMpduExp.ValueString()
	}

	if !o.LeafInterfacesWirelessCapabilitiesVhtMaxMpdu.IsNull() && !o.LeafInterfacesWirelessCapabilitiesVhtMaxMpdu.IsUnknown() {
		jsonData["max-mpdu"] = o.LeafInterfacesWirelessCapabilitiesVhtMaxMpdu.ValueString()
	}

	if !o.LeafInterfacesWirelessCapabilitiesVhtShortGi.IsNull() && !o.LeafInterfacesWirelessCapabilitiesVhtShortGi.IsUnknown() {
		jsonData["short-gi"] = o.LeafInterfacesWirelessCapabilitiesVhtShortGi.ValueString()
	}

	if !o.LeafInterfacesWirelessCapabilitiesVhtTxPowersave.IsNull() && !o.LeafInterfacesWirelessCapabilitiesVhtTxPowersave.IsUnknown() {
		jsonData["tx-powersave"] = o.LeafInterfacesWirelessCapabilitiesVhtTxPowersave.ValueString()
	}

	if !o.LeafInterfacesWirelessCapabilitiesVhtVhtCf.IsNull() && !o.LeafInterfacesWirelessCapabilitiesVhtVhtCf.IsUnknown() {
		jsonData["vht-cf"] = o.LeafInterfacesWirelessCapabilitiesVhtVhtCf.ValueString()
	}

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeInterfacesWirelessCapabilitiesVhtCenterChannelFreq).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesWirelessCapabilitiesVhtCenterChannelFreq)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["center-channel-freq"] = subData
	}

	if !reflect.ValueOf(o.NodeInterfacesWirelessCapabilitiesVhtStbc).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesWirelessCapabilitiesVhtStbc)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["stbc"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesWirelessCapabilitiesVht) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["antenna-count"]; ok {
		o.LeafInterfacesWirelessCapabilitiesVhtAntennaCount = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessCapabilitiesVhtAntennaCount = basetypes.NewStringNull()
	}

	if value, ok := jsonData["antenna-pattern-fixed"]; ok {
		o.LeafInterfacesWirelessCapabilitiesVhtAntennaPatternFixed = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessCapabilitiesVhtAntennaPatternFixed = basetypes.NewStringNull()
	}

	if value, ok := jsonData["beamform"]; ok {
		o.LeafInterfacesWirelessCapabilitiesVhtBeamform = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessCapabilitiesVhtBeamform = basetypes.NewStringNull()
	}

	if value, ok := jsonData["channel-set-width"]; ok {
		o.LeafInterfacesWirelessCapabilitiesVhtChannelSetWIDth = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessCapabilitiesVhtChannelSetWIDth = basetypes.NewStringNull()
	}

	if value, ok := jsonData["ldpc"]; ok {
		o.LeafInterfacesWirelessCapabilitiesVhtLdpc = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessCapabilitiesVhtLdpc = basetypes.NewStringNull()
	}

	if value, ok := jsonData["link-adaptation"]; ok {
		o.LeafInterfacesWirelessCapabilitiesVhtLinkAdaptation = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessCapabilitiesVhtLinkAdaptation = basetypes.NewStringNull()
	}

	if value, ok := jsonData["max-mpdu-exp"]; ok {
		o.LeafInterfacesWirelessCapabilitiesVhtMaxMpduExp = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessCapabilitiesVhtMaxMpduExp = basetypes.NewStringNull()
	}

	if value, ok := jsonData["max-mpdu"]; ok {
		o.LeafInterfacesWirelessCapabilitiesVhtMaxMpdu = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessCapabilitiesVhtMaxMpdu = basetypes.NewStringNull()
	}

	if value, ok := jsonData["short-gi"]; ok {
		o.LeafInterfacesWirelessCapabilitiesVhtShortGi = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessCapabilitiesVhtShortGi = basetypes.NewStringNull()
	}

	if value, ok := jsonData["tx-powersave"]; ok {
		o.LeafInterfacesWirelessCapabilitiesVhtTxPowersave = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessCapabilitiesVhtTxPowersave = basetypes.NewStringNull()
	}

	if value, ok := jsonData["vht-cf"]; ok {
		o.LeafInterfacesWirelessCapabilitiesVhtVhtCf = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessCapabilitiesVhtVhtCf = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := jsonData["center-channel-freq"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesWirelessCapabilitiesVhtCenterChannelFreq = &InterfacesWirelessCapabilitiesVhtCenterChannelFreq{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesWirelessCapabilitiesVhtCenterChannelFreq)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["stbc"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesWirelessCapabilitiesVhtStbc = &InterfacesWirelessCapabilitiesVhtStbc{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesWirelessCapabilitiesVhtStbc)
		if err != nil {
			return err
		}
	}

	return nil
}
