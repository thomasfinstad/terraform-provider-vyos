// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ServiceSnmpVthreeViewOID describes the resource data model.
type ServiceSnmpVthreeViewOID struct {
	// LeafNodes
	LeafServiceSnmpVthreeViewOIDExclude types.String `tfsdk:"exclude" json:"exclude,omitempty"`
	LeafServiceSnmpVthreeViewOIDMask    types.String `tfsdk:"mask" json:"mask,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceSnmpVthreeViewOID) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"exclude": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Exclude is an optional argument

`,
		},

		"mask": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Defines a bit-mask that is indicating which subidentifiers of the associated subtree OID should be regarded as significant

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServiceSnmpVthreeViewOID) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafServiceSnmpVthreeViewOIDExclude.IsNull() && !o.LeafServiceSnmpVthreeViewOIDExclude.IsUnknown() {
		jsonData["exclude"] = o.LeafServiceSnmpVthreeViewOIDExclude.ValueString()
	}

	if !o.LeafServiceSnmpVthreeViewOIDMask.IsNull() && !o.LeafServiceSnmpVthreeViewOIDMask.IsUnknown() {
		jsonData["mask"] = o.LeafServiceSnmpVthreeViewOIDMask.ValueString()
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
func (o *ServiceSnmpVthreeViewOID) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["exclude"]; ok {
		o.LeafServiceSnmpVthreeViewOIDExclude = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceSnmpVthreeViewOIDExclude = basetypes.NewStringNull()
	}

	if value, ok := jsonData["mask"]; ok {
		o.LeafServiceSnmpVthreeViewOIDMask = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceSnmpVthreeViewOIDMask = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
