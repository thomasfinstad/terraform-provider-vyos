// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// PolicyRouteMapRuleMatchEvpn describes the resource data model.
type PolicyRouteMapRuleMatchEvpn struct {
	// LeafNodes
	LeafPolicyRouteMapRuleMatchEvpnDefaultRoute types.String `tfsdk:"default_route" json:"default-route,omitempty"`
	LeafPolicyRouteMapRuleMatchEvpnRd           types.String `tfsdk:"rd" json:"rd,omitempty"`
	LeafPolicyRouteMapRuleMatchEvpnRouteType    types.String `tfsdk:"route_type" json:"route-type,omitempty"`
	LeafPolicyRouteMapRuleMatchEvpnVni          types.String `tfsdk:"vni" json:"vni,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRuleMatchEvpn) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"default_route": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Default EVPN type-5 route

`,
		},

		"rd": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route Distinguisher

|  Format  |  Description  |
|----------|---------------|
|  ASN:NN_OR_IP-ADDRESS:NN  |  Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)  |

`,
		},

		"route_type": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match route-type

|  Format  |  Description  |
|----------|---------------|
|  macip  |  mac-ip route  |
|  multicast  |  IMET route  |
|  prefix  |  Prefix route  |

`,
		},

		"vni": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Virtual Network Identifier

|  Format  |  Description  |
|----------|---------------|
|  u32:0-16777214  |  VXLAN virtual network identifier  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PolicyRouteMapRuleMatchEvpn) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafPolicyRouteMapRuleMatchEvpnDefaultRoute.IsNull() && !o.LeafPolicyRouteMapRuleMatchEvpnDefaultRoute.IsUnknown() {
		jsonData["default-route"] = o.LeafPolicyRouteMapRuleMatchEvpnDefaultRoute.ValueString()
	}

	if !o.LeafPolicyRouteMapRuleMatchEvpnRd.IsNull() && !o.LeafPolicyRouteMapRuleMatchEvpnRd.IsUnknown() {
		jsonData["rd"] = o.LeafPolicyRouteMapRuleMatchEvpnRd.ValueString()
	}

	if !o.LeafPolicyRouteMapRuleMatchEvpnRouteType.IsNull() && !o.LeafPolicyRouteMapRuleMatchEvpnRouteType.IsUnknown() {
		jsonData["route-type"] = o.LeafPolicyRouteMapRuleMatchEvpnRouteType.ValueString()
	}

	if !o.LeafPolicyRouteMapRuleMatchEvpnVni.IsNull() && !o.LeafPolicyRouteMapRuleMatchEvpnVni.IsUnknown() {
		jsonData["vni"] = o.LeafPolicyRouteMapRuleMatchEvpnVni.ValueString()
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
func (o *PolicyRouteMapRuleMatchEvpn) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["default-route"]; ok {
		o.LeafPolicyRouteMapRuleMatchEvpnDefaultRoute = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteMapRuleMatchEvpnDefaultRoute = basetypes.NewStringNull()
	}

	if value, ok := jsonData["rd"]; ok {
		o.LeafPolicyRouteMapRuleMatchEvpnRd = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteMapRuleMatchEvpnRd = basetypes.NewStringNull()
	}

	if value, ok := jsonData["route-type"]; ok {
		o.LeafPolicyRouteMapRuleMatchEvpnRouteType = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteMapRuleMatchEvpnRouteType = basetypes.NewStringNull()
	}

	if value, ok := jsonData["vni"]; ok {
		o.LeafPolicyRouteMapRuleMatchEvpnVni = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteMapRuleMatchEvpnVni = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
