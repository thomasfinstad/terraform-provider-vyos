// Package namedservicednsforwardingauthoritativedomainrecordsa code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicednsforwardingauthoritativedomainrecordsa

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r serviceDNSForwardingAuthoritativeDomainRecordsA) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	r.ResourceName = req.ProviderTypeName + "_service_dns_forwarding_authoritative_domain_records_a"
	resp.TypeName = r.ResourceName
}
