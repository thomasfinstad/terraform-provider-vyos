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
var _ resource.Resource = &service_ipoe_server_client_ipvsix_pool_delegate{}

// var _ resource.ResourceWithImportState = &service_ipoe_server_client_ipvsix_pool_delegate{}

// service_ipoe_server_client_ipvsix_pool_delegate defines the resource implementation.
type service_ipoe_server_client_ipvsix_pool_delegate struct {
	ResourceName string
	client       *client.Client
}

func (r *service_ipoe_server_client_ipvsix_pool_delegate) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// service_ipoe_server_client_ipvsix_pool_delegateModel describes the resource data model.
type service_ipoe_server_client_ipvsix_pool_delegateModel struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	Delegation_prefix types.String `tfsdk:"delegation_prefix"`

	// TagNodes

	// Nodes

}

func (m service_ipoe_server_client_ipvsix_pool_delegateModel) GetValues() (vyosPath []string, values map[string]attr.Value) {

	vyosPath = []string{
		"service",
		"ipoe-server",
		"client-ipv6-pool",
		"delegate",

		m.ID.ValueString(),
	}

	values = map[string]attr.Value{

		// LeafNodes
		"delegation_prefix": m.Delegation_prefix,

		// TagNodes

		// Nodes

	}

	return vyosPath, values
}

// Metadata method to define the resource type name.
func (r service_ipoe_server_client_ipvsix_pool_delegate) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	r.ResourceName = req.ProviderTypeName + "_service_ipoe_server_client_ipv6_pool_delegate"
	resp.TypeName = r.ResourceName
}

// service_ipoe_server_client_ipvsix_pool_delegateResource method to return the example resource reference
func service_ipoe_server_client_ipvsix_pool_delegateResource() resource.Resource {
	return &service_ipoe_server_client_ipvsix_pool_delegate{}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r service_ipoe_server_client_ipvsix_pool_delegate) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Internet Protocol over Ethernet (IPoE) Server

Pool of client IPv6 addresses

`,

		Attributes: map[string]schema.Attribute{
			"identifier": schema.StringAttribute{
				Required: true,
				MarkdownDescription: `Subnet used to delegate prefix through DHCPv6-PD (RFC3633)

|  Format  |  Description  |
|----------|---------------|
|  ipv6net  |  IPv6 address and prefix length  |
`,
			},

			"delegation_prefix": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Prefix length delegated to client

|  Format  |  Description  |
|----------|---------------|
|  u32:32-64  |  Delegated prefix length  |
`,
			},
		},
	}
}

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func (r service_ipoe_server_client_ipvsix_pool_delegate) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

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
func (r service_ipoe_server_client_ipvsix_pool_delegate) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *service_ipoe_server_client_ipvsix_pool_delegateModel

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
func (r service_ipoe_server_client_ipvsix_pool_delegate) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *service_ipoe_server_client_ipvsix_pool_delegateModel

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
func (r service_ipoe_server_client_ipvsix_pool_delegate) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *service_ipoe_server_client_ipvsix_pool_delegateModel

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
