// Package namedservicednsforwardingauthoritativedomainrecordsaaaa code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicednsforwardingauthoritativedomainrecordsaaaa

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r serviceDNSForwardingAuthoritativeDomainRecordsAaaa) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	r.ResourceName = req.ProviderTypeName + "_service_dns_forwarding_authoritative_domain_records_aaaa"
	resp.TypeName = r.ResourceName
}
