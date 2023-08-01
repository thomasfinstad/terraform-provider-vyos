// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ServiceSnmpVthreeViewOID describes the resource data model.
type ServiceSnmpVthreeViewOID struct {
	SelfIdentifier types.String `tfsdk:"oid_id" vyos:",self-id"`

	ParentIDServiceSnmpVthreeView types.String `tfsdk:"view" vyos:"view,parent-id"`

	// LeafNodes
	LeafServiceSnmpVthreeViewOIDExclude types.String `tfsdk:"exclude" vyos:"exclude,omitempty"`
	LeafServiceSnmpVthreeViewOIDMask    types.String `tfsdk:"mask" vyos:"mask,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceSnmpVthreeViewOID) GetVyosPath() []string {
	return []string{
		"service",

		"snmp",

		"v3",

		"view",
		o.ParentIDServiceSnmpVthreeView.ValueString(),

		"oid",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceSnmpVthreeViewOID) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, an amalgamation of the `oid_id` and the parents `*_id` fields seperated by dunder `__` starting with top level ancestor.",
		},
		"oid_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Specifies the oid

`,
		},

		"view_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Specifies the view with name viewname

`,
		},

		// LeafNodes

		"exclude": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Exclude is an optional argument

`,
		},

		"mask": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Defines a bit-mask that is indicating which subidentifiers of the associated subtree OID should be regarded as significant

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServiceSnmpVthreeViewOID) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ServiceSnmpVthreeViewOID) UnmarshalJSON(_ []byte) error {
	return nil
}
