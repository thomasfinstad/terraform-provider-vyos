// Package namedvrfnameprotocolsospfareavirtuallinkauthenticationmdfivekeyid code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsospfareavirtuallinkauthenticationmdfivekeyid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r vrfNameProtocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyID) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	r.ResourceName = req.ProviderTypeName + "_vrf_name_protocols_ospf_area_virtual_link_authentication_md5_key_id"
	resp.TypeName = r.ResourceName
}
