// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesMacsecSecURIty describes the resource data model.
type InterfacesMacsecSecURIty struct {
	// LeafNodes
	LeafInterfacesMacsecSecURItyCIPher       types.String `tfsdk:"cipher" json:"cipher,omitempty"`
	LeafInterfacesMacsecSecURItyEncrypt      types.String `tfsdk:"encrypt" json:"encrypt,omitempty"`
	LeafInterfacesMacsecSecURItyReplayWindow types.String `tfsdk:"replay_window" json:"replay-window,omitempty"`

	// TagNodes

	// Nodes
	NodeInterfacesMacsecSecURItyMka *InterfacesMacsecSecURItyMka `tfsdk:"mka" json:"mka,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesMacsecSecURIty) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"cipher": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Cipher suite used

|  Format  |  Description  |
|----------|---------------|
|  gcm-aes-128  |  Galois/Counter Mode of AES cipher with 128-bit key  |
|  gcm-aes-256  |  Galois/Counter Mode of AES cipher with 256-bit key  |

`,
		},

		"encrypt": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enable optional MACsec encryption

`,
		},

		"replay_window": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IEEE 802.1X/MACsec replay protection window

|  Format  |  Description  |
|----------|---------------|
|  u32:0  |  No replay window, strict check  |
|  u32:1-4294967295  |  Number of packets that could be misordered  |

`,
		},

		// TagNodes

		// Nodes

		"mka": schema.SingleNestedAttribute{
			Attributes: InterfacesMacsecSecURItyMka{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `MACsec Key Agreement protocol (MKA)

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesMacsecSecURIty) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesMacsecSecURItyCIPher.IsNull() && !o.LeafInterfacesMacsecSecURItyCIPher.IsUnknown() {
		jsonData["cipher"] = o.LeafInterfacesMacsecSecURItyCIPher.ValueString()
	}

	if !o.LeafInterfacesMacsecSecURItyEncrypt.IsNull() && !o.LeafInterfacesMacsecSecURItyEncrypt.IsUnknown() {
		jsonData["encrypt"] = o.LeafInterfacesMacsecSecURItyEncrypt.ValueString()
	}

	if !o.LeafInterfacesMacsecSecURItyReplayWindow.IsNull() && !o.LeafInterfacesMacsecSecURItyReplayWindow.IsUnknown() {
		jsonData["replay-window"] = o.LeafInterfacesMacsecSecURItyReplayWindow.ValueString()
	}

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeInterfacesMacsecSecURItyMka).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesMacsecSecURItyMka)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["mka"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesMacsecSecURIty) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["cipher"]; ok {
		o.LeafInterfacesMacsecSecURItyCIPher = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesMacsecSecURItyCIPher = basetypes.NewStringNull()
	}

	if value, ok := jsonData["encrypt"]; ok {
		o.LeafInterfacesMacsecSecURItyEncrypt = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesMacsecSecURItyEncrypt = basetypes.NewStringNull()
	}

	if value, ok := jsonData["replay-window"]; ok {
		o.LeafInterfacesMacsecSecURItyReplayWindow = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesMacsecSecURItyReplayWindow = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := jsonData["mka"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesMacsecSecURItyMka = &InterfacesMacsecSecURItyMka{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesMacsecSecURItyMka)
		if err != nil {
			return err
		}
	}

	return nil
}
