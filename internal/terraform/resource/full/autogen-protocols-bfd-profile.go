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
var _ resource.Resource = &protocols_bfd_profile{}

// var _ resource.ResourceWithImportState = &protocols_bfd_profile{}

// protocols_bfd_profile defines the resource implementation.
type protocols_bfd_profile struct {
	client *http.Client
}

// protocols bfd profileModel describes the resource data model.
type protocols_bfd_profileModel struct {
	ConfigurableAttribute types.String `tfsdk:"configurable_attribute"`
	Defaulted             types.String `tfsdk:"defaulted"`
	ID                    types.String `tfsdk:"id"`

	Interval types.String `tfsdk:interval`

	Echo_mode types.String `tfsdk:echo_mode`

	Passive types.String `tfsdk:passive`

	Shutdown types.String `tfsdk:shutdown`
}

// Metadata method to define the resource type name.
func (r *protocols_bfd_profile) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protocols bfd profile"
}

// protocols_bfd_profileResource method to return the example resource reference
func protocols_bfd_profileResource() resource.Resource {
	return &protocols_bfd_profile{}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r *protocols_bfd_profile) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example resource",

		Attributes: map[string]schema.Attribute{
			// TODO handle non-string types
			"echo_mode": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Enables the echo transmission mode

`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"passive": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Do not attempt to start sessions

`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"shutdown": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Disable this peer

`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			"interval": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					// TODO handle non-string types
					"receive": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Minimum interval of receiving control packets

Format: u32:10-60000
Interval in milliseconds
`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

						Default:  stringdefault.StaticString(`300`),
						Computed: true,
					},

					// TODO handle non-string types
					"transmit": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Minimum interval of transmitting control packets

Format: u32:10-60000
Interval in milliseconds
`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

						Default:  stringdefault.StaticString(`300`),
						Computed: true,
					},

					// TODO handle non-string types
					"multiplier": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Multiplier to determine packet loss

Format: u32:2-255
Remote transmission interval will be multiplied by this value
`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

						Default:  stringdefault.StaticString(`3`),
						Computed: true,
					},

					// TODO handle non-string types
					"echo_interval": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Echo receive transmission interval

Format: u32:10-60000
The minimal echo receive transmission interval that this system is capable of handling
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
				MarkdownDescription: `Configure timer intervals

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
func (r *protocols_bfd_profile) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *protocols_bfd_profileModel

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
func (r *protocols_bfd_profile) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *protocols_bfd_profileModel

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
func (r *protocols_bfd_profile) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *protocols_bfd_profileModel

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
func (r *protocols_bfd_profile) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *protocols_bfd_profileModel

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
