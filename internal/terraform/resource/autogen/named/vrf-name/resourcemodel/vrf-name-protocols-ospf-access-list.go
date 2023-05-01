// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsOspfAccessList describes the resource data model.
type VrfNameProtocolsOspfAccessList struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfAccessListExport types.String `tfsdk:"export" json:"export,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfAccessList) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"export": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Filter for outgoing routing update

|  Format  |  Description  |
|----------|---------------|
|  bgp  |  Filter BGP routes  |
|  connected  |  Filter connected routes  |
|  isis  |  Filter IS-IS routes  |
|  kernel  |  Filter Kernel routes  |
|  rip  |  Filter RIP routes  |
|  static  |  Filter static routes  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsOspfAccessList) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameProtocolsOspfAccessListExport.IsNull() && !o.LeafVrfNameProtocolsOspfAccessListExport.IsUnknown() {
		jsonData["export"] = o.LeafVrfNameProtocolsOspfAccessListExport.ValueString()
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
func (o *VrfNameProtocolsOspfAccessList) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["export"]; ok {
		o.LeafVrfNameProtocolsOspfAccessListExport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsOspfAccessListExport = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
