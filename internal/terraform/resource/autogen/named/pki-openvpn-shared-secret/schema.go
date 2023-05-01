// Package namedpkiopenvpnsharedsecret code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedpkiopenvpnsharedsecret

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r pkiOpenvpnSharedSecret) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `VyOS PKI configuration

OpenVPN keys

OpenVPN shared secret key

`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
