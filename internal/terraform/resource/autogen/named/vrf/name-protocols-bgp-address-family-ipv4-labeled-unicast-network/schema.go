// Package namedvrfnameprotocolsbgpaddressfamilyipvfourlabeledunicastnetwork code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsbgpaddressfamilyipvfourlabeledunicastnetwork

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r vrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastNetwork) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Virtual Routing and Forwarding

Virtual Routing and Forwarding instance

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  VRF instance name  |

Routing protocol parameters

Border Gateway Protocol (BGP)

BGP address-family parameters

Labeled Unicast IPv4 BGP settings

Import BGP network/prefix into labeled unicast IPv4 RIB

    |  Format  |  Description  |
    |----------|---------------|
    |  ipv4net  |  Labeled Unicast IPv4 BGP network/prefix  |

`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
