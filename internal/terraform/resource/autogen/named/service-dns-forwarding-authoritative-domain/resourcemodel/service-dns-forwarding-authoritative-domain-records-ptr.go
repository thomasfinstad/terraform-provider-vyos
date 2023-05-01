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

// ServiceDNSForwardingAuthoritativeDomainRecordsPtr describes the resource data model.
type ServiceDNSForwardingAuthoritativeDomainRecordsPtr struct {
	// LeafNodes
	LeafServiceDNSForwardingAuthoritativeDomainRecordsPtrTarget  types.String `tfsdk:"target"`
	LeafServiceDNSForwardingAuthoritativeDomainRecordsPtrTTL     types.String `tfsdk:"ttl"`
	LeafServiceDNSForwardingAuthoritativeDomainRecordsPtrDisable types.String `tfsdk:"disable"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsPtr) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"service", "dns", "forwarding", "authoritative-domain", "records", "ptr"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafServiceDNSForwardingAuthoritativeDomainRecordsPtrTarget.IsNull() || o.LeafServiceDNSForwardingAuthoritativeDomainRecordsPtrTarget.IsUnknown()) {
		vyosData["target"] = o.LeafServiceDNSForwardingAuthoritativeDomainRecordsPtrTarget.ValueString()
	}
	if !(o.LeafServiceDNSForwardingAuthoritativeDomainRecordsPtrTTL.IsNull() || o.LeafServiceDNSForwardingAuthoritativeDomainRecordsPtrTTL.IsUnknown()) {
		vyosData["ttl"] = o.LeafServiceDNSForwardingAuthoritativeDomainRecordsPtrTTL.ValueString()
	}
	if !(o.LeafServiceDNSForwardingAuthoritativeDomainRecordsPtrDisable.IsNull() || o.LeafServiceDNSForwardingAuthoritativeDomainRecordsPtrDisable.IsUnknown()) {
		vyosData["disable"] = o.LeafServiceDNSForwardingAuthoritativeDomainRecordsPtrDisable.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsPtr) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"service", "dns", "forwarding", "authoritative-domain", "records", "ptr"}})

	// Leafs
	if value, ok := vyosData["target"]; ok {
		o.LeafServiceDNSForwardingAuthoritativeDomainRecordsPtrTarget = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDNSForwardingAuthoritativeDomainRecordsPtrTarget = basetypes.NewStringNull()
	}
	if value, ok := vyosData["ttl"]; ok {
		o.LeafServiceDNSForwardingAuthoritativeDomainRecordsPtrTTL = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDNSForwardingAuthoritativeDomainRecordsPtrTTL = basetypes.NewStringNull()
	}
	if value, ok := vyosData["disable"]; ok {
		o.LeafServiceDNSForwardingAuthoritativeDomainRecordsPtrDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDNSForwardingAuthoritativeDomainRecordsPtrDisable = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"service", "dns", "forwarding", "authoritative-domain", "records", "ptr"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ServiceDNSForwardingAuthoritativeDomainRecordsPtr) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"target":  types.StringType,
		"ttl":     types.StringType,
		"disable": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDNSForwardingAuthoritativeDomainRecordsPtr) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"target": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Target DNS name

|  Format  |  Description  |
|----------|---------------|
|  name.example.com  |  An absolute DNS name  |

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
