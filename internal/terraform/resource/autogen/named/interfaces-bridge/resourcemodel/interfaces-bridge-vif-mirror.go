// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesBrIDgeVifMirror describes the resource data model.
type InterfacesBrIDgeVifMirror struct {
	// LeafNodes
	LeafInterfacesBrIDgeVifMirrorIngress types.String `tfsdk:"ingress" json:"ingress,omitempty"`
	LeafInterfacesBrIDgeVifMirrorEgress  types.String `tfsdk:"egress" json:"egress,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesBrIDgeVifMirror) ResourceSchemaAttributes() map[string]schema.Attribute {
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
func (o *InterfacesBrIDgeVifMirror) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesBrIDgeVifMirrorIngress.IsNull() && !o.LeafInterfacesBrIDgeVifMirrorIngress.IsUnknown() {
		jsonData["ingress"] = o.LeafInterfacesBrIDgeVifMirrorIngress.ValueString()
	}

	if !o.LeafInterfacesBrIDgeVifMirrorEgress.IsNull() && !o.LeafInterfacesBrIDgeVifMirrorEgress.IsUnknown() {
		jsonData["egress"] = o.LeafInterfacesBrIDgeVifMirrorEgress.ValueString()
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
func (o *InterfacesBrIDgeVifMirror) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["ingress"]; ok {
		o.LeafInterfacesBrIDgeVifMirrorIngress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeVifMirrorIngress = basetypes.NewStringNull()
	}

	if value, ok := jsonData["egress"]; ok {
		o.LeafInterfacesBrIDgeVifMirrorEgress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeVifMirrorEgress = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
