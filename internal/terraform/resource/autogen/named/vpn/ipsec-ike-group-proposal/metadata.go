// Package namedvpnipsecikegroupproposal code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnipsecikegroupproposal

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r vpnIPsecIkeGroupProposal) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	r.ResourceName = req.ProviderTypeName + "_vpn_ipsec_ike_group_proposal"
	resp.TypeName = r.ResourceName
}
