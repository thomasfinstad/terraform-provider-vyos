// Code generated by /workspaces/terraform-provider-vyos/tools/build-vyos-infterface-definition-structs/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/vyos/schema/interfacedefinition"
)

func dhcprelay() interfacedefinition.InterfaceDefinition {
	return interfacedefinition.InterfaceDefinition{
		XMLName: xml.Name{
			Local: "interfaceDefinition",
		},
		Node: []*interfacedefinition.Node{{
			XMLName: xml.Name{
				Local: "node",
			},
			NodeNameAttr: "service",
			Children: []*interfacedefinition.Children{{
				XMLName: xml.Name{
					Local: "children",
				},
				Node: []*interfacedefinition.Node{{
					XMLName: xml.Name{
						Local: "node",
					},
					NodeNameAttr: "dhcp-relay",
					OwnerAttr:    "${vyos_conf_scripts_dir}/dhcp_relay.py",
					Properties: []*interfacedefinition.Properties{{
						XMLName: xml.Name{
							Local: "properties",
						},
						Help:     []string{"Host Configuration Protocol (DHCP) relay agent"},
						Priority: []string{"910"},
					}},
					Children: []*interfacedefinition.Children{{
						XMLName: xml.Name{
							Local: "children",
						},
						Node: []*interfacedefinition.Node{{
							XMLName: xml.Name{
								Local: "node",
							},
							NodeNameAttr: "relay-options",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Relay options"},
							}},
							Children: []*interfacedefinition.Children{{
								XMLName: xml.Name{
									Local: "children",
								},
								LeafNode: []*interfacedefinition.LeafNode{{
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "hop-count",
									DefaultValue: []string{"10"},
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Policy to discard packets that have reached specified hop-count"},
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
											Description: "Hop count",
										}},
										ConstraintErrorMessage: []string{"hop-count must be a value between 1 and 255"},
									}},
								}, {
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "max-size",
									DefaultValue: []string{"576"},
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Maximum packet size to send to a DHCPv4/BOOTP server"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Validator: []*interfacedefinition.Validator{{
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr:     "numeric",
												ArgumentAttr: "--range 64-1400",
											}},
										}},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "u32:64-1400",
											Description: "Maximum packet size",
										}},
										ConstraintErrorMessage: []string{"max-size must be a value between 64 and 1400"},
									}},
								}, {
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "relay-agents-packets",
									DefaultValue: []string{"forward"},
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Policy to handle incoming DHCPv4 packets which already contain relay agent options"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Regex: []string{"(append|replace|forward|discard)"},
										}},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "append",
											Description: "append own relay options to packet",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "replace",
											Description: "replace existing agent option field",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "forward",
											Description: "forward packet unchanged",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "discard",
											Description: "discard packet (default action if giaddr not set in packet)",
										}},
										CompletionHelp: []*interfacedefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
											List: []string{"append replace forward discard"},
										}},
									}},
								}},
							}},
						}},
						LeafNode: []*interfacedefinition.LeafNode{{
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "interface",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Interface Name to use"},
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
									Description: "Interface name",
								}},
								CompletionHelp: []*interfacedefinition.CompletionHelp{{
									XMLName: xml.Name{
										Local: "completionHelp",
									},
									Script: []string{"${vyos_completion_dir}/list_interfaces --broadcast"},
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
							NodeNameAttr: "listen-interface",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Interface for DHCP Relay Agent to listen for requests"},
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
									Description: "Interface name",
								}},
								CompletionHelp: []*interfacedefinition.CompletionHelp{{
									XMLName: xml.Name{
										Local: "completionHelp",
									},
									Script: []string{"${vyos_completion_dir}/list_interfaces"},
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
							NodeNameAttr: "upstream-interface",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Interface for DHCP Relay Agent forward requests out"},
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
									Description: "Interface name",
								}},
								CompletionHelp: []*interfacedefinition.CompletionHelp{{
									XMLName: xml.Name{
										Local: "completionHelp",
									},
									Script: []string{"${vyos_completion_dir}/list_interfaces"},
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
							NodeNameAttr: "server",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"DHCP server address"},
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
									Description: "DHCP server IPv4 address",
								}},
								Multi: []*interfacedefinition.Multi{{
									XMLName: xml.Name{
										Local: "multi",
									},
								}},
							}},
						}},
					}},
				}},
			}},
		}},
	}
}
