// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesWirelessVifMirror describes the resource data model.
type InterfacesWirelessVifMirror struct {
	// LeafNodes
	LeafInterfacesWirelessVifMirrorIngress types.String `tfsdk:"ingress" json:"ingress,omitempty"`
	LeafInterfacesWirelessVifMirrorEgress  types.String `tfsdk:"egress" json:"egress,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWirelessVifMirror) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"ingress": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Mirror ingress traffic to destination interface

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Destination interface name  |

`,
		},

		"egress": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Mirror egress traffic to destination interface

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Destination interface name  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesWirelessVifMirror) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesWirelessVifMirrorIngress.IsNull() && !o.LeafInterfacesWirelessVifMirrorIngress.IsUnknown() {
		jsonData["ingress"] = o.LeafInterfacesWirelessVifMirrorIngress.ValueString()
	}

	if !o.LeafInterfacesWirelessVifMirrorEgress.IsNull() && !o.LeafInterfacesWirelessVifMirrorEgress.IsUnknown() {
		jsonData["egress"] = o.LeafInterfacesWirelessVifMirrorEgress.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesWirelessVifMirror) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["ingress"]; ok {
		o.LeafInterfacesWirelessVifMirrorIngress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifMirrorIngress = basetypes.NewStringNull()
	}

	if value, ok := jsonData["egress"]; ok {
		o.LeafInterfacesWirelessVifMirrorEgress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifMirrorEgress = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
