// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsRIPInterface describes the resource data model.
type ProtocolsRIPInterface struct {
	SelfIdentifier types.String `tfsdk:"interface_id" vyos:",self-id"`

	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeProtocolsRIPInterfaceSplitHorizon   *ProtocolsRIPInterfaceSplitHorizon   `tfsdk:"split_horizon" vyos:"split-horizon,omitempty"`
	NodeProtocolsRIPInterfaceAuthentication *ProtocolsRIPInterfaceAuthentication `tfsdk:"authentication" vyos:"authentication,omitempty"`
	NodeProtocolsRIPInterfaceReceive        *ProtocolsRIPInterfaceReceive        `tfsdk:"receive" vyos:"receive,omitempty"`
	NodeProtocolsRIPInterfaceSend           *ProtocolsRIPInterfaceSend           `tfsdk:"send" vyos:"send,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsRIPInterface) GetVyosPath() []string {
	return []string{
		"protocols",

		"rip",

		"interface",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsRIPInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, an amalgamation of the `interface_id` and the parents `*_id` fields seperated by dunder `__` starting with top level ancestor.",
		},
		"interface_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Interface name

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Interface name  |

`,
		},

		// LeafNodes

		// Nodes

		"split_horizon": schema.SingleNestedAttribute{
			Attributes: ProtocolsRIPInterfaceSplitHorizon{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Split horizon parameters

`,
		},

		"authentication": schema.SingleNestedAttribute{
			Attributes: ProtocolsRIPInterfaceAuthentication{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Authentication

`,
		},

		"receive": schema.SingleNestedAttribute{
			Attributes: ProtocolsRIPInterfaceReceive{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Advertisement reception

`,
		},

		"send": schema.SingleNestedAttribute{
			Attributes: ProtocolsRIPInterfaceSend{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Advertisement transmission

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsRIPInterface) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsRIPInterface) UnmarshalJSON(_ []byte) error {
	return nil
}
