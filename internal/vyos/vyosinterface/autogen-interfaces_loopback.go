// Code generated by /workspaces/terraform-provider-vyos/tools/build-vyos-infterface-definition-structs/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/vyos/schema/interfacedefinition"
)

func interfaces_loopback() interfacedefinition.InterfaceDefinition {
	return interfacedefinition.InterfaceDefinition{
		XMLName: xml.Name{
			Local: "interfaceDefinition",
		},
		Node: []*interfacedefinition.Node{{
			IsBaseNode: false,
			XMLName: xml.Name{
				Local: "node",
			},
			NodeNameAttr: "interfaces",
			Children: []*interfacedefinition.Children{{
				XMLName: xml.Name{
					Local: "children",
				},
				TagNode: []*interfacedefinition.TagNode{{
					IsBaseNode: false,
					XMLName: xml.Name{
						Local: "tagNode",
					},
					NodeNameAttr: "loopback",
					OwnerAttr:    "${vyos_conf_scripts_dir}/interfaces_loopback.py",
					Properties: []*interfacedefinition.Properties{{
						XMLName: xml.Name{
							Local: "properties",
						},
						Help: []string{"Loopback Interface"},
						Constraint: []*interfacedefinition.Constraint{{
							XMLName: xml.Name{
								Local: "constraint",
							},
							Regex: []string{"lo"},
						}},
						ValueHelp: []*interfacedefinition.ValueHelp{{
							XMLName: xml.Name{
								Local: "valueHelp",
							},
							Format:      "lo",
							Description: "Loopback interface",
						}},
						ConstraintErrorMessage: []string{"Loopback interface must be named lo"},
						Priority:               []string{"300"},
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
							NodeNameAttr: "ip",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"IPv4 routing parameters"},
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
									NodeNameAttr: "source-validation",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Source validation by reversed path (RFC3704)"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Regex: []string{"(strict|loose|disable)"},
										}},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "strict",
											Description: "Enable Strict Reverse Path Forwarding as defined in RFC3704",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "loose",
											Description: "Enable Loose Reverse Path Forwarding as defined in RFC3704",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "disable",
											Description: "No source validation",
										}},
										CompletionHelp: []*interfacedefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
											List: []string{"strict loose disable"},
										}},
									}},
								}},
							}},
						}, {
							IsBaseNode: false,
							XMLName: xml.Name{
								Local: "node",
							},
							NodeNameAttr: "mirror",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Mirror ingress/egress packets"},
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
									NodeNameAttr: "ingress",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Mirror ingress traffic to destination interface"},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "txt",
											Description: "Destination interface name",
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
									NodeNameAttr: "egress",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Mirror egress traffic to destination interface"},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "txt",
											Description: "Destination interface name",
										}},
										CompletionHelp: []*interfacedefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
											Script: []string{"${vyos_completion_dir}/list_interfaces"},
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
							NodeNameAttr: "address",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"IP address"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Validator: []*interfacedefinition.Validator{{
										XMLName: xml.Name{
											Local: "validator",
										},
										NameAttr: "ip-host",
									}},
								}},
								ValueHelp: []*interfacedefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "ipv4net",
									Description: "IPv4 address and prefix length",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "ipv6net",
									Description: "IPv6 address and prefix length",
								}},
								Multi: []*interfacedefinition.Multi{{
									XMLName: xml.Name{
										Local: "multi",
									},
								}},
							}},
						}, {
							IsBaseNode: false,
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "description",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Description"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Regex: []string{"[[:ascii:]]{0,256}"},
								}},
								ValueHelp: []*interfacedefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "txt",
									Description: "Description",
								}},
								ConstraintErrorMessage: []string{"Description too long (limit 256 characters)"},
							}},
						}, {
							IsBaseNode: false,
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "redirect",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Redirect incoming packet to destination"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Regex: []string{"(bond|br|dum|en|ersp|eth|gnv|ifb|ipoe|lan|l2tp|l2tpeth|macsec|peth|ppp|pppoe|pptp|sstp|sstpc|tun|veth|vti|vtun|vxlan|wg|wlan|wwan)[0-9]+(.\\d+)?|lo"},
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
									Description: "Destination interface name",
								}},
								CompletionHelp: []*interfacedefinition.CompletionHelp{{
									XMLName: xml.Name{
										Local: "completionHelp",
									},
									Script: []string{"${vyos_completion_dir}/list_interfaces"},
								}},
							}},
						}},
					}},
				}},
			}},
		}},
	}
}