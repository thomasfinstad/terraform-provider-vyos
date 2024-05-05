// Code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-vyos-infterface-definition-structs/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/vyos/schemadefinition"
)

func nat64() schemadefinition.InterfaceDefinition {
	return schemadefinition.InterfaceDefinition{
		XMLName: xml.Name{
			Local: "interfaceDefinition",
		},
		Node: []*schemadefinition.Node{{
			IsBaseNode: false,
			XMLName: xml.Name{
				Local: "node",
			},
			NodeNameAttr: "nat64",
			OwnerAttr:    "${vyos_conf_scripts_dir}/nat64.py",
			Properties: []*schemadefinition.Properties{{
				XMLName: xml.Name{
					Local: "properties",
				},
				Help:     []string{"Network Address Translation (NAT64) parameters"},
				Priority: []string{"501"},
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
					NodeNameAttr: "source",
					Properties: []*schemadefinition.Properties{{
						XMLName: xml.Name{
							Local: "properties",
						},
						Help: []string{"IPv6 source to IPv4 destination address translation"},
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
							NodeNameAttr: "rule",
							Properties: []*schemadefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Source NAT64 rule number"},
								Constraint: []*schemadefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Validator: []*schemadefinition.Validator{{
										XMLName: xml.Name{
											Local: "validator",
										},
										NameAttr:     "numeric",
										ArgumentAttr: "--range 1-999999",
									}},
								}},
								ValueHelp: []*schemadefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "u32:1-999999",
									Description: "Number for this rule",
								}},
								ConstraintErrorMessage: []string{"NAT64 rule number must be between 1 and 999999"},
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
									NodeNameAttr: "match",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Match"},
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
											NodeNameAttr: "mark",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Match fwmark value"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr:     "numeric",
														ArgumentAttr: "--range 1-2147483647",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "u32:1-2147483647",
													Description: "Fwmark value to match against",
												}},
											}},
										}},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "node",
									},
									NodeNameAttr: "source",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"IPv6 source prefix options"},
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
											NodeNameAttr: "prefix",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"IPv6 prefix to be translated"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "ipv6-prefix",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "ipv6net",
													Description: "IPv6 prefix",
												}},
											}},
										}},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "node",
									},
									NodeNameAttr: "translation",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Translated IPv4 address options"},
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
											NodeNameAttr: "pool",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Translation IPv4 pool number"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr:     "numeric",
														ArgumentAttr: "--range 1-999999",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "u32:1-999999",
													Description: "Number for this rule",
												}},
												ConstraintErrorMessage: []string{"NAT64 pool number must be between 1 and 999999"},
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
													NodeNameAttr: "protocol",
													Properties: []*schemadefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"Apply translation address to a specfic protocol"},
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
															NodeNameAttr: "tcp",
															Properties: []*schemadefinition.Properties{{
																XMLName: xml.Name{
																	Local: "properties",
																},
																Help: []string{"Transmission Control Protocol"},
																Valueless: []*schemadefinition.Valueless{{
																	XMLName: xml.Name{
																		Local: "valueless",
																	},
																}},
															}},
														}, {
															IsBaseNode: false,
															XMLName: xml.Name{
																Local: "leafNode",
															},
															NodeNameAttr: "udp",
															Properties: []*schemadefinition.Properties{{
																XMLName: xml.Name{
																	Local: "properties",
																},
																Help: []string{"User Datagram Protocol"},
																Valueless: []*schemadefinition.Valueless{{
																	XMLName: xml.Name{
																		Local: "valueless",
																	},
																}},
															}},
														}, {
															IsBaseNode: false,
															XMLName: xml.Name{
																Local: "leafNode",
															},
															NodeNameAttr: "icmp",
															Properties: []*schemadefinition.Properties{{
																XMLName: xml.Name{
																	Local: "properties",
																},
																Help: []string{"Internet Control Message Protocol"},
																Valueless: []*schemadefinition.Valueless{{
																	XMLName: xml.Name{
																		Local: "valueless",
																	},
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
													NodeNameAttr: "description",
													Properties: []*schemadefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"Description"},
														Constraint: []*schemadefinition.Constraint{{
															XMLName: xml.Name{
																Local: "constraint",
															},
															Regex: []string{".{0,255}"},
														}},
														ValueHelp: []*schemadefinition.ValueHelp{{
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "txt",
															Description: "Description",
														}},
														ConstraintErrorMessage: []string{"Description too long (limit 255 characters)"},
													}},
												}, {
													IsBaseNode: false,
													XMLName: xml.Name{
														Local: "leafNode",
													},
													NodeNameAttr: "disable",
													Properties: []*schemadefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"Disable instance"},
														Valueless: []*schemadefinition.Valueless{{
															XMLName: xml.Name{
																Local: "valueless",
															},
														}},
													}},
												}, {
													IsBaseNode: false,
													XMLName: xml.Name{
														Local: "leafNode",
													},
													NodeNameAttr: "port",
													Properties: []*schemadefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"Port number"},
														Constraint: []*schemadefinition.Constraint{{
															XMLName: xml.Name{
																Local: "constraint",
															},
															Validator: []*schemadefinition.Validator{{
																XMLName: xml.Name{
																	Local: "validator",
																},
																NameAttr: "port-range",
															}},
														}},
														ValueHelp: []*schemadefinition.ValueHelp{{
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "u32:1-65535",
															Description: "Numeric IP port",
														}, {
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "range",
															Description: "Numbered port range (e.g., 1001-1005)",
														}},
													}},
												}, {
													IsBaseNode: false,
													XMLName: xml.Name{
														Local: "leafNode",
													},
													NodeNameAttr: "address",
													Properties: []*schemadefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"IPv4 address or prefix to translate to"},
														Constraint: []*schemadefinition.Constraint{{
															XMLName: xml.Name{
																Local: "constraint",
															},
															Validator: []*schemadefinition.Validator{{
																XMLName: xml.Name{
																	Local: "validator",
																},
																NameAttr: "ipv4-address",
															}, {
																XMLName: xml.Name{
																	Local: "validator",
																},
																NameAttr: "ipv4-prefix",
															}},
														}},
														ValueHelp: []*schemadefinition.ValueHelp{{
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "ipv4",
															Description: "IPv4 address",
														}, {
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "ipv4net",
															Description: "IPv4 prefix",
														}},
													}},
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
									NodeNameAttr: "description",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Description"},
										Constraint: []*schemadefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Regex: []string{".{0,255}"},
										}},
										ValueHelp: []*schemadefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "txt",
											Description: "Description",
										}},
										ConstraintErrorMessage: []string{"Description too long (limit 255 characters)"},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "disable",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Disable instance"},
										Valueless: []*schemadefinition.Valueless{{
											XMLName: xml.Name{
												Local: "valueless",
											},
										}},
									}},
								}},
							}},
						}},
					}},
				}},
			}},
		}},
	}
}
