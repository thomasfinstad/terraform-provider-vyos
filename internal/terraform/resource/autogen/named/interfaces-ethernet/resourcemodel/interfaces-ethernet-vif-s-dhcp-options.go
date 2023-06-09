// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesEthernetVifSDhcpOptions describes the resource data model.
type InterfacesEthernetVifSDhcpOptions struct {
	// LeafNodes
	LeafInterfacesEthernetVifSDhcpOptionsClientID             types.String `tfsdk:"client_id" json:"client-id,omitempty"`
	LeafInterfacesEthernetVifSDhcpOptionsHostName             types.String `tfsdk:"host_name" json:"host-name,omitempty"`
	LeafInterfacesEthernetVifSDhcpOptionsMtu                  types.String `tfsdk:"mtu" json:"mtu,omitempty"`
	LeafInterfacesEthernetVifSDhcpOptionsVendorClassID        types.String `tfsdk:"vendor_class_id" json:"vendor-class-id,omitempty"`
	LeafInterfacesEthernetVifSDhcpOptionsNoDefaultRoute       types.String `tfsdk:"no_default_route" json:"no-default-route,omitempty"`
	LeafInterfacesEthernetVifSDhcpOptionsDefaultRouteDistance types.String `tfsdk:"default_route_distance" json:"default-route-distance,omitempty"`
	LeafInterfacesEthernetVifSDhcpOptionsReject               types.String `tfsdk:"reject" json:"reject,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesEthernetVifSDhcpOptions) ResourceSchemaAttributes() map[string]schema.Attribute {
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
func (o *InterfacesEthernetVifSDhcpOptions) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesEthernetVifSDhcpOptionsClientID.IsNull() && !o.LeafInterfacesEthernetVifSDhcpOptionsClientID.IsUnknown() {
		jsonData["client-id"] = o.LeafInterfacesEthernetVifSDhcpOptionsClientID.ValueString()
	}

	if !o.LeafInterfacesEthernetVifSDhcpOptionsHostName.IsNull() && !o.LeafInterfacesEthernetVifSDhcpOptionsHostName.IsUnknown() {
		jsonData["host-name"] = o.LeafInterfacesEthernetVifSDhcpOptionsHostName.ValueString()
	}

	if !o.LeafInterfacesEthernetVifSDhcpOptionsMtu.IsNull() && !o.LeafInterfacesEthernetVifSDhcpOptionsMtu.IsUnknown() {
		jsonData["mtu"] = o.LeafInterfacesEthernetVifSDhcpOptionsMtu.ValueString()
	}

	if !o.LeafInterfacesEthernetVifSDhcpOptionsVendorClassID.IsNull() && !o.LeafInterfacesEthernetVifSDhcpOptionsVendorClassID.IsUnknown() {
		jsonData["vendor-class-id"] = o.LeafInterfacesEthernetVifSDhcpOptionsVendorClassID.ValueString()
	}

	if !o.LeafInterfacesEthernetVifSDhcpOptionsNoDefaultRoute.IsNull() && !o.LeafInterfacesEthernetVifSDhcpOptionsNoDefaultRoute.IsUnknown() {
		jsonData["no-default-route"] = o.LeafInterfacesEthernetVifSDhcpOptionsNoDefaultRoute.ValueString()
	}

	if !o.LeafInterfacesEthernetVifSDhcpOptionsDefaultRouteDistance.IsNull() && !o.LeafInterfacesEthernetVifSDhcpOptionsDefaultRouteDistance.IsUnknown() {
		jsonData["default-route-distance"] = o.LeafInterfacesEthernetVifSDhcpOptionsDefaultRouteDistance.ValueString()
	}

	if !o.LeafInterfacesEthernetVifSDhcpOptionsReject.IsNull() && !o.LeafInterfacesEthernetVifSDhcpOptionsReject.IsUnknown() {
		jsonData["reject"] = o.LeafInterfacesEthernetVifSDhcpOptionsReject.ValueString()
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
func (o *InterfacesEthernetVifSDhcpOptions) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["client-id"]; ok {
		o.LeafInterfacesEthernetVifSDhcpOptionsClientID = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesEthernetVifSDhcpOptionsClientID = basetypes.NewStringNull()
	}

	if value, ok := jsonData["host-name"]; ok {
		o.LeafInterfacesEthernetVifSDhcpOptionsHostName = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesEthernetVifSDhcpOptionsHostName = basetypes.NewStringNull()
	}

	if value, ok := jsonData["mtu"]; ok {
		o.LeafInterfacesEthernetVifSDhcpOptionsMtu = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesEthernetVifSDhcpOptionsMtu = basetypes.NewStringNull()
	}

	if value, ok := jsonData["vendor-class-id"]; ok {
		o.LeafInterfacesEthernetVifSDhcpOptionsVendorClassID = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesEthernetVifSDhcpOptionsVendorClassID = basetypes.NewStringNull()
	}

	if value, ok := jsonData["no-default-route"]; ok {
		o.LeafInterfacesEthernetVifSDhcpOptionsNoDefaultRoute = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesEthernetVifSDhcpOptionsNoDefaultRoute = basetypes.NewStringNull()
	}

	if value, ok := jsonData["default-route-distance"]; ok {
		o.LeafInterfacesEthernetVifSDhcpOptionsDefaultRouteDistance = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesEthernetVifSDhcpOptionsDefaultRouteDistance = basetypes.NewStringNull()
	}

	if value, ok := jsonData["reject"]; ok {
		o.LeafInterfacesEthernetVifSDhcpOptionsReject = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesEthernetVifSDhcpOptionsReject = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
