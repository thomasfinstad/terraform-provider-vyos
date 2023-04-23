// Code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.

package resourcefull

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &container_registry{}

// var _ resource.ResourceWithImportState = &container_registry{}

// container_registry defines the resource implementation.
type container_registry struct {
	ResourceName string
	client       *client.Client
}

func (r *container_registry) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// container_registryModel describes the resource data model.
type container_registryModel struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	Disable types.String `tfsdk:"disable"`

	// TagNodes

	// Nodes
	Authentication types.List `tfsdk:"authentication"`
}

func (m container_registryModel) GetValues() (vyosPath []string, values map[string]attr.Value) {

	vyosPath = []string{
		"container",
		"registry",

		m.ID.ValueString(),
	}

	values = map[string]attr.Value{

		// LeafNodes
		"disable": m.Disable,

		// TagNodes

		// Nodes
		"authentication": m.Authentication,
	}

	return vyosPath, values
}

// Metadata method to define the resource type name.
func (r container_registry) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	r.ResourceName = req.ProviderTypeName + "_container_registry"
	resp.TypeName = r.ResourceName
}

// container_registryResource method to return the example resource reference
func container_registryResource() resource.Resource {
	return &container_registry{}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r container_registry) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Container applications

`,

		Attributes: map[string]schema.Attribute{
			"identifier": schema.StringAttribute{
				Required: true,
				MarkdownDescription: `Registry Name

`,
			},

			"disable": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Disable instance

`,
			},

			"authentication": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{

					"username": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Username used for authentication

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Username  |
`,
					},

					"password": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Password used for authentication

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Password  |
`,
					},
				},
				Optional: true,
				MarkdownDescription: `Authentication settings

`,
			},
		},
	}
}

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func (r container_registry) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

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
func (r container_registry) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *container_registryModel

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
func (r container_registry) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *container_registryModel

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
func (r container_registry) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *container_registryModel

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
