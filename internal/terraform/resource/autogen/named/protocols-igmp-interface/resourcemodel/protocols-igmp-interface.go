// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ProtocolsIgmpInterface describes the resource data model.
type ProtocolsIgmpInterface struct {
	// LeafNodes
	ProtocolsIgmpInterfaceVersion              customtypes.CustomStringValue `tfsdk:"version" json:"version,omitempty"`
	ProtocolsIgmpInterfaceQueryInterval        customtypes.CustomStringValue `tfsdk:"query_interval" json:"query-interval,omitempty"`
	ProtocolsIgmpInterfaceQueryMaxResponseTime customtypes.CustomStringValue `tfsdk:"query_max_response_time" json:"query-max-response-time,omitempty"`

	// TagNodes
	ProtocolsIgmpInterfaceJoin types.Map `tfsdk:"join" json:"join,omitempty"`

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ProtocolsIgmpInterface) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"version": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `IGMP version

|  Format  |  Description  |
|----------|---------------|
|  2  |  IGMP version 2  |
|  3  |  IGMP version 3  |
`,
		},

		"query_interval": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `IGMP host query interval

|  Format  |  Description  |
|----------|---------------|
|  u32:1-1800  |  Query interval in seconds  |
`,
		},

		"query_max_response_time": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `IGMP max query response time

|  Format  |  Description  |
|----------|---------------|
|  u32:10-250  |  Query response value in deci-seconds  |
`,
		},

		// TagNodes

		"join": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ProtocolsIgmpInterfaceJoin{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `IGMP join multicast group

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Multicast group address  |
`,
		},

		// Nodes

	}
}
