// Package namedinterfacespseudoethernetvifsvifc code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacespseudoethernetvifsvifc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r interfacesPseudoEthernetVifSVifC) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	r.ResourceName = req.ProviderTypeName + "_interfaces_pseudo_ethernet_vif_s_vif_c"
	resp.TypeName = r.ResourceName
}
