// Package namedservicepppoeserverauthenticationlocalusersusername code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicepppoeserverauthenticationlocalusersusername

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r servicePppoeServerAuthenticationLocalUsersUsername) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Point to Point over Ethernet (PPPoE) Server

Authentication for remote access PPPoE Server

Local user authentication for PPPoE server

User name for authentication

`,
		Attributes: r.model.ResourceAttributes(),
	}
}
