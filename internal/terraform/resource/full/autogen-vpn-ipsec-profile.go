package resourcefull

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &vpn_ipsec_profile{}

// var _ resource.ResourceWithImportState = &vpn_ipsec_profile{}

// vpn_ipsec_profile defines the resource implementation.
type vpn_ipsec_profile struct {
	client *http.Client
}

// vpn ipsec profileModel describes the resource data model.
type vpn_ipsec_profileModel struct {
	ConfigurableAttribute types.String `tfsdk:"configurable_attribute"`
	Defaulted             types.String `tfsdk:"defaulted"`
	ID                    types.String `tfsdk:"id"`

	Authentication types.String `tfsdk:authentication`

	Bind types.String `tfsdk:bind`

	Disable types.String `tfsdk:disable`

	Esp_group types.String `tfsdk:esp_group`

	Ike_group types.String `tfsdk:ike_group`
}

// Metadata method to define the resource type name.
func (r *vpn_ipsec_profile) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn ipsec profile"
}

// vpn_ipsec_profileResource method to return the example resource reference
func vpn_ipsec_profileResource() resource.Resource {
	return &vpn_ipsec_profile{}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r *vpn_ipsec_profile) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example resource",

		Attributes: map[string]schema.Attribute{
			// TODO handle non-string types
			"disable": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Disable instance

`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"esp_group": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Encapsulating Security Payloads (ESP) group name

`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"ike_group": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Internet Key Exchange (IKE) group name

`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			"authentication": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					// TODO handle non-string types
					"mode": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Authentication mode

Format: pre-shared-secret
Use a pre-shared secret key
`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},

					// TODO handle non-string types
					"pre_shared_secret": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Pre-shared secret key

Format: txt
Pre-shared secret key
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
				MarkdownDescription: `Authentication

`,
				// DeprecationMessage:  "",
				// Validators:          []validator.Map(nil),
				// PlanModifiers:       []planmodifier.Map(nil),
				// TODO investigate if node defaults can be handled
				// Default:             defaults.Map(nil),
			},

			"bind": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					// TODO handle non-string types
					"tunnel": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Tunnel interface associated with this profile

Format: txt
Associated interface to this profile
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
				MarkdownDescription: `DMVPN tunnel configuration

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
func (r *vpn_ipsec_profile) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *vpn_ipsec_profileModel

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
func (r *vpn_ipsec_profile) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *vpn_ipsec_profileModel

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
func (r *vpn_ipsec_profile) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *vpn_ipsec_profileModel

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
func (r *vpn_ipsec_profile) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *vpn_ipsec_profileModel

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
