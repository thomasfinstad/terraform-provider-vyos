// Package namedsystemsyslogglobalfacility code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedsystemsyslogglobalfacility

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r systemSyslogGlobalFacility) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	r.ResourceName = req.ProviderTypeName + "_system_syslog_global_facility"
	resp.TypeName = r.ResourceName
}
