// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// PolicyRouteRuleIcmp describes the resource data model.
type PolicyRouteRuleIcmp struct {
	// LeafNodes
	PolicyRouteRuleIcmpCode     customtypes.CustomStringValue `tfsdk:"code" json:"code,omitempty"`
	PolicyRouteRuleIcmpType     customtypes.CustomStringValue `tfsdk:"type" json:"type,omitempty"`
	PolicyRouteRuleIcmpTypeName customtypes.CustomStringValue `tfsdk:"type_name" json:"type-name,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o PolicyRouteRuleIcmp) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"code": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `ICMP code (0-255)

|  Format  |  Description  |
|----------|---------------|
|  u32:0-255  |  ICMP code (0-255)  |
`,
		},

		"type": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `ICMP type (0-255)

|  Format  |  Description  |
|----------|---------------|
|  u32:0-255  |  ICMP type (0-255)  |
`,
		},

		"type_name": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `ICMP type-name

|  Format  |  Description  |
|----------|---------------|
|  echo-reply  |  ICMP type 0: echo-reply  |
|  destination-unreachable  |  ICMP type 3: destination-unreachable  |
|  source-quench  |  ICMP type 4: source-quench  |
|  redirect  |  ICMP type 5: redirect  |
|  echo-request  |  ICMP type 8: echo-request  |
|  router-advertisement  |  ICMP type 9: router-advertisement  |
|  router-solicitation  |  ICMP type 10: router-solicitation  |
|  time-exceeded  |  ICMP type 11: time-exceeded  |
|  parameter-problem  |  ICMP type 12: parameter-problem  |
|  timestamp-request  |  ICMP type 13: timestamp-request  |
|  timestamp-reply  |  ICMP type 14: timestamp-reply  |
|  info-request  |  ICMP type 15: info-request  |
|  info-reply  |  ICMP type 16: info-reply  |
|  address-mask-request  |  ICMP type 17: address-mask-request  |
|  address-mask-reply  |  ICMP type 18: address-mask-reply  |
`,
		},

		// TagNodes

		// Nodes

	}
}
