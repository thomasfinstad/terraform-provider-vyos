// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &VrfNameProtocolsOspfAreaRange{}

// VrfNameProtocolsOspfAreaRange describes the resource data model.
type VrfNameProtocolsOspfAreaRange struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"range_id" vyos:"-,self-id"`

	ParentIDVrfName types.String `tfsdk:"name_id" vyos:"name,parent-id"`

	ParentIDVrfNameProtocolsOspfArea types.String `tfsdk:"area_id" vyos:"area,parent-id"`

	// LeafNodes
	LeafVrfNameProtocolsOspfAreaRangeCost         types.Number `tfsdk:"cost" vyos:"cost,omitempty"`
	LeafVrfNameProtocolsOspfAreaRangeNotAdvertise types.Bool   `tfsdk:"not_advertise" vyos:"not-advertise,omitempty"`
	LeafVrfNameProtocolsOspfAreaRangeSubstitute   types.String `tfsdk:"substitute" vyos:"substitute,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// SetID configures the resource ID
func (o *VrfNameProtocolsOspfAreaRange) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *VrfNameProtocolsOspfAreaRange) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VrfNameProtocolsOspfAreaRange) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"range",
		o.SelfIdentifier.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *VrfNameProtocolsOspfAreaRange) GetVyosParentPath() []string {
	return []string{
		"vrf",

		"name",
		o.ParentIDVrfName.ValueString(),

		"protocols",

		"ospf",

		"area",
		o.ParentIDVrfNameProtocolsOspfArea.ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *VrfNameProtocolsOspfAreaRange) GetVyosNamedParentPath() []string {
	return []string{
		"vrf",

		"name",
		o.ParentIDVrfName.ValueString(),

		"protocols",

		"ospf",

		"area",
		o.ParentIDVrfNameProtocolsOspfArea.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfAreaRange) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"range_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Summarize routes matching a prefix (border routers only)

    |  Format   &emsp;|  Description        |
    |-----------------|---------------------|
    |  ipv4net  &emsp;|  Area range prefix  |
`,
			Description: `Summarize routes matching a prefix (border routers only)

    |  Format   |  Description        |
    |-----------------|---------------------|
    |  ipv4net  |  Area range prefix  |
`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in range_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illigal character in  range_id, value must match: ^[a-zA-Z0-9-_]*$",
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

		"area_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `OSPF area settings

    |  Format  &emsp;|  Description                                  |
    |----------------|-----------------------------------------------|
    |  u32     &emsp;|  OSPF area number in decimal notation         |
    |  ipv4    &emsp;|  OSPF area number in dotted decimal notation  |
`,
			Description: `OSPF area settings

    |  Format  |  Description                                  |
    |----------------|-----------------------------------------------|
    |  u32     |  OSPF area number in decimal notation         |
    |  ipv4    |  OSPF area number in dotted decimal notation  |
`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in area_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illigal character in  area_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		// LeafNodes

		"cost": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Metric for this range

    |  Format      &emsp;|  Description            |
    |--------------------|-------------------------|
    |  0-16777215  &emsp;|  Metric for this range  |
`,
			Description: `Metric for this range

    |  Format      |  Description            |
    |--------------------|-------------------------|
    |  0-16777215  |  Metric for this range  |
`,
		},

		"not_advertise": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not advertise this range

`,
			Description: `Do not advertise this range

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"substitute": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Advertise area range as another prefix

    |  Format   &emsp;|  Description                             |
    |-----------------|------------------------------------------|
    |  ipv4net  &emsp;|  Advertise area range as another prefix  |
`,
			Description: `Advertise area range as another prefix

    |  Format   |  Description                             |
    |-----------------|------------------------------------------|
    |  ipv4net  |  Advertise area range as another prefix  |
`,
		},

		// Nodes

	}
}
