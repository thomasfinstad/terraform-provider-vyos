// Package namedqospolicyshaperhfscclassmatch code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedqospolicyshaperhfscclassmatch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r qosPolicyShaperHfscClassMatch) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	r.ResourceName = req.ProviderTypeName + "_qos_policy_shaper_hfsc_class_match"
	resp.TypeName = r.ResourceName
}
