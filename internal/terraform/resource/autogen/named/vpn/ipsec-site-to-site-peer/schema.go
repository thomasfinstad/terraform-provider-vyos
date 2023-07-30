// Package namedvpnipsecsitetositepeer code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnipsecsitetositepeer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r vpnIPsecSiteToSitePeer) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Virtual Private Network (VPN)

VPN IP security (IPsec) parameters

Site-to-site VPN

Connection name of the peer

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  Connection name of the peer  |

`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
