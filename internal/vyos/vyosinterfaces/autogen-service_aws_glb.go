// Code generated by /repo/tools/build-vyos-infterface-definition-structs/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/vyos/schemadefinition"
)

func service_aws_glb() schemadefinition.InterfaceDefinition {
	return schemadefinition.InterfaceDefinition{
		XMLName: xml.Name{
			Local: "interfaceDefinition",
		},
		Node: []*schemadefinition.Node{{
			IsBaseNode: false,
			XMLName: xml.Name{
				Local: "node",
			},
			NodeNameAttr: "service",
			Children: []*schemadefinition.Children{{
				XMLName: xml.Name{
					Local: "children",
				},
				Node: []*schemadefinition.Node{{
					IsBaseNode: false,
					XMLName: xml.Name{
						Local: "node",
					},
					NodeNameAttr: "aws",
					Properties: []*schemadefinition.Properties{{
						XMLName: xml.Name{
							Local: "properties",
						},
						Help: []string{"Amazon Web Service"},
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
							NodeNameAttr: "glb",
							OwnerAttr:    "${vyos_conf_scripts_dir}/service_aws_glb.py",
							Properties: []*schemadefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help:     []string{"Gateway load-balancer tunnel handler"},
								Priority: []string{"1280"},
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
									NodeNameAttr: "script",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Script executed on create or destroy tunnel"},
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
											NodeNameAttr: "on-create",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Script to run when interface is created"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "script",
													}},
												}},
											}},
										}, {
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "on-destroy",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Script to run when interface is destroyed"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "script",
													}},
												}},
											}},
										}},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "node",
									},
									NodeNameAttr: "status",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Status"},
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
											NodeNameAttr: "format",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Statistic format"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Regex: []string{"(simple|full)"},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "simple",
													Description: "Simple format",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "full",
													Description: "Full format",
												}},
												CompletionHelp: []*schemadefinition.CompletionHelp{{
													XMLName: xml.Name{
														Local: "completionHelp",
													},
													List: []string{"simple full"},
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
										}},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "node",
									},
									NodeNameAttr: "threads",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Threads settings"},
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
											NodeNameAttr: "tunnel",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Number of threads for each tunnel processor"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr:     "numeric",
														ArgumentAttr: "--range 1-256",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "u32:1-256",
													Description: "Number of threads",
												}},
											}},
										}, {
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "tunnel-affinity",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"List of cores worker threads"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr:     "numeric",
														ArgumentAttr: "--allow-range --range 0-255",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "<idN>-<idM>",
													Description: "CPU core id range (use '-' as delimiter)",
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
												Help: []string{"Number of threads for UDP receiver"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr:     "numeric",
														ArgumentAttr: "--range 1-256",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "u32:1-256",
													Description: "Number of threads",
												}},
											}},
										}, {
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "udp-affinity",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"List of cores worker threads"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr:     "numeric",
														ArgumentAttr: "--allow-range --range 0-255",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "<idN>-<idM>",
													Description: "CPU core id range (use '-' as delimiter)",
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
