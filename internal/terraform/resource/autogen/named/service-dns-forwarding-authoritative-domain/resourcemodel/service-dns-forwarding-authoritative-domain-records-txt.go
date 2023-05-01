// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ServiceDNSForwardingAuthoritativeDomainRecordsTxt describes the resource data model.
type ServiceDNSForwardingAuthoritativeDomainRecordsTxt struct {
	// LeafNodes
	LeafServiceDNSForwardingAuthoritativeDomainRecordsTxtValue   types.String `tfsdk:"value"`
	LeafServiceDNSForwardingAuthoritativeDomainRecordsTxtTTL     types.String `tfsdk:"ttl"`
	LeafServiceDNSForwardingAuthoritativeDomainRecordsTxtDisable types.String `tfsdk:"disable"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsTxt) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"service", "dns", "forwarding", "authoritative-domain", "records", "txt"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafServiceDNSForwardingAuthoritativeDomainRecordsTxtValue.IsNull() || o.LeafServiceDNSForwardingAuthoritativeDomainRecordsTxtValue.IsUnknown()) {
		vyosData["value"] = o.LeafServiceDNSForwardingAuthoritativeDomainRecordsTxtValue.ValueString()
	}
	if !(o.LeafServiceDNSForwardingAuthoritativeDomainRecordsTxtTTL.IsNull() || o.LeafServiceDNSForwardingAuthoritativeDomainRecordsTxtTTL.IsUnknown()) {
		vyosData["ttl"] = o.LeafServiceDNSForwardingAuthoritativeDomainRecordsTxtTTL.ValueString()
	}
	if !(o.LeafServiceDNSForwardingAuthoritativeDomainRecordsTxtDisable.IsNull() || o.LeafServiceDNSForwardingAuthoritativeDomainRecordsTxtDisable.IsUnknown()) {
		vyosData["disable"] = o.LeafServiceDNSForwardingAuthoritativeDomainRecordsTxtDisable.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsTxt) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"service", "dns", "forwarding", "authoritative-domain", "records", "txt"}})

	// Leafs
	if value, ok := vyosData["value"]; ok {
		o.LeafServiceDNSForwardingAuthoritativeDomainRecordsTxtValue = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDNSForwardingAuthoritativeDomainRecordsTxtValue = basetypes.NewStringNull()
	}
	if value, ok := vyosData["ttl"]; ok {
		o.LeafServiceDNSForwardingAuthoritativeDomainRecordsTxtTTL = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDNSForwardingAuthoritativeDomainRecordsTxtTTL = basetypes.NewStringNull()
	}
	if value, ok := vyosData["disable"]; ok {
		o.LeafServiceDNSForwardingAuthoritativeDomainRecordsTxtDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDNSForwardingAuthoritativeDomainRecordsTxtDisable = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"service", "dns", "forwarding", "authoritative-domain", "records", "txt"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ServiceDNSForwardingAuthoritativeDomainRecordsTxt) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"value":   types.StringType,
		"ttl":     types.StringType,
		"disable": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDNSForwardingAuthoritativeDomainRecordsTxt) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"value": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Record contents

|  Format  |  Description  |
|----------|---------------|
|  text  |  Record contents  |

`,
		},

		"ttl": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Time-to-live (TTL)

|  Format  |  Description  |
|----------|---------------|
|  u32:0-2147483647  |  TTL in seconds  |

`,

			// Default:          stringdefault.StaticString(`300`),
			Computed: true,
		},

		"disable": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
		},

		// TagNodes

		// Nodes

	}
}