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

var _ helpers.VyosTopResourceDataModel = &VrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistancePrefix{}

// VrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistancePrefix describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistancePrefix struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"prefix_id" vyos:"-,self-id"`

	ParentIDVrfName types.String `tfsdk:"name_id" vyos:"name,parent-id"`

	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistancePrefixDistance types.Number `tfsdk:"distance" vyos:"distance,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// SetID configures the resource ID
func (o *VrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistancePrefix) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *VrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistancePrefix) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistancePrefix) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"prefix",
		o.SelfIdentifier.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *VrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistancePrefix) GetVyosParentPath() []string {
	return []string{
		"vrf",

		"name",
		o.ParentIDVrfName.ValueString(),

		"protocols",

		"bgp",

		"address-family",

		"ipv4-multicast",

		"distance",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *VrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistancePrefix) GetVyosNamedParentPath() []string {
	return []string{
		"vrf",

		"name",
		o.ParentIDVrfName.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistancePrefix) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"prefix_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Administrative distance for a specific BGP prefix

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  Administrative distance for a specific BGP prefix  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in prefix_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illigal character in  prefix_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		"name_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Virtual Routing and Forwarding instance

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  VRF instance name  |

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

		"distance": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Administrative distance for prefix

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-255  &emsp; |  Administrative distance for external BGP routes  |

`,
		},

		// Nodes

	}
}
