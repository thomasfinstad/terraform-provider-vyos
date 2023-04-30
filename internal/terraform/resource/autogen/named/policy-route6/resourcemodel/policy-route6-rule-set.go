// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// PolicyRoutesixRuleSet describes the resource data model.
type PolicyRoutesixRuleSet struct {
	// LeafNodes
	PolicyRoutesixRuleSetConnectionMark customtypes.CustomStringValue `tfsdk:"connection_mark" json:"connection-mark,omitempty"`
	PolicyRoutesixRuleSetDscp           customtypes.CustomStringValue `tfsdk:"dscp" json:"dscp,omitempty"`
	PolicyRoutesixRuleSetMark           customtypes.CustomStringValue `tfsdk:"mark" json:"mark,omitempty"`
	PolicyRoutesixRuleSetTable          customtypes.CustomStringValue `tfsdk:"table" json:"table,omitempty"`
	PolicyRoutesixRuleSetTCPMss         customtypes.CustomStringValue `tfsdk:"tcp_mss" json:"tcp-mss,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o PolicyRoutesixRuleSet) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"connection_mark": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Connection marking

|  Format  |  Description  |
|----------|---------------|
|  u32:0-2147483647  |  Connection marking  |
`,
		},

		"dscp": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Packet Differentiated Services Codepoint (DSCP)

|  Format  |  Description  |
|----------|---------------|
|  u32:0-63  |  DSCP number  |
`,
		},

		"mark": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Packet marking

|  Format  |  Description  |
|----------|---------------|
|  u32:1-2147483647  |  Packet marking  |
`,
		},

		"table": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Routing table to forward packet with

|  Format  |  Description  |
|----------|---------------|
|  u32:1-200  |  Table number  |
|  main  |  Main table  |
`,
		},

		"tcp_mss": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `TCP Maximum Segment Size

|  Format  |  Description  |
|----------|---------------|
|  u32:500-1460  |  Explicitly set TCP MSS value  |
`,
		},

		// TagNodes

		// Nodes

	}
}
