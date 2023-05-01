// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ServiceDNSForwardingAuthoritativeDomainRecordsCname describes the resource data model.
type ServiceDNSForwardingAuthoritativeDomainRecordsCname struct {
	// LeafNodes
	LeafServiceDNSForwardingAuthoritativeDomainRecordsCnameTarget  types.String `tfsdk:"target" json:"target,omitempty"`
	LeafServiceDNSForwardingAuthoritativeDomainRecordsCnameTTL     types.String `tfsdk:"ttl" json:"ttl,omitempty"`
	LeafServiceDNSForwardingAuthoritativeDomainRecordsCnameDisable types.String `tfsdk:"disable" json:"disable,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDNSForwardingAuthoritativeDomainRecordsCname) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"target": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Target DNS name

|  Format  |  Description  |
|----------|---------------|
|  name.example.com  |  An absolute DNS name  |

`,
		},

		"ttl": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Time-to-live (TTL)

|  Format  |  Description  |
|----------|---------------|
|  u32:0-2147483647  |  TTL in seconds  |

`,

			// Default:          stringdefault.StaticString(`300`),
			Computed: true,
		},

		"disable": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsCname) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafServiceDNSForwardingAuthoritativeDomainRecordsCnameTarget.IsNull() && !o.LeafServiceDNSForwardingAuthoritativeDomainRecordsCnameTarget.IsUnknown() {
		jsonData["target"] = o.LeafServiceDNSForwardingAuthoritativeDomainRecordsCnameTarget.ValueString()
	}

	if !o.LeafServiceDNSForwardingAuthoritativeDomainRecordsCnameTTL.IsNull() && !o.LeafServiceDNSForwardingAuthoritativeDomainRecordsCnameTTL.IsUnknown() {
		jsonData["ttl"] = o.LeafServiceDNSForwardingAuthoritativeDomainRecordsCnameTTL.ValueString()
	}

	if !o.LeafServiceDNSForwardingAuthoritativeDomainRecordsCnameDisable.IsNull() && !o.LeafServiceDNSForwardingAuthoritativeDomainRecordsCnameDisable.IsUnknown() {
		jsonData["disable"] = o.LeafServiceDNSForwardingAuthoritativeDomainRecordsCnameDisable.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsCname) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["target"]; ok {
		o.LeafServiceDNSForwardingAuthoritativeDomainRecordsCnameTarget = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDNSForwardingAuthoritativeDomainRecordsCnameTarget = basetypes.NewStringNull()
	}

	if value, ok := jsonData["ttl"]; ok {
		o.LeafServiceDNSForwardingAuthoritativeDomainRecordsCnameTTL = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDNSForwardingAuthoritativeDomainRecordsCnameTTL = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable"]; ok {
		o.LeafServiceDNSForwardingAuthoritativeDomainRecordsCnameDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDNSForwardingAuthoritativeDomainRecordsCnameDisable = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
