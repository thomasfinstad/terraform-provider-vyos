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
var _ resource.Resource = &protocols_bfd_profile{}

// var _ resource.ResourceWithImportState = &protocols_bfd_profile{}

// protocols_bfd_profile defines the resource implementation.
type protocols_bfd_profile struct {
	ResourceName string
	client       *client.Client
}

func (r *protocols_bfd_profile) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// protocols_bfd_profileModel describes the resource data model.
type protocols_bfd_profileModel struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	Echo_mode types.String `tfsdk:"echo_mode"`
	Passive   types.String `tfsdk:"passive"`
	Shutdown  types.String `tfsdk:"shutdown"`

	// TagNodes

	// Nodes
	Interval types.List `tfsdk:"interval"`
}

func (m protocols_bfd_profileModel) GetValues() (vyosPath []string, values map[string]attr.Value) {

	vyosPath = []string{
		"protocols",
		"bfd",
		"profile",

		m.ID.ValueString(),
	}

	values = map[string]attr.Value{

		// LeafNodes
		"echo_mode": m.Echo_mode,
		"passive":   m.Passive,
		"shutdown":  m.Shutdown,

		// TagNodes

		// Nodes
		"interval": m.Interval,
	}

	return vyosPath, values
}

// Metadata method to define the resource type name.
func (r protocols_bfd_profile) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	r.ResourceName = req.ProviderTypeName + "_protocols_bfd_profile"
	resp.TypeName = r.ResourceName
}

// protocols_bfd_profileResource method to return the example resource reference
func protocols_bfd_profileResource() resource.Resource {
	return &protocols_bfd_profile{}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r protocols_bfd_profile) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Bidirectional Forwarding Detection (BFD)

`,

		Attributes: map[string]schema.Attribute{
			"identifier": schema.StringAttribute{
				Required: true,
				MarkdownDescription: `Configure BFD profile used by individual peer

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Name of BFD profile  |
`,
			},

			"echo_mode": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Enables the echo transmission mode

`,
			},

			"passive": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Do not attempt to start sessions

`,
			},

			"shutdown": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Disable this peer

`,
			},

			"interval": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{

					"receive": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Minimum interval of receiving control packets

|  Format  |  Description  |
|----------|---------------|
|  u32:10-60000  |  Interval in milliseconds  |
`,

						Default:  stringdefault.StaticString(`300`),
						Computed: true,
					},

					"transmit": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Minimum interval of transmitting control packets

|  Format  |  Description  |
|----------|---------------|
|  u32:10-60000  |  Interval in milliseconds  |
`,

						Default:  stringdefault.StaticString(`300`),
						Computed: true,
					},

					"multiplier": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Multiplier to determine packet loss

|  Format  |  Description  |
|----------|---------------|
|  u32:2-255  |  Remote transmission interval will be multiplied by this value  |
`,

						Default:  stringdefault.StaticString(`3`),
						Computed: true,
					},

					"echo_interval": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Echo receive transmission interval

|  Format  |  Description  |
|----------|---------------|
|  u32:10-60000  |  The minimal echo receive transmission interval that this system is capable of handling  |
`,
					},
				},
				Optional: true,
				MarkdownDescription: `Configure timer intervals

`,
			},
		},
	}
}

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func (r protocols_bfd_profile) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

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
func (r protocols_bfd_profile) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
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
func (r protocols_bfd_profile) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
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
func (r protocols_bfd_profile) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
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
