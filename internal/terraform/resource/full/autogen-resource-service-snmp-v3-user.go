// Code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.

package resourcefull

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &service_snmp_vthree_user{}

// var _ resource.ResourceWithImportState = &service_snmp_vthree_user{}

// service_snmp_vthree_user defines the resource implementation.
type service_snmp_vthree_user struct {
	ResourceName string
	client       *client.Client
}

func (r *service_snmp_vthree_user) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

// service_snmp_vthree_userModel describes the resource data model.
type service_snmp_vthree_userModel struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	Group types.String `tfsdk:"group"`
	Mode  types.String `tfsdk:"mode"`

	// TagNodes

	// Nodes
	Auth    types.List `tfsdk:"auth"`
	Privacy types.List `tfsdk:"privacy"`
}

func (m service_snmp_vthree_userModel) GetValues() (vyosPath []string, values map[string]attr.Value) {

	vyosPath = []string{
		"service",
		"snmp",
		"v3",
		"user",

		m.ID.ValueString(),
	}

	values = map[string]attr.Value{

		// LeafNodes
		"group": m.Group,
		"mode":  m.Mode,

		// TagNodes

		// Nodes
		"auth":    m.Auth,
		"privacy": m.Privacy,
	}

	return vyosPath, values
}

// Metadata method to define the resource type name.
func (r service_snmp_vthree_user) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	r.ResourceName = req.ProviderTypeName + "_service_snmp_v3_user"
	resp.TypeName = r.ResourceName
}

// service_snmp_vthree_userResource method to return the example resource reference
func service_snmp_vthree_userResource() resource.Resource {
	return &service_snmp_vthree_user{}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r service_snmp_vthree_user) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Simple Network Management Protocol (SNMP)

Simple Network Management Protocol (SNMP) v3

`,

		Attributes: map[string]schema.Attribute{
			"identifier": schema.StringAttribute{
				Required: true,
				MarkdownDescription: `Specifies the user with name username

`,
			},

			"group": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Specifies group for user name

`,
			},

			"mode": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Define access permission

|  Format  |  Description  |
|----------|---------------|
|  ro  |  Read-Only  |
|  rw  |  read write  |
`,

				Default:  stringdefault.StaticString(`ro`),
				Computed: true,
			},

			"auth": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{

					"encrypted_password": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Defines the encrypted key for authentication

`,
					},

					"plaintext_password": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Defines the clear text key for authentication

`,
					},

					"type": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Define used protocol

|  Format  |  Description  |
|----------|---------------|
|  md5  |  Message Digest 5  |
|  sha  |  Secure Hash Algorithm  |
`,

						Default:  stringdefault.StaticString(`md5`),
						Computed: true,
					},
				},
				Optional: true,
				MarkdownDescription: `Specifies the auth

`,
			},

			"privacy": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{

					"encrypted_password": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Defines the encrypted key for privacy protocol

`,
					},

					"plaintext_password": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Defines the clear text key for privacy protocol

`,
					},

					"type": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Defines the protocol for privacy

|  Format  |  Description  |
|----------|---------------|
|  des  |  Data Encryption Standard  |
|  aes  |  Advanced Encryption Standard  |
`,

						Default:  stringdefault.StaticString(`des`),
						Computed: true,
					},
				},
				Optional: true,
				MarkdownDescription: `Defines the privacy

`,
			},
		},
	}
}

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func (r service_snmp_vthree_user) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	ctx = context.WithValue(ctx, "crud_func", "Create")

	var data *firewall_nameModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create vyos api ops
	vyosOps := helpers.FromTerraformToVyos(ctx, data)
	for _, ops := range vyosOps {
		tflog.Error(ctx, "Vyos Ops generated", map[string]interface{}{"vyosOps": ops})
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	r.client.StageSet(ctx, vyosOps)
	response, err := r.client.CommitChanges(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create %s, got error: %s", r.ResourceName, err))
		return
	}
	if response != nil {
		tflog.Warn(ctx, "Got non-nil response from API", map[string]interface{}{"response": response})
	}

	// Save ID into the Terraform state.
	data.ID = types.StringValue(data.ID.ValueString())

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "created a resource")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read method to define the logic which refreshes the Terraform state for the resource.
func (r service_snmp_vthree_user) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *service_snmp_vthree_userModel

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
func (r service_snmp_vthree_user) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *service_snmp_vthree_userModel

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
func (r service_snmp_vthree_user) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *service_snmp_vthree_userModel

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
