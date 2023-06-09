// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsEigrp describes the resource data model.
type VrfNameProtocolsEigrp struct {
	// LeafNodes
	LeafVrfNameProtocolsEigrpLocalAs          types.String `tfsdk:"local_as" json:"local-as,omitempty"`
	LeafVrfNameProtocolsEigrpMaximumPaths     types.String `tfsdk:"maximum_paths" json:"maximum-paths,omitempty"`
	LeafVrfNameProtocolsEigrpNetwork          types.String `tfsdk:"network" json:"network,omitempty"`
	LeafVrfNameProtocolsEigrpPassiveInterface types.String `tfsdk:"passive_interface" json:"passive-interface,omitempty"`
	LeafVrfNameProtocolsEigrpRedistribute     types.String `tfsdk:"redistribute" json:"redistribute,omitempty"`
	LeafVrfNameProtocolsEigrpRouteMap         types.String `tfsdk:"route_map" json:"route-map,omitempty"`
	LeafVrfNameProtocolsEigrpRouterID         types.String `tfsdk:"router_id" json:"router-id,omitempty"`
	LeafVrfNameProtocolsEigrpVariance         types.String `tfsdk:"variance" json:"variance,omitempty"`

	// TagNodes

	// Nodes
	NodeVrfNameProtocolsEigrpMetric *VrfNameProtocolsEigrpMetric `tfsdk:"metric" json:"metric,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsEigrp) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"local_as": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Autonomous System Number (ASN)

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Autonomous System Number  |

`,
		},

		"maximum_paths": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Forward packets over multiple paths

|  Format  |  Description  |
|----------|---------------|
|  u32:1-32  |  Number of paths  |

`,
		},

		"network": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enable routing on an IP network

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  EIGRP network prefix  |

`,
		},

		"passive_interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Suppress routing updates on an interface

`,
		},

		"redistribute": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Redistribute information from another routing protocol

|  Format  |  Description  |
|----------|---------------|
|  bgp  |  Border Gateway Protocol (BGP)  |
|  connected  |  Connected routes  |
|  nhrp  |  Next Hop Resolution Protocol (NHRP)  |
|  ospf  |  Open Shortest Path First (OSPFv2)  |
|  rip  |  Routing Information Protocol (RIP)  |
|  babel  |  Babel routing protocol (Babel)  |
|  static  |  Statically configured routes  |
|  vnc  |  Virtual Network Control (VNC)  |

`,
		},

		"route_map": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specify route-map name to use

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Route map name  |

`,
		},

		"router_id": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Override default router identifier

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Router-ID in IP address format  |

`,
		},

		"variance": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Control load balancing variance

|  Format  |  Description  |
|----------|---------------|
|  u32:1-128  |  Metric variance multiplier  |

`,
		},

		// TagNodes

		// Nodes

		"metric": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsEigrpMetric{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Modify metrics and parameters for advertisement

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsEigrp) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameProtocolsEigrpLocalAs.IsNull() && !o.LeafVrfNameProtocolsEigrpLocalAs.IsUnknown() {
		jsonData["local-as"] = o.LeafVrfNameProtocolsEigrpLocalAs.ValueString()
	}

	if !o.LeafVrfNameProtocolsEigrpMaximumPaths.IsNull() && !o.LeafVrfNameProtocolsEigrpMaximumPaths.IsUnknown() {
		jsonData["maximum-paths"] = o.LeafVrfNameProtocolsEigrpMaximumPaths.ValueString()
	}

	if !o.LeafVrfNameProtocolsEigrpNetwork.IsNull() && !o.LeafVrfNameProtocolsEigrpNetwork.IsUnknown() {
		jsonData["network"] = o.LeafVrfNameProtocolsEigrpNetwork.ValueString()
	}

	if !o.LeafVrfNameProtocolsEigrpPassiveInterface.IsNull() && !o.LeafVrfNameProtocolsEigrpPassiveInterface.IsUnknown() {
		jsonData["passive-interface"] = o.LeafVrfNameProtocolsEigrpPassiveInterface.ValueString()
	}

	if !o.LeafVrfNameProtocolsEigrpRedistribute.IsNull() && !o.LeafVrfNameProtocolsEigrpRedistribute.IsUnknown() {
		jsonData["redistribute"] = o.LeafVrfNameProtocolsEigrpRedistribute.ValueString()
	}

	if !o.LeafVrfNameProtocolsEigrpRouteMap.IsNull() && !o.LeafVrfNameProtocolsEigrpRouteMap.IsUnknown() {
		jsonData["route-map"] = o.LeafVrfNameProtocolsEigrpRouteMap.ValueString()
	}

	if !o.LeafVrfNameProtocolsEigrpRouterID.IsNull() && !o.LeafVrfNameProtocolsEigrpRouterID.IsUnknown() {
		jsonData["router-id"] = o.LeafVrfNameProtocolsEigrpRouterID.ValueString()
	}

	if !o.LeafVrfNameProtocolsEigrpVariance.IsNull() && !o.LeafVrfNameProtocolsEigrpVariance.IsUnknown() {
		jsonData["variance"] = o.LeafVrfNameProtocolsEigrpVariance.ValueString()
	}

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeVrfNameProtocolsEigrpMetric).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameProtocolsEigrpMetric)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["metric"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsEigrp) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["local-as"]; ok {
		o.LeafVrfNameProtocolsEigrpLocalAs = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsEigrpLocalAs = basetypes.NewStringNull()
	}

	if value, ok := jsonData["maximum-paths"]; ok {
		o.LeafVrfNameProtocolsEigrpMaximumPaths = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsEigrpMaximumPaths = basetypes.NewStringNull()
	}

	if value, ok := jsonData["network"]; ok {
		o.LeafVrfNameProtocolsEigrpNetwork = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsEigrpNetwork = basetypes.NewStringNull()
	}

	if value, ok := jsonData["passive-interface"]; ok {
		o.LeafVrfNameProtocolsEigrpPassiveInterface = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsEigrpPassiveInterface = basetypes.NewStringNull()
	}

	if value, ok := jsonData["redistribute"]; ok {
		o.LeafVrfNameProtocolsEigrpRedistribute = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsEigrpRedistribute = basetypes.NewStringNull()
	}

	if value, ok := jsonData["route-map"]; ok {
		o.LeafVrfNameProtocolsEigrpRouteMap = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsEigrpRouteMap = basetypes.NewStringNull()
	}

	if value, ok := jsonData["router-id"]; ok {
		o.LeafVrfNameProtocolsEigrpRouterID = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsEigrpRouterID = basetypes.NewStringNull()
	}

	if value, ok := jsonData["variance"]; ok {
		o.LeafVrfNameProtocolsEigrpVariance = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsEigrpVariance = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := jsonData["metric"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVrfNameProtocolsEigrpMetric = &VrfNameProtocolsEigrpMetric{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameProtocolsEigrpMetric)
		if err != nil {
			return err
		}
	}

	return nil
}
