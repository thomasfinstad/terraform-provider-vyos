// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ServiceDNSForwardingAuthoritativeDomainRecordsSpf describes the resource data model.
type ServiceDNSForwardingAuthoritativeDomainRecordsSpf struct {
	// LeafNodes
	ServiceDNSForwardingAuthoritativeDomainRecordsSpfValue   customtypes.CustomStringValue `tfsdk:"value" json:"value,omitempty"`
	ServiceDNSForwardingAuthoritativeDomainRecordsSpfTTL     customtypes.CustomStringValue `tfsdk:"ttl" json:"ttl,omitempty"`
	ServiceDNSForwardingAuthoritativeDomainRecordsSpfDisable customtypes.CustomStringValue `tfsdk:"disable" json:"disable,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ServiceDNSForwardingAuthoritativeDomainRecordsSpf) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"value": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Record contents

|  Format  |  Description  |
|----------|---------------|
|  text  |  Record contents  |
`,
		},

		"ttl": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Time-to-live (TTL)

|  Format  |  Description  |
|----------|---------------|
|  u32:0-2147483647  |  TTL in seconds  |
`,

			// Default:          stringdefault.StaticString(`300`),
			Computed: true,
		},

		"disable": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Disable instance

`,
		},

		// TagNodes

		// Nodes

	}
}
