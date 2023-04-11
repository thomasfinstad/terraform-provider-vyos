// Code generated by /workspaces/terraform-provider-vyos/tools/build-vyos-infterface-definition-structs/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/vyos/schema/interfacedefinition"
)

func vpnsstp() interfacedefinition.InterfaceDefinition {
	return interfacedefinition.InterfaceDefinition{
		XMLName: xml.Name{
			Local: "interfaceDefinition",
		},
		Node: []*interfacedefinition.Node{{
			XMLName: xml.Name{
				Local: "node",
			},
			NodeNameAttr: "vpn",
			Children: []*interfacedefinition.Children{{
				XMLName: xml.Name{
					Local: "children",
				},
				Node: []*interfacedefinition.Node{{
					XMLName: xml.Name{
						Local: "node",
					},
					NodeNameAttr: "sstp",
					OwnerAttr:    "${vyos_conf_scripts_dir}/vpn_sstp.py",
					Properties: []*interfacedefinition.Properties{{
						XMLName: xml.Name{
							Local: "properties",
						},
						Help:     []string{"Secure Socket Tunneling Protocol (SSTP) server"},
						Priority: []string{"901"},
					}},
					Children: []*interfacedefinition.Children{{
						XMLName: xml.Name{
							Local: "children",
						},
						Node: []*interfacedefinition.Node{{
							XMLName: xml.Name{
								Local: "node",
							},
							NodeNameAttr: "authentication",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Authentication for remote access SSTP Server"},
							}},
							Children: []*interfacedefinition.Children{{
								XMLName: xml.Name{
									Local: "children",
								},
								Node: []*interfacedefinition.Node{{
									XMLName: xml.Name{
										Local: "node",
									},
									NodeNameAttr: "local-users",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Local user authentication for PPPoE server"},
									}},
									Children: []*interfacedefinition.Children{{
										XMLName: xml.Name{
											Local: "children",
										},
										TagNode: []*interfacedefinition.TagNode{{
											XMLName: xml.Name{
												Local: "tagNode",
											},
											NodeNameAttr: "username",
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"User name for authentication"},
											}},
											Children: []*interfacedefinition.Children{{
												XMLName: xml.Name{
													Local: "children",
												},
												Node: []*interfacedefinition.Node{{
													XMLName: xml.Name{
														Local: "node",
													},
													NodeNameAttr: "rate-limit",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"Upload/Download speed limits"},
													}},
													Children: []*interfacedefinition.Children{{
														XMLName: xml.Name{
															Local: "children",
														},
														LeafNode: []*interfacedefinition.LeafNode{{
															XMLName: xml.Name{
																Local: "leafNode",
															},
															NodeNameAttr: "upload",
															Properties: []*interfacedefinition.Properties{{
																XMLName: xml.Name{
																	Local: "properties",
																},
																Help: []string{"Upload bandwidth limit in kbits/sec"},
																Constraint: []*interfacedefinition.Constraint{{
																	XMLName: xml.Name{
																		Local: "constraint",
																	},
																	Validator: []*interfacedefinition.Validator{{
																		XMLName: xml.Name{
																			Local: "validator",
																		},
																		NameAttr:     "numeric",
																		ArgumentAttr: "--range 1-10000000",
																	}},
																}},
															}},
														}, {
															XMLName: xml.Name{
																Local: "leafNode",
															},
															NodeNameAttr: "download",
															Properties: []*interfacedefinition.Properties{{
																XMLName: xml.Name{
																	Local: "properties",
																},
																Help: []string{"Download bandwidth limit in kbits/sec"},
																Constraint: []*interfacedefinition.Constraint{{
																	XMLName: xml.Name{
																		Local: "constraint",
																	},
																	Validator: []*interfacedefinition.Validator{{
																		XMLName: xml.Name{
																			Local: "validator",
																		},
																		NameAttr:     "numeric",
																		ArgumentAttr: "--range 1-10000000",
																	}},
																}},
															}},
														}},
													}},
												}},
												LeafNode: []*interfacedefinition.LeafNode{{
													XMLName: xml.Name{
														Local: "leafNode",
													},
													NodeNameAttr: "disable",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"Disable instance"},
														Valueless: []*interfacedefinition.Valueless{{
															XMLName: xml.Name{
																Local: "valueless",
															},
														}},
													}},
												}, {
													XMLName: xml.Name{
														Local: "leafNode",
													},
													NodeNameAttr: "password",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"Password for authentication"},
													}},
												}, {
													XMLName: xml.Name{
														Local: "leafNode",
													},
													NodeNameAttr: "static-ip",
													DefaultValue: []string{"&"},
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"Static client IP address"},
														Constraint: []*interfacedefinition.Constraint{{
															XMLName: xml.Name{
																Local: "constraint",
															},
															Validator: []*interfacedefinition.Validator{{
																XMLName: xml.Name{
																	Local: "validator",
																},
																NameAttr: "ipv4-address",
															}},
														}},
													}},
												}},
											}},
										}},
									}},
								}, {
									XMLName: xml.Name{
										Local: "node",
									},
									NodeNameAttr: "radius",
									Children: []*interfacedefinition.Children{{
										XMLName: xml.Name{
											Local: "children",
										},
										Node: []*interfacedefinition.Node{{
											XMLName: xml.Name{
												Local: "node",
											},
											NodeNameAttr: "rate-limit",
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Upload/Download speed limits"},
											}},
											Children: []*interfacedefinition.Children{{
												XMLName: xml.Name{
													Local: "children",
												},
												LeafNode: []*interfacedefinition.LeafNode{{
													XMLName: xml.Name{
														Local: "leafNode",
													},
													NodeNameAttr: "attribute",
													DefaultValue: []string{"Filter-Id"},
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"RADIUS attribute that contains rate information"},
													}},
												}, {
													XMLName: xml.Name{
														Local: "leafNode",
													},
													NodeNameAttr: "vendor",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"Vendor dictionary"},
														Constraint: []*interfacedefinition.Constraint{{
															XMLName: xml.Name{
																Local: "constraint",
															},
															Validator: []*interfacedefinition.Validator{{
																XMLName: xml.Name{
																	Local: "validator",
																},
																NameAttr: "accel-radius-dictionary",
															}},
														}},
														CompletionHelp: []*interfacedefinition.CompletionHelp{{
															XMLName: xml.Name{
																Local: "completionHelp",
															},
															List: []string{"alcatel cisco microsoft mikrotik"},
														}},
													}},
												}, {
													XMLName: xml.Name{
														Local: "leafNode",
													},
													NodeNameAttr: "enable",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"Enable bandwidth shaping via RADIUS"},
														Valueless: []*interfacedefinition.Valueless{{
															XMLName: xml.Name{
																Local: "valueless",
															},
														}},
													}},
												}, {
													XMLName: xml.Name{
														Local: "leafNode",
													},
													NodeNameAttr: "multiplier",
													DefaultValue: []string{"1"},
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"Shaper multiplier"},
														Constraint: []*interfacedefinition.Constraint{{
															XMLName: xml.Name{
																Local: "constraint",
															},
															Validator: []*interfacedefinition.Validator{{
																XMLName: xml.Name{
																	Local: "validator",
																},
																NameAttr:     "numeric",
																ArgumentAttr: "--range 0.001-1000 --float",
															}},
														}},
														ValueHelp: []*interfacedefinition.ValueHelp{{
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "<0.001-1000>",
															Description: "Shaper multiplier",
														}},
														ConstraintErrorMessage: []string{"Multiplier needs to be between 0.001 and 1000"},
													}},
												}},
											}},
										}},
									}},
								}},
								LeafNode: []*interfacedefinition.LeafNode{{
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "mode",
									DefaultValue: []string{"local"},
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Authentication mode used by this server"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Regex: []string{"(local|radius|noauth)"},
										}},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "local",
											Description: "Use local username/password configuration",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "radius",
											Description: "Use RADIUS server for user autentication",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "noauth",
											Description: "Authentication disabled",
										}},
										CompletionHelp: []*interfacedefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
											List: []string{"local radius noauth"},
										}},
									}},
								}, {
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "protocols",
									DefaultValue: []string{"pap chap mschap mschap-v2"},
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Authentication protocol for remote access peer SSTP VPN"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Regex: []string{"(pap|chap|mschap|mschap-v2)"},
										}},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "pap",
											Description: "Authentication via PAP (Password Authentication Protocol)",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "chap",
											Description: "Authentication via CHAP (Challenge Handshake Authentication Protocol)",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "mschap",
											Description: "Authentication via MS-CHAP (Microsoft Challenge Handshake Authentication Protocol)",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "mschap-v2",
											Description: "Authentication via MS-CHAPv2 (Microsoft Challenge Handshake Authentication Protocol, version 2)",
										}},
										CompletionHelp: []*interfacedefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
											List: []string{"pap chap mschap mschap-v2"},
										}},
										Multi: []*interfacedefinition.Multi{{
											XMLName: xml.Name{
												Local: "multi",
											},
										}},
									}},
								}},
							}},
						}, {
							XMLName: xml.Name{
								Local: "node",
							},
							NodeNameAttr: "client-ip-pool",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Client IP pools and gateway setting"},
							}},
							Children: []*interfacedefinition.Children{{
								XMLName: xml.Name{
									Local: "children",
								},
								LeafNode: []*interfacedefinition.LeafNode{{
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "subnet",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Client IP subnet (CIDR notation)"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Validator: []*interfacedefinition.Validator{{
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr: "ipv4-prefix",
											}},
										}},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "ipv4net",
											Description: "IPv4 address and prefix length",
										}},
										ConstraintErrorMessage: []string{"Not a valid CIDR formatted prefix"},
										Multi: []*interfacedefinition.Multi{{
											XMLName: xml.Name{
												Local: "multi",
											},
										}},
									}},
								}},
							}},
						}, {
							XMLName: xml.Name{
								Local: "node",
							},
							NodeNameAttr: "client-ipv6-pool",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Pool of client IPv6 addresses"},
							}},
							Children: []*interfacedefinition.Children{{
								XMLName: xml.Name{
									Local: "children",
								},
								TagNode: []*interfacedefinition.TagNode{{
									XMLName: xml.Name{
										Local: "tagNode",
									},
									NodeNameAttr: "prefix",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Pool of addresses used to assign to clients"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Validator: []*interfacedefinition.Validator{{
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr: "ipv6-prefix",
											}},
										}},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "ipv6net",
											Description: "IPv6 address and prefix length",
										}},
									}},
									Children: []*interfacedefinition.Children{{
										XMLName: xml.Name{
											Local: "children",
										},
										LeafNode: []*interfacedefinition.LeafNode{{
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "mask",
											DefaultValue: []string{"64"},
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Prefix length used for individual client"},
												Constraint: []*interfacedefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*interfacedefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr:     "numeric",
														ArgumentAttr: "--range 48-128",
													}},
												}},
												ValueHelp: []*interfacedefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "u32:48-128",
													Description: "Client prefix length",
												}},
											}},
										}},
									}},
								}, {
									XMLName: xml.Name{
										Local: "tagNode",
									},
									NodeNameAttr: "delegate",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Subnet used to delegate prefix through DHCPv6-PD (RFC3633)"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Validator: []*interfacedefinition.Validator{{
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr: "ipv6-prefix",
											}},
										}},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "ipv6net",
											Description: "IPv6 address and prefix length",
										}},
									}},
									Children: []*interfacedefinition.Children{{
										XMLName: xml.Name{
											Local: "children",
										},
										LeafNode: []*interfacedefinition.LeafNode{{
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "delegation-prefix",
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Prefix length delegated to client"},
												Constraint: []*interfacedefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*interfacedefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr:     "numeric",
														ArgumentAttr: "--range 32-64",
													}},
												}},
												ValueHelp: []*interfacedefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "u32:32-64",
													Description: "Delegated prefix length",
												}},
											}},
										}},
									}},
								}},
							}},
						}, {
							XMLName: xml.Name{
								Local: "node",
							},
							NodeNameAttr: "ppp-options",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"PPP (Point-to-Point Protocol) settings"},
							}},
							Children: []*interfacedefinition.Children{{
								XMLName: xml.Name{
									Local: "children",
								},
								LeafNode: []*interfacedefinition.LeafNode{{
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "mppe",
									DefaultValue: []string{"prefer"},
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Specifies mppe negotiation preferences"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Regex: []string{"(require|prefer|deny)"},
										}},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "require",
											Description: "send mppe request, if client rejects, drop the connection",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "prefer",
											Description: "send mppe request, if client rejects continue",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "deny",
											Description: "drop all mppe",
										}},
										CompletionHelp: []*interfacedefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
											List: []string{"require prefer deny"},
										}},
									}},
								}, {
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "ipv4",
									DefaultValue: []string{"allow"},
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"IPv4 negotiation algorithm"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Regex: []string{"(deny|allow)"},
										}},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "deny",
											Description: "Do not negotiate IPv4",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "allow",
											Description: "Negotiate IPv4 only if client requests",
										}},
										ConstraintErrorMessage: []string{"invalid value"},
										CompletionHelp: []*interfacedefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
											List: []string{"deny allow"},
										}},
									}},
								}, {
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "ipv6",
									DefaultValue: []string{"deny"},
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"IPv6 (IPCP6) negotiation algorithm"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Regex: []string{"(deny|allow|prefer|require)"},
										}},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "deny",
											Description: "Do not negotiate IPv6",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "allow",
											Description: "Negotiate IPv6 only if client requests",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "prefer",
											Description: "Ask client for IPv6 negotiation, do not fail if it rejects",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "require",
											Description: "Require IPv6 negotiation",
										}},
										ConstraintErrorMessage: []string{"invalid value"},
										CompletionHelp: []*interfacedefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
											List: []string{"deny allow prefer require"},
										}},
									}},
								}, {
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "lcp-echo-interval",
									DefaultValue: []string{"30"},
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"LCP echo-requests/sec"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Validator: []*interfacedefinition.Validator{{
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr:     "numeric",
												ArgumentAttr: "--positive",
											}},
										}},
									}},
								}, {
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "lcp-echo-failure",
									DefaultValue: []string{"3"},
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Maximum number of Echo-Requests may be sent without valid reply"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Validator: []*interfacedefinition.Validator{{
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr:     "numeric",
												ArgumentAttr: "--positive",
											}},
										}},
									}},
								}, {
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "lcp-echo-timeout",
									DefaultValue: []string{"0"},
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Timeout in seconds to wait for any peer activity. If this option specified it turns on adaptive lcp echo functionality and \"lcp-echo-failure\" is not used."},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Validator: []*interfacedefinition.Validator{{
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr:     "numeric",
												ArgumentAttr: "--positive",
											}},
										}},
									}},
								}},
							}},
						}, {
							XMLName: xml.Name{
								Local: "node",
							},
							NodeNameAttr: "ssl",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"SSL Certificate, SSL Key and CA"},
							}},
							Children: []*interfacedefinition.Children{{
								XMLName: xml.Name{
									Local: "children",
								},
								LeafNode: []*interfacedefinition.LeafNode{{
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "ca-certificate",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Certificate Authority in PKI configuration"},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "txt",
											Description: "Name of CA in PKI configuration",
										}},
										CompletionHelp: []*interfacedefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
											Path: []string{"pki ca"},
										}},
									}},
								}, {
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "certificate",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Certificate in PKI configuration"},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "txt",
											Description: "Name of certificate in PKI configuration",
										}},
										CompletionHelp: []*interfacedefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
											Path: []string{"pki certificate"},
										}},
									}},
								}},
							}},
						}},
						LeafNode: []*interfacedefinition.LeafNode{{
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "mtu",
							DefaultValue: []string{"1500"},
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Maximum Transmission Unit (MTU)"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Validator: []*interfacedefinition.Validator{{
										XMLName: xml.Name{
											Local: "validator",
										},
										NameAttr:     "numeric",
										ArgumentAttr: "--range 68-1500",
									}},
								}},
								ValueHelp: []*interfacedefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "u32:68-1500",
									Description: "Maximum Transmission Unit in byte",
								}},
								ConstraintErrorMessage: []string{"MTU must be between 68 and 1500"},
							}},
						}, {
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "gateway-address",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Gateway IP address"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Validator: []*interfacedefinition.Validator{{
										XMLName: xml.Name{
											Local: "validator",
										},
										NameAttr: "ipv4-address",
									}},
								}},
								ValueHelp: []*interfacedefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "ipv4",
									Description: "Default Gateway send to the client",
								}},
								ConstraintErrorMessage: []string{"invalid IPv4 address"},
							}},
						}, {
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "name-server",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Domain Name Servers (DNS) addresses"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Validator: []*interfacedefinition.Validator{{
										XMLName: xml.Name{
											Local: "validator",
										},
										NameAttr: "ipv4-address",
									}, {
										XMLName: xml.Name{
											Local: "validator",
										},
										NameAttr: "ipv6-address",
									}},
								}},
								ValueHelp: []*interfacedefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "ipv4",
									Description: "Domain Name Server (DNS) IPv4 address",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "ipv6",
									Description: "Domain Name Server (DNS) IPv6 address",
								}},
								Multi: []*interfacedefinition.Multi{{
									XMLName: xml.Name{
										Local: "multi",
									},
								}},
							}},
						}, {
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "port",
							DefaultValue: []string{"443"},
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Port number used by connection"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Validator: []*interfacedefinition.Validator{{
										XMLName: xml.Name{
											Local: "validator",
										},
										NameAttr:     "numeric",
										ArgumentAttr: "--range 1-65535",
									}},
								}},
								ValueHelp: []*interfacedefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "u32:1-65535",
									Description: "Numeric IP port",
								}},
								ConstraintErrorMessage: []string{"Port number must be in range 1 to 65535"},
							}},
						}},
					}},
				}},
			}},
		}},
	}
}
