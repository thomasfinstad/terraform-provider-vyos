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
var _ resource.Resource = &system_login_user{}

// var _ resource.ResourceWithImportState = &system_login_user{}

// system_login_user defines the resource implementation.
type system_login_user struct {
	client *http.Client
}

// system login userModel describes the resource data model.
type system_login_userModel struct {
	ConfigurableAttribute types.String `tfsdk:"configurable_attribute"`
	Defaulted             types.String `tfsdk:"defaulted"`
	ID                    types.String `tfsdk:"id"`

	Authentication types.String `tfsdk:authentication`

	Full_name types.String `tfsdk:full_name`

	Home_directory types.String `tfsdk:home_directory`
}

// Metadata method to define the resource type name.
func (r *system_login_user) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system login user"
}

// system_login_userResource method to return the example resource reference
func system_login_userResource() resource.Resource {
	return &system_login_user{}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r *system_login_user) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example resource",

		Attributes: map[string]schema.Attribute{
			// TODO handle non-string types
			"full_name": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Full name of the user (use quotes for names with spaces)

`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"home_directory": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Home directory

`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			"authentication": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"public_keys": schema.MapNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								// TODO handle non-string types
								"key": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Public key value (Base64 encoded)

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"options": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Optional public key options

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"type": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `SSH public key type

Format: ssh-dss
Digital Signature Algorithm (DSA) key support
Format: ssh-rsa
Key pair based on RSA algorithm
Format: ecdsa-sha2-nistp256
Elliptic Curve DSA with NIST P-256 curve
Format: ecdsa-sha2-nistp384
Elliptic Curve DSA with NIST P-384 curve
Format: ecdsa-sha2-nistp521
Elliptic Curve DSA with NIST P-521 curve
Format: ssh-ed25519
Edwards-curve DSA with elliptic curve 25519
Format: sk-ecdsa-sha2-nistp256@openssh.com
Elliptic Curve DSA security key
Format: sk-ssh-ed25519@openssh.com
Elliptic curve 25519 security key
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
						MarkdownDescription: `Remote access public keys

Format: txt
Key identifier used by ssh-keygen (usually of form user@host)
`,
						// DeprecationMessage:  "",
						// Validators:          []validator.Map(nil),
						// PlanModifiers:       []planmodifier.Map(nil),
						// TODO investigate if tagnode defaults can be handled
						// Default:             defaults.Map(nil),
					},

					// TODO handle non-string types
					"encrypted_password": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Encrypted password

`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

						Default:  stringdefault.StaticString(`!`),
						Computed: true,
					},

					// TODO handle non-string types
					"plaintext_password": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Plaintext password used for encryption

`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},

					"otp": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							// TODO handle non-string types
							"rate_limit": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Limit number of logins (rate-limit) per rate-time

Format: u32:1-10
Number of attempts
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

								Default:  stringdefault.StaticString(`3`),
								Computed: true,
							},

							// TODO handle non-string types
							"rate_time": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Limit number of logins (rate-limit) per rate-time

Format: u32:15-600
Time interval
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

								Default:  stringdefault.StaticString(`30`),
								Computed: true,
							},

							// TODO handle non-string types
							"window_size": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Set window of concurrently valid codes

Format: u32:1-21
Window size
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

								Default:  stringdefault.StaticString(`3`),
								Computed: true,
							},

							// TODO handle non-string types
							"key": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Key/secret the token algorithm (see RFC4226)

Format: txt
Base32 encoded key/token
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
						MarkdownDescription: `One-Time-Pad (two-factor) authentication parameters

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
				MarkdownDescription: `Authentication settings

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
func (r *system_login_user) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *system_login_userModel

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
func (r *system_login_user) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *system_login_userModel

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
func (r *system_login_user) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *system_login_userModel

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
func (r *system_login_user) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *system_login_userModel

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
