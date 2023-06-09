// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesPseudoEthernetDhcpOptions describes the resource data model.
type InterfacesPseudoEthernetDhcpOptions struct {
	// LeafNodes
	LeafInterfacesPseudoEthernetDhcpOptionsClientID             types.String `tfsdk:"client_id" json:"client-id,omitempty"`
	LeafInterfacesPseudoEthernetDhcpOptionsHostName             types.String `tfsdk:"host_name" json:"host-name,omitempty"`
	LeafInterfacesPseudoEthernetDhcpOptionsMtu                  types.String `tfsdk:"mtu" json:"mtu,omitempty"`
	LeafInterfacesPseudoEthernetDhcpOptionsVendorClassID        types.String `tfsdk:"vendor_class_id" json:"vendor-class-id,omitempty"`
	LeafInterfacesPseudoEthernetDhcpOptionsNoDefaultRoute       types.String `tfsdk:"no_default_route" json:"no-default-route,omitempty"`
	LeafInterfacesPseudoEthernetDhcpOptionsDefaultRouteDistance types.String `tfsdk:"default_route_distance" json:"default-route-distance,omitempty"`
	LeafInterfacesPseudoEthernetDhcpOptionsReject               types.String `tfsdk:"reject" json:"reject,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesPseudoEthernetDhcpOptions) ResourceSchemaAttributes() map[string]schema.Attribute {
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
func (o *InterfacesPseudoEthernetDhcpOptions) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesPseudoEthernetDhcpOptionsClientID.IsNull() && !o.LeafInterfacesPseudoEthernetDhcpOptionsClientID.IsUnknown() {
		jsonData["client-id"] = o.LeafInterfacesPseudoEthernetDhcpOptionsClientID.ValueString()
	}

	if !o.LeafInterfacesPseudoEthernetDhcpOptionsHostName.IsNull() && !o.LeafInterfacesPseudoEthernetDhcpOptionsHostName.IsUnknown() {
		jsonData["host-name"] = o.LeafInterfacesPseudoEthernetDhcpOptionsHostName.ValueString()
	}

	if !o.LeafInterfacesPseudoEthernetDhcpOptionsMtu.IsNull() && !o.LeafInterfacesPseudoEthernetDhcpOptionsMtu.IsUnknown() {
		jsonData["mtu"] = o.LeafInterfacesPseudoEthernetDhcpOptionsMtu.ValueString()
	}

	if !o.LeafInterfacesPseudoEthernetDhcpOptionsVendorClassID.IsNull() && !o.LeafInterfacesPseudoEthernetDhcpOptionsVendorClassID.IsUnknown() {
		jsonData["vendor-class-id"] = o.LeafInterfacesPseudoEthernetDhcpOptionsVendorClassID.ValueString()
	}

	if !o.LeafInterfacesPseudoEthernetDhcpOptionsNoDefaultRoute.IsNull() && !o.LeafInterfacesPseudoEthernetDhcpOptionsNoDefaultRoute.IsUnknown() {
		jsonData["no-default-route"] = o.LeafInterfacesPseudoEthernetDhcpOptionsNoDefaultRoute.ValueString()
	}

	if !o.LeafInterfacesPseudoEthernetDhcpOptionsDefaultRouteDistance.IsNull() && !o.LeafInterfacesPseudoEthernetDhcpOptionsDefaultRouteDistance.IsUnknown() {
		jsonData["default-route-distance"] = o.LeafInterfacesPseudoEthernetDhcpOptionsDefaultRouteDistance.ValueString()
	}

	if !o.LeafInterfacesPseudoEthernetDhcpOptionsReject.IsNull() && !o.LeafInterfacesPseudoEthernetDhcpOptionsReject.IsUnknown() {
		jsonData["reject"] = o.LeafInterfacesPseudoEthernetDhcpOptionsReject.ValueString()
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
func (o *InterfacesPseudoEthernetDhcpOptions) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["client-id"]; ok {
		o.LeafInterfacesPseudoEthernetDhcpOptionsClientID = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesPseudoEthernetDhcpOptionsClientID = basetypes.NewStringNull()
	}

	if value, ok := jsonData["host-name"]; ok {
		o.LeafInterfacesPseudoEthernetDhcpOptionsHostName = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesPseudoEthernetDhcpOptionsHostName = basetypes.NewStringNull()
	}

	if value, ok := jsonData["mtu"]; ok {
		o.LeafInterfacesPseudoEthernetDhcpOptionsMtu = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesPseudoEthernetDhcpOptionsMtu = basetypes.NewStringNull()
	}

	if value, ok := jsonData["vendor-class-id"]; ok {
		o.LeafInterfacesPseudoEthernetDhcpOptionsVendorClassID = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesPseudoEthernetDhcpOptionsVendorClassID = basetypes.NewStringNull()
	}

	if value, ok := jsonData["no-default-route"]; ok {
		o.LeafInterfacesPseudoEthernetDhcpOptionsNoDefaultRoute = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesPseudoEthernetDhcpOptionsNoDefaultRoute = basetypes.NewStringNull()
	}

	if value, ok := jsonData["default-route-distance"]; ok {
		o.LeafInterfacesPseudoEthernetDhcpOptionsDefaultRouteDistance = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesPseudoEthernetDhcpOptionsDefaultRouteDistance = basetypes.NewStringNull()
	}

	if value, ok := jsonData["reject"]; ok {
		o.LeafInterfacesPseudoEthernetDhcpOptionsReject = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesPseudoEthernetDhcpOptionsReject = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
