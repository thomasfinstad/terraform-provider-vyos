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

// ServiceDNSForwardingDomain describes the resource data model.
type ServiceDNSForwardingDomain struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafServiceDNSForwardingDomainServer           types.String `tfsdk:"server"`
	LeafServiceDNSForwardingDomainAddnta           types.String `tfsdk:"addnta"`
	LeafServiceDNSForwardingDomainRecursionDesired types.String `tfsdk:"recursion_desired"`

	// TagNodes

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceDNSForwardingDomain) GetVyosPath() []string {
	return []string{
		"service",
		"dns",
		"forwarding",
		"domain",
		o.ID.ValueString(),
	}
}

// TerraformToVyos converts terraform data to vyos data
func (o *ServiceDNSForwardingDomain) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"service", "dns", "forwarding", "domain"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafServiceDNSForwardingDomainServer.IsNull() || o.LeafServiceDNSForwardingDomainServer.IsUnknown()) {
		vyosData["server"] = o.LeafServiceDNSForwardingDomainServer.ValueString()
	}
	if !(o.LeafServiceDNSForwardingDomainAddnta.IsNull() || o.LeafServiceDNSForwardingDomainAddnta.IsUnknown()) {
		vyosData["addnta"] = o.LeafServiceDNSForwardingDomainAddnta.ValueString()
	}
	if !(o.LeafServiceDNSForwardingDomainRecursionDesired.IsNull() || o.LeafServiceDNSForwardingDomainRecursionDesired.IsUnknown()) {
		vyosData["recursion-desired"] = o.LeafServiceDNSForwardingDomainRecursionDesired.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ServiceDNSForwardingDomain) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"service", "dns", "forwarding", "domain"}})

	// Leafs
	if value, ok := vyosData["server"]; ok {
		o.LeafServiceDNSForwardingDomainServer = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDNSForwardingDomainServer = basetypes.NewStringNull()
	}
	if value, ok := vyosData["addnta"]; ok {
		o.LeafServiceDNSForwardingDomainAddnta = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDNSForwardingDomainAddnta = basetypes.NewStringNull()
	}
	if value, ok := vyosData["recursion-desired"]; ok {
		o.LeafServiceDNSForwardingDomainRecursionDesired = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDNSForwardingDomainRecursionDesired = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"service", "dns", "forwarding", "domain"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ServiceDNSForwardingDomain) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"server":            types.StringType,
		"addnta":            types.StringType,
		"recursion_desired": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDNSForwardingDomain) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Domain to forward to a custom DNS server

`,
		},

		// LeafNodes

		"server": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Domain Name Server (DNS) to forward queries to

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Domain Name Server (DNS) IPv4 address  |
|  ipv6  |  Domain Name Server (DNS) IPv6 address  |

`,
		},

		"addnta": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Add NTA (negative trust anchor) for this domain (must be set if the domain does not support DNSSEC)

`,
		},

		"recursion_desired": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set the "recursion desired" bit in requests to the upstream nameserver

`,
		},

		// TagNodes

		// Nodes

	}
}
