// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ServiceDNSForwardingAuthoritativeDomainRecordsSrv describes the resource data model.
type ServiceDNSForwardingAuthoritativeDomainRecordsSrv struct {
	// LeafNodes
	LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvTTL     types.String `tfsdk:"ttl" json:"ttl,omitempty"`
	LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvDisable types.String `tfsdk:"disable" json:"disable,omitempty"`

	// TagNodes
	TagServiceDNSForwardingAuthoritativeDomainRecordsSrvEntry *map[string]ServiceDNSForwardingAuthoritativeDomainRecordsSrvEntry `tfsdk:"entry" json:"entry,omitempty"`

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDNSForwardingAuthoritativeDomainRecordsSrv) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

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

		"entry": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ServiceDNSForwardingAuthoritativeDomainRecordsSrvEntry{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Service entry

|  Format  |  Description  |
|----------|---------------|
|  u32:0-65535  |  Entry number  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsSrv) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvTTL.IsNull() && !o.LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvTTL.IsUnknown() {
		jsonData["ttl"] = o.LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvTTL.ValueString()
	}

	if !o.LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvDisable.IsNull() && !o.LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvDisable.IsUnknown() {
		jsonData["disable"] = o.LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvDisable.ValueString()
	}

	// Tags

	if !reflect.ValueOf(o.TagServiceDNSForwardingAuthoritativeDomainRecordsSrvEntry).IsZero() {
		subJSONStr, err := json.Marshal(o.TagServiceDNSForwardingAuthoritativeDomainRecordsSrvEntry)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["entry"] = subData
	}

	// Nodes

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsSrv) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["ttl"]; ok {
		o.LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvTTL = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvTTL = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable"]; ok {
		o.LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvDisable = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := jsonData["entry"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.TagServiceDNSForwardingAuthoritativeDomainRecordsSrvEntry = &map[string]ServiceDNSForwardingAuthoritativeDomainRecordsSrvEntry{}

		err = json.Unmarshal(subJSONStr, o.TagServiceDNSForwardingAuthoritativeDomainRecordsSrvEntry)
		if err != nil {
			return err
		}
	}

	// Nodes

	return nil
}
