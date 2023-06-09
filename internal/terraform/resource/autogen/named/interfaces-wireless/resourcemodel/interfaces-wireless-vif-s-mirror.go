// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesWirelessVifSMirror describes the resource data model.
type InterfacesWirelessVifSMirror struct {
	// LeafNodes
	LeafInterfacesWirelessVifSMirrorIngress types.String `tfsdk:"ingress" json:"ingress,omitempty"`
	LeafInterfacesWirelessVifSMirrorEgress  types.String `tfsdk:"egress" json:"egress,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWirelessVifSMirror) ResourceSchemaAttributes() map[string]schema.Attribute {
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
func (o *InterfacesWirelessVifSMirror) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesWirelessVifSMirrorIngress.IsNull() && !o.LeafInterfacesWirelessVifSMirrorIngress.IsUnknown() {
		jsonData["ingress"] = o.LeafInterfacesWirelessVifSMirrorIngress.ValueString()
	}

	if !o.LeafInterfacesWirelessVifSMirrorEgress.IsNull() && !o.LeafInterfacesWirelessVifSMirrorEgress.IsUnknown() {
		jsonData["egress"] = o.LeafInterfacesWirelessVifSMirrorEgress.ValueString()
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
func (o *InterfacesWirelessVifSMirror) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["ingress"]; ok {
		o.LeafInterfacesWirelessVifSMirrorIngress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifSMirrorIngress = basetypes.NewStringNull()
	}

	if value, ok := jsonData["egress"]; ok {
		o.LeafInterfacesWirelessVifSMirrorEgress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifSMirrorEgress = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
