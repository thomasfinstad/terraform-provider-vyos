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
var _ resource.Resource = &protocols_ospf_area{}

// var _ resource.ResourceWithImportState = &protocols_ospf_area{}

// protocols_ospf_area defines the resource implementation.
type protocols_ospf_area struct {
	client *http.Client
}

// protocols ospf areaModel describes the resource data model.
type protocols_ospf_areaModel struct {
	ConfigurableAttribute types.String `tfsdk:"configurable_attribute"`
	Defaulted             types.String `tfsdk:"defaulted"`
	ID                    types.String `tfsdk:"id"`

	Range types.Map `tfsdk:range`

	Virtual_link types.Map `tfsdk:virtual_link`

	Area_type types.String `tfsdk:area_type`

	Authentication types.String `tfsdk:authentication`

	Network types.String `tfsdk:network`

	Shortcut types.String `tfsdk:shortcut`

	Export_list types.String `tfsdk:export_list`

	Import_list types.String `tfsdk:import_list`
}

// Metadata method to define the resource type name.
func (r *protocols_ospf_area) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protocols ospf area"
}

// protocols_ospf_areaResource method to return the example resource reference
func protocols_ospf_areaResource() resource.Resource {
	return &protocols_ospf_area{}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r *protocols_ospf_area) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example resource",

		Attributes: map[string]schema.Attribute{
			"range": schema.MapNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						// TODO handle non-string types
						"cost": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Metric for this range

Format: u32:0-16777215
Metric for this range
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"not_advertise": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Do not advertise this range

`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"substitute": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Advertise area range as another prefix

Format: ipv4net
Advertise area range as another prefix
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// CustomType:    basetypes.ObjectTypable(nil),
						// Validators:    []validator.Object(nil),
						// PlanModifiers: []planmodifier.Object(nil),
					},
				},
				// CustomType:          basetypes.MapTypable(nil),
				// Required:            false,
				Optional: true,
				// Computed:            false,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Summarize routes matching a prefix (border routers only)

Format: ipv4net
Area range prefix
`,
				// DeprecationMessage:  "",
				// Validators:          []validator.Map(nil),
				// PlanModifiers:       []planmodifier.Map(nil),
				// TODO investigate if tagnode defaults can be handled
				// Default:             defaults.Map(nil),
			},

			"virtual_link": schema.MapNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						// TODO handle non-string types
						"dead_interval": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Interval after which a neighbor is declared dead

Format: u32:1-65535
Neighbor dead interval (seconds)
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

							Default:  stringdefault.StaticString(`40`),
							Computed: true,
						},

						// TODO handle non-string types
						"hello_interval": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Interval between hello packets

Format: u32:1-65535
Hello interval (seconds)
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

							Default:  stringdefault.StaticString(`10`),
							Computed: true,
						},

						// TODO handle non-string types
						"retransmit_interval": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Interval between retransmitting lost link state advertisements

Format: u32:1-65535
Retransmit interval (seconds)
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

							Default:  stringdefault.StaticString(`5`),
							Computed: true,
						},

						// TODO handle non-string types
						"transmit_delay": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Link state transmit delay

Format: u32:1-65535
Link state transmit delay (seconds)
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

							Default:  stringdefault.StaticString(`1`),
							Computed: true,
						},

						"authentication": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								// TODO handle non-string types
								"plaintext_password": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Plain text password

Format: txt
Plain text password (8 characters or less)
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								"md5": schema.SingleNestedAttribute{
									Attributes: map[string]schema.Attribute{
										"key_id": schema.MapNestedAttribute{
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													// TODO handle non-string types
													"md5_key": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `MD5 authentication type

Format: txt
MD5 Key (16 characters or less)
`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// CustomType:    basetypes.ObjectTypable(nil),
													// Validators:    []validator.Object(nil),
													// PlanModifiers: []planmodifier.Object(nil),
												},
											},
											// CustomType:          basetypes.MapTypable(nil),
											// Required:            false,
											Optional: true,
											// Computed:            false,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `MD5 key id

Format: u32:1-255
MD5 key id
`,
											// DeprecationMessage:  "",
											// Validators:          []validator.Map(nil),
											// PlanModifiers:       []planmodifier.Map(nil),
											// TODO investigate if tagnode defaults can be handled
											// Default:             defaults.Map(nil),
										},
									},
									// CustomType:          basetypes.MapTypable(nil),
									// Required:            false,
									Optional: true,
									// Computed:            false,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `MD5 key id

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
							MarkdownDescription: `Authentication

`,
							// DeprecationMessage:  "",
							// Validators:          []validator.Map(nil),
							// PlanModifiers:       []planmodifier.Map(nil),
							// TODO investigate if node defaults can be handled
							// Default:             defaults.Map(nil),
						},

						// CustomType:    basetypes.ObjectTypable(nil),
						// Validators:    []validator.Object(nil),
						// PlanModifiers: []planmodifier.Object(nil),
					},
				},
				// CustomType:          basetypes.MapTypable(nil),
				// Required:            false,
				Optional: true,
				// Computed:            false,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Virtual link

Format: ipv4
OSPF area in dotted decimal notation
`,
				// DeprecationMessage:  "",
				// Validators:          []validator.Map(nil),
				// PlanModifiers:       []planmodifier.Map(nil),
				// TODO investigate if tagnode defaults can be handled
				// Default:             defaults.Map(nil),
			},

			// TODO handle non-string types
			"authentication": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `OSPF area authentication type

Format: plaintext-password
Use plain-text authentication
Format: md5
Use MD5 authentication
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"network": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `OSPF network

Format: ipv4net
OSPF network
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"shortcut": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Area shortcut mode

Format: default
Set default
Format: disable
Disable shortcutting mode
Format: enable
Enable shortcutting mode
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"export_list": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Set the filter for networks announced to other areas

Format: u32
Access-list number
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"import_list": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Set the filter for networks from other areas announced

Format: u32
Access-list number
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			"area_type": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					// TODO handle non-string types
					"normal": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Normal OSPF area

`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},

					"nssa": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							// TODO handle non-string types
							"default_cost": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Summary-default cost of an NSSA area

Format: u32:0-16777215
Summary default cost
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"no_summary": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Do not inject inter-area routes into stub

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"translate": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Configure NSSA-ABR

Format: always
Always translate LSA types
Format: candidate
Translate for election
Format: never
Never translate LSA types
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

								Default:  stringdefault.StaticString(`candidate`),
								Computed: true,
							},
						},
						// CustomType:          basetypes.MapTypable(nil),
						// Required:            false,
						Optional: true,
						// Computed:            false,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Not-So-Stubby OSPF area

`,
						// DeprecationMessage:  "",
						// Validators:          []validator.Map(nil),
						// PlanModifiers:       []planmodifier.Map(nil),
						// TODO investigate if node defaults can be handled
						// Default:             defaults.Map(nil),
					},

					"stub": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							// TODO handle non-string types
							"default_cost": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Summary-default cost

Format: u32:0-16777215
Summary default cost
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"no_summary": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Do not inject inter-area routes into the stub

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
						MarkdownDescription: `Stub OSPF area

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
				MarkdownDescription: `Area type

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
func (r *protocols_ospf_area) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *protocols_ospf_areaModel

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
func (r *protocols_ospf_area) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *protocols_ospf_areaModel

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
func (r *protocols_ospf_area) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *protocols_ospf_areaModel

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
func (r *protocols_ospf_area) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *protocols_ospf_areaModel

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