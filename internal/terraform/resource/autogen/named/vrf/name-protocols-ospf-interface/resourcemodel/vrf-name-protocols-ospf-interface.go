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

var _ helpers.VyosTopResourceDataModel = &VrfNameProtocolsOspfInterface{}

// VrfNameProtocolsOspfInterface describes the resource data model.
type VrfNameProtocolsOspfInterface struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"interface_id" vyos:"-,self-id"`

	ParentIDVrfName types.String `tfsdk:"name_id" vyos:"name,parent-id"`

	// LeafNodes
	LeafVrfNameProtocolsOspfInterfaceArea               types.String `tfsdk:"area" vyos:"area,omitempty"`
	LeafVrfNameProtocolsOspfInterfaceDeadInterval       types.Number `tfsdk:"dead_interval" vyos:"dead-interval,omitempty"`
	LeafVrfNameProtocolsOspfInterfaceHelloInterval      types.Number `tfsdk:"hello_interval" vyos:"hello-interval,omitempty"`
	LeafVrfNameProtocolsOspfInterfaceRetransmitInterval types.Number `tfsdk:"retransmit_interval" vyos:"retransmit-interval,omitempty"`
	LeafVrfNameProtocolsOspfInterfaceTransmitDelay      types.Number `tfsdk:"transmit_delay" vyos:"transmit-delay,omitempty"`
	LeafVrfNameProtocolsOspfInterfaceCost               types.Number `tfsdk:"cost" vyos:"cost,omitempty"`
	LeafVrfNameProtocolsOspfInterfaceMtuIgnore          types.Bool   `tfsdk:"mtu_ignore" vyos:"mtu-ignore,omitempty"`
	LeafVrfNameProtocolsOspfInterfacePriority           types.Number `tfsdk:"priority" vyos:"priority,omitempty"`
	LeafVrfNameProtocolsOspfInterfaceBandwIDth          types.Number `tfsdk:"bandwidth" vyos:"bandwidth,omitempty"`
	LeafVrfNameProtocolsOspfInterfaceHelloMultIPlier    types.Number `tfsdk:"hello_multiplier" vyos:"hello-multiplier,omitempty"`
	LeafVrfNameProtocolsOspfInterfaceNetwork            types.String `tfsdk:"network" vyos:"network,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsOspfInterfaceAuthentication *VrfNameProtocolsOspfInterfaceAuthentication `tfsdk:"authentication" vyos:"authentication,omitempty"`
	NodeVrfNameProtocolsOspfInterfaceBfd            *VrfNameProtocolsOspfInterfaceBfd            `tfsdk:"bfd" vyos:"bfd,omitempty"`
	NodeVrfNameProtocolsOspfInterfaceLdpSync        *VrfNameProtocolsOspfInterfaceLdpSync        `tfsdk:"ldp_sync" vyos:"ldp-sync,omitempty"`
	NodeVrfNameProtocolsOspfInterfacePassive        *VrfNameProtocolsOspfInterfacePassive        `tfsdk:"passive" vyos:"passive,omitempty"`
}

// SetID configures the resource ID
func (o *VrfNameProtocolsOspfInterface) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *VrfNameProtocolsOspfInterface) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VrfNameProtocolsOspfInterface) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"interface",
		o.SelfIdentifier.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *VrfNameProtocolsOspfInterface) GetVyosParentPath() []string {
	return []string{
		"vrf",

		"name",
		o.ParentIDVrfName.ValueString(),

		"protocols",

		"ospf",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *VrfNameProtocolsOspfInterface) GetVyosNamedParentPath() []string {
	return []string{
		"vrf",

		"name",
		o.ParentIDVrfName.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"interface_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Interface configuration

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Interface name  |
`,
			Description: `Interface configuration

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Interface name  |
`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in interface_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illigal character in  interface_id, value must match: ^[a-zA-Z0-9-_]*$",
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
						"illigal character in  name_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		// LeafNodes

		"area": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enable OSPF on this interface

    |  Format  |  Description                          |
    |----------|---------------------------------------|
    |  u32     |  OSPF area ID as decimal notation     |
    |  ipv4    |  OSPF area ID in IP address notation  |
`,
			Description: `Enable OSPF on this interface

    |  Format  |  Description                          |
    |----------|---------------------------------------|
    |  u32     |  OSPF area ID as decimal notation     |
    |  ipv4    |  OSPF area ID in IP address notation  |
`,
		},

		"dead_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interval after which a neighbor is declared dead

    |  Format   |  Description                       |
    |-----------|------------------------------------|
    |  1-65535  |  Neighbor dead interval (seconds)  |
`,
			Description: `Interval after which a neighbor is declared dead

    |  Format   |  Description                       |
    |-----------|------------------------------------|
    |  1-65535  |  Neighbor dead interval (seconds)  |
`,

			// Default:          stringdefault.StaticString(`40`),
			Computed: true,
		},

		"hello_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interval between hello packets

    |  Format   |  Description               |
    |-----------|----------------------------|
    |  1-65535  |  Hello interval (seconds)  |
`,
			Description: `Interval between hello packets

    |  Format   |  Description               |
    |-----------|----------------------------|
    |  1-65535  |  Hello interval (seconds)  |
`,

			// Default:          stringdefault.StaticString(`10`),
			Computed: true,
		},

		"retransmit_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interval between retransmitting lost link state advertisements

    |  Format   |  Description                    |
    |-----------|---------------------------------|
    |  1-65535  |  Retransmit interval (seconds)  |
`,
			Description: `Interval between retransmitting lost link state advertisements

    |  Format   |  Description                    |
    |-----------|---------------------------------|
    |  1-65535  |  Retransmit interval (seconds)  |
`,

			// Default:          stringdefault.StaticString(`5`),
			Computed: true,
		},

		"transmit_delay": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Link state transmit delay

    |  Format   |  Description                          |
    |-----------|---------------------------------------|
    |  1-65535  |  Link state transmit delay (seconds)  |
`,
			Description: `Link state transmit delay

    |  Format   |  Description                          |
    |-----------|---------------------------------------|
    |  1-65535  |  Link state transmit delay (seconds)  |
`,

			// Default:          stringdefault.StaticString(`1`),
			Computed: true,
		},

		"cost": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interface cost

    |  Format   |  Description          |
    |-----------|-----------------------|
    |  1-65535  |  OSPF interface cost  |
`,
			Description: `Interface cost

    |  Format   |  Description          |
    |-----------|-----------------------|
    |  1-65535  |  OSPF interface cost  |
`,
		},

		"mtu_ignore": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable Maximum Transmission Unit (MTU) mismatch detection

`,
			Description: `Disable Maximum Transmission Unit (MTU) mismatch detection

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"priority": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Router priority

    |  Format  |  Description                |
    |----------|-----------------------------|
    |  0-255   |  OSPF router priority cost  |
`,
			Description: `Router priority

    |  Format  |  Description                |
    |----------|-----------------------------|
    |  0-255   |  OSPF router priority cost  |
`,

			// Default:          stringdefault.StaticString(`1`),
			Computed: true,
		},

		"bandwidth": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interface bandwidth (Mbit/s)

    |  Format    |  Description                                           |
    |------------|--------------------------------------------------------|
    |  1-100000  |  Bandwidth in Megabit/sec (for calculating OSPF cost)  |
`,
			Description: `Interface bandwidth (Mbit/s)

    |  Format    |  Description                                           |
    |------------|--------------------------------------------------------|
    |  1-100000  |  Bandwidth in Megabit/sec (for calculating OSPF cost)  |
`,
		},

		"hello_multiplier": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Hello multiplier factor

    |  Format  |  Description                           |
    |----------|----------------------------------------|
    |  1-10    |  Number of Hellos to send each second  |
`,
			Description: `Hello multiplier factor

    |  Format  |  Description                           |
    |----------|----------------------------------------|
    |  1-10    |  Number of Hellos to send each second  |
`,
		},

		"network": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Network type

    |  Format               |  Description                       |
    |-----------------------|------------------------------------|
    |  broadcast            |  Broadcast network type            |
    |  non-broadcast        |  Non-broadcast network type        |
    |  point-to-multipoint  |  Point-to-multipoint network type  |
    |  point-to-point       |  Point-to-point network type       |
`,
			Description: `Network type

    |  Format               |  Description                       |
    |-----------------------|------------------------------------|
    |  broadcast            |  Broadcast network type            |
    |  non-broadcast        |  Non-broadcast network type        |
    |  point-to-multipoint  |  Point-to-multipoint network type  |
    |  point-to-point       |  Point-to-point network type       |
`,
		},

		// Nodes

		"authentication": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfInterfaceAuthentication{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Authentication

`,
			Description: `Authentication

`,
		},

		"bfd": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfInterfaceBfd{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Enable Bidirectional Forwarding Detection (BFD)

`,
			Description: `Enable Bidirectional Forwarding Detection (BFD)

`,
		},

		"ldp_sync": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfInterfaceLdpSync{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `LDP-IGP synchronization configuration for interface

`,
			Description: `LDP-IGP synchronization configuration for interface

`,
		},

		"passive": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfInterfacePassive{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Suppress routing updates on an interface

`,
			Description: `Suppress routing updates on an interface

`,
		},
	}
}
