// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesWirelessVifDhcpOptions describes the resource data model.
type InterfacesWirelessVifDhcpOptions struct {
	// LeafNodes
	LeafInterfacesWirelessVifDhcpOptionsClientID             types.String `tfsdk:"client_id" json:"client-id,omitempty"`
	LeafInterfacesWirelessVifDhcpOptionsHostName             types.String `tfsdk:"host_name" json:"host-name,omitempty"`
	LeafInterfacesWirelessVifDhcpOptionsMtu                  types.String `tfsdk:"mtu" json:"mtu,omitempty"`
	LeafInterfacesWirelessVifDhcpOptionsVendorClassID        types.String `tfsdk:"vendor_class_id" json:"vendor-class-id,omitempty"`
	LeafInterfacesWirelessVifDhcpOptionsNoDefaultRoute       types.String `tfsdk:"no_default_route" json:"no-default-route,omitempty"`
	LeafInterfacesWirelessVifDhcpOptionsDefaultRouteDistance types.String `tfsdk:"default_route_distance" json:"default-route-distance,omitempty"`
	LeafInterfacesWirelessVifDhcpOptionsReject               types.String `tfsdk:"reject" json:"reject,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWirelessVifDhcpOptions) ResourceSchemaAttributes() map[string]schema.Attribute {
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
func (o *InterfacesWirelessVifDhcpOptions) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesWirelessVifDhcpOptionsClientID.IsNull() && !o.LeafInterfacesWirelessVifDhcpOptionsClientID.IsUnknown() {
		jsonData["client-id"] = o.LeafInterfacesWirelessVifDhcpOptionsClientID.ValueString()
	}

	if !o.LeafInterfacesWirelessVifDhcpOptionsHostName.IsNull() && !o.LeafInterfacesWirelessVifDhcpOptionsHostName.IsUnknown() {
		jsonData["host-name"] = o.LeafInterfacesWirelessVifDhcpOptionsHostName.ValueString()
	}

	if !o.LeafInterfacesWirelessVifDhcpOptionsMtu.IsNull() && !o.LeafInterfacesWirelessVifDhcpOptionsMtu.IsUnknown() {
		jsonData["mtu"] = o.LeafInterfacesWirelessVifDhcpOptionsMtu.ValueString()
	}

	if !o.LeafInterfacesWirelessVifDhcpOptionsVendorClassID.IsNull() && !o.LeafInterfacesWirelessVifDhcpOptionsVendorClassID.IsUnknown() {
		jsonData["vendor-class-id"] = o.LeafInterfacesWirelessVifDhcpOptionsVendorClassID.ValueString()
	}

	if !o.LeafInterfacesWirelessVifDhcpOptionsNoDefaultRoute.IsNull() && !o.LeafInterfacesWirelessVifDhcpOptionsNoDefaultRoute.IsUnknown() {
		jsonData["no-default-route"] = o.LeafInterfacesWirelessVifDhcpOptionsNoDefaultRoute.ValueString()
	}

	if !o.LeafInterfacesWirelessVifDhcpOptionsDefaultRouteDistance.IsNull() && !o.LeafInterfacesWirelessVifDhcpOptionsDefaultRouteDistance.IsUnknown() {
		jsonData["default-route-distance"] = o.LeafInterfacesWirelessVifDhcpOptionsDefaultRouteDistance.ValueString()
	}

	if !o.LeafInterfacesWirelessVifDhcpOptionsReject.IsNull() && !o.LeafInterfacesWirelessVifDhcpOptionsReject.IsUnknown() {
		jsonData["reject"] = o.LeafInterfacesWirelessVifDhcpOptionsReject.ValueString()
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
func (o *InterfacesWirelessVifDhcpOptions) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["client-id"]; ok {
		o.LeafInterfacesWirelessVifDhcpOptionsClientID = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifDhcpOptionsClientID = basetypes.NewStringNull()
	}

	if value, ok := jsonData["host-name"]; ok {
		o.LeafInterfacesWirelessVifDhcpOptionsHostName = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifDhcpOptionsHostName = basetypes.NewStringNull()
	}

	if value, ok := jsonData["mtu"]; ok {
		o.LeafInterfacesWirelessVifDhcpOptionsMtu = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifDhcpOptionsMtu = basetypes.NewStringNull()
	}

	if value, ok := jsonData["vendor-class-id"]; ok {
		o.LeafInterfacesWirelessVifDhcpOptionsVendorClassID = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifDhcpOptionsVendorClassID = basetypes.NewStringNull()
	}

	if value, ok := jsonData["no-default-route"]; ok {
		o.LeafInterfacesWirelessVifDhcpOptionsNoDefaultRoute = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifDhcpOptionsNoDefaultRoute = basetypes.NewStringNull()
	}

	if value, ok := jsonData["default-route-distance"]; ok {
		o.LeafInterfacesWirelessVifDhcpOptionsDefaultRouteDistance = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifDhcpOptionsDefaultRouteDistance = basetypes.NewStringNull()
	}

	if value, ok := jsonData["reject"]; ok {
		o.LeafInterfacesWirelessVifDhcpOptionsReject = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifDhcpOptionsReject = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
