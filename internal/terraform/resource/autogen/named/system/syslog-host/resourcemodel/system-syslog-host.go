// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// SystemSyslogHost describes the resource data model.
type SystemSyslogHost struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"host_id" vyos:",self-id"`

	// LeafNodes
	LeafSystemSyslogHostPort types.Number `tfsdk:"port" vyos:"port,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagSystemSyslogHostFacility bool `tfsdk:"-" vyos:"facility,child"`

	// Nodes
	NodeSystemSyslogHostFormat *SystemSyslogHostFormat `tfsdk:"format" vyos:"format,omitempty"`
}

// SetID configures the resource ID
func (o *SystemSyslogHost) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *SystemSyslogHost) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"system",

		"syslog",

		"host",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemSyslogHost) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"host_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Logging to a remote host

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  Remote syslog server IPv4 address  |
    |  hostname  &emsp; |  Remote syslog server FQDN  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"port": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Port number used by connection

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Numeric IP port  |

`,
		},

		// Nodes

		"format": schema.SingleNestedAttribute{
			Attributes: SystemSyslogHostFormat{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Logging format

`,
		},
	}
}
