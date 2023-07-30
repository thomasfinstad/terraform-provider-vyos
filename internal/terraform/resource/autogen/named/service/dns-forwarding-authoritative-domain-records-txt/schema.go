// Package namedservicednsforwardingauthoritativedomainrecordstxt code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicednsforwardingauthoritativedomainrecordstxt

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r serviceDNSForwardingAuthoritativeDomainRecordsTxt) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Domain Name System related services

DNS forwarding

Domain to host authoritative records for

    |  Format  |  Description  |
    |----------|---------------|
    |  text  |  An absolute DNS name  |

DNS zone records

"TXT" record

    |  Format  |  Description  |
    |----------|---------------|
    |  text  |  A DNS name relative to the root record  |
    |  @  |  Root record  |

`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
