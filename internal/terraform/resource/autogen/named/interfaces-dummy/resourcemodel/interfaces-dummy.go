// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesDummy describes the resource data model.
type InterfacesDummy struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafInterfacesDummyAddress     types.String `tfsdk:"address" json:"address,omitempty"`
	LeafInterfacesDummyDescrIPtion types.String `tfsdk:"description" json:"description,omitempty"`
	LeafInterfacesDummyDisable     types.String `tfsdk:"disable" json:"disable,omitempty"`
	LeafInterfacesDummyMtu         types.String `tfsdk:"mtu" json:"mtu,omitempty"`
	LeafInterfacesDummyNetns       types.String `tfsdk:"netns" json:"netns,omitempty"`
	LeafInterfacesDummyRedirect    types.String `tfsdk:"redirect" json:"redirect,omitempty"`
	LeafInterfacesDummyVrf         types.String `tfsdk:"vrf" json:"vrf,omitempty"`

	// TagNodes

	// Nodes
	NodeInterfacesDummyIP     *InterfacesDummyIP     `tfsdk:"ip" json:"ip,omitempty"`
	NodeInterfacesDummyIPvsix *InterfacesDummyIPvsix `tfsdk:"ipv6" json:"ipv6,omitempty"`
	NodeInterfacesDummyMirror *InterfacesDummyMirror `tfsdk:"mirror" json:"mirror,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesDummy) GetVyosPath() []string {
	return []string{
		"interfaces",
		"dummy",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesDummy) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Dummy Interface

|  Format  |  Description  |
|----------|---------------|
|  dumN  |  Dummy interface name  |

`,
		},

		// LeafNodes

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  IPv4 address and prefix length  |
|  ipv6net  |  IPv6 address and prefix length  |

`,
		},

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |

`,
		},

		"disable": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Administratively disable interface

`,
		},

		"mtu": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum Transmission Unit (MTU)

|  Format  |  Description  |
|----------|---------------|
|  u32:68-16000  |  Maximum Transmission Unit in byte  |

`,

			// Default:          stringdefault.StaticString(`1500`),
			Computed: true,
		},

		"netns": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Network namespace name

|  Format  |  Description  |
|----------|---------------|
|  text  |  Network namespace name  |

`,
		},

		"redirect": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Redirect incoming packet to destination

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Destination interface name  |

`,
		},

		"vrf": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VRF instance name

|  Format  |  Description  |
|----------|---------------|
|  txt  |  VRF instance name  |

`,
		},

		// TagNodes

		// Nodes

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesDummyIP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesDummyIPvsix{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesDummyMirror{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesDummy) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesDummyAddress.IsNull() && !o.LeafInterfacesDummyAddress.IsUnknown() {
		jsonData["address"] = o.LeafInterfacesDummyAddress.ValueString()
	}

	if !o.LeafInterfacesDummyDescrIPtion.IsNull() && !o.LeafInterfacesDummyDescrIPtion.IsUnknown() {
		jsonData["description"] = o.LeafInterfacesDummyDescrIPtion.ValueString()
	}

	if !o.LeafInterfacesDummyDisable.IsNull() && !o.LeafInterfacesDummyDisable.IsUnknown() {
		jsonData["disable"] = o.LeafInterfacesDummyDisable.ValueString()
	}

	if !o.LeafInterfacesDummyMtu.IsNull() && !o.LeafInterfacesDummyMtu.IsUnknown() {
		jsonData["mtu"] = o.LeafInterfacesDummyMtu.ValueString()
	}

	if !o.LeafInterfacesDummyNetns.IsNull() && !o.LeafInterfacesDummyNetns.IsUnknown() {
		jsonData["netns"] = o.LeafInterfacesDummyNetns.ValueString()
	}

	if !o.LeafInterfacesDummyRedirect.IsNull() && !o.LeafInterfacesDummyRedirect.IsUnknown() {
		jsonData["redirect"] = o.LeafInterfacesDummyRedirect.ValueString()
	}

	if !o.LeafInterfacesDummyVrf.IsNull() && !o.LeafInterfacesDummyVrf.IsUnknown() {
		jsonData["vrf"] = o.LeafInterfacesDummyVrf.ValueString()
	}

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeInterfacesDummyIP).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesDummyIP)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["ip"] = subData
	}

	if !reflect.ValueOf(o.NodeInterfacesDummyIPvsix).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesDummyIPvsix)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["ipv6"] = subData
	}

	if !reflect.ValueOf(o.NodeInterfacesDummyMirror).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesDummyMirror)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["mirror"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesDummy) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["address"]; ok {
		o.LeafInterfacesDummyAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesDummyAddress = basetypes.NewStringNull()
	}

	if value, ok := jsonData["description"]; ok {
		o.LeafInterfacesDummyDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesDummyDescrIPtion = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable"]; ok {
		o.LeafInterfacesDummyDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesDummyDisable = basetypes.NewStringNull()
	}

	if value, ok := jsonData["mtu"]; ok {
		o.LeafInterfacesDummyMtu = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesDummyMtu = basetypes.NewStringNull()
	}

	if value, ok := jsonData["netns"]; ok {
		o.LeafInterfacesDummyNetns = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesDummyNetns = basetypes.NewStringNull()
	}

	if value, ok := jsonData["redirect"]; ok {
		o.LeafInterfacesDummyRedirect = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesDummyRedirect = basetypes.NewStringNull()
	}

	if value, ok := jsonData["vrf"]; ok {
		o.LeafInterfacesDummyVrf = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesDummyVrf = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := jsonData["ip"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesDummyIP = &InterfacesDummyIP{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesDummyIP)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["ipv6"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesDummyIPvsix = &InterfacesDummyIPvsix{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesDummyIPvsix)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["mirror"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesDummyMirror = &InterfacesDummyMirror{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesDummyMirror)
		if err != nil {
			return err
		}
	}

	return nil
}
