// Package resourcemodel code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &VrfName{}

// VrfName describes the resource data model.
type VrfName struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"name_id" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafVrfNameDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafVrfNameDisable     types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafVrfNameTable       types.Number `tfsdk:"table" vyos:"table,omitempty"`
	LeafVrfNameVni         types.Number `tfsdk:"vni" vyos:"vni,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameIP        *VrfNameIP        `tfsdk:"ip" vyos:"ip,omitempty"`
	NodeVrfNameIPvsix    *VrfNameIPvsix    `tfsdk:"ipv6" vyos:"ipv6,omitempty"`
	NodeVrfNameProtocols *VrfNameProtocols `tfsdk:"protocols" vyos:"protocols,omitempty"`
}

// SetID configures the resource ID
func (o *VrfName) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *VrfName) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *VrfName) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VrfName) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"name",
		o.SelfIdentifier.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *VrfName) GetVyosParentPath() []string {
	return []string{
		"vrf",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *VrfName) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfName) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"name_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Virtual Routing and Forwarding instance

    |  Format  |  Description        |
    |----------|---------------------|
    |  txt     |  VRF instance name  |
`,
			Description: `Virtual Routing and Forwarding instance

    |  Format  |  Description        |
    |----------|---------------------|
    |  txt     |  VRF instance name  |
`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in name_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illegal character in  name_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Description  |
`,
			Description: `Description

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Description  |
`,
		},

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Administratively disable interface

`,
			Description: `Administratively disable interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"table": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Routing table associated with this instance

    |  Format     |  Description       |
    |-------------|--------------------|
    |  100-65535  |  Routing table ID  |
`,
			Description: `Routing table associated with this instance

    |  Format     |  Description       |
    |-------------|--------------------|
    |  100-65535  |  Routing table ID  |
`,
		},

		"vni": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Virtual Network Identifier

    |  Format      |  Description                       |
    |--------------|------------------------------------|
    |  0-16777214  |  VXLAN virtual network identifier  |
`,
			Description: `Virtual Network Identifier

    |  Format      |  Description                       |
    |--------------|------------------------------------|
    |  0-16777214  |  VXLAN virtual network identifier  |
`,
		},

		// Nodes

		"ip": schema.SingleNestedAttribute{
			Attributes: VrfNameIP{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
			Description: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: VrfNameIPvsix{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
			Description: `IPv6 routing parameters

`,
		},

		"protocols": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocols{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Routing protocol parameters

`,
			Description: `Routing protocol parameters

`,
		},
	}
}
