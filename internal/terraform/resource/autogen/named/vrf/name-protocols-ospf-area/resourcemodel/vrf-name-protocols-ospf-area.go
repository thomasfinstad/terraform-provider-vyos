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

// VrfNameProtocolsOspfArea describes the resource data model.
type VrfNameProtocolsOspfArea struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"area_id" vyos:"-,self-id"`

	ParentIDVrfName types.String `tfsdk:"name_id" vyos:"name,parent-id"`

	// LeafNodes
	LeafVrfNameProtocolsOspfAreaAuthentication types.String `tfsdk:"authentication" vyos:"authentication,omitempty"`
	LeafVrfNameProtocolsOspfAreaNetwork        types.List   `tfsdk:"network" vyos:"network,omitempty"`
	LeafVrfNameProtocolsOspfAreaShortcut       types.String `tfsdk:"shortcut" vyos:"shortcut,omitempty"`
	LeafVrfNameProtocolsOspfAreaExportList     types.Number `tfsdk:"export_list" vyos:"export-list,omitempty"`
	LeafVrfNameProtocolsOspfAreaImportList     types.Number `tfsdk:"import_list" vyos:"import-list,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagVrfNameProtocolsOspfAreaRange       bool `tfsdk:"-" vyos:"range,child"`
	ExistsTagVrfNameProtocolsOspfAreaVirtualLink bool `tfsdk:"-" vyos:"virtual-link,child"`

	// Nodes
	NodeVrfNameProtocolsOspfAreaAreaType *VrfNameProtocolsOspfAreaAreaType `tfsdk:"area_type" vyos:"area-type,omitempty"`
}

// SetID configures the resource ID
func (o *VrfNameProtocolsOspfArea) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VrfNameProtocolsOspfArea) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"vrf",

		"name",
		o.ParentIDVrfName.ValueString(),

		"protocols",

		"ospf",

		"area",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfArea) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"area_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `OSPF area settings

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  u32  &emsp; |  OSPF area number in decimal notation  |
    |  ipv4  &emsp; |  OSPF area number in dotted decimal notation  |

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

		"authentication": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `OSPF area authentication type

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  plaintext-password  &emsp; |  Use plain-text authentication  |
    |  md5  &emsp; |  Use MD5 authentication  |

`,
		},

		"network": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `OSPF network

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  OSPF network  |

`,
		},

		"shortcut": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Area shortcut mode

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  default  &emsp; |  Set default  |
    |  disable  &emsp; |  Disable shortcutting mode  |
    |  enable  &emsp; |  Enable shortcutting mode  |

`,
		},

		"export_list": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set the filter for networks announced to other areas

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  u32  &emsp; |  Access-list number  |

`,
		},

		"import_list": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set the filter for networks from other areas announced

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  u32  &emsp; |  Access-list number  |

`,
		},

		// Nodes

		"area_type": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfAreaAreaType{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Area type

`,
		},
	}
}
