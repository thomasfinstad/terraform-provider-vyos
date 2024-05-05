// Code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-vyos-infterface-definition-structs/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/vyos/schemadefinition"
)

func protocols_rpki() schemadefinition.InterfaceDefinition {
	return schemadefinition.InterfaceDefinition{
		XMLName: xml.Name{
			Local: "interfaceDefinition",
		},
		Node: []*schemadefinition.Node{{
			IsBaseNode: false,
			XMLName: xml.Name{
				Local: "node",
			},
			NodeNameAttr: "protocols",
			Children: []*schemadefinition.Children{{
				XMLName: xml.Name{
					Local: "children",
				},
				Node: []*schemadefinition.Node{{
					IsBaseNode: false,
					XMLName: xml.Name{
						Local: "node",
					},
					NodeNameAttr: "rpki",
					OwnerAttr:    "${vyos_conf_scripts_dir}/protocols_rpki.py",
					Properties: []*schemadefinition.Properties{{
						XMLName: xml.Name{
							Local: "properties",
						},
						Help:     []string{"Resource Public Key Infrastructure (RPKI)"},
						Priority: []string{"819"},
					}},
					Children: []*schemadefinition.Children{{
						XMLName: xml.Name{
							Local: "children",
						},
						TagNode: []*schemadefinition.TagNode{{
							IsBaseNode: false,
							XMLName: xml.Name{
								Local: "tagNode",
							},
							NodeNameAttr: "cache",
							Properties: []*schemadefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"RPKI cache server address"},
								Constraint: []*schemadefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Validator: []*schemadefinition.Validator{{
										XMLName: xml.Name{
											Local: "validator",
										},
										NameAttr: "ip-address",
									}, {
										XMLName: xml.Name{
											Local: "validator",
										},
										NameAttr: "fqdn",
									}},
								}},
								ValueHelp: []*schemadefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "ipv4",
									Description: "IP address of RPKI server",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "ipv6",
									Description: "IPv6 address of RPKI server",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "hostname",
									Description: "Fully qualified domain name of RPKI server",
								}},
							}},
							Children: []*schemadefinition.Children{{
								XMLName: xml.Name{
									Local: "children",
								},
								Node: []*schemadefinition.Node{{
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "node",
									},
									NodeNameAttr: "ssh",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"RPKI SSH connection settings"},
									}},
									Children: []*schemadefinition.Children{{
										XMLName: xml.Name{
											Local: "children",
										},
										LeafNode: []*schemadefinition.LeafNode{{
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "key",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"OpenSSH key in PKI configuration"},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "txt",
													Description: "Name of OpenSSH key in PKI configuration",
												}},
												CompletionHelp: []*schemadefinition.CompletionHelp{{
													XMLName: xml.Name{
														Local: "completionHelp",
													},
													Path: []string{"pki openssh"},
												}},
											}},
										}, {
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "username",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Username used for authentication"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Regex: []string{"[[:ascii:]]{1,128}"},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "txt",
													Description: "Username",
												}},
												ConstraintErrorMessage: []string{"Username is limited to ASCII characters only, with a total length of 128"},
											}},
										}},
									}},
								}},
								LeafNode: []*schemadefinition.LeafNode{{
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "port",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Port number used by connection"},
										Constraint: []*schemadefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Validator: []*schemadefinition.Validator{{
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr:     "numeric",
												ArgumentAttr: "--range 1-65535",
											}},
										}},
										ValueHelp: []*schemadefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "u32:1-65535",
											Description: "Numeric IP port",
										}},
										ConstraintErrorMessage: []string{"Port number must be in range 1 to 65535"},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "preference",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Preference of the cache server"},
										Constraint: []*schemadefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Validator: []*schemadefinition.Validator{{
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr:     "numeric",
												ArgumentAttr: "--range 1-255",
											}},
										}},
										ValueHelp: []*schemadefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "u32:1-255",
											Description: "Preference of the cache server",
										}},
									}},
								}},
							}},
						}},
						LeafNode: []*schemadefinition.LeafNode{{
							IsBaseNode: false,
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "expire-interval",
							DefaultValue: []string{"7200"},
							Properties: []*schemadefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Interval to wait before expiring the cache"},
								Constraint: []*schemadefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Validator: []*schemadefinition.Validator{{
										XMLName: xml.Name{
											Local: "validator",
										},
										NameAttr:     "numeric",
										ArgumentAttr: "--range 600-172800",
									}},
								}},
								ValueHelp: []*schemadefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "u32:600-172800",
									Description: "Interval in seconds",
								}},
							}},
						}, {
							IsBaseNode: false,
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "polling-period",
							DefaultValue: []string{"300"},
							Properties: []*schemadefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Cache polling interval"},
								Constraint: []*schemadefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Validator: []*schemadefinition.Validator{{
										XMLName: xml.Name{
											Local: "validator",
										},
										NameAttr:     "numeric",
										ArgumentAttr: "--range 1-86400",
									}},
								}},
								ValueHelp: []*schemadefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "u32:1-86400",
									Description: "Interval in seconds",
								}},
							}},
						}, {
							IsBaseNode: false,
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "retry-interval",
							DefaultValue: []string{"600"},
							Properties: []*schemadefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Retry interval to connect to the cache server"},
								Constraint: []*schemadefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Validator: []*schemadefinition.Validator{{
										XMLName: xml.Name{
											Local: "validator",
										},
										NameAttr:     "numeric",
										ArgumentAttr: "--range 1-7200",
									}},
								}},
								ValueHelp: []*schemadefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "u32:1-7200",
									Description: "Interval in seconds",
								}},
							}},
						}},
					}},
				}},
			}},
		}},
	}
}
