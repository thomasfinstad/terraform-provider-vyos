// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicast describes the resource data model.
type ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicast struct {
	// LeafNodes
	LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAddpathTxAll         types.String `tfsdk:"addpath_tx_all" json:"addpath-tx-all,omitempty"`
	LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAddpathTxPerAs       types.String `tfsdk:"addpath_tx_per_as" json:"addpath-tx-per-as,omitempty"`
	LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAsOverrIDe           types.String `tfsdk:"as_override" json:"as-override,omitempty"`
	LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastMaximumPrefix        types.String `tfsdk:"maximum_prefix" json:"maximum-prefix,omitempty"`
	LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastMaximumPrefixOut     types.String `tfsdk:"maximum_prefix_out" json:"maximum-prefix-out,omitempty"`
	LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRemovePrivateAs      types.String `tfsdk:"remove_private_as" json:"remove-private-as,omitempty"`
	LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteReflectorClient types.String `tfsdk:"route_reflector_client" json:"route-reflector-client,omitempty"`
	LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteServerClient    types.String `tfsdk:"route_server_client" json:"route-server-client,omitempty"`
	LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastUnsuppressMap        types.String `tfsdk:"unsuppress_map" json:"unsuppress-map,omitempty"`
	LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastWeight               types.String `tfsdk:"weight" json:"weight,omitempty"`

	// TagNodes

	// Nodes
	NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapability             *ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapability             `tfsdk:"capability" json:"capability,omitempty"`
	NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastPrefixList             *ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastPrefixList             `tfsdk:"prefix_list" json:"prefix-list,omitempty"`
	NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastConditionallyAdvertise *ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastConditionallyAdvertise `tfsdk:"conditionally_advertise" json:"conditionally-advertise,omitempty"`
	NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAllowasIn              *ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAllowasIn              `tfsdk:"allowas_in" json:"allowas-in,omitempty"`
	NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchanged     *ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchanged     `tfsdk:"attribute_unchanged" json:"attribute-unchanged,omitempty"`
	NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDisableSendCommunity   *ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDisableSendCommunity   `tfsdk:"disable_send_community" json:"disable-send-community,omitempty"`
	NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDistributeList         *ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDistributeList         `tfsdk:"distribute_list" json:"distribute-list,omitempty"`
	NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastFilterList             *ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastFilterList             `tfsdk:"filter_list" json:"filter-list,omitempty"`
	NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastNexthopSelf            *ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastNexthopSelf            `tfsdk:"nexthop_self" json:"nexthop-self,omitempty"`
	NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteMap               *ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteMap               `tfsdk:"route_map" json:"route-map,omitempty"`
	NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastSoftReconfiguration    *ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastSoftReconfiguration    `tfsdk:"soft_reconfiguration" json:"soft-reconfiguration,omitempty"`
	NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDefaultOriginate       *ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDefaultOriginate       `tfsdk:"default_originate" json:"default-originate,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicast) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"addpath_tx_all": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Use addpath to advertise all paths to a neighbor

`,
		},

		"addpath_tx_per_as": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Use addpath to advertise the bestpath per each neighboring AS

`,
		},

		"as_override": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Override ASN in outbound updates to configured neighbor local-as

`,
		},

		"maximum_prefix": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum number of prefixes to accept from this peer

|  Format  |  Description  |
|----------|---------------|
|  u32:1-4294967295  |  Prefix limit  |

`,
		},

		"maximum_prefix_out": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum number of prefixes to be sent to this peer

|  Format  |  Description  |
|----------|---------------|
|  u32:1-4294967295  |  Prefix limit  |

`,
		},

		"remove_private_as": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Remove private AS numbers from AS path in outbound route updates

`,
		},

		"route_reflector_client": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Peer is a route reflector client

`,
		},

		"route_server_client": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Peer is a route server client

`,
		},

		"unsuppress_map": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route-map to selectively unsuppress suppressed routes

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Route map name  |

`,
		},

		"weight": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Default weight for routes from this peer

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Default weight  |

`,
		},

		// TagNodes

		// Nodes

		"capability": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapability{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Advertise capabilities to this neighbor (IPv4)

`,
		},

		"prefix_list": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastPrefixList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4-Prefix-list to filter route updates to/from this peer

`,
		},

		"conditionally_advertise": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastConditionallyAdvertise{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Use route-map to conditionally advertise routes

`,
		},

		"allowas_in": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAllowasIn{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Accept route that contains the local-as in the as-path

`,
		},

		"attribute_unchanged": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchanged{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `BGP attributes are sent unchanged

`,
		},

		"disable_send_community": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDisableSendCommunity{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Disable sending community attributes to this peer

`,
		},

		"distribute_list": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDistributeList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Access-list to filter route updates to/from this peer-group

`,
		},

		"filter_list": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastFilterList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `as-path-list to filter route updates to/from this peer

`,
		},

		"nexthop_self": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastNexthopSelf{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Disable the next hop calculation for this peer

`,
		},

		"route_map": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteMap{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
		},

		"soft_reconfiguration": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastSoftReconfiguration{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Soft reconfiguration for peer

`,
		},

		"default_originate": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDefaultOriginate{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Originate default route to this peer

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicast) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAddpathTxAll.IsNull() && !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAddpathTxAll.IsUnknown() {
		jsonData["addpath-tx-all"] = o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAddpathTxAll.ValueString()
	}

	if !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAddpathTxPerAs.IsNull() && !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAddpathTxPerAs.IsUnknown() {
		jsonData["addpath-tx-per-as"] = o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAddpathTxPerAs.ValueString()
	}

	if !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAsOverrIDe.IsNull() && !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAsOverrIDe.IsUnknown() {
		jsonData["as-override"] = o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAsOverrIDe.ValueString()
	}

	if !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastMaximumPrefix.IsNull() && !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastMaximumPrefix.IsUnknown() {
		jsonData["maximum-prefix"] = o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastMaximumPrefix.ValueString()
	}

	if !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastMaximumPrefixOut.IsNull() && !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastMaximumPrefixOut.IsUnknown() {
		jsonData["maximum-prefix-out"] = o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastMaximumPrefixOut.ValueString()
	}

	if !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRemovePrivateAs.IsNull() && !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRemovePrivateAs.IsUnknown() {
		jsonData["remove-private-as"] = o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRemovePrivateAs.ValueString()
	}

	if !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteReflectorClient.IsNull() && !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteReflectorClient.IsUnknown() {
		jsonData["route-reflector-client"] = o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteReflectorClient.ValueString()
	}

	if !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteServerClient.IsNull() && !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteServerClient.IsUnknown() {
		jsonData["route-server-client"] = o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteServerClient.ValueString()
	}

	if !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastUnsuppressMap.IsNull() && !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastUnsuppressMap.IsUnknown() {
		jsonData["unsuppress-map"] = o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastUnsuppressMap.ValueString()
	}

	if !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastWeight.IsNull() && !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastWeight.IsUnknown() {
		jsonData["weight"] = o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastWeight.ValueString()
	}

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapability).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapability)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["capability"] = subData
	}

	if !reflect.ValueOf(o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastPrefixList).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastPrefixList)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["prefix-list"] = subData
	}

	if !reflect.ValueOf(o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastConditionallyAdvertise).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastConditionallyAdvertise)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["conditionally-advertise"] = subData
	}

	if !reflect.ValueOf(o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAllowasIn).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAllowasIn)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["allowas-in"] = subData
	}

	if !reflect.ValueOf(o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchanged).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchanged)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["attribute-unchanged"] = subData
	}

	if !reflect.ValueOf(o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDisableSendCommunity).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDisableSendCommunity)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["disable-send-community"] = subData
	}

	if !reflect.ValueOf(o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDistributeList).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDistributeList)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["distribute-list"] = subData
	}

	if !reflect.ValueOf(o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastFilterList).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastFilterList)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["filter-list"] = subData
	}

	if !reflect.ValueOf(o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastNexthopSelf).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastNexthopSelf)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["nexthop-self"] = subData
	}

	if !reflect.ValueOf(o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteMap).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteMap)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["route-map"] = subData
	}

	if !reflect.ValueOf(o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastSoftReconfiguration).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastSoftReconfiguration)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["soft-reconfiguration"] = subData
	}

	if !reflect.ValueOf(o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDefaultOriginate).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDefaultOriginate)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["default-originate"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicast) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["addpath-tx-all"]; ok {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAddpathTxAll = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAddpathTxAll = basetypes.NewStringNull()
	}

	if value, ok := jsonData["addpath-tx-per-as"]; ok {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAddpathTxPerAs = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAddpathTxPerAs = basetypes.NewStringNull()
	}

	if value, ok := jsonData["as-override"]; ok {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAsOverrIDe = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAsOverrIDe = basetypes.NewStringNull()
	}

	if value, ok := jsonData["maximum-prefix"]; ok {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastMaximumPrefix = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastMaximumPrefix = basetypes.NewStringNull()
	}

	if value, ok := jsonData["maximum-prefix-out"]; ok {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastMaximumPrefixOut = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastMaximumPrefixOut = basetypes.NewStringNull()
	}

	if value, ok := jsonData["remove-private-as"]; ok {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRemovePrivateAs = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRemovePrivateAs = basetypes.NewStringNull()
	}

	if value, ok := jsonData["route-reflector-client"]; ok {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteReflectorClient = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteReflectorClient = basetypes.NewStringNull()
	}

	if value, ok := jsonData["route-server-client"]; ok {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteServerClient = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteServerClient = basetypes.NewStringNull()
	}

	if value, ok := jsonData["unsuppress-map"]; ok {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastUnsuppressMap = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastUnsuppressMap = basetypes.NewStringNull()
	}

	if value, ok := jsonData["weight"]; ok {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastWeight = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastWeight = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := jsonData["capability"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapability = &ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapability{}

		err = json.Unmarshal(subJSONStr, o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapability)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["prefix-list"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastPrefixList = &ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastPrefixList{}

		err = json.Unmarshal(subJSONStr, o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastPrefixList)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["conditionally-advertise"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastConditionallyAdvertise = &ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastConditionallyAdvertise{}

		err = json.Unmarshal(subJSONStr, o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastConditionallyAdvertise)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["allowas-in"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAllowasIn = &ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAllowasIn{}

		err = json.Unmarshal(subJSONStr, o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAllowasIn)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["attribute-unchanged"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchanged = &ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchanged{}

		err = json.Unmarshal(subJSONStr, o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchanged)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["disable-send-community"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDisableSendCommunity = &ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDisableSendCommunity{}

		err = json.Unmarshal(subJSONStr, o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDisableSendCommunity)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["distribute-list"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDistributeList = &ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDistributeList{}

		err = json.Unmarshal(subJSONStr, o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDistributeList)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["filter-list"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastFilterList = &ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastFilterList{}

		err = json.Unmarshal(subJSONStr, o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastFilterList)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["nexthop-self"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastNexthopSelf = &ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastNexthopSelf{}

		err = json.Unmarshal(subJSONStr, o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastNexthopSelf)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["route-map"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteMap = &ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteMap{}

		err = json.Unmarshal(subJSONStr, o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteMap)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["soft-reconfiguration"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastSoftReconfiguration = &ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastSoftReconfiguration{}

		err = json.Unmarshal(subJSONStr, o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastSoftReconfiguration)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["default-originate"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDefaultOriginate = &ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDefaultOriginate{}

		err = json.Unmarshal(subJSONStr, o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDefaultOriginate)
		if err != nil {
			return err
		}
	}

	return nil
}
