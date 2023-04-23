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
var _ resource.Resource = &policy_access_listsix{}

// var _ resource.ResourceWithImportState = &policy_access_listsix{}

// policy_access_listsix defines the resource implementation.
type policy_access_listsix struct {
	ResourceName string
	client       *client.Client
}

func (r *policy_access_listsix) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// policy_access_listsixModel describes the resource data model.
type policy_access_listsixModel struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	Description types.String `tfsdk:"description"`

	// TagNodes
	Rule types.Map `tfsdk:"rule"`

	// Nodes

}

func (m policy_access_listsixModel) GetValues() (vyosPath []string, values map[string]attr.Value) {

	vyosPath = []string{
		"policy",
		"access-list6",

		m.ID.ValueString(),
	}

	values = map[string]attr.Value{

		// LeafNodes
		"description": m.Description,

		// TagNodes
		"rule": m.Rule,

		// Nodes

	}

	return vyosPath, values
}

// Metadata method to define the resource type name.
func (r policy_access_listsix) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	r.ResourceName = req.ProviderTypeName + "_policy_access_list6"
	resp.TypeName = r.ResourceName
}

// policy_access_listsixResource method to return the example resource reference
func policy_access_listsixResource() resource.Resource {
	return &policy_access_listsix{}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r policy_access_listsix) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Routing policy

`,

		Attributes: map[string]schema.Attribute{
			"identifier": schema.StringAttribute{
				Required: true,
				MarkdownDescription: `IPv6 access-list filter

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Name of IPv6 access-list  |
`,
			},

			"rule": schema.MapNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"action": schema.StringAttribute{

							Optional: true,
							MarkdownDescription: `Action to take on entries matching this rule

|  Format  |  Description  |
|----------|---------------|
|  permit  |  Permit matching entries  |
|  deny  |  Deny matching entries  |
`,
						},

						"description": schema.StringAttribute{

							Optional: true,
							MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |
`,
						},

						"source": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{

								"any": schema.StringAttribute{

									Optional: true,
									MarkdownDescription: `Any IP address to match

`,
								},

								"exact_match": schema.StringAttribute{

									Optional: true,
									MarkdownDescription: `Exact match of the network prefixes

`,
								},

								"network": schema.StringAttribute{

									Optional: true,
									MarkdownDescription: `Network/netmask to match

|  Format  |  Description  |
|----------|---------------|
|  ipv6net  |  IPv6 address and prefix length  |
`,
								},
							},
							Optional: true,
							MarkdownDescription: `Source IPv6 network to match

`,
						},
					},
				},
				Optional: true,
				MarkdownDescription: `Rule for this access-list6

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Access-list6 rule number  |
`,
			},

			"description": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |
`,
			},
		},
	}
}

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func (r policy_access_listsix) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

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
func (r policy_access_listsix) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *policy_access_listsixModel

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
func (r policy_access_listsix) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *policy_access_listsixModel

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
func (r policy_access_listsix) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *policy_access_listsixModel

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
