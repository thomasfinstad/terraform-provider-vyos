// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ServiceWebproxyURLFilteringSquIDguardRule describes the resource data model.
type ServiceWebproxyURLFilteringSquIDguardRule struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafServiceWebproxyURLFilteringSquIDguardRuleAllowCategory     types.String `tfsdk:"allow_category" json:"allow-category,omitempty"`
	LeafServiceWebproxyURLFilteringSquIDguardRuleAllowIPaddrURL    types.String `tfsdk:"allow_ipaddr_url" json:"allow-ipaddr-url,omitempty"`
	LeafServiceWebproxyURLFilteringSquIDguardRuleBlockCategory     types.String `tfsdk:"block_category" json:"block-category,omitempty"`
	LeafServiceWebproxyURLFilteringSquIDguardRuleDefaultAction     types.String `tfsdk:"default_action" json:"default-action,omitempty"`
	LeafServiceWebproxyURLFilteringSquIDguardRuleEnableSafeSearch  types.String `tfsdk:"enable_safe_search" json:"enable-safe-search,omitempty"`
	LeafServiceWebproxyURLFilteringSquIDguardRuleLocalBlockKeyword types.String `tfsdk:"local_block_keyword" json:"local-block-keyword,omitempty"`
	LeafServiceWebproxyURLFilteringSquIDguardRuleLocalBlockURL     types.String `tfsdk:"local_block_url" json:"local-block-url,omitempty"`
	LeafServiceWebproxyURLFilteringSquIDguardRuleLocalBlock        types.String `tfsdk:"local_block" json:"local-block,omitempty"`
	LeafServiceWebproxyURLFilteringSquIDguardRuleLocalOkURL        types.String `tfsdk:"local_ok_url" json:"local-ok-url,omitempty"`
	LeafServiceWebproxyURLFilteringSquIDguardRuleLocalOk           types.String `tfsdk:"local_ok" json:"local-ok,omitempty"`
	LeafServiceWebproxyURLFilteringSquIDguardRuleLog               types.String `tfsdk:"log" json:"log,omitempty"`
	LeafServiceWebproxyURLFilteringSquIDguardRuleRedirectURL       types.String `tfsdk:"redirect_url" json:"redirect-url,omitempty"`
	LeafServiceWebproxyURLFilteringSquIDguardRuleSourceGroup       types.String `tfsdk:"source_group" json:"source-group,omitempty"`
	LeafServiceWebproxyURLFilteringSquIDguardRuleTimePeriod        types.String `tfsdk:"time_period" json:"time-period,omitempty"`

	// TagNodes

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceWebproxyURLFilteringSquIDguardRule) GetVyosPath() []string {
	return []string{
		"service",
		"webproxy",
		"url-filtering",
		"squidguard",
		"rule",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceWebproxyURLFilteringSquIDguardRule) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `URL filter rule for a source-group

|  Format  |  Description  |
|----------|---------------|
|  u32:1-1024  |  Rule Number  |

`,
		},

		// LeafNodes

		"allow_category": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Category to allow

`,
		},

		"allow_ipaddr_url": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Allow IP address URLs

`,
		},

		"block_category": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Category to block

`,
		},

		"default_action": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Default action (default: allow)

|  Format  |  Description  |
|----------|---------------|
|  allow  |  Default filter action is allow)  |
|  block  |  Default filter action is block  |

`,
		},

		"enable_safe_search": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enable safe-mode search on popular search engines

`,
		},

		"local_block_keyword": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Local keyword to block

|  Format  |  Description  |
|----------|---------------|
|  keyword  |  Keyword (or regex) to block  |

`,
		},

		"local_block_url": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Local URL to block

|  Format  |  Description  |
|----------|---------------|
|  url  |  Local URL to block (without "http://")  |

`,
		},

		"local_block": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Local site to block

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IP address of site to block  |

`,
		},

		"local_ok_url": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Local URL to allow

|  Format  |  Description  |
|----------|---------------|
|  url  |  Local URL to allow (without "http://")  |

`,
		},

		"local_ok": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Local site to allow

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IP address of site to allow  |

`,
		},

		"log": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Log block category

`,
		},

		"redirect_url": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Redirect URL for filtered websites

|  Format  |  Description  |
|----------|---------------|
|  url  |  URL for redirect  |

`,
		},

		"source_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Source-group for this rule

|  Format  |  Description  |
|----------|---------------|
|  group  |  Source group identifier for this rule  |

`,
		},

		"time_period": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Time-period for this rule

|  Format  |  Description  |
|----------|---------------|
|  period  |  Time period for this rule  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServiceWebproxyURLFilteringSquIDguardRule) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafServiceWebproxyURLFilteringSquIDguardRuleAllowCategory.IsNull() && !o.LeafServiceWebproxyURLFilteringSquIDguardRuleAllowCategory.IsUnknown() {
		jsonData["allow-category"] = o.LeafServiceWebproxyURLFilteringSquIDguardRuleAllowCategory.ValueString()
	}

	if !o.LeafServiceWebproxyURLFilteringSquIDguardRuleAllowIPaddrURL.IsNull() && !o.LeafServiceWebproxyURLFilteringSquIDguardRuleAllowIPaddrURL.IsUnknown() {
		jsonData["allow-ipaddr-url"] = o.LeafServiceWebproxyURLFilteringSquIDguardRuleAllowIPaddrURL.ValueString()
	}

	if !o.LeafServiceWebproxyURLFilteringSquIDguardRuleBlockCategory.IsNull() && !o.LeafServiceWebproxyURLFilteringSquIDguardRuleBlockCategory.IsUnknown() {
		jsonData["block-category"] = o.LeafServiceWebproxyURLFilteringSquIDguardRuleBlockCategory.ValueString()
	}

	if !o.LeafServiceWebproxyURLFilteringSquIDguardRuleDefaultAction.IsNull() && !o.LeafServiceWebproxyURLFilteringSquIDguardRuleDefaultAction.IsUnknown() {
		jsonData["default-action"] = o.LeafServiceWebproxyURLFilteringSquIDguardRuleDefaultAction.ValueString()
	}

	if !o.LeafServiceWebproxyURLFilteringSquIDguardRuleEnableSafeSearch.IsNull() && !o.LeafServiceWebproxyURLFilteringSquIDguardRuleEnableSafeSearch.IsUnknown() {
		jsonData["enable-safe-search"] = o.LeafServiceWebproxyURLFilteringSquIDguardRuleEnableSafeSearch.ValueString()
	}

	if !o.LeafServiceWebproxyURLFilteringSquIDguardRuleLocalBlockKeyword.IsNull() && !o.LeafServiceWebproxyURLFilteringSquIDguardRuleLocalBlockKeyword.IsUnknown() {
		jsonData["local-block-keyword"] = o.LeafServiceWebproxyURLFilteringSquIDguardRuleLocalBlockKeyword.ValueString()
	}

	if !o.LeafServiceWebproxyURLFilteringSquIDguardRuleLocalBlockURL.IsNull() && !o.LeafServiceWebproxyURLFilteringSquIDguardRuleLocalBlockURL.IsUnknown() {
		jsonData["local-block-url"] = o.LeafServiceWebproxyURLFilteringSquIDguardRuleLocalBlockURL.ValueString()
	}

	if !o.LeafServiceWebproxyURLFilteringSquIDguardRuleLocalBlock.IsNull() && !o.LeafServiceWebproxyURLFilteringSquIDguardRuleLocalBlock.IsUnknown() {
		jsonData["local-block"] = o.LeafServiceWebproxyURLFilteringSquIDguardRuleLocalBlock.ValueString()
	}

	if !o.LeafServiceWebproxyURLFilteringSquIDguardRuleLocalOkURL.IsNull() && !o.LeafServiceWebproxyURLFilteringSquIDguardRuleLocalOkURL.IsUnknown() {
		jsonData["local-ok-url"] = o.LeafServiceWebproxyURLFilteringSquIDguardRuleLocalOkURL.ValueString()
	}

	if !o.LeafServiceWebproxyURLFilteringSquIDguardRuleLocalOk.IsNull() && !o.LeafServiceWebproxyURLFilteringSquIDguardRuleLocalOk.IsUnknown() {
		jsonData["local-ok"] = o.LeafServiceWebproxyURLFilteringSquIDguardRuleLocalOk.ValueString()
	}

	if !o.LeafServiceWebproxyURLFilteringSquIDguardRuleLog.IsNull() && !o.LeafServiceWebproxyURLFilteringSquIDguardRuleLog.IsUnknown() {
		jsonData["log"] = o.LeafServiceWebproxyURLFilteringSquIDguardRuleLog.ValueString()
	}

	if !o.LeafServiceWebproxyURLFilteringSquIDguardRuleRedirectURL.IsNull() && !o.LeafServiceWebproxyURLFilteringSquIDguardRuleRedirectURL.IsUnknown() {
		jsonData["redirect-url"] = o.LeafServiceWebproxyURLFilteringSquIDguardRuleRedirectURL.ValueString()
	}

	if !o.LeafServiceWebproxyURLFilteringSquIDguardRuleSourceGroup.IsNull() && !o.LeafServiceWebproxyURLFilteringSquIDguardRuleSourceGroup.IsUnknown() {
		jsonData["source-group"] = o.LeafServiceWebproxyURLFilteringSquIDguardRuleSourceGroup.ValueString()
	}

	if !o.LeafServiceWebproxyURLFilteringSquIDguardRuleTimePeriod.IsNull() && !o.LeafServiceWebproxyURLFilteringSquIDguardRuleTimePeriod.IsUnknown() {
		jsonData["time-period"] = o.LeafServiceWebproxyURLFilteringSquIDguardRuleTimePeriod.ValueString()
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
func (o *ServiceWebproxyURLFilteringSquIDguardRule) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["allow-category"]; ok {
		o.LeafServiceWebproxyURLFilteringSquIDguardRuleAllowCategory = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceWebproxyURLFilteringSquIDguardRuleAllowCategory = basetypes.NewStringNull()
	}

	if value, ok := jsonData["allow-ipaddr-url"]; ok {
		o.LeafServiceWebproxyURLFilteringSquIDguardRuleAllowIPaddrURL = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceWebproxyURLFilteringSquIDguardRuleAllowIPaddrURL = basetypes.NewStringNull()
	}

	if value, ok := jsonData["block-category"]; ok {
		o.LeafServiceWebproxyURLFilteringSquIDguardRuleBlockCategory = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceWebproxyURLFilteringSquIDguardRuleBlockCategory = basetypes.NewStringNull()
	}

	if value, ok := jsonData["default-action"]; ok {
		o.LeafServiceWebproxyURLFilteringSquIDguardRuleDefaultAction = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceWebproxyURLFilteringSquIDguardRuleDefaultAction = basetypes.NewStringNull()
	}

	if value, ok := jsonData["enable-safe-search"]; ok {
		o.LeafServiceWebproxyURLFilteringSquIDguardRuleEnableSafeSearch = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceWebproxyURLFilteringSquIDguardRuleEnableSafeSearch = basetypes.NewStringNull()
	}

	if value, ok := jsonData["local-block-keyword"]; ok {
		o.LeafServiceWebproxyURLFilteringSquIDguardRuleLocalBlockKeyword = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceWebproxyURLFilteringSquIDguardRuleLocalBlockKeyword = basetypes.NewStringNull()
	}

	if value, ok := jsonData["local-block-url"]; ok {
		o.LeafServiceWebproxyURLFilteringSquIDguardRuleLocalBlockURL = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceWebproxyURLFilteringSquIDguardRuleLocalBlockURL = basetypes.NewStringNull()
	}

	if value, ok := jsonData["local-block"]; ok {
		o.LeafServiceWebproxyURLFilteringSquIDguardRuleLocalBlock = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceWebproxyURLFilteringSquIDguardRuleLocalBlock = basetypes.NewStringNull()
	}

	if value, ok := jsonData["local-ok-url"]; ok {
		o.LeafServiceWebproxyURLFilteringSquIDguardRuleLocalOkURL = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceWebproxyURLFilteringSquIDguardRuleLocalOkURL = basetypes.NewStringNull()
	}

	if value, ok := jsonData["local-ok"]; ok {
		o.LeafServiceWebproxyURLFilteringSquIDguardRuleLocalOk = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceWebproxyURLFilteringSquIDguardRuleLocalOk = basetypes.NewStringNull()
	}

	if value, ok := jsonData["log"]; ok {
		o.LeafServiceWebproxyURLFilteringSquIDguardRuleLog = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceWebproxyURLFilteringSquIDguardRuleLog = basetypes.NewStringNull()
	}

	if value, ok := jsonData["redirect-url"]; ok {
		o.LeafServiceWebproxyURLFilteringSquIDguardRuleRedirectURL = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceWebproxyURLFilteringSquIDguardRuleRedirectURL = basetypes.NewStringNull()
	}

	if value, ok := jsonData["source-group"]; ok {
		o.LeafServiceWebproxyURLFilteringSquIDguardRuleSourceGroup = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceWebproxyURLFilteringSquIDguardRuleSourceGroup = basetypes.NewStringNull()
	}

	if value, ok := jsonData["time-period"]; ok {
		o.LeafServiceWebproxyURLFilteringSquIDguardRuleTimePeriod = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceWebproxyURLFilteringSquIDguardRuleTimePeriod = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
