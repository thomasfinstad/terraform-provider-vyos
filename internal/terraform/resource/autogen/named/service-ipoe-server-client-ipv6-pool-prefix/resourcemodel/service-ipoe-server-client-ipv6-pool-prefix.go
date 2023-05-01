// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ServiceIPoeServerClientIPvsixPoolPrefix describes the resource data model.
type ServiceIPoeServerClientIPvsixPoolPrefix struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafServiceIPoeServerClientIPvsixPoolPrefixMask types.String `tfsdk:"mask" json:"mask,omitempty"`

	// TagNodes

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceIPoeServerClientIPvsixPoolPrefix) GetVyosPath() []string {
	return []string{
		"service",
		"ipoe-server",
		"client-ipv6-pool",
		"prefix",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceIPoeServerClientIPvsixPoolPrefix) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Pool of addresses used to assign to clients

|  Format  |  Description  |
|----------|---------------|
|  ipv6net  |  IPv6 address and prefix length  |

`,
		},

		// LeafNodes

		"mask": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Prefix length used for individual client

|  Format  |  Description  |
|----------|---------------|
|  u32:48-128  |  Client prefix length  |

`,

			// Default:          stringdefault.StaticString(`64`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServiceIPoeServerClientIPvsixPoolPrefix) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafServiceIPoeServerClientIPvsixPoolPrefixMask.IsNull() && !o.LeafServiceIPoeServerClientIPvsixPoolPrefixMask.IsUnknown() {
		jsonData["mask"] = o.LeafServiceIPoeServerClientIPvsixPoolPrefixMask.ValueString()
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
func (o *ServiceIPoeServerClientIPvsixPoolPrefix) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["mask"]; ok {
		o.LeafServiceIPoeServerClientIPvsixPoolPrefixMask = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceIPoeServerClientIPvsixPoolPrefixMask = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
