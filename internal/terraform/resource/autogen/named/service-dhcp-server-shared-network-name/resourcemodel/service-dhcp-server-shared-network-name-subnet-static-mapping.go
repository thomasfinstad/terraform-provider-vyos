// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ServiceDhcpServerSharedNetworkNameSubnetStaticMapping describes the resource data model.
type ServiceDhcpServerSharedNetworkNameSubnetStaticMapping struct {
	// LeafNodes
	LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingDisable                 types.String `tfsdk:"disable" json:"disable,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingIPAddress               types.String `tfsdk:"ip_address" json:"ip-address,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingMacAddress              types.String `tfsdk:"mac_address" json:"mac-address,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingStaticMappingParameters types.String `tfsdk:"static_mapping_parameters" json:"static-mapping-parameters,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDhcpServerSharedNetworkNameSubnetStaticMapping) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"disable": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
		},

		"ip_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Fixed IP address of static mapping

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address used in static mapping  |

`,
		},

		"mac_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Media Access Control (MAC) address

|  Format  |  Description  |
|----------|---------------|
|  macaddr  |  Hardware (MAC) address  |

`,
		},

		"static_mapping_parameters": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Additional static-mapping parameters for DHCP server. Will be placed inside the "host" block of the mapping. You must use the syntax of dhcpd.conf in this text-field. Using this without proper knowledge may result in a crashed DHCP server. Check system log to look for errors.

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServiceDhcpServerSharedNetworkNameSubnetStaticMapping) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingDisable.IsNull() && !o.LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingDisable.IsUnknown() {
		jsonData["disable"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingDisable.ValueString()
	}

	if !o.LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingIPAddress.IsNull() && !o.LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingIPAddress.IsUnknown() {
		jsonData["ip-address"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingIPAddress.ValueString()
	}

	if !o.LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingMacAddress.IsNull() && !o.LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingMacAddress.IsUnknown() {
		jsonData["mac-address"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingMacAddress.ValueString()
	}

	if !o.LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingStaticMappingParameters.IsNull() && !o.LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingStaticMappingParameters.IsUnknown() {
		jsonData["static-mapping-parameters"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingStaticMappingParameters.ValueString()
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
func (o *ServiceDhcpServerSharedNetworkNameSubnetStaticMapping) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["disable"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingDisable = basetypes.NewStringNull()
	}

	if value, ok := jsonData["ip-address"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingIPAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingIPAddress = basetypes.NewStringNull()
	}

	if value, ok := jsonData["mac-address"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingMacAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingMacAddress = basetypes.NewStringNull()
	}

	if value, ok := jsonData["static-mapping-parameters"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingStaticMappingParameters = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingStaticMappingParameters = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
