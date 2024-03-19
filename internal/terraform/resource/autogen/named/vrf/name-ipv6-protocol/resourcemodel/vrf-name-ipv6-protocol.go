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

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &VrfNameIPvsixProtocol{}

// VrfNameIPvsixProtocol describes the resource data model.
type VrfNameIPvsixProtocol struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"protocol_id" vyos:"-,self-id"`

	ParentIDVrfName types.String `tfsdk:"name_id" vyos:"name,parent-id"`

	// LeafNodes
	LeafVrfNameIPvsixProtocolRouteMap types.String `tfsdk:"route_map" vyos:"route-map,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// SetID configures the resource ID
func (o *VrfNameIPvsixProtocol) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *VrfNameIPvsixProtocol) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VrfNameIPvsixProtocol) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"protocol",
		o.SelfIdentifier.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *VrfNameIPvsixProtocol) GetVyosParentPath() []string {
	return []string{
		"vrf",

		"name",
		o.ParentIDVrfName.ValueString(),

		"ipv6",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *VrfNameIPvsixProtocol) GetVyosNamedParentPath() []string {
	return []string{
		"vrf",

		"name",
		o.ParentIDVrfName.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameIPvsixProtocol) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"protocol_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Filter routing info exchanged between routing protocol and zebra

    |  Format     &emsp;|  Description                                          |
    |-------------------|-------------------------------------------------------|
    |  any        &emsp;|  Any of the above protocols                           |
    |  babel      &emsp;|  Babel routing protocol                               |
    |  bgp        &emsp;|  Border Gateway Protocol                              |
    |  connected  &emsp;|  Connected routes (directly attached subnet or host)  |
    |  isis       &emsp;|  Intermediate System to Intermediate System           |
    |  kernel     &emsp;|  Kernel routes (not installed via the zebra RIB)      |
    |  ospfv3     &emsp;|  Open Shortest Path First (OSPFv3)                    |
    |  ripng      &emsp;|  Routing Information Protocol next-generation         |
    |  static     &emsp;|  Statically configured routes                         |
`,
			Description: `Filter routing info exchanged between routing protocol and zebra

    |  Format     |  Description                                          |
    |-------------------|-------------------------------------------------------|
    |  any        |  Any of the above protocols                           |
    |  babel      |  Babel routing protocol                               |
    |  bgp        |  Border Gateway Protocol                              |
    |  connected  |  Connected routes (directly attached subnet or host)  |
    |  isis       |  Intermediate System to Intermediate System           |
    |  kernel     |  Kernel routes (not installed via the zebra RIB)      |
    |  ospfv3     |  Open Shortest Path First (OSPFv3)                    |
    |  ripng      |  Routing Information Protocol next-generation         |
    |  static     |  Statically configured routes                         |
`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in protocol_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illigal character in  protocol_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		"name_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Virtual Routing and Forwarding instance

    |  Format  &emsp;|  Description        |
    |----------------|---------------------|
    |  txt     &emsp;|  VRF instance name  |
`,
			Description: `Virtual Routing and Forwarding instance

    |  Format  |  Description        |
    |----------------|---------------------|
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
						"illigal character in  name_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		// LeafNodes

		"route_map": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------------|------------------|
    |  txt     &emsp;|  Route map name  |
`,
			Description: `Specify route-map name to use

    |  Format  |  Description     |
    |----------------|------------------|
    |  txt     |  Route map name  |
`,
		},

		// Nodes

	}
}
