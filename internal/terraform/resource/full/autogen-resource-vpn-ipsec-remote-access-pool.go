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
var _ resource.Resource = &vpn_ipsec_remote_access_pool{}

// var _ resource.ResourceWithImportState = &vpn_ipsec_remote_access_pool{}

// vpn_ipsec_remote_access_pool defines the resource implementation.
type vpn_ipsec_remote_access_pool struct {
	ResourceName string
	client       *client.Client
}

func (r *vpn_ipsec_remote_access_pool) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// vpn_ipsec_remote_access_poolModel describes the resource data model.
type vpn_ipsec_remote_access_poolModel struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	Exclude     types.String `tfsdk:"exclude"`
	Prefix      types.String `tfsdk:"prefix"`
	Name_server types.String `tfsdk:"name_server"`

	// TagNodes

	// Nodes

}

func (m vpn_ipsec_remote_access_poolModel) GetValues() (vyosPath []string, values map[string]attr.Value) {

	vyosPath = []string{
		"vpn",
		"ipsec",
		"remote-access",
		"pool",

		m.ID.ValueString(),
	}

	values = map[string]attr.Value{

		// LeafNodes
		"exclude":     m.Exclude,
		"prefix":      m.Prefix,
		"name_server": m.Name_server,

		// TagNodes

		// Nodes

	}

	return vyosPath, values
}

// Metadata method to define the resource type name.
func (r vpn_ipsec_remote_access_pool) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	r.ResourceName = req.ProviderTypeName + "_vpn_ipsec_remote_access_pool"
	resp.TypeName = r.ResourceName
}

// vpn_ipsec_remote_access_poolResource method to return the example resource reference
func vpn_ipsec_remote_access_poolResource() resource.Resource {
	return &vpn_ipsec_remote_access_pool{}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r vpn_ipsec_remote_access_pool) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Virtual Private Network (VPN)

VPN IP security (IPsec) parameters

IKEv2 remote access VPN

`,

		Attributes: map[string]schema.Attribute{
			"identifier": schema.StringAttribute{
				Required: true,
				MarkdownDescription: `IP address pool for remote access users

`,
			},

			"exclude": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Local IPv4 or IPv6 pool prefix exclusions

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  Local IPv4 pool prefix exclusion  |
|  ipv6net  |  Local IPv6 pool prefix exclusion  |
`,
			},

			"prefix": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Local IPv4 or IPv6 pool prefix

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  Local IPv4 pool prefix  |
|  ipv6net  |  Local IPv6 pool prefix  |
`,
			},

			"name_server": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Domain Name Servers (DNS) addresses

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Domain Name Server (DNS) IPv4 address  |
|  ipv6  |  Domain Name Server (DNS) IPv6 address  |
`,
			},
		},
	}
}

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func (r vpn_ipsec_remote_access_pool) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

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
func (r vpn_ipsec_remote_access_pool) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *vpn_ipsec_remote_access_poolModel

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
func (r vpn_ipsec_remote_access_pool) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *vpn_ipsec_remote_access_poolModel

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
func (r vpn_ipsec_remote_access_pool) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *vpn_ipsec_remote_access_poolModel

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
