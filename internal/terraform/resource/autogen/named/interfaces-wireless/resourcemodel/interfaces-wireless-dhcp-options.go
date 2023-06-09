// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesWirelessDhcpOptions describes the resource data model.
type InterfacesWirelessDhcpOptions struct {
	// LeafNodes
	LeafInterfacesWirelessDhcpOptionsClientID             types.String `tfsdk:"client_id" json:"client-id,omitempty"`
	LeafInterfacesWirelessDhcpOptionsHostName             types.String `tfsdk:"host_name" json:"host-name,omitempty"`
	LeafInterfacesWirelessDhcpOptionsMtu                  types.String `tfsdk:"mtu" json:"mtu,omitempty"`
	LeafInterfacesWirelessDhcpOptionsVendorClassID        types.String `tfsdk:"vendor_class_id" json:"vendor-class-id,omitempty"`
	LeafInterfacesWirelessDhcpOptionsNoDefaultRoute       types.String `tfsdk:"no_default_route" json:"no-default-route,omitempty"`
	LeafInterfacesWirelessDhcpOptionsDefaultRouteDistance types.String `tfsdk:"default_route_distance" json:"default-route-distance,omitempty"`
	LeafInterfacesWirelessDhcpOptionsReject               types.String `tfsdk:"reject" json:"reject,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWirelessDhcpOptions) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"client_id": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Identifier used by client to identify itself to the DHCP server

`,
		},

		"host_name": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Override system host-name sent to DHCP server

`,
		},

		"mtu": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Use MTU value from DHCP server - ignore interface setting

`,
		},

		"vendor_class_id": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Identify the vendor client type to the DHCP server

`,
		},

		"no_default_route": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Do not install default route to system

`,
		},

		"default_route_distance": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Distance for installed default route

|  Format  |  Description  |
|----------|---------------|
|  u32:1-255  |  Distance for the default route from DHCP server  |

`,

			// Default:          stringdefault.StaticString(`210`),
			Computed: true,
		},

		"reject": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP addresses or subnets from which to reject DHCP leases

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address to match  |
|  ipv4net  |  IPv4 prefix to match  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesWirelessDhcpOptions) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesWirelessDhcpOptionsClientID.IsNull() && !o.LeafInterfacesWirelessDhcpOptionsClientID.IsUnknown() {
		jsonData["client-id"] = o.LeafInterfacesWirelessDhcpOptionsClientID.ValueString()
	}

	if !o.LeafInterfacesWirelessDhcpOptionsHostName.IsNull() && !o.LeafInterfacesWirelessDhcpOptionsHostName.IsUnknown() {
		jsonData["host-name"] = o.LeafInterfacesWirelessDhcpOptionsHostName.ValueString()
	}

	if !o.LeafInterfacesWirelessDhcpOptionsMtu.IsNull() && !o.LeafInterfacesWirelessDhcpOptionsMtu.IsUnknown() {
		jsonData["mtu"] = o.LeafInterfacesWirelessDhcpOptionsMtu.ValueString()
	}

	if !o.LeafInterfacesWirelessDhcpOptionsVendorClassID.IsNull() && !o.LeafInterfacesWirelessDhcpOptionsVendorClassID.IsUnknown() {
		jsonData["vendor-class-id"] = o.LeafInterfacesWirelessDhcpOptionsVendorClassID.ValueString()
	}

	if !o.LeafInterfacesWirelessDhcpOptionsNoDefaultRoute.IsNull() && !o.LeafInterfacesWirelessDhcpOptionsNoDefaultRoute.IsUnknown() {
		jsonData["no-default-route"] = o.LeafInterfacesWirelessDhcpOptionsNoDefaultRoute.ValueString()
	}

	if !o.LeafInterfacesWirelessDhcpOptionsDefaultRouteDistance.IsNull() && !o.LeafInterfacesWirelessDhcpOptionsDefaultRouteDistance.IsUnknown() {
		jsonData["default-route-distance"] = o.LeafInterfacesWirelessDhcpOptionsDefaultRouteDistance.ValueString()
	}

	if !o.LeafInterfacesWirelessDhcpOptionsReject.IsNull() && !o.LeafInterfacesWirelessDhcpOptionsReject.IsUnknown() {
		jsonData["reject"] = o.LeafInterfacesWirelessDhcpOptionsReject.ValueString()
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
func (o *InterfacesWirelessDhcpOptions) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["client-id"]; ok {
		o.LeafInterfacesWirelessDhcpOptionsClientID = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessDhcpOptionsClientID = basetypes.NewStringNull()
	}

	if value, ok := jsonData["host-name"]; ok {
		o.LeafInterfacesWirelessDhcpOptionsHostName = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessDhcpOptionsHostName = basetypes.NewStringNull()
	}

	if value, ok := jsonData["mtu"]; ok {
		o.LeafInterfacesWirelessDhcpOptionsMtu = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessDhcpOptionsMtu = basetypes.NewStringNull()
	}

	if value, ok := jsonData["vendor-class-id"]; ok {
		o.LeafInterfacesWirelessDhcpOptionsVendorClassID = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessDhcpOptionsVendorClassID = basetypes.NewStringNull()
	}

	if value, ok := jsonData["no-default-route"]; ok {
		o.LeafInterfacesWirelessDhcpOptionsNoDefaultRoute = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessDhcpOptionsNoDefaultRoute = basetypes.NewStringNull()
	}

	if value, ok := jsonData["default-route-distance"]; ok {
		o.LeafInterfacesWirelessDhcpOptionsDefaultRouteDistance = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessDhcpOptionsDefaultRouteDistance = basetypes.NewStringNull()
	}

	if value, ok := jsonData["reject"]; ok {
		o.LeafInterfacesWirelessDhcpOptionsReject = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessDhcpOptionsReject = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
