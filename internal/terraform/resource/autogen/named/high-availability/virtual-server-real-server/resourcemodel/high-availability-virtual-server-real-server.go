// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// HighAvailabilityVirtualServerRealServer describes the resource data model.
type HighAvailabilityVirtualServerRealServer struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"real_server_id" vyos:"-,self-id"`

	ParentIDHighAvailabilityVirtualServer types.String `tfsdk:"virtual_server_id" vyos:"virtual-server,parent-id"`

	// LeafNodes
	LeafHighAvailabilityVirtualServerRealServerPort              types.Number `tfsdk:"port" vyos:"port,omitempty"`
	LeafHighAvailabilityVirtualServerRealServerConnectionTimeout types.Number `tfsdk:"connection_timeout" vyos:"connection-timeout,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeHighAvailabilityVirtualServerRealServerHealthCheck *HighAvailabilityVirtualServerRealServerHealthCheck `tfsdk:"health_check" vyos:"health-check,omitempty"`
}

// SetID configures the resource ID
func (o *HighAvailabilityVirtualServerRealServer) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *HighAvailabilityVirtualServerRealServer) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"high-availability",

		"virtual-server",
		o.ParentIDHighAvailabilityVirtualServer.ValueString(),

		"real-server",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o HighAvailabilityVirtualServerRealServer) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"real_server_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Real server address

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in real_server_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illigal character in  real_server_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		"virtual_server_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Load-balancing virtual server alias

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in virtual_server_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illigal character in  virtual_server_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		// LeafNodes

		"port": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Port number used by connection

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-65535  &emsp; |  Numeric IP port  |

`,
		},

		"connection_timeout": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Server connection timeout

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-86400  &emsp; |  Connection timeout to remote server  |

`,
		},

		// Nodes

		"health_check": schema.SingleNestedAttribute{
			Attributes: HighAvailabilityVirtualServerRealServerHealthCheck{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Health check script

`,
		},
	}
}
