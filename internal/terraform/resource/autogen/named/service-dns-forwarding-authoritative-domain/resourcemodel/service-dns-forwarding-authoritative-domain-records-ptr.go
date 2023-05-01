// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ServiceDNSForwardingAuthoritativeDomainRecordsPtr describes the resource data model.
type ServiceDNSForwardingAuthoritativeDomainRecordsPtr struct {
	// LeafNodes
	LeafServiceDNSForwardingAuthoritativeDomainRecordsPtrTarget  types.String `tfsdk:"target" json:"target,omitempty"`
	LeafServiceDNSForwardingAuthoritativeDomainRecordsPtrTTL     types.String `tfsdk:"ttl" json:"ttl,omitempty"`
	LeafServiceDNSForwardingAuthoritativeDomainRecordsPtrDisable types.String `tfsdk:"disable" json:"disable,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDNSForwardingAuthoritativeDomainRecordsPtr) ResourceSchemaAttributes() map[string]schema.Attribute {
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
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsPtr) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafServiceDNSForwardingAuthoritativeDomainRecordsPtrTarget.IsNull() && !o.LeafServiceDNSForwardingAuthoritativeDomainRecordsPtrTarget.IsUnknown() {
		jsonData["target"] = o.LeafServiceDNSForwardingAuthoritativeDomainRecordsPtrTarget.ValueString()
	}

	if !o.LeafServiceDNSForwardingAuthoritativeDomainRecordsPtrTTL.IsNull() && !o.LeafServiceDNSForwardingAuthoritativeDomainRecordsPtrTTL.IsUnknown() {
		jsonData["ttl"] = o.LeafServiceDNSForwardingAuthoritativeDomainRecordsPtrTTL.ValueString()
	}

	if !o.LeafServiceDNSForwardingAuthoritativeDomainRecordsPtrDisable.IsNull() && !o.LeafServiceDNSForwardingAuthoritativeDomainRecordsPtrDisable.IsUnknown() {
		jsonData["disable"] = o.LeafServiceDNSForwardingAuthoritativeDomainRecordsPtrDisable.ValueString()
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
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsPtr) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["target"]; ok {
		o.LeafServiceDNSForwardingAuthoritativeDomainRecordsPtrTarget = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDNSForwardingAuthoritativeDomainRecordsPtrTarget = basetypes.NewStringNull()
	}

	if value, ok := jsonData["ttl"]; ok {
		o.LeafServiceDNSForwardingAuthoritativeDomainRecordsPtrTTL = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDNSForwardingAuthoritativeDomainRecordsPtrTTL = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable"]; ok {
		o.LeafServiceDNSForwardingAuthoritativeDomainRecordsPtrDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDNSForwardingAuthoritativeDomainRecordsPtrDisable = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
