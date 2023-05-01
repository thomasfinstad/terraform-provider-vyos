// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// FirewallIPvsixNameRuleIcmpvsix describes the resource data model.
type FirewallIPvsixNameRuleIcmpvsix struct {
	// LeafNodes
	LeafFirewallIPvsixNameRuleIcmpvsixCode     types.String `tfsdk:"code" json:"code,omitempty"`
	LeafFirewallIPvsixNameRuleIcmpvsixType     types.String `tfsdk:"type" json:"type,omitempty"`
	LeafFirewallIPvsixNameRuleIcmpvsixTypeName types.String `tfsdk:"type_name" json:"type-name,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixNameRuleIcmpvsix) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"code": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `ICMPv6 code

|  Format  |  Description  |
|----------|---------------|
|  u32:0-255  |  ICMPv6 code (0-255)  |

`,
		},

		"type": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `ICMPv6 type

|  Format  |  Description  |
|----------|---------------|
|  u32:0-255  |  ICMPv6 type (0-255)  |

`,
		},

		"type_name": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `ICMPv6 type-name

|  Format  |  Description  |
|----------|---------------|
|  destination-unreachable  |  ICMPv6 type 1: destination-unreachable  |
|  packet-too-big  |  ICMPv6 type 2: packet-too-big  |
|  time-exceeded  |  ICMPv6 type 3: time-exceeded  |
|  echo-request  |  ICMPv6 type 128: echo-request  |
|  echo-reply  |  ICMPv6 type 129: echo-reply  |
|  mld-listener-query  |  ICMPv6 type 130: mld-listener-query  |
|  mld-listener-report  |  ICMPv6 type 131: mld-listener-report  |
|  mld-listener-reduction  |  ICMPv6 type 132: mld-listener-reduction  |
|  nd-router-solicit  |  ICMPv6 type 133: nd-router-solicit  |
|  nd-router-advert  |  ICMPv6 type 134: nd-router-advert  |
|  nd-neighbor-solicit  |  ICMPv6 type 135: nd-neighbor-solicit  |
|  nd-neighbor-advert  |  ICMPv6 type 136: nd-neighbor-advert  |
|  nd-redirect  |  ICMPv6 type 137: nd-redirect  |
|  parameter-problem  |  ICMPv6 type 4: parameter-problem  |
|  router-renumbering  |  ICMPv6 type 138: router-renumbering  |
|  ind-neighbor-solicit  |  ICMPv6 type 141: ind-neighbor-solicit  |
|  ind-neighbor-advert  |  ICMPv6 type 142: ind-neighbor-advert  |
|  mld2-listener-report  |  ICMPv6 type 143: mld2-listener-report  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *FirewallIPvsixNameRuleIcmpvsix) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafFirewallIPvsixNameRuleIcmpvsixCode.IsNull() && !o.LeafFirewallIPvsixNameRuleIcmpvsixCode.IsUnknown() {
		jsonData["code"] = o.LeafFirewallIPvsixNameRuleIcmpvsixCode.ValueString()
	}

	if !o.LeafFirewallIPvsixNameRuleIcmpvsixType.IsNull() && !o.LeafFirewallIPvsixNameRuleIcmpvsixType.IsUnknown() {
		jsonData["type"] = o.LeafFirewallIPvsixNameRuleIcmpvsixType.ValueString()
	}

	if !o.LeafFirewallIPvsixNameRuleIcmpvsixTypeName.IsNull() && !o.LeafFirewallIPvsixNameRuleIcmpvsixTypeName.IsUnknown() {
		jsonData["type-name"] = o.LeafFirewallIPvsixNameRuleIcmpvsixTypeName.ValueString()
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
func (o *FirewallIPvsixNameRuleIcmpvsix) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["code"]; ok {
		o.LeafFirewallIPvsixNameRuleIcmpvsixCode = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRuleIcmpvsixCode = basetypes.NewStringNull()
	}

	if value, ok := jsonData["type"]; ok {
		o.LeafFirewallIPvsixNameRuleIcmpvsixType = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRuleIcmpvsixType = basetypes.NewStringNull()
	}

	if value, ok := jsonData["type-name"]; ok {
		o.LeafFirewallIPvsixNameRuleIcmpvsixTypeName = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRuleIcmpvsixTypeName = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
