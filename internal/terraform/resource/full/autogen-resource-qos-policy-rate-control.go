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
var _ resource.Resource = &qos_policy_rate_control{}

// var _ resource.ResourceWithImportState = &qos_policy_rate_control{}

// qos_policy_rate_control defines the resource implementation.
type qos_policy_rate_control struct {
	ResourceName string
	client       *client.Client
}

func (r *qos_policy_rate_control) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// qos_policy_rate_controlModel describes the resource data model.
type qos_policy_rate_controlModel struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	Description types.String `tfsdk:"description"`
	Bandwidth   types.String `tfsdk:"bandwidth"`
	Burst       types.String `tfsdk:"burst"`
	Latency     types.String `tfsdk:"latency"`

	// TagNodes

	// Nodes

}

func (m qos_policy_rate_controlModel) GetValues() (vyosPath []string, values map[string]attr.Value) {

	vyosPath = []string{
		"qos",
		"policy",
		"rate-control",

		m.ID.ValueString(),
	}

	values = map[string]attr.Value{

		// LeafNodes
		"description": m.Description,
		"bandwidth":   m.Bandwidth,
		"burst":       m.Burst,
		"latency":     m.Latency,

		// TagNodes

		// Nodes

	}

	return vyosPath, values
}

// Metadata method to define the resource type name.
func (r qos_policy_rate_control) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	r.ResourceName = req.ProviderTypeName + "_qos_policy_rate_control"
	resp.TypeName = r.ResourceName
}

// qos_policy_rate_controlResource method to return the example resource reference
func qos_policy_rate_controlResource() resource.Resource {
	return &qos_policy_rate_control{}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r qos_policy_rate_control) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Quality of Service (QoS)

Service Policy definitions

`,

		Attributes: map[string]schema.Attribute{
			"identifier": schema.StringAttribute{
				Required: true,
				MarkdownDescription: `Rate limiting policy (Token Bucket Filter)

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Policy name  |
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

			"bandwidth": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Available bandwidth for this policy

|  Format  |  Description  |
|----------|---------------|
|  <number>  |  Bits per second  |
|  <number>bit  |  Bits per second  |
|  <number>kbit  |  Kilobits per second  |
|  <number>mbit  |  Megabits per second  |
|  <number>gbit  |  Gigabits per second  |
|  <number>tbit  |  Terabits per second  |
|  <number>%  |  Percentage of interface link speed  |
`,
			},

			"burst": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Burst size for this class

|  Format  |  Description  |
|----------|---------------|
|  <number>  |  Bytes  |
|  <number><suffix>  |  Bytes with scaling suffix (kb, mb, gb)  |
`,

				Default:  stringdefault.StaticString(`15k`),
				Computed: true,
			},

			"latency": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Maximum latency

|  Format  |  Description  |
|----------|---------------|
|  <number>  |  Time in milliseconds  |
`,

				Default:  stringdefault.StaticString(`50`),
				Computed: true,
			},
		},
	}
}

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func (r qos_policy_rate_control) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

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
func (r qos_policy_rate_control) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *qos_policy_rate_controlModel

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
func (r qos_policy_rate_control) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *qos_policy_rate_controlModel

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
func (r qos_policy_rate_control) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *qos_policy_rate_controlModel

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
