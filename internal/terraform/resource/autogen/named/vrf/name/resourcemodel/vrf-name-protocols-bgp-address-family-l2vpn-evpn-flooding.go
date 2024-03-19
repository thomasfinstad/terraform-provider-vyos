// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnFlooding{}

// VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnFlooding describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnFlooding struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnFloodingDisable            types.Bool `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnFloodingHeadEndReplication types.Bool `tfsdk:"head_end_replication" vyos:"head-end-replication,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnFlooding) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
			Description: `Disable instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"head_end_replication": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Flood BUM packets using head-end replication

`,
			Description: `Flood BUM packets using head-end replication

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
