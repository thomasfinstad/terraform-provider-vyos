// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesBrIDgeDhcpOptions describes the resource data model.
type InterfacesBrIDgeDhcpOptions struct {
	// LeafNodes
	LeafInterfacesBrIDgeDhcpOptionsClientID             types.String `tfsdk:"client_id" json:"client-id,omitempty"`
	LeafInterfacesBrIDgeDhcpOptionsHostName             types.String `tfsdk:"host_name" json:"host-name,omitempty"`
	LeafInterfacesBrIDgeDhcpOptionsMtu                  types.String `tfsdk:"mtu" json:"mtu,omitempty"`
	LeafInterfacesBrIDgeDhcpOptionsVendorClassID        types.String `tfsdk:"vendor_class_id" json:"vendor-class-id,omitempty"`
	LeafInterfacesBrIDgeDhcpOptionsNoDefaultRoute       types.String `tfsdk:"no_default_route" json:"no-default-route,omitempty"`
	LeafInterfacesBrIDgeDhcpOptionsDefaultRouteDistance types.String `tfsdk:"default_route_distance" json:"default-route-distance,omitempty"`
	LeafInterfacesBrIDgeDhcpOptionsReject               types.String `tfsdk:"reject" json:"reject,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesBrIDgeDhcpOptions) ResourceSchemaAttributes() map[string]schema.Attribute {
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
func (o *InterfacesBrIDgeDhcpOptions) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesBrIDgeDhcpOptionsClientID.IsNull() && !o.LeafInterfacesBrIDgeDhcpOptionsClientID.IsUnknown() {
		jsonData["client-id"] = o.LeafInterfacesBrIDgeDhcpOptionsClientID.ValueString()
	}

	if !o.LeafInterfacesBrIDgeDhcpOptionsHostName.IsNull() && !o.LeafInterfacesBrIDgeDhcpOptionsHostName.IsUnknown() {
		jsonData["host-name"] = o.LeafInterfacesBrIDgeDhcpOptionsHostName.ValueString()
	}

	if !o.LeafInterfacesBrIDgeDhcpOptionsMtu.IsNull() && !o.LeafInterfacesBrIDgeDhcpOptionsMtu.IsUnknown() {
		jsonData["mtu"] = o.LeafInterfacesBrIDgeDhcpOptionsMtu.ValueString()
	}

	if !o.LeafInterfacesBrIDgeDhcpOptionsVendorClassID.IsNull() && !o.LeafInterfacesBrIDgeDhcpOptionsVendorClassID.IsUnknown() {
		jsonData["vendor-class-id"] = o.LeafInterfacesBrIDgeDhcpOptionsVendorClassID.ValueString()
	}

	if !o.LeafInterfacesBrIDgeDhcpOptionsNoDefaultRoute.IsNull() && !o.LeafInterfacesBrIDgeDhcpOptionsNoDefaultRoute.IsUnknown() {
		jsonData["no-default-route"] = o.LeafInterfacesBrIDgeDhcpOptionsNoDefaultRoute.ValueString()
	}

	if !o.LeafInterfacesBrIDgeDhcpOptionsDefaultRouteDistance.IsNull() && !o.LeafInterfacesBrIDgeDhcpOptionsDefaultRouteDistance.IsUnknown() {
		jsonData["default-route-distance"] = o.LeafInterfacesBrIDgeDhcpOptionsDefaultRouteDistance.ValueString()
	}

	if !o.LeafInterfacesBrIDgeDhcpOptionsReject.IsNull() && !o.LeafInterfacesBrIDgeDhcpOptionsReject.IsUnknown() {
		jsonData["reject"] = o.LeafInterfacesBrIDgeDhcpOptionsReject.ValueString()
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
func (o *InterfacesBrIDgeDhcpOptions) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["client-id"]; ok {
		o.LeafInterfacesBrIDgeDhcpOptionsClientID = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeDhcpOptionsClientID = basetypes.NewStringNull()
	}

	if value, ok := jsonData["host-name"]; ok {
		o.LeafInterfacesBrIDgeDhcpOptionsHostName = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeDhcpOptionsHostName = basetypes.NewStringNull()
	}

	if value, ok := jsonData["mtu"]; ok {
		o.LeafInterfacesBrIDgeDhcpOptionsMtu = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeDhcpOptionsMtu = basetypes.NewStringNull()
	}

	if value, ok := jsonData["vendor-class-id"]; ok {
		o.LeafInterfacesBrIDgeDhcpOptionsVendorClassID = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeDhcpOptionsVendorClassID = basetypes.NewStringNull()
	}

	if value, ok := jsonData["no-default-route"]; ok {
		o.LeafInterfacesBrIDgeDhcpOptionsNoDefaultRoute = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeDhcpOptionsNoDefaultRoute = basetypes.NewStringNull()
	}

	if value, ok := jsonData["default-route-distance"]; ok {
		o.LeafInterfacesBrIDgeDhcpOptionsDefaultRouteDistance = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeDhcpOptionsDefaultRouteDistance = basetypes.NewStringNull()
	}

	if value, ok := jsonData["reject"]; ok {
		o.LeafInterfacesBrIDgeDhcpOptionsReject = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeDhcpOptionsReject = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
