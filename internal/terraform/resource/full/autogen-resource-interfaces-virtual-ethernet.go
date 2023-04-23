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
var _ resource.Resource = &interfaces_virtual_ethernet{}

// var _ resource.ResourceWithImportState = &interfaces_virtual_ethernet{}

// interfaces_virtual_ethernet defines the resource implementation.
type interfaces_virtual_ethernet struct {
	ResourceName string
	client       *client.Client
}

func (r *interfaces_virtual_ethernet) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// interfaces_virtual_ethernetModel describes the resource data model.
type interfaces_virtual_ethernetModel struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	Address     types.String `tfsdk:"address"`
	Description types.String `tfsdk:"description"`
	Disable     types.String `tfsdk:"disable"`
	Vrf         types.String `tfsdk:"vrf"`
	Peer_name   types.String `tfsdk:"peer_name"`

	// TagNodes

	// Nodes
	Dhcp_options     types.List `tfsdk:"dhcp_options"`
	Dhcpvsix_options types.List `tfsdk:"dhcpv6_options"`
}

func (m interfaces_virtual_ethernetModel) GetValues() (vyosPath []string, values map[string]attr.Value) {

	vyosPath = []string{
		"interfaces",
		"virtual-ethernet",

		m.ID.ValueString(),
	}

	values = map[string]attr.Value{

		// LeafNodes
		"address":     m.Address,
		"description": m.Description,
		"disable":     m.Disable,
		"vrf":         m.Vrf,
		"peer_name":   m.Peer_name,

		// TagNodes

		// Nodes
		"dhcp_options":   m.Dhcp_options,
		"dhcpv6_options": m.Dhcpvsix_options,
	}

	return vyosPath, values
}

// Metadata method to define the resource type name.
func (r interfaces_virtual_ethernet) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	r.ResourceName = req.ProviderTypeName + "_interfaces_virtual_ethernet"
	resp.TypeName = r.ResourceName
}

// interfaces_virtual_ethernetResource method to return the example resource reference
func interfaces_virtual_ethernetResource() resource.Resource {
	return &interfaces_virtual_ethernet{}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r interfaces_virtual_ethernet) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: ``,

		Attributes: map[string]schema.Attribute{
			"identifier": schema.StringAttribute{
				Required: true,
				MarkdownDescription: `Virtual Ethernet (veth) Interface

|  Format  |  Description  |
|----------|---------------|
|  vethN  |  Virtual Ethernet interface name  |
`,
			},

			"address": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `IP address

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  IPv4 address and prefix length  |
|  ipv6net  |  IPv6 address and prefix length  |
|  dhcp  |  Dynamic Host Configuration Protocol  |
|  dhcpv6  |  Dynamic Host Configuration Protocol for IPv6  |
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

			"disable": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Administratively disable interface

`,
			},

			"vrf": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `VRF instance name

|  Format  |  Description  |
|----------|---------------|
|  txt  |  VRF instance name  |
`,
			},

			"peer_name": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Virtual ethernet peer interface name

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Name of peer interface  |
`,
			},

			"dhcp_options": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{

					"client_id": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Identifier used by client to identify itself to the DHCP server

`,
					},

					"host_name": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Override system host-name sent to DHCP server

`,
					},

					"mtu": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Use MTU value from DHCP server - ignore interface setting

`,
					},

					"vendor_class_id": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Identify the vendor client type to the DHCP server

`,
					},

					"no_default_route": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Do not install default route to system

`,
					},

					"default_route_distance": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Distance for installed default route

|  Format  |  Description  |
|----------|---------------|
|  u32:1-255  |  Distance for the default route from DHCP server  |
`,

						Default:  stringdefault.StaticString(`210`),
						Computed: true,
					},

					"reject": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `IP addresses or subnets from which to reject DHCP leases

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address to match  |
|  ipv4net  |  IPv4 prefix to match  |
`,
					},
				},
				Optional: true,
				MarkdownDescription: `DHCP client settings/options

`,
			},

			"dhcpv6_options": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{

					"pd": schema.MapNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"interface": schema.MapNestedAttribute{
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"address": schema.StringAttribute{

												Optional: true,
												MarkdownDescription: `Local interface address assigned to interface (default: EUI-64)

|  Format  |  Description  |
|----------|---------------|
|  >0  |  Used to form IPv6 interface address  |
`,
											},

											"sla_id": schema.StringAttribute{

												Optional: true,
												MarkdownDescription: `Interface site-Level aggregator (SLA)

|  Format  |  Description  |
|----------|---------------|
|  u32:0-65535  |  Decimal integer which fits in the length of SLA IDs  |
`,
											},
										},
									},
									Optional: true,
									MarkdownDescription: `Delegate IPv6 prefix from provider to this interface

`,
								},

								"length": schema.StringAttribute{

									Optional: true,
									MarkdownDescription: `Request IPv6 prefix length from peer

|  Format  |  Description  |
|----------|---------------|
|  u32:32-64  |  Length of delegated prefix  |
`,

									Default:  stringdefault.StaticString(`64`),
									Computed: true,
								},
							},
						},
						Optional: true,
						MarkdownDescription: `DHCPv6 prefix delegation interface statement

|  Format  |  Description  |
|----------|---------------|
|  instance number  |  Prefix delegation instance (>= 0)  |
`,
					},

					"duid": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `DHCP unique identifier (DUID) to be sent by dhcpv6 client

|  Format  |  Description  |
|----------|---------------|
|  duid  |  DHCP unique identifier (DUID)  |
`,
					},

					"parameters_only": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Acquire only config parameters, no address

`,
					},

					"rapid_commit": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Wait for immediate reply instead of advertisements

`,
					},

					"temporary": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `IPv6 temporary address

`,
					},
				},
				Optional: true,
				MarkdownDescription: `DHCPv6 client settings/options

`,
			},
		},
	}
}

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func (r interfaces_virtual_ethernet) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

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
func (r interfaces_virtual_ethernet) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *interfaces_virtual_ethernetModel

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
func (r interfaces_virtual_ethernet) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *interfaces_virtual_ethernetModel

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
func (r interfaces_virtual_ethernet) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *interfaces_virtual_ethernetModel

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
