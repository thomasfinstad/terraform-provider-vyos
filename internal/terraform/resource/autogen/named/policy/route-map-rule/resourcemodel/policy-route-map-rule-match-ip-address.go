// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &PolicyRouteMapRuleMatchIPAddress{}

// PolicyRouteMapRuleMatchIPAddress describes the resource data model.
type PolicyRouteMapRuleMatchIPAddress struct {
	// LeafNodes
	LeafPolicyRouteMapRuleMatchIPAddressAccessList types.Number `tfsdk:"access_list" vyos:"access-list,omitempty"`
	LeafPolicyRouteMapRuleMatchIPAddressPrefixList types.String `tfsdk:"prefix_list" vyos:"prefix-list,omitempty"`
	LeafPolicyRouteMapRuleMatchIPAddressPrefixLen  types.Number `tfsdk:"prefix_len" vyos:"prefix-len,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRuleMatchIPAddress) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"access_list": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `IP access-list to match

    |  Format     &emsp;|  Description                               |
    |-------------------|--------------------------------------------|
    |  1-99       &emsp;|  IP standard access list                   |
    |  100-199    &emsp;|  IP extended access list                   |
    |  1300-1999  &emsp;|  IP standard access list (expanded range)  |
    |  2000-2699  &emsp;|  IP extended access list (expanded range)  |
`,
			Description: `IP access-list to match

    |  Format     |  Description                               |
    |-------------------|--------------------------------------------|
    |  1-99       |  IP standard access list                   |
    |  100-199    |  IP extended access list                   |
    |  1300-1999  |  IP standard access list (expanded range)  |
    |  2000-2699  |  IP extended access list (expanded range)  |
`,
		},

		"prefix_list": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP prefix-list to match

`,
			Description: `IP prefix-list to match

`,
		},

		"prefix_len": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `IP prefix-length to match (can be used for kernel routes only)

    |  Format  &emsp;|  Description    |
    |----------------|-----------------|
    |  0-32    &emsp;|  Prefix length  |
`,
			Description: `IP prefix-length to match (can be used for kernel routes only)

    |  Format  |  Description    |
    |----------------|-----------------|
    |  0-32    |  Prefix length  |
`,
		},

		// Nodes

	}
}
