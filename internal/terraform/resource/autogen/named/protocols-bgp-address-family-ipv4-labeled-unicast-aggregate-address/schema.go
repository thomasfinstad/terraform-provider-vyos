// Package namedprotocolsbgpaddressfamilyipvfourlabeledunicastaggregateaddress code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsbgpaddressfamilyipvfourlabeledunicastaggregateaddress

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r protocolsBgpAddressFamilyIPvfourLabeledUnicastAggregateAddress) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Border Gateway Protocol (BGP)

BGP address-family parameters

Labeled Unicast IPv4 BGP settings

BGP aggregate network/prefix

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  BGP aggregate network/prefix  |

`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
