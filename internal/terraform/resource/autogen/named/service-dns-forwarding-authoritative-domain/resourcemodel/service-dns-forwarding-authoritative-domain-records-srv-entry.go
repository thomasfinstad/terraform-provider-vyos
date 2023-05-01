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

// ServiceDNSForwardingAuthoritativeDomainRecordsSrvEntry describes the resource data model.
type ServiceDNSForwardingAuthoritativeDomainRecordsSrvEntry struct {
	// LeafNodes
	LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvEntryHostname types.String `tfsdk:"hostname"`
	LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvEntryPort     types.String `tfsdk:"port"`
	LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvEntryPriority types.String `tfsdk:"priority"`
	LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvEntryWeight   types.String `tfsdk:"weight"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsSrvEntry) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"service", "dns", "forwarding", "authoritative-domain", "records", "srv", "entry"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvEntryHostname.IsNull() || o.LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvEntryHostname.IsUnknown()) {
		vyosData["hostname"] = o.LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvEntryHostname.ValueString()
	}
	if !(o.LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvEntryPort.IsNull() || o.LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvEntryPort.IsUnknown()) {
		vyosData["port"] = o.LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvEntryPort.ValueString()
	}
	if !(o.LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvEntryPriority.IsNull() || o.LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvEntryPriority.IsUnknown()) {
		vyosData["priority"] = o.LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvEntryPriority.ValueString()
	}
	if !(o.LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvEntryWeight.IsNull() || o.LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvEntryWeight.IsUnknown()) {
		vyosData["weight"] = o.LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvEntryWeight.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsSrvEntry) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"service", "dns", "forwarding", "authoritative-domain", "records", "srv", "entry"}})

	// Leafs
	if value, ok := vyosData["hostname"]; ok {
		o.LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvEntryHostname = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvEntryHostname = basetypes.NewStringNull()
	}
	if value, ok := vyosData["port"]; ok {
		o.LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvEntryPort = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvEntryPort = basetypes.NewStringNull()
	}
	if value, ok := vyosData["priority"]; ok {
		o.LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvEntryPriority = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvEntryPriority = basetypes.NewStringNull()
	}
	if value, ok := vyosData["weight"]; ok {
		o.LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvEntryWeight = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvEntryWeight = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"service", "dns", "forwarding", "authoritative-domain", "records", "srv", "entry"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ServiceDNSForwardingAuthoritativeDomainRecordsSrvEntry) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"hostname": types.StringType,
		"port":     types.StringType,
		"priority": types.StringType,
		"weight":   types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDNSForwardingAuthoritativeDomainRecordsSrvEntry) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"hostname": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Server hostname

|  Format  |  Description  |
|----------|---------------|
|  name.example.com  |  An absolute DNS name  |

`,
		},

		"port": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Port number

|  Format  |  Description  |
|----------|---------------|
|  u32:0-65535  |  TCP/UDP port number  |

`,
		},

		"priority": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Entry priority

|  Format  |  Description  |
|----------|---------------|
|  u32:0-65535  |  Entry priority (lower numbers are higher priority)  |

`,

			// Default:          stringdefault.StaticString(`10`),
			Computed: true,
		},

		"weight": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Entry weight

|  Format  |  Description  |
|----------|---------------|
|  u32:0-65535  |  Entry weight  |

`,

			// Default:          stringdefault.StaticString(`0`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
