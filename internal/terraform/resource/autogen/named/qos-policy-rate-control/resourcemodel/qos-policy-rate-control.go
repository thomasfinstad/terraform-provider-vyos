// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// QosPolicyRateControl describes the resource data model.
type QosPolicyRateControl struct {
	// LeafNodes
	QosPolicyRateControlDescrIPtion customtypes.CustomStringValue `tfsdk:"description" json:"description,omitempty"`
	QosPolicyRateControlBandwIDth   customtypes.CustomStringValue `tfsdk:"bandwidth" json:"bandwidth,omitempty"`
	QosPolicyRateControlBurst       customtypes.CustomStringValue `tfsdk:"burst" json:"burst,omitempty"`
	QosPolicyRateControlLatency     customtypes.CustomStringValue `tfsdk:"latency" json:"latency,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o QosPolicyRateControl) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"description": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |
`,
		},

		"bandwidth": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Available bandwidth for this policy

|  Format  |  Description  |
|----------|---------------|
|  <number>  |  Bits per second  |
|  <number>bit  |  Bits per second  |
|  <number>kbit  |  Kilobits per second  |
|  <number>mbit  |  Megabits per second  |
|  <number>gbit  |  Gigabits per second  |
|  <number>tbit  |  Terabits per second  |
|  <number>%  |  Percentage of interface link speed  |
`,
		},

		"burst": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Burst size for this class

|  Format  |  Description  |
|----------|---------------|
|  <number>  |  Bytes  |
|  <number><suffix>  |  Bytes with scaling suffix (kb, mb, gb)  |
`,

			// Default:          stringdefault.StaticString(`15k`),
			Computed: true,
		},

		"latency": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Maximum latency

|  Format  |  Description  |
|----------|---------------|
|  <number>  |  Time in milliseconds  |
`,

			// Default:          stringdefault.StaticString(`50`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
