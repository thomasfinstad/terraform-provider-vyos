// Code generated by /workspaces/terraform-provider-vyos/tools/build-vyos-infterface-definition-structs/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/vyos/schema/interfacedefinition"
)

func protocolsfailover() interfacedefinition.InterfaceDefinition {
	return interfacedefinition.InterfaceDefinition{
		XMLName: xml.Name{
			Local: "interfaceDefinition",
		},
		Node: []*interfacedefinition.Node{{
			IsBaseNode: false,
			XMLName: xml.Name{
				Local: "node",
			},
			NodeNameAttr: "protocols",
			Children: []*interfacedefinition.Children{{
				XMLName: xml.Name{
					Local: "children",
				},
				Node: []*interfacedefinition.Node{{
					IsBaseNode: false,
					XMLName: xml.Name{
						Local: "node",
					},
					NodeNameAttr: "failover",
					OwnerAttr:    "${vyos_conf_scripts_dir}/protocols_failover.py",
					Properties: []*interfacedefinition.Properties{{
						XMLName: xml.Name{
							Local: "properties",
						},
						Help:     []string{"Failover Routing"},
						Priority: []string{"490"},
					}},
					Children: []*interfacedefinition.Children{{
						XMLName: xml.Name{
							Local: "children",
						},
						TagNode: []*interfacedefinition.TagNode{{
							IsBaseNode: false,
							XMLName: xml.Name{
								Local: "tagNode",
							},
							NodeNameAttr: "route",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Failover IPv4 route"},
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
									Description: "IPv4 failover route",
								}},
							}},
							Children: []*interfacedefinition.Children{{
								XMLName: xml.Name{
									Local: "children",
								},
								TagNode: []*interfacedefinition.TagNode{{
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "tagNode",
									},
									NodeNameAttr: "next-hop",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Next-hop IPv4 router address"},
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
											Description: "Next-hop router address",
										}},
									}},
									Children: []*interfacedefinition.Children{{
										XMLName: xml.Name{
											Local: "children",
										},
										Node: []*interfacedefinition.Node{{
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "node",
											},
											NodeNameAttr: "check",
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Check target options"},
											}},
											Children: []*interfacedefinition.Children{{
												XMLName: xml.Name{
													Local: "children",
												},
												LeafNode: []*interfacedefinition.LeafNode{{
													IsBaseNode: false,
													XMLName: xml.Name{
														Local: "leafNode",
													},
													NodeNameAttr: "port",
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
												}, {
													IsBaseNode: false,
													XMLName: xml.Name{
														Local: "leafNode",
													},
													NodeNameAttr: "target",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"Check target address"},
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
															Description: "Address to check",
														}},
													}},
												}, {
													IsBaseNode: false,
													XMLName: xml.Name{
														Local: "leafNode",
													},
													NodeNameAttr: "timeout",
													DefaultValue: []string{"10"},
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"Timeout between checks"},
														Constraint: []*interfacedefinition.Constraint{{
															XMLName: xml.Name{
																Local: "constraint",
															},
															Validator: []*interfacedefinition.Validator{{
																XMLName: xml.Name{
																	Local: "validator",
																},
																NameAttr:     "numeric",
																ArgumentAttr: "--range 1-255",
															}},
														}},
														ValueHelp: []*interfacedefinition.ValueHelp{{
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "u32:1-300",
															Description: "Timeout in seconds between checks",
														}},
													}},
												}, {
													IsBaseNode: false,
													XMLName: xml.Name{
														Local: "leafNode",
													},
													NodeNameAttr: "type",
													DefaultValue: []string{"icmp"},
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"Check type"},
														Constraint: []*interfacedefinition.Constraint{{
															XMLName: xml.Name{
																Local: "constraint",
															},
															Regex: []string{"(arp|icmp|tcp)"},
														}},
														ValueHelp: []*interfacedefinition.ValueHelp{{
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "arp",
															Description: "Check target by ARP",
														}, {
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "icmp",
															Description: "Check target by ICMP",
														}, {
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "tcp",
															Description: "Check target by TCP",
														}},
														CompletionHelp: []*interfacedefinition.CompletionHelp{{
															XMLName: xml.Name{
																Local: "completionHelp",
															},
															List: []string{"arp icmp tcp"},
														}},
													}},
												}},
											}},
										}},
										LeafNode: []*interfacedefinition.LeafNode{{
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "interface",
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Gateway interface name"},
												Constraint: []*interfacedefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Regex: []string{"(bond|br|dum|en|ersp|eth|gnv|ifb|lan|l2tp|l2tpeth|macsec|peth|ppp|pppoe|pptp|sstp|tun|veth|vti|vtun|vxlan|wg|wlan|wwan)[0-9]+(.\\d+)?|lo"},
													Validator: []*interfacedefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "file-path --lookup-path /sys/class/net --directory",
													}},
												}},
												ValueHelp: []*interfacedefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "txt",
													Description: "Gateway interface name",
												}},
												CompletionHelp: []*interfacedefinition.CompletionHelp{{
													XMLName: xml.Name{
														Local: "completionHelp",
													},
													Script: []string{"${vyos_completion_dir}/list_interfaces"},
												}},
											}},
										}, {
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "metric",
											DefaultValue: []string{"1"},
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Route metric for this gateway"},
												Constraint: []*interfacedefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*interfacedefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr:     "numeric",
														ArgumentAttr: "--range 1-255",
													}},
												}},
												ValueHelp: []*interfacedefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "u32:1-255",
													Description: "Route metric",
												}},
											}},
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
