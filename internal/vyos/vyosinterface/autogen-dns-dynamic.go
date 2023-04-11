// Code generated by /workspaces/terraform-provider-vyos/tools/build-vyos-infterface-definition-structs/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/vyos/schema/interfacedefinition"
)

func dnsdynamic() interfacedefinition.InterfaceDefinition {
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
					NodeNameAttr: "dns",
					Properties: []*interfacedefinition.Properties{{
						XMLName: xml.Name{
							Local: "properties",
						},
						Help: []string{"Domain Name System related services"},
					}},
					Children: []*interfacedefinition.Children{{
						XMLName: xml.Name{
							Local: "children",
						},
						Node: []*interfacedefinition.Node{{
							XMLName: xml.Name{
								Local: "node",
							},
							NodeNameAttr: "dynamic",
							OwnerAttr:    "${vyos_conf_scripts_dir}/dynamic_dns.py",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Dynamic DNS"},
							}},
							Children: []*interfacedefinition.Children{{
								XMLName: xml.Name{
									Local: "children",
								},
								TagNode: []*interfacedefinition.TagNode{{
									XMLName: xml.Name{
										Local: "tagNode",
									},
									NodeNameAttr: "interface",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Interface to send DDNS updates for"},
										CompletionHelp: []*interfacedefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
											Script: []string{"${vyos_completion_dir}/list_interfaces"},
										}},
									}},
									Children: []*interfacedefinition.Children{{
										XMLName: xml.Name{
											Local: "children",
										},
										Node: []*interfacedefinition.Node{{
											XMLName: xml.Name{
												Local: "node",
											},
											NodeNameAttr: "use-web",
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Web check used for obtaining the external IP address"},
											}},
											Children: []*interfacedefinition.Children{{
												XMLName: xml.Name{
													Local: "children",
												},
												LeafNode: []*interfacedefinition.LeafNode{{
													XMLName: xml.Name{
														Local: "leafNode",
													},
													NodeNameAttr: "skip",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"Skip everything before this on the given URL"},
													}},
												}, {
													XMLName: xml.Name{
														Local: "leafNode",
													},
													NodeNameAttr: "url",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"URL to obtain the current external IP address"},
													}},
												}},
											}},
										}},
										TagNode: []*interfacedefinition.TagNode{{
											XMLName: xml.Name{
												Local: "tagNode",
											},
											NodeNameAttr: "rfc2136",
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"RFC2136 Update name"},
											}},
											Children: []*interfacedefinition.Children{{
												XMLName: xml.Name{
													Local: "children",
												},
												LeafNode: []*interfacedefinition.LeafNode{{
													XMLName: xml.Name{
														Local: "leafNode",
													},
													NodeNameAttr: "key",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"File containing the secret key shared with remote DNS server"},
														ValueHelp: []*interfacedefinition.ValueHelp{{
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "filename",
															Description: "File in /config/auth directory",
														}},
													}},
												}, {
													XMLName: xml.Name{
														Local: "leafNode",
													},
													NodeNameAttr: "record",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"Record to be updated"},
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
														Help: []string{"Server to be updated"},
													}},
												}, {
													XMLName: xml.Name{
														Local: "leafNode",
													},
													NodeNameAttr: "ttl",
													DefaultValue: []string{"600"},
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"Time To Live (default: 600)"},
														Constraint: []*interfacedefinition.Constraint{{
															XMLName: xml.Name{
																Local: "constraint",
															},
															Validator: []*interfacedefinition.Validator{{
																XMLName: xml.Name{
																	Local: "validator",
																},
																NameAttr:     "numeric",
																ArgumentAttr: "--range 1-86400",
															}},
														}},
														ValueHelp: []*interfacedefinition.ValueHelp{{
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "u32:1-86400",
															Description: "DNS forwarding cache size",
														}},
													}},
												}, {
													XMLName: xml.Name{
														Local: "leafNode",
													},
													NodeNameAttr: "zone",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"Zone to be updated"},
													}},
												}},
											}},
										}, {
											XMLName: xml.Name{
												Local: "tagNode",
											},
											NodeNameAttr: "service",
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Service being used for Dynamic DNS"},
												Constraint: []*interfacedefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Regex: []string{"(custom|afraid|changeip|cloudflare|dnspark|dslreports|dyndns|easydns|namecheap|noip|sitelutions|zoneedit|\\w+)"},
												}},
												ValueHelp: []*interfacedefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "txt",
													Description: "Dynanmic DNS service with a custom name",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "afraid",
													Description: "afraid.org Services",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "changeip",
													Description: "changeip.com Services",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "cloudflare",
													Description: "cloudflare.com Services",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "dnspark",
													Description: "dnspark.com Services",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "dslreports",
													Description: "dslreports.com Services",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "dyndns",
													Description: "dyndns.com Services",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "easydns",
													Description: "easydns.com Services",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "namecheap",
													Description: "namecheap.com Services",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "noip",
													Description: "noip.com Services",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "sitelutions",
													Description: "sitelutions.com Services",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "zoneedit",
													Description: "zoneedit.com Services",
												}},
												ConstraintErrorMessage: []string{"You can use only predefined list of services or word characters (_, a-z, A-Z, 0-9) as service name"},
												CompletionHelp: []*interfacedefinition.CompletionHelp{{
													XMLName: xml.Name{
														Local: "completionHelp",
													},
													List: []string{"afraid changeip cloudflare dnspark dslreports dyndns easydns namecheap noip sitelutions zoneedit"},
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
													NodeNameAttr: "host-name",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"Hostname registered with DDNS service"},
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
													NodeNameAttr: "login",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"Login for DDNS service"},
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
														Help: []string{"Password for DDNS service"},
													}},
												}, {
													XMLName: xml.Name{
														Local: "leafNode",
													},
													NodeNameAttr: "protocol",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"ddclient protocol used for DDNS service"},
														Constraint: []*interfacedefinition.Constraint{{
															XMLName: xml.Name{
																Local: "constraint",
															},
															Regex: []string{"(changeip|cloudflare|dnsmadeeasy|dnspark|dondominio|dslreports1|dtdns|duckdns|dyndns2|easydns|freedns|freemyip|googledomains|hammernode1|namecheap|nfsn|noip|sitelutions|woima|yandex|zoneedit1)"},
														}},
														ValueHelp: []*interfacedefinition.ValueHelp{{
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "changeip",
															Description: "ChangeIP protocol",
														}, {
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "cloudflare",
															Description: "Cloudflare protocol",
														}, {
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "dnsmadeeasy",
															Description: "DNS Made Easy protocol",
														}, {
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "dnspark",
															Description: "DNS Park protocol",
														}, {
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "dondominio",
															Description: "DonDominio protocol",
														}, {
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "dslreports1",
															Description: "DslReports protocol",
														}, {
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "dtdns",
															Description: "DtDNS protocol",
														}, {
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "duckdns",
															Description: "DuckDNS protocol",
														}, {
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "dyndns2",
															Description: "DynDNS protocol v2",
														}, {
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "easydns",
															Description: "easyDNS protocol",
														}, {
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "freedns",
															Description: "FreeDNS protocol",
														}, {
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "freemyip",
															Description: "freemyip protocol",
														}, {
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "googledomains",
															Description: "Google domains protocol",
														}, {
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "hammernode1",
															Description: "Hammernode protocol",
														}, {
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "namecheap",
															Description: "Namecheap protocol",
														}, {
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "nfsn",
															Description: "NearlyFreeSpeech DNS protocol",
														}, {
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "noip",
															Description: "No-IP protocol",
														}, {
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "sitelutions",
															Description: "Sitelutions protocol",
														}, {
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "woima",
															Description: "WOIMA protocol",
														}, {
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "yandex",
															Description: "Yandex.DNS protocol",
														}, {
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "zoneedit1",
															Description: "Zoneedit protocol",
														}},
														ConstraintErrorMessage: []string{"Please choose from the list of allowed protocols"},
														CompletionHelp: []*interfacedefinition.CompletionHelp{{
															XMLName: xml.Name{
																Local: "completionHelp",
															},
															List: []string{"changeip cloudflare dnsmadeeasy dnspark dondominio dslreports1 dtdns duckdns dyndns2 easydns freedns freemyip googledomains hammernode1 namecheap nfsn noip sitelutions woima yandex zoneedit1"},
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
														Help: []string{"Remote server to connect to"},
														ValueHelp: []*interfacedefinition.ValueHelp{{
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "ipv4",
															Description: "Server IPv4 address",
														}, {
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "hostname",
															Description: "Server hostname/FQDN",
														}},
													}},
												}, {
													XMLName: xml.Name{
														Local: "leafNode",
													},
													NodeNameAttr: "zone",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"DNS zone to update (only available with CloudFlare)"},
													}},
												}},
											}},
										}},
										LeafNode: []*interfacedefinition.LeafNode{{
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "ipv6-enable",
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Allow explicit IPv6 addresses for Dynamic DNS for this interface"},
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
					}},
				}},
			}},
		}},
	}
}
