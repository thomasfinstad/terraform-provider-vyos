package resourcefull

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &interfaces_dummy{}

// var _ resource.ResourceWithImportState = &interfaces_dummy{}

// interfaces_dummy defines the resource implementation.
type interfaces_dummy struct {
	client *http.Client
}

// interfaces dummyModel describes the resource data model.
type interfaces_dummyModel struct {
	ConfigurableAttribute types.String `tfsdk:"configurable_attribute"`
	Defaulted             types.String `tfsdk:"defaulted"`
	ID                    types.String `tfsdk:"id"`

	Ip types.String `tfsdk:ip`

	Ipv_ types.String `tfsdk:ipv6`

	Mirror types.String `tfsdk:mirror`

	Address types.String `tfsdk:address`

	Description types.String `tfsdk:description`

	Disable types.String `tfsdk:disable`

	Mtu types.String `tfsdk:mtu`

	Netns types.String `tfsdk:netns`

	Redirect types.String `tfsdk:redirect`

	Vrf types.String `tfsdk:vrf`
}

// Metadata method to define the resource type name.
func (r *interfaces_dummy) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interfaces dummy"
}

// interfaces_dummyResource method to return the example resource reference
func interfaces_dummyResource() resource.Resource {
	return &interfaces_dummy{}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r *interfaces_dummy) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example resource",

		Attributes: map[string]schema.Attribute{
			// TODO handle non-string types
			"address": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `IP address

Format: ipv4net
IPv4 address and prefix length
Format: ipv6net
IPv6 address and prefix length
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"description": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Description

Format: txt
Description
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"disable": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Administratively disable interface

`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"mtu": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Maximum Transmission Unit (MTU)

Format: u32:68-16000
Maximum Transmission Unit in byte
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

				Default:  stringdefault.StaticString(`1500`),
				Computed: true,
			},

			// TODO handle non-string types
			"netns": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Network namespace name

Format: text
Network namespace name
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"redirect": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Redirect incoming packet to destination

Format: txt
Destination interface name
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"vrf": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `VRF instance name

Format: txt
VRF instance name
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			"ip": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					// TODO handle non-string types
					"source_validation": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Source validation by reversed path (RFC3704)

Format: strict
Enable Strict Reverse Path Forwarding as defined in RFC3704
Format: loose
Enable Loose Reverse Path Forwarding as defined in RFC3704
Format: disable
No source validation
`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},

					// TODO handle non-string types
					"disable_forwarding": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Disable IP forwarding on this interface

`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},
				},
				// CustomType:          basetypes.MapTypable(nil),
				// Required:            false,
				Optional: true,
				// Computed:            false,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `IPv4 routing parameters

`,
				// DeprecationMessage:  "",
				// Validators:          []validator.Map(nil),
				// PlanModifiers:       []planmodifier.Map(nil),
				// TODO investigate if node defaults can be handled
				// Default:             defaults.Map(nil),
			},

			"ipv6": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					// TODO handle non-string types
					"disable_forwarding": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Disable IP forwarding on this interface

`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},

					"address": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							// TODO handle non-string types
							"eui64": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Prefix for IPv6 address with MAC-based EUI-64

Format: <h:h:h:h:h:h:h:h/64>
IPv6 /64 network
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"no_default_link_local": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Remove the default link-local address from the interface

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},
						},
						// CustomType:          basetypes.MapTypable(nil),
						// Required:            false,
						Optional: true,
						// Computed:            false,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `IPv6 address configuration modes

`,
						// DeprecationMessage:  "",
						// Validators:          []validator.Map(nil),
						// PlanModifiers:       []planmodifier.Map(nil),
						// TODO investigate if node defaults can be handled
						// Default:             defaults.Map(nil),
					},
				},
				// CustomType:          basetypes.MapTypable(nil),
				// Required:            false,
				Optional: true,
				// Computed:            false,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `IPv6 routing parameters

`,
				// DeprecationMessage:  "",
				// Validators:          []validator.Map(nil),
				// PlanModifiers:       []planmodifier.Map(nil),
				// TODO investigate if node defaults can be handled
				// Default:             defaults.Map(nil),
			},

			"mirror": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					// TODO handle non-string types
					"ingress": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Mirror ingress traffic to destination interface

Format: txt
Destination interface name
`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},

					// TODO handle non-string types
					"egress": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Mirror egress traffic to destination interface

Format: txt
Destination interface name
`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},
				},
				// CustomType:          basetypes.MapTypable(nil),
				// Required:            false,
				Optional: true,
				// Computed:            false,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Mirror ingress/egress packets

`,
				// DeprecationMessage:  "",
				// Validators:          []validator.Map(nil),
				// PlanModifiers:       []planmodifier.Map(nil),
				// TODO investigate if node defaults can be handled
				// Default:             defaults.Map(nil),
			},
		},
	}
}

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func (r *interfaces_dummy) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *interfaces_dummyModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create example, got error: %s", err))
	//     return
	// }

	// For the purposes of this example code, hardcoding a response value to
	// save into the Terraform state.
	data.ID = types.StringValue("example-id")

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "created a resource")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read method to define the logic which refreshes the Terraform state for the resource.
func (r *interfaces_dummy) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *interfaces_dummyModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read example, got error: %s", err))
	//     return
	// }

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update method to define the logic which updates the resource and sets the updated Terraform state on success.
func (r *interfaces_dummy) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *interfaces_dummyModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update example, got error: %s", err))
	//     return
	// }

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete method to define the logic which deletes the resource and removes the Terraform state on success.
func (r *interfaces_dummy) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *interfaces_dummyModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete example, got error: %s", err))
	//     return
	// }
}
