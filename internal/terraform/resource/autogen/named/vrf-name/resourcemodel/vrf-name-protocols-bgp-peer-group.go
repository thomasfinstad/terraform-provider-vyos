// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsBgpPeerGroup describes the resource data model.
type VrfNameProtocolsBgpPeerGroup struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpPeerGroupDescrIPtion                  types.String `tfsdk:"description" json:"description,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupDisableCapabilityNegotiation types.String `tfsdk:"disable_capability_negotiation" json:"disable-capability-negotiation,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupDisableConnectedCheck        types.String `tfsdk:"disable_connected_check" json:"disable-connected-check,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupEbgpMultihop                 types.String `tfsdk:"ebgp_multihop" json:"ebgp-multihop,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupGracefulRestart              types.String `tfsdk:"graceful_restart" json:"graceful-restart,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupOverrIDeCapability           types.String `tfsdk:"override_capability" json:"override-capability,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupPassive                      types.String `tfsdk:"passive" json:"passive,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupPassword                     types.String `tfsdk:"password" json:"password,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupRemoteAs                     types.String `tfsdk:"remote_as" json:"remote-as,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupShutdown                     types.String `tfsdk:"shutdown" json:"shutdown,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupUpdateSource                 types.String `tfsdk:"update_source" json:"update-source,omitempty"`

	// TagNodes
	TagVrfNameProtocolsBgpPeerGroupLocalAs   *map[string]VrfNameProtocolsBgpPeerGroupLocalAs   `tfsdk:"local_as" json:"local-as,omitempty"`
	TagVrfNameProtocolsBgpPeerGroupLocalRole *map[string]VrfNameProtocolsBgpPeerGroupLocalRole `tfsdk:"local_role" json:"local-role,omitempty"`

	// Nodes
	NodeVrfNameProtocolsBgpPeerGroupAddressFamily *VrfNameProtocolsBgpPeerGroupAddressFamily `tfsdk:"address_family" json:"address-family,omitempty"`
	NodeVrfNameProtocolsBgpPeerGroupBfd           *VrfNameProtocolsBgpPeerGroupBfd           `tfsdk:"bfd" json:"bfd,omitempty"`
	NodeVrfNameProtocolsBgpPeerGroupCapability    *VrfNameProtocolsBgpPeerGroupCapability    `tfsdk:"capability" json:"capability,omitempty"`
	NodeVrfNameProtocolsBgpPeerGroupTTLSecURIty   *VrfNameProtocolsBgpPeerGroupTTLSecURIty   `tfsdk:"ttl_security" json:"ttl-security,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpPeerGroup) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |

`,
		},

		"disable_capability_negotiation": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable capability negotiation with this neighbor

`,
		},

		"disable_connected_check": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable check to see if eBGP peer address is a connected route

`,
		},

		"ebgp_multihop": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Allow this EBGP neighbor to not be on a directly connected network

|  Format  |  Description  |
|----------|---------------|
|  u32:1-255  |  Number of hops  |

`,
		},

		"graceful_restart": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `BGP graceful restart functionality

|  Format  |  Description  |
|----------|---------------|
|  enable  |  Enable BGP graceful restart at peer level  |
|  disable  |  Disable BGP graceful restart at peer level  |
|  restart-helper  |  Enable BGP graceful restart helper only functionality  |

`,
		},

		"override_capability": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Ignore capability negotiation with specified neighbor

`,
		},

		"passive": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Do not initiate a session with this neighbor

`,
		},

		"password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `BGP MD5 password

`,
		},

		"remote_as": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Neighbor BGP AS number

|  Format  |  Description  |
|----------|---------------|
|  u32:1-4294967294  |  Neighbor AS number  |
|  external  |  Any AS different from the local AS  |
|  internal  |  Neighbor AS number  |

`,
		},

		"shutdown": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Administratively shutdown this neighbor

`,
		},

		"update_source": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Source IP of routing updates

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address of route source  |
|  ipv6  |  IPv6 address of route source  |
|  txt  |  Interface as route source  |

`,
		},

		// TagNodes

		"local_as": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: VrfNameProtocolsBgpPeerGroupLocalAs{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Specify alternate ASN for this BGP process

|  Format  |  Description  |
|----------|---------------|
|  u32:1-4294967294  |  Autonomous System Number (ASN)  |

`,
		},

		"local_role": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: VrfNameProtocolsBgpPeerGroupLocalRole{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Local role for BGP neighbor (RFC9234)

|  Format  |  Description  |
|----------|---------------|
|  customer  |  Using Transit  |
|  peer  |  Public/Private Peering  |
|  provider  |  Providing Transit  |
|  rs-client  |  RS Client  |
|  rs-server  |  Route Server  |

`,
		},

		// Nodes

		"address_family": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamily{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Address-family parameters

`,
		},

		"bfd": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupBfd{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Enable Bidirectional Forwarding Detection (BFD) support

`,
		},

		"capability": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupCapability{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Advertise capabilities to this peer-group

`,
		},

		"ttl_security": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupTTLSecURIty{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Ttl security mechanism

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsBgpPeerGroup) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameProtocolsBgpPeerGroupDescrIPtion.IsNull() && !o.LeafVrfNameProtocolsBgpPeerGroupDescrIPtion.IsUnknown() {
		jsonData["description"] = o.LeafVrfNameProtocolsBgpPeerGroupDescrIPtion.ValueString()
	}

	if !o.LeafVrfNameProtocolsBgpPeerGroupDisableCapabilityNegotiation.IsNull() && !o.LeafVrfNameProtocolsBgpPeerGroupDisableCapabilityNegotiation.IsUnknown() {
		jsonData["disable-capability-negotiation"] = o.LeafVrfNameProtocolsBgpPeerGroupDisableCapabilityNegotiation.ValueString()
	}

	if !o.LeafVrfNameProtocolsBgpPeerGroupDisableConnectedCheck.IsNull() && !o.LeafVrfNameProtocolsBgpPeerGroupDisableConnectedCheck.IsUnknown() {
		jsonData["disable-connected-check"] = o.LeafVrfNameProtocolsBgpPeerGroupDisableConnectedCheck.ValueString()
	}

	if !o.LeafVrfNameProtocolsBgpPeerGroupEbgpMultihop.IsNull() && !o.LeafVrfNameProtocolsBgpPeerGroupEbgpMultihop.IsUnknown() {
		jsonData["ebgp-multihop"] = o.LeafVrfNameProtocolsBgpPeerGroupEbgpMultihop.ValueString()
	}

	if !o.LeafVrfNameProtocolsBgpPeerGroupGracefulRestart.IsNull() && !o.LeafVrfNameProtocolsBgpPeerGroupGracefulRestart.IsUnknown() {
		jsonData["graceful-restart"] = o.LeafVrfNameProtocolsBgpPeerGroupGracefulRestart.ValueString()
	}

	if !o.LeafVrfNameProtocolsBgpPeerGroupOverrIDeCapability.IsNull() && !o.LeafVrfNameProtocolsBgpPeerGroupOverrIDeCapability.IsUnknown() {
		jsonData["override-capability"] = o.LeafVrfNameProtocolsBgpPeerGroupOverrIDeCapability.ValueString()
	}

	if !o.LeafVrfNameProtocolsBgpPeerGroupPassive.IsNull() && !o.LeafVrfNameProtocolsBgpPeerGroupPassive.IsUnknown() {
		jsonData["passive"] = o.LeafVrfNameProtocolsBgpPeerGroupPassive.ValueString()
	}

	if !o.LeafVrfNameProtocolsBgpPeerGroupPassword.IsNull() && !o.LeafVrfNameProtocolsBgpPeerGroupPassword.IsUnknown() {
		jsonData["password"] = o.LeafVrfNameProtocolsBgpPeerGroupPassword.ValueString()
	}

	if !o.LeafVrfNameProtocolsBgpPeerGroupRemoteAs.IsNull() && !o.LeafVrfNameProtocolsBgpPeerGroupRemoteAs.IsUnknown() {
		jsonData["remote-as"] = o.LeafVrfNameProtocolsBgpPeerGroupRemoteAs.ValueString()
	}

	if !o.LeafVrfNameProtocolsBgpPeerGroupShutdown.IsNull() && !o.LeafVrfNameProtocolsBgpPeerGroupShutdown.IsUnknown() {
		jsonData["shutdown"] = o.LeafVrfNameProtocolsBgpPeerGroupShutdown.ValueString()
	}

	if !o.LeafVrfNameProtocolsBgpPeerGroupUpdateSource.IsNull() && !o.LeafVrfNameProtocolsBgpPeerGroupUpdateSource.IsUnknown() {
		jsonData["update-source"] = o.LeafVrfNameProtocolsBgpPeerGroupUpdateSource.ValueString()
	}

	// Tags

	if !reflect.ValueOf(o.TagVrfNameProtocolsBgpPeerGroupLocalAs).IsZero() {
		subJSONStr, err := json.Marshal(o.TagVrfNameProtocolsBgpPeerGroupLocalAs)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["local-as"] = subData
	}

	if !reflect.ValueOf(o.TagVrfNameProtocolsBgpPeerGroupLocalRole).IsZero() {
		subJSONStr, err := json.Marshal(o.TagVrfNameProtocolsBgpPeerGroupLocalRole)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["local-role"] = subData
	}

	// Nodes

	if !reflect.ValueOf(o.NodeVrfNameProtocolsBgpPeerGroupAddressFamily).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameProtocolsBgpPeerGroupAddressFamily)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["address-family"] = subData
	}

	if !reflect.ValueOf(o.NodeVrfNameProtocolsBgpPeerGroupBfd).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameProtocolsBgpPeerGroupBfd)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["bfd"] = subData
	}

	if !reflect.ValueOf(o.NodeVrfNameProtocolsBgpPeerGroupCapability).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameProtocolsBgpPeerGroupCapability)
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

	if !reflect.ValueOf(o.NodeVrfNameProtocolsBgpPeerGroupTTLSecURIty).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameProtocolsBgpPeerGroupTTLSecURIty)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["ttl-security"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsBgpPeerGroup) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["description"]; ok {
		o.LeafVrfNameProtocolsBgpPeerGroupDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpPeerGroupDescrIPtion = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable-capability-negotiation"]; ok {
		o.LeafVrfNameProtocolsBgpPeerGroupDisableCapabilityNegotiation = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpPeerGroupDisableCapabilityNegotiation = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable-connected-check"]; ok {
		o.LeafVrfNameProtocolsBgpPeerGroupDisableConnectedCheck = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpPeerGroupDisableConnectedCheck = basetypes.NewStringNull()
	}

	if value, ok := jsonData["ebgp-multihop"]; ok {
		o.LeafVrfNameProtocolsBgpPeerGroupEbgpMultihop = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpPeerGroupEbgpMultihop = basetypes.NewStringNull()
	}

	if value, ok := jsonData["graceful-restart"]; ok {
		o.LeafVrfNameProtocolsBgpPeerGroupGracefulRestart = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpPeerGroupGracefulRestart = basetypes.NewStringNull()
	}

	if value, ok := jsonData["override-capability"]; ok {
		o.LeafVrfNameProtocolsBgpPeerGroupOverrIDeCapability = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpPeerGroupOverrIDeCapability = basetypes.NewStringNull()
	}

	if value, ok := jsonData["passive"]; ok {
		o.LeafVrfNameProtocolsBgpPeerGroupPassive = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpPeerGroupPassive = basetypes.NewStringNull()
	}

	if value, ok := jsonData["password"]; ok {
		o.LeafVrfNameProtocolsBgpPeerGroupPassword = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpPeerGroupPassword = basetypes.NewStringNull()
	}

	if value, ok := jsonData["remote-as"]; ok {
		o.LeafVrfNameProtocolsBgpPeerGroupRemoteAs = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpPeerGroupRemoteAs = basetypes.NewStringNull()
	}

	if value, ok := jsonData["shutdown"]; ok {
		o.LeafVrfNameProtocolsBgpPeerGroupShutdown = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpPeerGroupShutdown = basetypes.NewStringNull()
	}

	if value, ok := jsonData["update-source"]; ok {
		o.LeafVrfNameProtocolsBgpPeerGroupUpdateSource = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpPeerGroupUpdateSource = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := jsonData["local-as"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.TagVrfNameProtocolsBgpPeerGroupLocalAs = &map[string]VrfNameProtocolsBgpPeerGroupLocalAs{}

		err = json.Unmarshal(subJSONStr, o.TagVrfNameProtocolsBgpPeerGroupLocalAs)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["local-role"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.TagVrfNameProtocolsBgpPeerGroupLocalRole = &map[string]VrfNameProtocolsBgpPeerGroupLocalRole{}

		err = json.Unmarshal(subJSONStr, o.TagVrfNameProtocolsBgpPeerGroupLocalRole)
		if err != nil {
			return err
		}
	}

	// Nodes
	if value, ok := jsonData["address-family"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVrfNameProtocolsBgpPeerGroupAddressFamily = &VrfNameProtocolsBgpPeerGroupAddressFamily{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameProtocolsBgpPeerGroupAddressFamily)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["bfd"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVrfNameProtocolsBgpPeerGroupBfd = &VrfNameProtocolsBgpPeerGroupBfd{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameProtocolsBgpPeerGroupBfd)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["capability"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVrfNameProtocolsBgpPeerGroupCapability = &VrfNameProtocolsBgpPeerGroupCapability{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameProtocolsBgpPeerGroupCapability)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["ttl-security"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVrfNameProtocolsBgpPeerGroupTTLSecURIty = &VrfNameProtocolsBgpPeerGroupTTLSecURIty{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameProtocolsBgpPeerGroupTTLSecURIty)
		if err != nil {
			return err
		}
	}

	return nil
}
