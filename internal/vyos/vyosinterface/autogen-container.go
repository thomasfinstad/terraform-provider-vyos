// Code generated by /workspaces/terraform-provider-vyos/tools/build-vyos-infterface-definition-structs/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/vyos/schema/interfacedefinition"
)

func container() interfacedefinition.InterfaceDefinition {
	return interfacedefinition.InterfaceDefinition{
		XMLName: xml.Name{
			Local: "interfaceDefinition",
		},
		Node: []*interfacedefinition.Node{{
			XMLName: xml.Name{
				Local: "node",
			},
			NodeNameAttr: "container",
			OwnerAttr:    "${vyos_conf_scripts_dir}/container.py",
			Properties: []*interfacedefinition.Properties{{
				XMLName: xml.Name{
					Local: "properties",
				},
				Help:     []string{"Container applications"},
				Priority: []string{"1280"},
			}},
			Children: []*interfacedefinition.Children{{
				XMLName: xml.Name{
					Local: "children",
				},
				TagNode: []*interfacedefinition.TagNode{{
					XMLName: xml.Name{
						Local: "tagNode",
					},
					NodeNameAttr: "name",
					Properties: []*interfacedefinition.Properties{{
						XMLName: xml.Name{
							Local: "properties",
						},
						Help: []string{"Container name"},
						Constraint: []*interfacedefinition.Constraint{{
							XMLName: xml.Name{
								Local: "constraint",
							},
							Regex: []string{"[-a-zA-Z0-9]+"},
						}},
						ConstraintErrorMessage: []string{"Container name must be alphanumeric and can contain hyphens"},
					}},
					Children: []*interfacedefinition.Children{{
						XMLName: xml.Name{
							Local: "children",
						},
						TagNode: []*interfacedefinition.TagNode{{
							XMLName: xml.Name{
								Local: "tagNode",
							},
							NodeNameAttr: "device",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Add a host device to the container"},
							}},
							Children: []*interfacedefinition.Children{{
								XMLName: xml.Name{
									Local: "children",
								},
								LeafNode: []*interfacedefinition.LeafNode{{
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "source",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Source device (Example: \"/dev/x\")"},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "txt",
											Description: "Source device",
										}},
									}},
								}, {
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "destination",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Destination container device (Example: \"/dev/x\")"},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "txt",
											Description: "Destination container device",
										}},
									}},
								}},
							}},
						}, {
							XMLName: xml.Name{
								Local: "tagNode",
							},
							NodeNameAttr: "environment",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Add custom environment variables"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Regex: []string{"[-_a-zA-Z0-9]+"},
								}},
								ConstraintErrorMessage: []string{"Environment variable name must be alphanumeric and can contain hyphen and underscores"},
							}},
							Children: []*interfacedefinition.Children{{
								XMLName: xml.Name{
									Local: "children",
								},
								LeafNode: []*interfacedefinition.LeafNode{{
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "value",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Set environment option value"},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "txt",
											Description: "Set environment option value",
										}},
									}},
								}},
							}},
						}, {
							XMLName: xml.Name{
								Local: "tagNode",
							},
							NodeNameAttr: "network",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Attach user defined network to container"},
								CompletionHelp: []*interfacedefinition.CompletionHelp{{
									XMLName: xml.Name{
										Local: "completionHelp",
									},
									Path: []string{"container network"},
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
									NodeNameAttr: "address",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Assign static IP address to container"},
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
											Description: "IPv4 address",
										}},
									}},
								}},
							}},
						}, {
							XMLName: xml.Name{
								Local: "tagNode",
							},
							NodeNameAttr: "port",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Publish port to the container"},
							}},
							Children: []*interfacedefinition.Children{{
								XMLName: xml.Name{
									Local: "children",
								},
								LeafNode: []*interfacedefinition.LeafNode{{
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "source",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Source host port"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Validator: []*interfacedefinition.Validator{{
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr: "port-range",
											}},
										}},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "u32:1-65535",
											Description: "Source host port",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "start-end",
											Description: "Source host port range (e.g. 10025-10030)",
										}},
									}},
								}, {
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "destination",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Destination container port"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Validator: []*interfacedefinition.Validator{{
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr: "port-range",
											}},
										}},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "u32:1-65535",
											Description: "Destination container port",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "start-end",
											Description: "Destination container port range (e.g. 10025-10030)",
										}},
									}},
								}, {
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "protocol",
									DefaultValue: []string{"tcp"},
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Transport protocol used for port mapping"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Regex: []string{"(tcp|udp)"},
										}},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "tcp",
											Description: "Use Transmission Control Protocol for given port",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "udp",
											Description: "Use User Datagram Protocol for given port",
										}},
										CompletionHelp: []*interfacedefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
											List: []string{"tcp udp"},
										}},
									}},
								}},
							}},
						}, {
							XMLName: xml.Name{
								Local: "tagNode",
							},
							NodeNameAttr: "volume",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Mount a volume into the container"},
							}},
							Children: []*interfacedefinition.Children{{
								XMLName: xml.Name{
									Local: "children",
								},
								LeafNode: []*interfacedefinition.LeafNode{{
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "source",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Source host directory"},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "txt",
											Description: "Source host directory",
										}},
									}},
								}, {
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "destination",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Destination container directory"},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "txt",
											Description: "Destination container directory",
										}},
									}},
								}, {
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "mode",
									DefaultValue: []string{"rw"},
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Volume access mode ro/rw"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Regex: []string{"(ro|rw)"},
										}},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "ro",
											Description: "Volume mounted into the container as read-only",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "rw",
											Description: "Volume mounted into the container as read-write",
										}},
										CompletionHelp: []*interfacedefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
											List: []string{"ro rw"},
										}},
									}},
								}},
							}},
						}},
						LeafNode: []*interfacedefinition.LeafNode{{
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "allow-host-networks",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Allow host networks in container"},
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
							NodeNameAttr: "cap-add",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Container capabilities/permissions"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Regex: []string{"(net-admin|net-bind-service|net-raw|setpcap|sys-admin|sys-time)"},
								}},
								ValueHelp: []*interfacedefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "net-admin",
									Description: "Network operations (interface, firewall, routing tables)",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "net-bind-service",
									Description: "Bind a socket to privileged ports (port numbers less than 1024)",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "net-raw",
									Description: "Permission to create raw network sockets",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "setpcap",
									Description: "Capability sets (from bounded or inherited set)",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "sys-admin",
									Description: "Administation operations (quotactl, mount, sethostname, setdomainame)",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "sys-time",
									Description: "Permission to set system clock",
								}},
								CompletionHelp: []*interfacedefinition.CompletionHelp{{
									XMLName: xml.Name{
										Local: "completionHelp",
									},
									List: []string{"net-admin net-bind-service net-raw setpcap sys-admin sys-time"},
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
							NodeNameAttr: "entrypoint",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Override the default ENTRYPOINT from the image"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Regex: []string{"[ !#-%&(-~]+"},
								}},
								ConstraintErrorMessage: []string{"Entrypoint must be ascii characters, use &quot; and &apos for double and single quotes respectively"},
							}},
						}, {
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "host-name",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Container host name"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Regex: []string{"[A-Za-z0-9][-.A-Za-z0-9]&[A-Za-z0-9]"},
								}},
								ConstraintErrorMessage: []string{"Host-name must be alphanumeric and can contain hyphens"},
							}},
						}, {
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "image",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Image name in the hub-registry"},
							}},
						}, {
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "command",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Override the default CMD from the image"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Regex: []string{"[ !#-%&(-~]+"},
								}},
								ConstraintErrorMessage: []string{"Command must be ascii characters, use &quot; and &apos for double and single quotes respectively"},
							}},
						}, {
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "arguments",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"The command's arguments for this container"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Regex: []string{"[ !#-%&(-~]+"},
								}},
								ConstraintErrorMessage: []string{"The command's arguments must be ascii characters, use &quot; and &apos for double and single quotes respectively"},
							}},
						}, {
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "memory",
							DefaultValue: []string{"512"},
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Memory (RAM) available to this container"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Validator: []*interfacedefinition.Validator{{
										XMLName: xml.Name{
											Local: "validator",
										},
										NameAttr:     "numeric",
										ArgumentAttr: "--range 0-16384",
									}},
								}},
								ValueHelp: []*interfacedefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "u32:0",
									Description: "Unlimited",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "u32:1-16384",
									Description: "Container memory in megabytes (MB)",
								}},
								ConstraintErrorMessage: []string{"Container memory must be in range 0 to 16384 MB"},
							}},
						}, {
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "shared-memory",
							DefaultValue: []string{"64"},
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Shared memory available to this container"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Validator: []*interfacedefinition.Validator{{
										XMLName: xml.Name{
											Local: "validator",
										},
										NameAttr:     "numeric",
										ArgumentAttr: "--range 0-8192",
									}},
								}},
								ValueHelp: []*interfacedefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "u32:0",
									Description: "Unlimited",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "u32:1-8192",
									Description: "Container memory in megabytes (MB)",
								}},
								ConstraintErrorMessage: []string{"Container memory must be in range 0 to 8192 MB"},
							}},
						}, {
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "restart",
							DefaultValue: []string{"on-failure"},
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Restart options for container"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Regex: []string{"(no|on-failure|always)"},
								}},
								ValueHelp: []*interfacedefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "no",
									Description: "Do not restart containers on exit",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "on-failure",
									Description: "Restart containers when they exit with a non-zero exit code, retrying indefinitely",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "always",
									Description: "Restart containers when they exit, regardless of status, retrying indefinitely",
								}},
								CompletionHelp: []*interfacedefinition.CompletionHelp{{
									XMLName: xml.Name{
										Local: "completionHelp",
									},
									List: []string{"no on-failure always"},
								}},
							}},
						}},
					}},
				}, {
					XMLName: xml.Name{
						Local: "tagNode",
					},
					NodeNameAttr: "network",
					Properties: []*interfacedefinition.Properties{{
						XMLName: xml.Name{
							Local: "properties",
						},
						Help: []string{"Network name"},
						Constraint: []*interfacedefinition.Constraint{{
							XMLName: xml.Name{
								Local: "constraint",
							},
							Regex: []string{"[-_a-zA-Z0-9]{1,11}"},
						}},
						ConstraintErrorMessage: []string{"Network name cannot be longer than 11 characters"},
					}},
					Children: []*interfacedefinition.Children{{
						XMLName: xml.Name{
							Local: "children",
						},
						LeafNode: []*interfacedefinition.LeafNode{{
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "description",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Network description"},
							}},
						}, {
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "prefix",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Prefix which allocated to that network"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Validator: []*interfacedefinition.Validator{{
										XMLName: xml.Name{
											Local: "validator",
										},
										NameAttr: "ipv4-prefix",
									}, {
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
									Format:      "ipv4net",
									Description: "IPv4 network prefix",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "ipv6net",
									Description: "IPv6 network prefix",
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
						Local: "tagNode",
					},
					NodeNameAttr: "registry",
					DefaultValue: []string{"docker.io quay.io"},
					Properties: []*interfacedefinition.Properties{{
						XMLName: xml.Name{
							Local: "properties",
						},
						Help: []string{"Registry Name"},
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
								Help: []string{"Authentication settings"},
							}},
							Children: []*interfacedefinition.Children{{
								XMLName: xml.Name{
									Local: "children",
								},
								LeafNode: []*interfacedefinition.LeafNode{{
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "username",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Username used for authentication"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Regex: []string{"[[:ascii:]]{1,128}"},
										}},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "txt",
											Description: "Username",
										}},
										ConstraintErrorMessage: []string{"Username is limited to ASCII characters only, with a total length of 128"},
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
										Help: []string{"Password used for authentication"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Regex: []string{"[[:ascii:]]{1,128}"},
										}},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "txt",
											Description: "Password",
										}},
										ConstraintErrorMessage: []string{"Password is limited to ASCII characters only, with a total length of 128"},
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
						}},
					}},
				}},
			}},
		}},
	}
}