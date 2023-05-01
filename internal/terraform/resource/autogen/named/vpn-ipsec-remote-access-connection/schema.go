// Package namedvpnipsecremoteaccessconnection code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnipsecremoteaccessconnection

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r vpnIPsecRemoteAccessConnection) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Virtual Private Network (VPN)

VPN IP security (IPsec) parameters

IKEv2 remote access VPN

IKEv2 VPN connection name

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Connection name  |

`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
