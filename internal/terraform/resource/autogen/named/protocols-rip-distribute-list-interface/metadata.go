// Package namedprotocolsripdistributelistinterface code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsripdistributelistinterface

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r protocolsRIPDistributeListInterface) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	r.ResourceName = req.ProviderTypeName + "_protocols_rip_distribute_list_interface"
	resp.TypeName = r.ResourceName
}
