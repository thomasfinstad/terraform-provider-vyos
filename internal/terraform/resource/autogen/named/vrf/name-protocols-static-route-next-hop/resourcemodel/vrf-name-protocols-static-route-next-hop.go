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

var _ helpers.VyosTopResourceDataModel = &VrfNameProtocolsStaticRouteNextHop{}

// VrfNameProtocolsStaticRouteNextHop describes the resource data model.
type VrfNameProtocolsStaticRouteNextHop struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"next_hop_id" vyos:"-,self-id"`

	ParentIDVrfName types.String `tfsdk:"name_id" vyos:"name,parent-id"`

	ParentIDVrfNameProtocolsStaticRoute types.String `tfsdk:"route_id" vyos:"route,parent-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafVrfNameProtocolsStaticRouteNextHopDisable   types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafVrfNameProtocolsStaticRouteNextHopDistance  types.Number `tfsdk:"distance" vyos:"distance,omitempty"`
	LeafVrfNameProtocolsStaticRouteNextHopInterface types.String `tfsdk:"interface" vyos:"interface,omitempty"`
	LeafVrfNameProtocolsStaticRouteNextHopVrf       types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsStaticRouteNextHopBfd *VrfNameProtocolsStaticRouteNextHopBfd `tfsdk:"bfd" vyos:"bfd,omitempty"`
}

// SetID configures the resource ID
func (o *VrfNameProtocolsStaticRouteNextHop) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *VrfNameProtocolsStaticRouteNextHop) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *VrfNameProtocolsStaticRouteNextHop) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VrfNameProtocolsStaticRouteNextHop) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"next-hop",
		o.SelfIdentifier.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *VrfNameProtocolsStaticRouteNextHop) GetVyosParentPath() []string {
	return []string{
		"vrf",

		"name",
		o.ParentIDVrfName.ValueString(),

		"protocols",

		"static",

		"route",
		o.ParentIDVrfNameProtocolsStaticRoute.ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *VrfNameProtocolsStaticRouteNextHop) GetVyosNamedParentPath() []string {
	return []string{
		"vrf",

		"name",
		o.ParentIDVrfName.ValueString(),

		"protocols",

		"static",

		"route",
		o.ParentIDVrfNameProtocolsStaticRoute.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsStaticRouteNextHop) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"next_hop_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Next-hop IPv4 router address

    |  Format  |  Description              |
    |----------|---------------------------|
    |  ipv4    |  Next-hop router address  |
`,
			Description: `Next-hop IPv4 router address

    |  Format  |  Description              |
    |----------|---------------------------|
    |  ipv4    |  Next-hop router address  |
`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in next_hop_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illegal character in  next_hop_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
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

		"route_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Static IPv4 route

    |  Format   |  Description        |
    |-----------|---------------------|
    |  ipv4net  |  IPv4 static route  |
`,
			Description: `Static IPv4 route

    |  Format   |  Description        |
    |-----------|---------------------|
    |  ipv4net  |  IPv4 static route  |
`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in route_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illegal character in  route_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

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

		"distance": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Distance for this route

    |  Format  |  Description              |
    |----------|---------------------------|
    |  1-255   |  Distance for this route  |
`,
			Description: `Distance for this route

    |  Format  |  Description              |
    |----------|---------------------------|
    |  1-255   |  Distance for this route  |
`,
		},

		"interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Gateway interface name

    |  Format  |  Description             |
    |----------|--------------------------|
    |  txt     |  Gateway interface name  |
`,
			Description: `Gateway interface name

    |  Format  |  Description             |
    |----------|--------------------------|
    |  txt     |  Gateway interface name  |
`,
		},

		"vrf": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VRF to leak route

    |  Format  |  Description             |
    |----------|--------------------------|
    |  txt     |  Name of VRF to leak to  |
`,
			Description: `VRF to leak route

    |  Format  |  Description             |
    |----------|--------------------------|
    |  txt     |  Name of VRF to leak to  |
`,
		},

		// Nodes

		"bfd": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsStaticRouteNextHopBfd{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `BFD monitoring

`,
			Description: `BFD monitoring

`,
		},
	}
}
