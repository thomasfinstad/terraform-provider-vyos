package resourcefull

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &policy_route_map{}

// var _ resource.ResourceWithImportState = &policy_route_map{}

// policy_route_map defines the resource implementation.
type policy_route_map struct {
	client *http.Client
}

// policy route-mapModel describes the resource data model.
type policy_route_mapModel struct {
	ConfigurableAttribute types.String `tfsdk:"configurable_attribute"`
	Defaulted             types.String `tfsdk:"defaulted"`
	ID                    types.String `tfsdk:"id"`

	Rule types.Map `tfsdk:rule`

	Description types.String `tfsdk:description`
}

// Metadata method to define the resource type name.
func (r *policy_route_map) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policy route-map"
}

// policy_route_mapResource method to return the example resource reference
func policy_route_mapResource() resource.Resource {
	return &policy_route_map{}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r *policy_route_map) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example resource",

		Attributes: map[string]schema.Attribute{
			"rule": schema.MapNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						// TODO handle non-string types
						"action": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Action to take on entries matching this rule

Format: permit
Permit matching entries
Format: deny
Deny matching entries
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"call": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Call another route-map on match

Format: txt
Route map name
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"continue": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Jump to a different rule in this route-map on a match

Format: u32:1-65535
Rule number
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"description": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Description

Format: txt
Description
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						"match": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								// TODO handle non-string types
								"as_path": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `BGP as-path-list to match

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"extcommunity": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `BGP extended community to match

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"interface": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Interface to use

Format: txt
Interface name
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"local_preference": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Local Preference

Format: u32:0-4294967295
Local Preference
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"metric": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Metric of route to match

Format: u32:1-65535
Route metric
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"origin": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `BGP origin code to match

Format: egp
Exterior gateway protocol origin
Format: igp
Interior gateway protocol origin
Format: incomplete
Incomplete origin
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"peer": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Peer address to match

Format: ipv4
Peer IP address
Format: ipv6
Peer IPv6 address
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"rpki": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Match RPKI validation result

Format: invalid
Match invalid entries
Format: notfound
Match notfound entries
Format: valid
Match valid entries
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"tag": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Route tag to match

Format: u32:1-65535
Route tag
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								"community": schema.SingleNestedAttribute{
									Attributes: map[string]schema.Attribute{
										// TODO handle non-string types
										"community_list": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `BGP community-list to match

`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"exact_match": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Community-list to exactly match

`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},
									},
									// CustomType:          basetypes.MapTypable(nil),
									// Required:            false,
									Optional: true,
									// Computed:            false,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `BGP community-list to match

`,
									// DeprecationMessage:  "",
									// Validators:          []validator.Map(nil),
									// PlanModifiers:       []planmodifier.Map(nil),
									// TODO investigate if node defaults can be handled
									// Default:             defaults.Map(nil),
								},

								"evpn": schema.SingleNestedAttribute{
									Attributes: map[string]schema.Attribute{
										// TODO handle non-string types
										"default_route": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Default EVPN type-5 route

`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"rd": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Route Distinguisher

Format: ASN:NN_OR_IP-ADDRESS:NN
Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"route_type": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Match route-type

Format: macip
mac-ip route
Format: multicast
IMET route
Format: prefix
Prefix route
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"vni": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Virtual Network Identifier

Format: u32:0-16777214
VXLAN virtual network identifier
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},
									},
									// CustomType:          basetypes.MapTypable(nil),
									// Required:            false,
									Optional: true,
									// Computed:            false,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Ethernet Virtual Private Network

`,
									// DeprecationMessage:  "",
									// Validators:          []validator.Map(nil),
									// PlanModifiers:       []planmodifier.Map(nil),
									// TODO investigate if node defaults can be handled
									// Default:             defaults.Map(nil),
								},

								"ip": schema.SingleNestedAttribute{
									Attributes: map[string]schema.Attribute{
										"address": schema.SingleNestedAttribute{
											Attributes: map[string]schema.Attribute{
												// TODO handle non-string types
												"access_list": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `IP access-list to match

Format: u32:1-99
IP standard access list
Format: u32:100-199
IP extended access list
Format: u32:1300-1999
IP standard access list (expanded range)
Format: u32:2000-2699
IP extended access list (expanded range)
`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

												},

												// TODO handle non-string types
												"prefix_list": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `IP prefix-list to match

`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

												},

												// TODO handle non-string types
												"prefix_len": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `IP prefix-length to match (can be used for kernel routes only)

Format: u32:0-32
Prefix length
`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

												},
											},
											// CustomType:          basetypes.MapTypable(nil),
											// Required:            false,
											Optional: true,
											// Computed:            false,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `IP address of route to match

`,
											// DeprecationMessage:  "",
											// Validators:          []validator.Map(nil),
											// PlanModifiers:       []planmodifier.Map(nil),
											// TODO investigate if node defaults can be handled
											// Default:             defaults.Map(nil),
										},

										"nexthop": schema.SingleNestedAttribute{
											Attributes: map[string]schema.Attribute{
												// TODO handle non-string types
												"address": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `IP address to match

Format: ipv4
Nexthop IP address
`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

												},

												// TODO handle non-string types
												"access_list": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `IP access-list to match

Format: u32:1-99
IP standard access list
Format: u32:100-199
IP extended access list
Format: u32:1300-1999
IP standard access list (expanded range)
Format: u32:2000-2699
IP extended access list (expanded range)
`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

												},

												// TODO handle non-string types
												"prefix_len": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `IP prefix-length to match

Format: u32:0-32
Prefix length
`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

												},

												// TODO handle non-string types
												"prefix_list": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `IP prefix-list to match

`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

												},

												// TODO handle non-string types
												"type": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `Match type

Format: blackhole
Blackhole
`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

												},
											},
											// CustomType:          basetypes.MapTypable(nil),
											// Required:            false,
											Optional: true,
											// Computed:            false,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `IP next-hop of route to match

`,
											// DeprecationMessage:  "",
											// Validators:          []validator.Map(nil),
											// PlanModifiers:       []planmodifier.Map(nil),
											// TODO investigate if node defaults can be handled
											// Default:             defaults.Map(nil),
										},

										"route_source": schema.SingleNestedAttribute{
											Attributes: map[string]schema.Attribute{
												// TODO handle non-string types
												"access_list": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `IP access-list to match

Format: u32:1-99
IP standard access list
Format: u32:100-199
IP extended access list
Format: u32:1300-1999
IP standard access list (expanded range)
Format: u32:2000-2699
IP extended access list (expanded range)
`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

												},

												// TODO handle non-string types
												"prefix_list": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `IP prefix-list to match

`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

												},
											},
											// CustomType:          basetypes.MapTypable(nil),
											// Required:            false,
											Optional: true,
											// Computed:            false,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Match advertising source address of route

`,
											// DeprecationMessage:  "",
											// Validators:          []validator.Map(nil),
											// PlanModifiers:       []planmodifier.Map(nil),
											// TODO investigate if node defaults can be handled
											// Default:             defaults.Map(nil),
										},
									},
									// CustomType:          basetypes.MapTypable(nil),
									// Required:            false,
									Optional: true,
									// Computed:            false,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `IP prefix parameters to match

`,
									// DeprecationMessage:  "",
									// Validators:          []validator.Map(nil),
									// PlanModifiers:       []planmodifier.Map(nil),
									// TODO investigate if node defaults can be handled
									// Default:             defaults.Map(nil),
								},

								"ipv6": schema.SingleNestedAttribute{
									Attributes: map[string]schema.Attribute{
										"address": schema.SingleNestedAttribute{
											Attributes: map[string]schema.Attribute{
												// TODO handle non-string types
												"access_list": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `IPv6 access-list to match

Format: txt
IPV6 access list name
`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

												},

												// TODO handle non-string types
												"prefix_list": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `IPv6 prefix-list to match

`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

												},

												// TODO handle non-string types
												"prefix_len": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `IPv6 prefix-length to match (can be used for kernel routes only)

Format: u32:0-128
Prefix length
`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

												},
											},
											// CustomType:          basetypes.MapTypable(nil),
											// Required:            false,
											Optional: true,
											// Computed:            false,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `IPv6 address of route to match

`,
											// DeprecationMessage:  "",
											// Validators:          []validator.Map(nil),
											// PlanModifiers:       []planmodifier.Map(nil),
											// TODO investigate if node defaults can be handled
											// Default:             defaults.Map(nil),
										},

										"nexthop": schema.SingleNestedAttribute{
											Attributes: map[string]schema.Attribute{
												// TODO handle non-string types
												"address": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `IPv6 address of next-hop

Format: ipv6
Nexthop IPv6 address
`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

												},

												// TODO handle non-string types
												"access_list": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `IPv6 access-list to match

Format: txt
IPV6 access list name
`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

												},

												// TODO handle non-string types
												"prefix_list": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `IPv6 prefix-list to match

`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

												},

												// TODO handle non-string types
												"type": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `Match type

Format: blackhole
Blackhole
`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

												},
											},
											// CustomType:          basetypes.MapTypable(nil),
											// Required:            false,
											Optional: true,
											// Computed:            false,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `IPv6 next-hop of route to match

`,
											// DeprecationMessage:  "",
											// Validators:          []validator.Map(nil),
											// PlanModifiers:       []planmodifier.Map(nil),
											// TODO investigate if node defaults can be handled
											// Default:             defaults.Map(nil),
										},
									},
									// CustomType:          basetypes.MapTypable(nil),
									// Required:            false,
									Optional: true,
									// Computed:            false,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `IPv6 prefix parameters to match

`,
									// DeprecationMessage:  "",
									// Validators:          []validator.Map(nil),
									// PlanModifiers:       []planmodifier.Map(nil),
									// TODO investigate if node defaults can be handled
									// Default:             defaults.Map(nil),
								},

								"large_community": schema.SingleNestedAttribute{
									Attributes: map[string]schema.Attribute{
										// TODO handle non-string types
										"large_community_list": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `BGP large-community-list to match

`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},
									},
									// CustomType:          basetypes.MapTypable(nil),
									// Required:            false,
									Optional: true,
									// Computed:            false,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Match BGP large communities

`,
									// DeprecationMessage:  "",
									// Validators:          []validator.Map(nil),
									// PlanModifiers:       []planmodifier.Map(nil),
									// TODO investigate if node defaults can be handled
									// Default:             defaults.Map(nil),
								},
							},
							// CustomType:          basetypes.MapTypable(nil),
							// Required:            false,
							Optional: true,
							// Computed:            false,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Route parameters to match

`,
							// DeprecationMessage:  "",
							// Validators:          []validator.Map(nil),
							// PlanModifiers:       []planmodifier.Map(nil),
							// TODO investigate if node defaults can be handled
							// Default:             defaults.Map(nil),
						},

						"on_match": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								// TODO handle non-string types
								"goto": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Rule number to goto on match

Format: u32:1-65535
Rule number
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"next": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Next sequence number to goto on match

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},
							},
							// CustomType:          basetypes.MapTypable(nil),
							// Required:            false,
							Optional: true,
							// Computed:            false,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Exit policy on matches

`,
							// DeprecationMessage:  "",
							// Validators:          []validator.Map(nil),
							// PlanModifiers:       []planmodifier.Map(nil),
							// TODO investigate if node defaults can be handled
							// Default:             defaults.Map(nil),
						},

						"set": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								// TODO handle non-string types
								"atomic_aggregate": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `BGP atomic aggregate attribute

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"distance": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Locally significant administrative distance

Format: u32:0-255
Distance value
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"ip_next_hop": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Nexthop IP address

Format: ipv4
IP address
Format: unchanged
Set the BGP nexthop address as unchanged
Format: peer-address
Set the BGP nexthop address to the address of the peer
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"local_preference": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `BGP local preference attribute

Format: u32:0-4294967295
Local preference value
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"metric": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Destination routing protocol metric

Format: <+/-metric>
Add or subtract metric
Format: u32:0-4294967295
Metric value
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"metric_type": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Open Shortest Path First (OSPF) external metric-type

Format: type-1
OSPF external type 1 metric
Format: type-2
OSPF external type 2 metric
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"origin": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Border Gateway Protocl (BGP) origin code

Format: igp
Interior gateway protocol origin
Format: egp
Exterior gateway protocol origin
Format: incomplete
Incomplete origin
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"originator_id": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `BGP originator ID attribute

Format: ipv4
Orignator IP address
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"src": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Source address for route

Format: ipv4
IPv4 address
Format: ipv6
IPv6 address
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"table": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Set prefixes to table

Format: u32:1-200
Table value
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"tag": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Tag value for routing protocol

Format: u32:1-65535
Tag value
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"weight": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `BGP weight attribute

Format: u32:0-4294967295
BGP weight
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								"aggregator": schema.SingleNestedAttribute{
									Attributes: map[string]schema.Attribute{
										// TODO handle non-string types
										"as": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `AS number of an aggregation

Format: u32:1-4294967295
Rule number
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"ip": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `IP address of an aggregation

Format: ipv4
IP address
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},
									},
									// CustomType:          basetypes.MapTypable(nil),
									// Required:            false,
									Optional: true,
									// Computed:            false,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `BGP aggregator attribute

`,
									// DeprecationMessage:  "",
									// Validators:          []validator.Map(nil),
									// PlanModifiers:       []planmodifier.Map(nil),
									// TODO investigate if node defaults can be handled
									// Default:             defaults.Map(nil),
								},

								"as_path": schema.SingleNestedAttribute{
									Attributes: map[string]schema.Attribute{
										// TODO handle non-string types
										"exclude": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Remove/exclude from the as-path attribute

Format: u32
AS number
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"prepend": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Prepend to the as-path

Format: u32
AS number
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"prepend_last_as": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Use the last AS-number in the as-path

Format: u32:1-10
Number of times to insert
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},
									},
									// CustomType:          basetypes.MapTypable(nil),
									// Required:            false,
									Optional: true,
									// Computed:            false,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Transform BGP AS_PATH attribute

`,
									// DeprecationMessage:  "",
									// Validators:          []validator.Map(nil),
									// PlanModifiers:       []planmodifier.Map(nil),
									// TODO investigate if node defaults can be handled
									// Default:             defaults.Map(nil),
								},

								"community": schema.SingleNestedAttribute{
									Attributes: map[string]schema.Attribute{
										// TODO handle non-string types
										"add": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Add communities to a prefix

Format: <AS:VAL>
Community number in <0-65535:0-65535> format
Format: local-as
Well-known communities value NO_EXPORT_SUBCONFED 0xFFFFFF03
Format: no-advertise
Well-known communities value NO_ADVERTISE 0xFFFFFF02
Format: no-export
Well-known communities value NO_EXPORT 0xFFFFFF01
Format: internet
Well-known communities value 0
Format: graceful-shutdown
Well-known communities value GRACEFUL_SHUTDOWN 0xFFFF0000
Format: accept-own
Well-known communities value ACCEPT_OWN 0xFFFF0001
Format: route-filter-translated-v4
Well-known communities value ROUTE_FILTER_TRANSLATED_v4 0xFFFF0002
Format: route-filter-v4
Well-known communities value ROUTE_FILTER_v4 0xFFFF0003
Format: route-filter-translated-v6
Well-known communities value ROUTE_FILTER_TRANSLATED_v6 0xFFFF0004
Format: route-filter-v6
Well-known communities value ROUTE_FILTER_v6 0xFFFF0005
Format: llgr-stale
Well-known communities value LLGR_STALE 0xFFFF0006
Format: no-llgr
Well-known communities value NO_LLGR 0xFFFF0007
Format: accept-own-nexthop
Well-known communities value accept-own-nexthop 0xFFFF0008
Format: blackhole
Well-known communities value BLACKHOLE 0xFFFF029A
Format: no-peer
Well-known communities value NOPEER 0xFFFFFF04
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"replace": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Set communities for a prefix

Format: <AS:VAL>
Community number in <0-65535:0-65535> format
Format: local-as
Well-known communities value NO_EXPORT_SUBCONFED 0xFFFFFF03
Format: no-advertise
Well-known communities value NO_ADVERTISE 0xFFFFFF02
Format: no-export
Well-known communities value NO_EXPORT 0xFFFFFF01
Format: internet
Well-known communities value 0
Format: graceful-shutdown
Well-known communities value GRACEFUL_SHUTDOWN 0xFFFF0000
Format: accept-own
Well-known communities value ACCEPT_OWN 0xFFFF0001
Format: route-filter-translated-v4
Well-known communities value ROUTE_FILTER_TRANSLATED_v4 0xFFFF0002
Format: route-filter-v4
Well-known communities value ROUTE_FILTER_v4 0xFFFF0003
Format: route-filter-translated-v6
Well-known communities value ROUTE_FILTER_TRANSLATED_v6 0xFFFF0004
Format: route-filter-v6
Well-known communities value ROUTE_FILTER_v6 0xFFFF0005
Format: llgr-stale
Well-known communities value LLGR_STALE 0xFFFF0006
Format: no-llgr
Well-known communities value NO_LLGR 0xFFFF0007
Format: accept-own-nexthop
Well-known communities value accept-own-nexthop 0xFFFF0008
Format: blackhole
Well-known communities value BLACKHOLE 0xFFFF029A
Format: no-peer
Well-known communities value NOPEER 0xFFFFFF04
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"none": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Completely remove communities attribute from a prefix

`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"delete": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Remove communities defined in a list from a prefix

Format: txt
Community-list
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},
									},
									// CustomType:          basetypes.MapTypable(nil),
									// Required:            false,
									Optional: true,
									// Computed:            false,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `BGP community attribute

`,
									// DeprecationMessage:  "",
									// Validators:          []validator.Map(nil),
									// PlanModifiers:       []planmodifier.Map(nil),
									// TODO investigate if node defaults can be handled
									// Default:             defaults.Map(nil),
								},

								"large_community": schema.SingleNestedAttribute{
									Attributes: map[string]schema.Attribute{
										// TODO handle non-string types
										"add": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Add large communities to a prefix ;

Format: <GA:LDP1:LDP2>
Community in format <0-4294967295:0-4294967295:0-4294967295>
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"replace": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Set large communities for a prefix

Format: <GA:LDP1:LDP2>
Community in format <0-4294967295:0-4294967295:0-4294967295>
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"none": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Completely remove communities attribute from a prefix

`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"delete": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Remove communities defined in a list from a prefix

Format: txt
Community-list
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},
									},
									// CustomType:          basetypes.MapTypable(nil),
									// Required:            false,
									Optional: true,
									// Computed:            false,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `BGP large community attribute

`,
									// DeprecationMessage:  "",
									// Validators:          []validator.Map(nil),
									// PlanModifiers:       []planmodifier.Map(nil),
									// TODO investigate if node defaults can be handled
									// Default:             defaults.Map(nil),
								},

								"extcommunity": schema.SingleNestedAttribute{
									Attributes: map[string]schema.Attribute{
										// TODO handle non-string types
										"bandwidth": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Bandwidth value in Mbps

Format: u32:1-25600
Bandwidth value in Mbps
Format: cumulative
Cumulative bandwidth of all multipaths (outbound-only)
Format: num-multipaths
Internally computed bandwidth based on number of multipaths (outbound-only)
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"bandwidth_non_transitive": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `The link bandwidth extended community is encoded as non-transitive

`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"rt": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Set route target value

Format: ASN:NN
based on autonomous system number in format <0-65535:0-4294967295>
Format: IP:NN
Based on a router-id IP address in format <IP:0-65535>
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"soo": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Set Site of Origin value

Format: ASN:NN
based on autonomous system number in format <0-65535:0-4294967295>
Format: IP:NN
Based on a router-id IP address in format <IP:0-65535>
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"none": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Completely remove communities attribute from a prefix

`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},
									},
									// CustomType:          basetypes.MapTypable(nil),
									// Required:            false,
									Optional: true,
									// Computed:            false,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `BGP extended community attribute

`,
									// DeprecationMessage:  "",
									// Validators:          []validator.Map(nil),
									// PlanModifiers:       []planmodifier.Map(nil),
									// TODO investigate if node defaults can be handled
									// Default:             defaults.Map(nil),
								},

								"evpn": schema.SingleNestedAttribute{
									Attributes: map[string]schema.Attribute{
										"gateway": schema.SingleNestedAttribute{
											Attributes: map[string]schema.Attribute{
												// TODO handle non-string types
												"ipv4": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `Set gateway IPv4 address

Format: ipv4
Gateway IPv4 address
`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

												},

												// TODO handle non-string types
												"ipv6": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `Set gateway IPv6 address

Format: ipv6
Gateway IPv6 address
`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

												},
											},
											// CustomType:          basetypes.MapTypable(nil),
											// Required:            false,
											Optional: true,
											// Computed:            false,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Set gateway IP for prefix advertisement route

`,
											// DeprecationMessage:  "",
											// Validators:          []validator.Map(nil),
											// PlanModifiers:       []planmodifier.Map(nil),
											// TODO investigate if node defaults can be handled
											// Default:             defaults.Map(nil),
										},
									},
									// CustomType:          basetypes.MapTypable(nil),
									// Required:            false,
									Optional: true,
									// Computed:            false,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Ethernet Virtual Private Network

`,
									// DeprecationMessage:  "",
									// Validators:          []validator.Map(nil),
									// PlanModifiers:       []planmodifier.Map(nil),
									// TODO investigate if node defaults can be handled
									// Default:             defaults.Map(nil),
								},

								"ipv6_next_hop": schema.SingleNestedAttribute{
									Attributes: map[string]schema.Attribute{
										// TODO handle non-string types
										"global": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Nexthop IPv6 global address

Format: ipv6
IPv6 address and prefix length
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"local": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Nexthop IPv6 local address

Format: ipv6
IPv6 address and prefix length
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"peer_address": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Use peer address (for BGP only)

`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"prefer_global": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Prefer global address as the nexthop

`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},
									},
									// CustomType:          basetypes.MapTypable(nil),
									// Required:            false,
									Optional: true,
									// Computed:            false,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Nexthop IPv6 address

`,
									// DeprecationMessage:  "",
									// Validators:          []validator.Map(nil),
									// PlanModifiers:       []planmodifier.Map(nil),
									// TODO investigate if node defaults can be handled
									// Default:             defaults.Map(nil),
								},

								"l3vpn_nexthop": schema.SingleNestedAttribute{
									Attributes: map[string]schema.Attribute{
										"encapsulation": schema.SingleNestedAttribute{
											Attributes: map[string]schema.Attribute{
												// TODO handle non-string types
												"gre": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `Accept L3VPN traffic over GRE encapsulation

`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

												},
											},
											// CustomType:          basetypes.MapTypable(nil),
											// Required:            false,
											Optional: true,
											// Computed:            false,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Encapsulation options (for BGP only)

`,
											// DeprecationMessage:  "",
											// Validators:          []validator.Map(nil),
											// PlanModifiers:       []planmodifier.Map(nil),
											// TODO investigate if node defaults can be handled
											// Default:             defaults.Map(nil),
										},
									},
									// CustomType:          basetypes.MapTypable(nil),
									// Required:            false,
									Optional: true,
									// Computed:            false,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Next hop Information

`,
									// DeprecationMessage:  "",
									// Validators:          []validator.Map(nil),
									// PlanModifiers:       []planmodifier.Map(nil),
									// TODO investigate if node defaults can be handled
									// Default:             defaults.Map(nil),
								},
							},
							// CustomType:          basetypes.MapTypable(nil),
							// Required:            false,
							Optional: true,
							// Computed:            false,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Route parameters

`,
							// DeprecationMessage:  "",
							// Validators:          []validator.Map(nil),
							// PlanModifiers:       []planmodifier.Map(nil),
							// TODO investigate if node defaults can be handled
							// Default:             defaults.Map(nil),
						},

						// CustomType:    basetypes.ObjectTypable(nil),
						// Validators:    []validator.Object(nil),
						// PlanModifiers: []planmodifier.Object(nil),
					},
				},
				// CustomType:          basetypes.MapTypable(nil),
				// Required:            false,
				Optional: true,
				// Computed:            false,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Rule for this route-map

Format: u32:1-65535
Route-map rule number
`,
				// DeprecationMessage:  "",
				// Validators:          []validator.Map(nil),
				// PlanModifiers:       []planmodifier.Map(nil),
				// TODO investigate if tagnode defaults can be handled
				// Default:             defaults.Map(nil),
			},

			// TODO handle non-string types
			"description": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Description

Format: txt
Description
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},
		},
	}
}

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func (r *policy_route_map) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *policy_route_mapModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create example, got error: %s", err))
	//     return
	// }

	// For the purposes of this example code, hardcoding a response value to
	// save into the Terraform state.
	data.ID = types.StringValue("example-id")

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "created a resource")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read method to define the logic which refreshes the Terraform state for the resource.
func (r *policy_route_map) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *policy_route_mapModel

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
func (r *policy_route_map) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *policy_route_mapModel

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
func (r *policy_route_map) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *policy_route_mapModel

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
