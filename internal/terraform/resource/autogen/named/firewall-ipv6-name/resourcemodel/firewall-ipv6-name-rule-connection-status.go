// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// FirewallIPvsixNameRuleConnectionStatus describes the resource data model.
type FirewallIPvsixNameRuleConnectionStatus struct {
	// LeafNodes
	FirewallIPvsixNameRuleConnectionStatusNat customtypes.CustomStringValue `tfsdk:"nat" json:"nat,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o FirewallIPvsixNameRuleConnectionStatus) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"nat": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `NAT connection status

|  Format  |  Description  |
|----------|---------------|
|  destination  |  Match connections that are subject to destination NAT  |
|  source  |  Match connections that are subject to source NAT  |
`,
		},

		// TagNodes

		// Nodes

	}
}
