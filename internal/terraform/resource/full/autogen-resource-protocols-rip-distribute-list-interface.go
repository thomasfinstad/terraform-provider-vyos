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
var _ resource.Resource = &protocols_rip_distribute_list_interface{}

// var _ resource.ResourceWithImportState = &protocols_rip_distribute_list_interface{}

// protocols_rip_distribute_list_interface defines the resource implementation.
type protocols_rip_distribute_list_interface struct {
	ResourceName string
	client       *client.Client
}

func (r *protocols_rip_distribute_list_interface) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// protocols_rip_distribute_list_interfaceModel describes the resource data model.
type protocols_rip_distribute_list_interfaceModel struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes

	// TagNodes

	// Nodes
	Access_list types.List `tfsdk:"access_list"`
	Prefix_list types.List `tfsdk:"prefix_list"`
}

func (m protocols_rip_distribute_list_interfaceModel) GetValues() (vyosPath []string, values map[string]attr.Value) {

	vyosPath = []string{
		"protocols",
		"rip",
		"distribute-list",
		"interface",

		m.ID.ValueString(),
	}

	values = map[string]attr.Value{

		// LeafNodes

		// TagNodes

		// Nodes
		"access_list": m.Access_list,
		"prefix_list": m.Prefix_list,
	}

	return vyosPath, values
}

// Metadata method to define the resource type name.
func (r protocols_rip_distribute_list_interface) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	r.ResourceName = req.ProviderTypeName + "_protocols_rip_distribute_list_interface"
	resp.TypeName = r.ResourceName
}

// protocols_rip_distribute_list_interfaceResource method to return the example resource reference
func protocols_rip_distribute_list_interfaceResource() resource.Resource {
	return &protocols_rip_distribute_list_interface{}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r protocols_rip_distribute_list_interface) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Routing Information Protocol (RIP) parameters

Filter networks in routing updates

`,

		Attributes: map[string]schema.Attribute{
			"identifier": schema.StringAttribute{
				Required: true,
				MarkdownDescription: `Apply filtering to an interface

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Apply filtering to an interface  |
`,
			},

			"access_list": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{

					"in": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Access list to apply to input packets

|  Format  |  Description  |
|----------|---------------|
|  u32  |  Access list to apply to input packets  |
`,
					},

					"out": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Access list to apply to output packets

|  Format  |  Description  |
|----------|---------------|
|  u32  |  Access list to apply to output packets  |
`,
					},
				},
				Optional: true,
				MarkdownDescription: `Access-list

`,
			},

			"prefix_list": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{

					"in": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Prefix-list to apply to input packets

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Prefix-list to apply to input packets  |
`,
					},

					"out": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Prefix-list to apply to output packets

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Prefix-list to apply to output packets  |
`,
					},
				},
				Optional: true,
				MarkdownDescription: `Prefix-list

`,
			},
		},
	}
}

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func (r protocols_rip_distribute_list_interface) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

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
func (r protocols_rip_distribute_list_interface) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *protocols_rip_distribute_list_interfaceModel

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
func (r protocols_rip_distribute_list_interface) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *protocols_rip_distribute_list_interfaceModel

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
func (r protocols_rip_distribute_list_interface) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *protocols_rip_distribute_list_interfaceModel

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
