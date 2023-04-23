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
var _ resource.Resource = &interfaces_macsec{}

// var _ resource.ResourceWithImportState = &interfaces_macsec{}

// interfaces_macsec defines the resource implementation.
type interfaces_macsec struct {
	ResourceName string
	client       *client.Client
}

func (r *interfaces_macsec) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// interfaces_macsecModel describes the resource data model.
type interfaces_macsecModel struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	Address          types.String `tfsdk:"address"`
	Description      types.String `tfsdk:"description"`
	Disable          types.String `tfsdk:"disable"`
	Mtu              types.String `tfsdk:"mtu"`
	Source_interface types.String `tfsdk:"source_interface"`
	Redirect         types.String `tfsdk:"redirect"`
	Vrf              types.String `tfsdk:"vrf"`

	// TagNodes

	// Nodes
	Dhcp_options     types.List `tfsdk:"dhcp_options"`
	Dhcpvsix_options types.List `tfsdk:"dhcpv6_options"`
	Ip               types.List `tfsdk:"ip"`
	Ipvsix           types.List `tfsdk:"ipv6"`
	Mirror           types.List `tfsdk:"mirror"`
	Security         types.List `tfsdk:"security"`
}

func (m interfaces_macsecModel) GetValues() (vyosPath []string, values map[string]attr.Value) {

	vyosPath = []string{
		"interfaces",
		"macsec",

		m.ID.ValueString(),
	}

	values = map[string]attr.Value{

		// LeafNodes
		"address":          m.Address,
		"description":      m.Description,
		"disable":          m.Disable,
		"mtu":              m.Mtu,
		"source_interface": m.Source_interface,
		"redirect":         m.Redirect,
		"vrf":              m.Vrf,

		// TagNodes

		// Nodes
		"dhcp_options":   m.Dhcp_options,
		"dhcpv6_options": m.Dhcpvsix_options,
		"ip":             m.Ip,
		"ipv6":           m.Ipvsix,
		"mirror":         m.Mirror,
		"security":       m.Security,
	}

	return vyosPath, values
}

// Metadata method to define the resource type name.
func (r interfaces_macsec) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	r.ResourceName = req.ProviderTypeName + "_interfaces_macsec"
	resp.TypeName = r.ResourceName
}

// interfaces_macsecResource method to return the example resource reference
func interfaces_macsecResource() resource.Resource {
	return &interfaces_macsec{}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r interfaces_macsec) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: ``,

		Attributes: map[string]schema.Attribute{
			"identifier": schema.StringAttribute{
				Required: true,
				MarkdownDescription: `MACsec Interface (802.1ae)

|  Format  |  Description  |
|----------|---------------|
|  macsecN  |  MACsec interface name  |
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

			"mtu": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Maximum Transmission Unit (MTU)

|  Format  |  Description  |
|----------|---------------|
|  u32:68-16000  |  Maximum Transmission Unit in byte  |
`,

				Default:  stringdefault.StaticString(`1460`),
				Computed: true,
			},

			"source_interface": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Physical interface the traffic will go through

|  Format  |  Description  |
|----------|---------------|
|  interface  |  Physical interface used for traffic forwarding  |
`,
			},

			"redirect": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Redirect incoming packet to destination

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Destination interface name  |
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

			"ip": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{

					"adjust_mss": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Adjust TCP MSS value

|  Format  |  Description  |
|----------|---------------|
|  clamp-mss-to-pmtu  |  Automatically sets the MSS to the proper value  |
|  u32:536-65535  |  TCP Maximum segment size in bytes  |
`,
					},

					"arp_cache_timeout": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `ARP cache entry timeout in seconds

|  Format  |  Description  |
|----------|---------------|
|  u32:1-86400  |  ARP cache entry timout in seconds  |
`,

						Default:  stringdefault.StaticString(`30`),
						Computed: true,
					},

					"disable_arp_filter": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Disable ARP filter on this interface

`,
					},

					"disable_forwarding": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Disable IP forwarding on this interface

`,
					},

					"enable_directed_broadcast": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Enable directed broadcast forwarding on this interface

`,
					},

					"enable_arp_accept": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Enable ARP accept on this interface

`,
					},

					"enable_arp_announce": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Enable ARP announce on this interface

`,
					},

					"enable_arp_ignore": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Enable ARP ignore on this interface

`,
					},

					"enable_proxy_arp": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Enable proxy-arp on this interface

`,
					},

					"proxy_arp_pvlan": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Enable private VLAN proxy ARP on this interface

`,
					},

					"source_validation": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Source validation by reversed path (RFC3704)

|  Format  |  Description  |
|----------|---------------|
|  strict  |  Enable Strict Reverse Path Forwarding as defined in RFC3704  |
|  loose  |  Enable Loose Reverse Path Forwarding as defined in RFC3704  |
|  disable  |  No source validation  |
`,
					},
				},
				Optional: true,
				MarkdownDescription: `IPv4 routing parameters

`,
			},

			"ipv6": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{

					"adjust_mss": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Adjust TCP MSS value

|  Format  |  Description  |
|----------|---------------|
|  clamp-mss-to-pmtu  |  Automatically sets the MSS to the proper value  |
|  u32:536-65535  |  TCP Maximum segment size in bytes  |
`,
					},

					"disable_forwarding": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Disable IP forwarding on this interface

`,
					},

					"dup_addr_detect_transmits": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Number of NS messages to send while performing DAD (default: 1)

|  Format  |  Description  |
|----------|---------------|
|  u32:0  |  Disable Duplicate Address Dectection (DAD)  |
|  u32:1-n  |  Number of NS messages to send while performing DAD  |
`,
					},

					"address": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{

							"autoconf": schema.StringAttribute{

								Optional: true,
								MarkdownDescription: `Enable acquisition of IPv6 address using stateless autoconfig (SLAAC)

`,
							},

							"eui64": schema.StringAttribute{

								Optional: true,
								MarkdownDescription: `Prefix for IPv6 address with MAC-based EUI-64

|  Format  |  Description  |
|----------|---------------|
|  <h:h:h:h:h:h:h:h/64>  |  IPv6 /64 network  |
`,
							},

							"no_default_link_local": schema.StringAttribute{

								Optional: true,
								MarkdownDescription: `Remove the default link-local address from the interface

`,
							},
						},
						Optional: true,
						MarkdownDescription: `IPv6 address configuration modes

`,
					},
				},
				Optional: true,
				MarkdownDescription: `IPv6 routing parameters

`,
			},

			"mirror": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{

					"ingress": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Mirror ingress traffic to destination interface

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Destination interface name  |
`,
					},

					"egress": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Mirror egress traffic to destination interface

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Destination interface name  |
`,
					},
				},
				Optional: true,
				MarkdownDescription: `Mirror ingress/egress packets

`,
			},

			"security": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{

					"cipher": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Cipher suite used

|  Format  |  Description  |
|----------|---------------|
|  gcm-aes-128  |  Galois/Counter Mode of AES cipher with 128-bit key  |
|  gcm-aes-256  |  Galois/Counter Mode of AES cipher with 256-bit key  |
`,
					},

					"encrypt": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Enable optional MACsec encryption

`,
					},

					"replay_window": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `IEEE 802.1X/MACsec replay protection window

|  Format  |  Description  |
|----------|---------------|
|  u32:0  |  No replay window, strict check  |
|  u32:1-4294967295  |  Number of packets that could be misordered  |
`,
					},

					"mka": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{

							"cak": schema.StringAttribute{

								Optional: true,
								MarkdownDescription: `Secure Connectivity Association Key

|  Format  |  Description  |
|----------|---------------|
|  txt  |  16-byte (128-bit) hex-string (32 hex-digits) for gcm-aes-128 or 32-byte (256-bit) hex-string (64 hex-digits) for gcm-aes-256  |
`,
							},

							"ckn": schema.StringAttribute{

								Optional: true,
								MarkdownDescription: `Secure Connectivity Association Key Name

|  Format  |  Description  |
|----------|---------------|
|  txt  |  1..32-bytes (8..256 bit) hex-string (2..64 hex-digits)  |
`,
							},

							"priority": schema.StringAttribute{

								Optional: true,
								MarkdownDescription: `Priority of MACsec Key Agreement protocol (MKA) actor

|  Format  |  Description  |
|----------|---------------|
|  u32:0-255  |  MACsec Key Agreement protocol (MKA) priority  |
`,

								Default:  stringdefault.StaticString(`255`),
								Computed: true,
							},
						},
						Optional: true,
						MarkdownDescription: `MACsec Key Agreement protocol (MKA)

`,
					},
				},
				Optional: true,
				MarkdownDescription: `Security/Encryption Settings

`,
			},
		},
	}
}

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func (r interfaces_macsec) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

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
func (r interfaces_macsec) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *interfaces_macsecModel

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
func (r interfaces_macsec) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *interfaces_macsecModel

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
func (r interfaces_macsec) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *interfaces_macsecModel

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
