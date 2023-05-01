// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesBondingVifSMirror describes the resource data model.
type InterfacesBondingVifSMirror struct {
	// LeafNodes
	LeafInterfacesBondingVifSMirrorIngress types.String `tfsdk:"ingress" json:"ingress,omitempty"`
	LeafInterfacesBondingVifSMirrorEgress  types.String `tfsdk:"egress" json:"egress,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesBondingVifSMirror) ResourceSchemaAttributes() map[string]schema.Attribute {
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
func (o *InterfacesBondingVifSMirror) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesBondingVifSMirrorIngress.IsNull() && !o.LeafInterfacesBondingVifSMirrorIngress.IsUnknown() {
		jsonData["ingress"] = o.LeafInterfacesBondingVifSMirrorIngress.ValueString()
	}

	if !o.LeafInterfacesBondingVifSMirrorEgress.IsNull() && !o.LeafInterfacesBondingVifSMirrorEgress.IsUnknown() {
		jsonData["egress"] = o.LeafInterfacesBondingVifSMirrorEgress.ValueString()
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
func (o *InterfacesBondingVifSMirror) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["ingress"]; ok {
		o.LeafInterfacesBondingVifSMirrorIngress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingVifSMirrorIngress = basetypes.NewStringNull()
	}

	if value, ok := jsonData["egress"]; ok {
		o.LeafInterfacesBondingVifSMirrorEgress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingVifSMirrorEgress = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
