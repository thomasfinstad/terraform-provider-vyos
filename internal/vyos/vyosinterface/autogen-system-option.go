// Code generated by /workspaces/terraform-provider-vyos/tools/build-vyos-infterface-definition-structs/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/vyos/schema/interfacedefinition"
)

func systemoption() interfacedefinition.InterfaceDefinition {
	return interfacedefinition.InterfaceDefinition{
		XMLName: xml.Name{
			Local: "interfaceDefinition",
		},
		Node: []*interfacedefinition.Node{{
			XMLName: xml.Name{
				Local: "node",
			},
			NodeNameAttr: "system",
			Children: []*interfacedefinition.Children{{
				XMLName: xml.Name{
					Local: "children",
				},
				Node: []*interfacedefinition.Node{{
					XMLName: xml.Name{
						Local: "node",
					},
					NodeNameAttr: "option",
					OwnerAttr:    "${vyos_conf_scripts_dir}/system-option.py",
					Properties: []*interfacedefinition.Properties{{
						XMLName: xml.Name{
							Local: "properties",
						},
						Help:     []string{"System Options"},
						Priority: []string{"9999"},
					}},
					Children: []*interfacedefinition.Children{{
						XMLName: xml.Name{
							Local: "children",
						},
						Node: []*interfacedefinition.Node{{
							XMLName: xml.Name{
								Local: "node",
							},
							NodeNameAttr: "http-client",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Global options used for HTTP client"},
							}},
							Children: []*interfacedefinition.Children{{
								XMLName: xml.Name{
									Local: "children",
								},
								LeafNode: []*interfacedefinition.LeafNode{{
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "source-interface",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Interface used to establish connection"},
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
											Format:      "interface",
											Description: "Interface name",
										}},
										CompletionHelp: []*interfacedefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
											Script: []string{"${vyos_completion_dir}/list_interfaces"},
										}},
									}},
								}, {
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "source-address",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Source IP address used to initiate connection"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Validator: []*interfacedefinition.Validator{{
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr: "ip-address",
											}},
										}},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "ipv4",
											Description: "IPv4 source address",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "ipv6",
											Description: "IPv6 source address",
										}},
										CompletionHelp: []*interfacedefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
											Script: []string{"${vyos_completion_dir}/list_local_ips.sh --both"},
										}},
									}},
								}},
							}},
						}, {
							XMLName: xml.Name{
								Local: "node",
							},
							NodeNameAttr: "ssh-client",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Global options used for SSH client"},
							}},
							Children: []*interfacedefinition.Children{{
								XMLName: xml.Name{
									Local: "children",
								},
								LeafNode: []*interfacedefinition.LeafNode{{
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "source-address",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Source IP address used to initiate connection"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Validator: []*interfacedefinition.Validator{{
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr: "ip-address",
											}},
										}},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "ipv4",
											Description: "IPv4 source address",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "ipv6",
											Description: "IPv6 source address",
										}},
										CompletionHelp: []*interfacedefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
											Script: []string{"${vyos_completion_dir}/list_local_ips.sh --both"},
										}},
									}},
								}, {
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "source-interface",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Interface used to establish connection"},
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
											Format:      "interface",
											Description: "Interface name",
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
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "ctrl-alt-delete",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"System action on Ctrl-Alt-Delete keystroke"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Regex: []string{"(ignore|reboot|poweroff)"},
								}},
								ValueHelp: []*interfacedefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "ignore",
									Description: "Ignore key sequence",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "reboot",
									Description: "Reboot system",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "poweroff",
									Description: "Poweroff system",
								}},
								ConstraintErrorMessage: []string{"Must be ignore, reboot, or poweroff"},
								CompletionHelp: []*interfacedefinition.CompletionHelp{{
									XMLName: xml.Name{
										Local: "completionHelp",
									},
									List: []string{"ignore reboot poweroff"},
								}},
							}},
						}, {
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "keyboard-layout",
							DefaultValue: []string{"us"},
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"System keyboard layout, type ISO2"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Regex: []string{"(us|uk|fr|de|es|fi|jp106|no|dk|dvorak)"},
								}},
								ValueHelp: []*interfacedefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "us",
									Description: "United States",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "uk",
									Description: "United Kingdom",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "fr",
									Description: "France",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "de",
									Description: "Germany",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "es",
									Description: "Spain",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "fi",
									Description: "Finland",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "jp106",
									Description: "Japan",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "no",
									Description: "Norway",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "dk",
									Description: "Denmark",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "dvorak",
									Description: "Dvorak",
								}},
								ConstraintErrorMessage: []string{"Invalid keyboard layout"},
								CompletionHelp: []*interfacedefinition.CompletionHelp{{
									XMLName: xml.Name{
										Local: "completionHelp",
									},
									List: []string{"us uk fr de es fi jp106 no dk dvorak"},
								}},
							}},
						}, {
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "performance",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Tune system performance"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Regex: []string{"(throughput|latency)"},
								}},
								ValueHelp: []*interfacedefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "throughput",
									Description: "Tune for maximum network throughput",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "latency",
									Description: "Tune for low network latency",
								}},
								CompletionHelp: []*interfacedefinition.CompletionHelp{{
									XMLName: xml.Name{
										Local: "completionHelp",
									},
									List: []string{"throughput latency"},
								}},
							}},
						}, {
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "reboot-on-panic",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Reboot system on kernel panic"},
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
							NodeNameAttr: "startup-beep",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"plays sound via system speaker when you can login"},
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
							NodeNameAttr: "root-partition-auto-resize",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Enable root partition auto-extention on system boot"},
								Valueless: []*interfacedefinition.Valueless{{
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
	}
}
