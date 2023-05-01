// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ServiceDhcpServerSharedNetworkNameSubnet describes the resource data model.
type ServiceDhcpServerSharedNetworkNameSubnet struct {
	// LeafNodes
	LeafServiceDhcpServerSharedNetworkNameSubnetBootfileName        types.String `tfsdk:"bootfile_name"`
	LeafServiceDhcpServerSharedNetworkNameSubnetBootfileServer      types.String `tfsdk:"bootfile_server"`
	LeafServiceDhcpServerSharedNetworkNameSubnetBootfileSize        types.String `tfsdk:"bootfile_size"`
	LeafServiceDhcpServerSharedNetworkNameSubnetClientPrefixLength  types.String `tfsdk:"client_prefix_length"`
	LeafServiceDhcpServerSharedNetworkNameSubnetDefaultRouter       types.String `tfsdk:"default_router"`
	LeafServiceDhcpServerSharedNetworkNameSubnetDomainName          types.String `tfsdk:"domain_name"`
	LeafServiceDhcpServerSharedNetworkNameSubnetDomainSearch        types.String `tfsdk:"domain_search"`
	LeafServiceDhcpServerSharedNetworkNameSubnetDescrIPtion         types.String `tfsdk:"description"`
	LeafServiceDhcpServerSharedNetworkNameSubnetNameServer          types.String `tfsdk:"name_server"`
	LeafServiceDhcpServerSharedNetworkNameSubnetEnableFailover      types.String `tfsdk:"enable_failover"`
	LeafServiceDhcpServerSharedNetworkNameSubnetExclude             types.String `tfsdk:"exclude"`
	LeafServiceDhcpServerSharedNetworkNameSubnetIPForwarding        types.String `tfsdk:"ip_forwarding"`
	LeafServiceDhcpServerSharedNetworkNameSubnetLease               types.String `tfsdk:"lease"`
	LeafServiceDhcpServerSharedNetworkNameSubnetNtpServer           types.String `tfsdk:"ntp_server"`
	LeafServiceDhcpServerSharedNetworkNameSubnetPingCheck           types.String `tfsdk:"ping_check"`
	LeafServiceDhcpServerSharedNetworkNameSubnetPopServer           types.String `tfsdk:"pop_server"`
	LeafServiceDhcpServerSharedNetworkNameSubnetServerIDentifier    types.String `tfsdk:"server_identifier"`
	LeafServiceDhcpServerSharedNetworkNameSubnetSMTPServer          types.String `tfsdk:"smtp_server"`
	LeafServiceDhcpServerSharedNetworkNameSubnetIPvsixOnlyPreferred types.String `tfsdk:"ipv6_only_preferred"`
	LeafServiceDhcpServerSharedNetworkNameSubnetSubnetParameters    types.String `tfsdk:"subnet_parameters"`
	LeafServiceDhcpServerSharedNetworkNameSubnetTftpServerName      types.String `tfsdk:"tftp_server_name"`
	LeafServiceDhcpServerSharedNetworkNameSubnetTimeOffset          types.String `tfsdk:"time_offset"`
	LeafServiceDhcpServerSharedNetworkNameSubnetTimeServer          types.String `tfsdk:"time_server"`
	LeafServiceDhcpServerSharedNetworkNameSubnetWinsServer          types.String `tfsdk:"wins_server"`
	LeafServiceDhcpServerSharedNetworkNameSubnetWpadURL             types.String `tfsdk:"wpad_url"`

	// TagNodes
	TagServiceDhcpServerSharedNetworkNameSubnetRange         types.Map `tfsdk:"range"`
	TagServiceDhcpServerSharedNetworkNameSubnetStaticMapping types.Map `tfsdk:"static_mapping"`
	TagServiceDhcpServerSharedNetworkNameSubnetStaticRoute   types.Map `tfsdk:"static_route"`

	// Nodes
	NodeServiceDhcpServerSharedNetworkNameSubnetVendorOption types.Object `tfsdk:"vendor_option"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *ServiceDhcpServerSharedNetworkNameSubnet) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"service", "dhcp-server", "shared-network-name", "subnet"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafServiceDhcpServerSharedNetworkNameSubnetBootfileName.IsNull() || o.LeafServiceDhcpServerSharedNetworkNameSubnetBootfileName.IsUnknown()) {
		vyosData["bootfile-name"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetBootfileName.ValueString()
	}
	if !(o.LeafServiceDhcpServerSharedNetworkNameSubnetBootfileServer.IsNull() || o.LeafServiceDhcpServerSharedNetworkNameSubnetBootfileServer.IsUnknown()) {
		vyosData["bootfile-server"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetBootfileServer.ValueString()
	}
	if !(o.LeafServiceDhcpServerSharedNetworkNameSubnetBootfileSize.IsNull() || o.LeafServiceDhcpServerSharedNetworkNameSubnetBootfileSize.IsUnknown()) {
		vyosData["bootfile-size"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetBootfileSize.ValueString()
	}
	if !(o.LeafServiceDhcpServerSharedNetworkNameSubnetClientPrefixLength.IsNull() || o.LeafServiceDhcpServerSharedNetworkNameSubnetClientPrefixLength.IsUnknown()) {
		vyosData["client-prefix-length"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetClientPrefixLength.ValueString()
	}
	if !(o.LeafServiceDhcpServerSharedNetworkNameSubnetDefaultRouter.IsNull() || o.LeafServiceDhcpServerSharedNetworkNameSubnetDefaultRouter.IsUnknown()) {
		vyosData["default-router"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetDefaultRouter.ValueString()
	}
	if !(o.LeafServiceDhcpServerSharedNetworkNameSubnetDomainName.IsNull() || o.LeafServiceDhcpServerSharedNetworkNameSubnetDomainName.IsUnknown()) {
		vyosData["domain-name"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetDomainName.ValueString()
	}
	if !(o.LeafServiceDhcpServerSharedNetworkNameSubnetDomainSearch.IsNull() || o.LeafServiceDhcpServerSharedNetworkNameSubnetDomainSearch.IsUnknown()) {
		vyosData["domain-search"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetDomainSearch.ValueString()
	}
	if !(o.LeafServiceDhcpServerSharedNetworkNameSubnetDescrIPtion.IsNull() || o.LeafServiceDhcpServerSharedNetworkNameSubnetDescrIPtion.IsUnknown()) {
		vyosData["description"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetDescrIPtion.ValueString()
	}
	if !(o.LeafServiceDhcpServerSharedNetworkNameSubnetNameServer.IsNull() || o.LeafServiceDhcpServerSharedNetworkNameSubnetNameServer.IsUnknown()) {
		vyosData["name-server"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetNameServer.ValueString()
	}
	if !(o.LeafServiceDhcpServerSharedNetworkNameSubnetEnableFailover.IsNull() || o.LeafServiceDhcpServerSharedNetworkNameSubnetEnableFailover.IsUnknown()) {
		vyosData["enable-failover"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetEnableFailover.ValueString()
	}
	if !(o.LeafServiceDhcpServerSharedNetworkNameSubnetExclude.IsNull() || o.LeafServiceDhcpServerSharedNetworkNameSubnetExclude.IsUnknown()) {
		vyosData["exclude"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetExclude.ValueString()
	}
	if !(o.LeafServiceDhcpServerSharedNetworkNameSubnetIPForwarding.IsNull() || o.LeafServiceDhcpServerSharedNetworkNameSubnetIPForwarding.IsUnknown()) {
		vyosData["ip-forwarding"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetIPForwarding.ValueString()
	}
	if !(o.LeafServiceDhcpServerSharedNetworkNameSubnetLease.IsNull() || o.LeafServiceDhcpServerSharedNetworkNameSubnetLease.IsUnknown()) {
		vyosData["lease"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetLease.ValueString()
	}
	if !(o.LeafServiceDhcpServerSharedNetworkNameSubnetNtpServer.IsNull() || o.LeafServiceDhcpServerSharedNetworkNameSubnetNtpServer.IsUnknown()) {
		vyosData["ntp-server"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetNtpServer.ValueString()
	}
	if !(o.LeafServiceDhcpServerSharedNetworkNameSubnetPingCheck.IsNull() || o.LeafServiceDhcpServerSharedNetworkNameSubnetPingCheck.IsUnknown()) {
		vyosData["ping-check"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetPingCheck.ValueString()
	}
	if !(o.LeafServiceDhcpServerSharedNetworkNameSubnetPopServer.IsNull() || o.LeafServiceDhcpServerSharedNetworkNameSubnetPopServer.IsUnknown()) {
		vyosData["pop-server"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetPopServer.ValueString()
	}
	if !(o.LeafServiceDhcpServerSharedNetworkNameSubnetServerIDentifier.IsNull() || o.LeafServiceDhcpServerSharedNetworkNameSubnetServerIDentifier.IsUnknown()) {
		vyosData["server-identifier"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetServerIDentifier.ValueString()
	}
	if !(o.LeafServiceDhcpServerSharedNetworkNameSubnetSMTPServer.IsNull() || o.LeafServiceDhcpServerSharedNetworkNameSubnetSMTPServer.IsUnknown()) {
		vyosData["smtp-server"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetSMTPServer.ValueString()
	}
	if !(o.LeafServiceDhcpServerSharedNetworkNameSubnetIPvsixOnlyPreferred.IsNull() || o.LeafServiceDhcpServerSharedNetworkNameSubnetIPvsixOnlyPreferred.IsUnknown()) {
		vyosData["ipv6-only-preferred"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetIPvsixOnlyPreferred.ValueString()
	}
	if !(o.LeafServiceDhcpServerSharedNetworkNameSubnetSubnetParameters.IsNull() || o.LeafServiceDhcpServerSharedNetworkNameSubnetSubnetParameters.IsUnknown()) {
		vyosData["subnet-parameters"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetSubnetParameters.ValueString()
	}
	if !(o.LeafServiceDhcpServerSharedNetworkNameSubnetTftpServerName.IsNull() || o.LeafServiceDhcpServerSharedNetworkNameSubnetTftpServerName.IsUnknown()) {
		vyosData["tftp-server-name"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetTftpServerName.ValueString()
	}
	if !(o.LeafServiceDhcpServerSharedNetworkNameSubnetTimeOffset.IsNull() || o.LeafServiceDhcpServerSharedNetworkNameSubnetTimeOffset.IsUnknown()) {
		vyosData["time-offset"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetTimeOffset.ValueString()
	}
	if !(o.LeafServiceDhcpServerSharedNetworkNameSubnetTimeServer.IsNull() || o.LeafServiceDhcpServerSharedNetworkNameSubnetTimeServer.IsUnknown()) {
		vyosData["time-server"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetTimeServer.ValueString()
	}
	if !(o.LeafServiceDhcpServerSharedNetworkNameSubnetWinsServer.IsNull() || o.LeafServiceDhcpServerSharedNetworkNameSubnetWinsServer.IsUnknown()) {
		vyosData["wins-server"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetWinsServer.ValueString()
	}
	if !(o.LeafServiceDhcpServerSharedNetworkNameSubnetWpadURL.IsNull() || o.LeafServiceDhcpServerSharedNetworkNameSubnetWpadURL.IsUnknown()) {
		vyosData["wpad-url"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetWpadURL.ValueString()
	}

	// Tags
	if !(o.TagServiceDhcpServerSharedNetworkNameSubnetRange.IsNull() || o.TagServiceDhcpServerSharedNetworkNameSubnetRange.IsUnknown()) {
		subModel := make(map[string]ServiceDhcpServerSharedNetworkNameSubnetRange)
		diags.Append(o.TagServiceDhcpServerSharedNetworkNameSubnetRange.ElementsAs(ctx, &subModel, false)...)

		subData := make(map[string]interface{})
		for k, v := range subModel {
			subData[k] = v.TerraformToVyos(ctx, diags)
		}
		vyosData["range"] = subData
	}
	if !(o.TagServiceDhcpServerSharedNetworkNameSubnetStaticMapping.IsNull() || o.TagServiceDhcpServerSharedNetworkNameSubnetStaticMapping.IsUnknown()) {
		subModel := make(map[string]ServiceDhcpServerSharedNetworkNameSubnetStaticMapping)
		diags.Append(o.TagServiceDhcpServerSharedNetworkNameSubnetStaticMapping.ElementsAs(ctx, &subModel, false)...)

		subData := make(map[string]interface{})
		for k, v := range subModel {
			subData[k] = v.TerraformToVyos(ctx, diags)
		}
		vyosData["static-mapping"] = subData
	}
	if !(o.TagServiceDhcpServerSharedNetworkNameSubnetStaticRoute.IsNull() || o.TagServiceDhcpServerSharedNetworkNameSubnetStaticRoute.IsUnknown()) {
		subModel := make(map[string]ServiceDhcpServerSharedNetworkNameSubnetStaticRoute)
		diags.Append(o.TagServiceDhcpServerSharedNetworkNameSubnetStaticRoute.ElementsAs(ctx, &subModel, false)...)

		subData := make(map[string]interface{})
		for k, v := range subModel {
			subData[k] = v.TerraformToVyos(ctx, diags)
		}
		vyosData["static-route"] = subData
	}

	// Nodes
	if !(o.NodeServiceDhcpServerSharedNetworkNameSubnetVendorOption.IsNull() || o.NodeServiceDhcpServerSharedNetworkNameSubnetVendorOption.IsUnknown()) {
		var subModel ServiceDhcpServerSharedNetworkNameSubnetVendorOption
		diags.Append(o.NodeServiceDhcpServerSharedNetworkNameSubnetVendorOption.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["vendor-option"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ServiceDhcpServerSharedNetworkNameSubnet) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"service", "dhcp-server", "shared-network-name", "subnet"}})

	// Leafs
	if value, ok := vyosData["bootfile-name"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetBootfileName = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetBootfileName = basetypes.NewStringNull()
	}
	if value, ok := vyosData["bootfile-server"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetBootfileServer = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetBootfileServer = basetypes.NewStringNull()
	}
	if value, ok := vyosData["bootfile-size"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetBootfileSize = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetBootfileSize = basetypes.NewStringNull()
	}
	if value, ok := vyosData["client-prefix-length"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetClientPrefixLength = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetClientPrefixLength = basetypes.NewStringNull()
	}
	if value, ok := vyosData["default-router"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetDefaultRouter = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetDefaultRouter = basetypes.NewStringNull()
	}
	if value, ok := vyosData["domain-name"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetDomainName = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetDomainName = basetypes.NewStringNull()
	}
	if value, ok := vyosData["domain-search"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetDomainSearch = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetDomainSearch = basetypes.NewStringNull()
	}
	if value, ok := vyosData["description"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetDescrIPtion = basetypes.NewStringNull()
	}
	if value, ok := vyosData["name-server"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetNameServer = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetNameServer = basetypes.NewStringNull()
	}
	if value, ok := vyosData["enable-failover"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetEnableFailover = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetEnableFailover = basetypes.NewStringNull()
	}
	if value, ok := vyosData["exclude"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetExclude = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetExclude = basetypes.NewStringNull()
	}
	if value, ok := vyosData["ip-forwarding"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetIPForwarding = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetIPForwarding = basetypes.NewStringNull()
	}
	if value, ok := vyosData["lease"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetLease = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetLease = basetypes.NewStringNull()
	}
	if value, ok := vyosData["ntp-server"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetNtpServer = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetNtpServer = basetypes.NewStringNull()
	}
	if value, ok := vyosData["ping-check"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetPingCheck = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetPingCheck = basetypes.NewStringNull()
	}
	if value, ok := vyosData["pop-server"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetPopServer = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetPopServer = basetypes.NewStringNull()
	}
	if value, ok := vyosData["server-identifier"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetServerIDentifier = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetServerIDentifier = basetypes.NewStringNull()
	}
	if value, ok := vyosData["smtp-server"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetSMTPServer = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetSMTPServer = basetypes.NewStringNull()
	}
	if value, ok := vyosData["ipv6-only-preferred"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetIPvsixOnlyPreferred = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetIPvsixOnlyPreferred = basetypes.NewStringNull()
	}
	if value, ok := vyosData["subnet-parameters"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetSubnetParameters = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetSubnetParameters = basetypes.NewStringNull()
	}
	if value, ok := vyosData["tftp-server-name"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetTftpServerName = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetTftpServerName = basetypes.NewStringNull()
	}
	if value, ok := vyosData["time-offset"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetTimeOffset = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetTimeOffset = basetypes.NewStringNull()
	}
	if value, ok := vyosData["time-server"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetTimeServer = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetTimeServer = basetypes.NewStringNull()
	}
	if value, ok := vyosData["wins-server"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetWinsServer = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetWinsServer = basetypes.NewStringNull()
	}
	if value, ok := vyosData["wpad-url"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetWpadURL = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetWpadURL = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := vyosData["range"]; ok {
		data, d := types.MapValueFrom(ctx, types.ObjectType{AttrTypes: ServiceDhcpServerSharedNetworkNameSubnetRange{}.AttributeTypes()}, value.(map[string]interface{}))
		diags.Append(d...)
		o.TagServiceDhcpServerSharedNetworkNameSubnetRange = data
	} else {
		o.TagServiceDhcpServerSharedNetworkNameSubnetRange = basetypes.NewMapNull(types.ObjectType{})
	}
	if value, ok := vyosData["static-mapping"]; ok {
		data, d := types.MapValueFrom(ctx, types.ObjectType{AttrTypes: ServiceDhcpServerSharedNetworkNameSubnetStaticMapping{}.AttributeTypes()}, value.(map[string]interface{}))
		diags.Append(d...)
		o.TagServiceDhcpServerSharedNetworkNameSubnetStaticMapping = data
	} else {
		o.TagServiceDhcpServerSharedNetworkNameSubnetStaticMapping = basetypes.NewMapNull(types.ObjectType{})
	}
	if value, ok := vyosData["static-route"]; ok {
		data, d := types.MapValueFrom(ctx, types.ObjectType{AttrTypes: ServiceDhcpServerSharedNetworkNameSubnetStaticRoute{}.AttributeTypes()}, value.(map[string]interface{}))
		diags.Append(d...)
		o.TagServiceDhcpServerSharedNetworkNameSubnetStaticRoute = data
	} else {
		o.TagServiceDhcpServerSharedNetworkNameSubnetStaticRoute = basetypes.NewMapNull(types.ObjectType{})
	}

	// Nodes
	if value, ok := vyosData["vendor-option"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, ServiceDhcpServerSharedNetworkNameSubnetVendorOption{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeServiceDhcpServerSharedNetworkNameSubnetVendorOption = data

	} else {
		o.NodeServiceDhcpServerSharedNetworkNameSubnetVendorOption = basetypes.NewObjectNull(ServiceDhcpServerSharedNetworkNameSubnetVendorOption{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"service", "dhcp-server", "shared-network-name", "subnet"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ServiceDhcpServerSharedNetworkNameSubnet) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"bootfile_name":        types.StringType,
		"bootfile_server":      types.StringType,
		"bootfile_size":        types.StringType,
		"client_prefix_length": types.StringType,
		"default_router":       types.StringType,
		"domain_name":          types.StringType,
		"domain_search":        types.StringType,
		"description":          types.StringType,
		"name_server":          types.StringType,
		"enable_failover":      types.StringType,
		"exclude":              types.StringType,
		"ip_forwarding":        types.StringType,
		"lease":                types.StringType,
		"ntp_server":           types.StringType,
		"ping_check":           types.StringType,
		"pop_server":           types.StringType,
		"server_identifier":    types.StringType,
		"smtp_server":          types.StringType,
		"ipv6_only_preferred":  types.StringType,
		"subnet_parameters":    types.StringType,
		"tftp_server_name":     types.StringType,
		"time_offset":          types.StringType,
		"time_server":          types.StringType,
		"wins_server":          types.StringType,
		"wpad_url":             types.StringType,

		// Tags
		"range":          types.MapType{ElemType: types.ObjectType{AttrTypes: ServiceDhcpServerSharedNetworkNameSubnetRange{}.AttributeTypes()}},
		"static_mapping": types.MapType{ElemType: types.ObjectType{AttrTypes: ServiceDhcpServerSharedNetworkNameSubnetStaticMapping{}.AttributeTypes()}},
		"static_route":   types.MapType{ElemType: types.ObjectType{AttrTypes: ServiceDhcpServerSharedNetworkNameSubnetStaticRoute{}.AttributeTypes()}},

		// Nodes
		"vendor_option": types.ObjectType{AttrTypes: ServiceDhcpServerSharedNetworkNameSubnetVendorOption{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDhcpServerSharedNetworkNameSubnet) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"bootfile_name": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Bootstrap file name

`,
		},

		"bootfile_server": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Server from which the initial boot file is to be loaded

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Bootfile server IPv4 address  |
|  hostname  |  Bootfile server FQDN  |

`,
		},

		"bootfile_size": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Bootstrap file size

|  Format  |  Description  |
|----------|---------------|
|  u32:1-16  |  Bootstrap file size in 512 byte blocks  |

`,
		},

		"client_prefix_length": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specifies the clients subnet mask as per RFC 950. If unset, subnet declaration is used.

|  Format  |  Description  |
|----------|---------------|
|  u32:0-32  |  DHCP client prefix length must be 0 to 32  |

`,
		},

		"default_router": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address of default router

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Default router IPv4 address  |

`,
		},

		"domain_name": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Client Domain Name

`,
		},

		"domain_search": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Client Domain Name search list

`,
		},

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |

`,
		},

		"name_server": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Domain Name Servers (DNS) addresses

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Domain Name Server (DNS) IPv4 address  |

`,
		},

		"enable_failover": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enable DHCP failover support for this subnet

`,
		},

		"exclude": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address to exclude from DHCP lease range

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address to exclude from lease range  |

`,
		},

		"ip_forwarding": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enable IP forwarding on client

`,
		},

		"lease": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Lease timeout in seconds

|  Format  |  Description  |
|----------|---------------|
|  u32  |  DHCP lease time in seconds  |

`,

			// Default:          stringdefault.StaticString(`86400`),
			Computed: true,
		},

		"ntp_server": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address of NTP server

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  NTP server IPv4 address  |

`,
		},

		"ping_check": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Sends ICMP Echo request to the address being assigned

`,
		},

		"pop_server": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address of POP3 server

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  POP3 server IPv4 address  |

`,
		},

		"server_identifier": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Address for DHCP server identifier

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  DHCP server identifier IPv4 address  |

`,
		},

		"smtp_server": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address of SMTP server

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  SMTP server IPv4 address  |

`,
		},

		"ipv6_only_preferred": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable IPv4 on IPv6 only hosts (RFC 8925)

|  Format  |  Description  |
|----------|---------------|
|  u32  |  Seconds  |

`,
		},

		"subnet_parameters": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Additional subnet parameters for DHCP server. You must use the syntax of dhcpd.conf in this text-field. Using this without proper knowledge may result in a crashed DHCP server. Check system log to look for errors.

`,
		},

		"tftp_server_name": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `TFTP server name

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  TFTP server IPv4 address  |
|  hostname  |  TFTP server FQDN  |

`,
		},

		"time_offset": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Client subnet offset in seconds from Coordinated Universal Time (UTC)

|  Format  |  Description  |
|----------|---------------|
|  [-]N  |  Time offset (number, may be negative)  |

`,
		},

		"time_server": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address of time server

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Time server IPv4 address  |

`,
		},

		"wins_server": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address for Windows Internet Name Service (WINS) server

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  WINS server IPv4 address  |

`,
		},

		"wpad_url": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Web Proxy Autodiscovery (WPAD) URL

`,
		},

		// TagNodes

		"range": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ServiceDhcpServerSharedNetworkNameSubnetRange{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `DHCP lease range

`,
		},

		"static_mapping": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ServiceDhcpServerSharedNetworkNameSubnetStaticMapping{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Name of static mapping

`,
		},

		"static_route": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ServiceDhcpServerSharedNetworkNameSubnetStaticRoute{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Classless static route destination subnet

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  IPv4 address and prefix length  |

`,
		},

		// Nodes

		"vendor_option": schema.SingleNestedAttribute{
			Attributes: ServiceDhcpServerSharedNetworkNameSubnetVendorOption{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Vendor Specific Options

`,
		},
	}
}
