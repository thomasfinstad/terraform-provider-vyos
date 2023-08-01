// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsRpkiCache describes the resource data model.
type ProtocolsRpkiCache struct {
	SelfIdentifier types.String `tfsdk:"cache_id" vyos:",self-id"`

	// LeafNodes
	LeafProtocolsRpkiCachePort       types.Number `tfsdk:"port" vyos:"port,omitempty"`
	LeafProtocolsRpkiCachePreference types.Number `tfsdk:"preference" vyos:"preference,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeProtocolsRpkiCacheTCP *ProtocolsRpkiCacheTCP `tfsdk:"ssh" vyos:"ssh,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsRpkiCache) GetVyosPath() []string {
	return []string{
		"protocols",

		"rpki",

		"cache",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsRpkiCache) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, an amalgamation of the `cache_id` and the parents `*_id` fields seperated by dunder `__` starting with top level ancestor.",
		},
		"cache_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `RPKI cache server address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IP address of RPKI server  |
    |  ipv6  &emsp; |  IPv6 address of RPKI server  |
    |  hostname  &emsp; |  Fully qualified domain name of RPKI server  |

`,
		},

		// LeafNodes

		"port": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Port number used by connection

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Numeric IP port  |

`,
		},

		"preference": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Preference of the cache server

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-255  &emsp; |  Preference of the cache server  |

`,
		},

		// Nodes

		"ssh": schema.SingleNestedAttribute{
			Attributes: ProtocolsRpkiCacheTCP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `RPKI SSH connection settings

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsRpkiCache) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsRpkiCache) UnmarshalJSON(_ []byte) error {
	return nil
}
