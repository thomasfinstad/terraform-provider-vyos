// Package namedservicewebproxycachepeer code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicewebproxycachepeer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r serviceWebproxyCachePeer) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Webproxy service settings

Specify other caches in a hierarchy

    |  Format  |  Description  |
    |----------|---------------|
    |  hostname  |  Cache peers FQDN  |

`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
