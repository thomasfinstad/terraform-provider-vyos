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

// ContainerName describes the resource data model.
type ContainerName struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafContainerNameAllowHostNetworks types.String `tfsdk:"allow_host_networks"`
	LeafContainerNameCapAdd            types.String `tfsdk:"cap_add"`
	LeafContainerNameDescrIPtion       types.String `tfsdk:"description"`
	LeafContainerNameDisable           types.String `tfsdk:"disable"`
	LeafContainerNameEntrypoint        types.String `tfsdk:"entrypoint"`
	LeafContainerNameHostName          types.String `tfsdk:"host_name"`
	LeafContainerNameImage             types.String `tfsdk:"image"`
	LeafContainerNameCommand           types.String `tfsdk:"command"`
	LeafContainerNameArguments         types.String `tfsdk:"arguments"`
	LeafContainerNameMemory            types.String `tfsdk:"memory"`
	LeafContainerNameSharedMemory      types.String `tfsdk:"shared_memory"`
	LeafContainerNameRestart           types.String `tfsdk:"restart"`

	// TagNodes
	TagContainerNameDevice      types.Map `tfsdk:"device"`
	TagContainerNameEnvironment types.Map `tfsdk:"environment"`
	TagContainerNameNetwork     types.Map `tfsdk:"network"`
	TagContainerNamePort        types.Map `tfsdk:"port"`
	TagContainerNameVolume      types.Map `tfsdk:"volume"`

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ContainerName) GetVyosPath() []string {
	return []string{
		"container",
		"name",
		o.ID.ValueString(),
	}
}

// TerraformToVyos converts terraform data to vyos data
func (o *ContainerName) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"container", "name"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafContainerNameAllowHostNetworks.IsNull() || o.LeafContainerNameAllowHostNetworks.IsUnknown()) {
		vyosData["allow-host-networks"] = o.LeafContainerNameAllowHostNetworks.ValueString()
	}
	if !(o.LeafContainerNameCapAdd.IsNull() || o.LeafContainerNameCapAdd.IsUnknown()) {
		vyosData["cap-add"] = o.LeafContainerNameCapAdd.ValueString()
	}
	if !(o.LeafContainerNameDescrIPtion.IsNull() || o.LeafContainerNameDescrIPtion.IsUnknown()) {
		vyosData["description"] = o.LeafContainerNameDescrIPtion.ValueString()
	}
	if !(o.LeafContainerNameDisable.IsNull() || o.LeafContainerNameDisable.IsUnknown()) {
		vyosData["disable"] = o.LeafContainerNameDisable.ValueString()
	}
	if !(o.LeafContainerNameEntrypoint.IsNull() || o.LeafContainerNameEntrypoint.IsUnknown()) {
		vyosData["entrypoint"] = o.LeafContainerNameEntrypoint.ValueString()
	}
	if !(o.LeafContainerNameHostName.IsNull() || o.LeafContainerNameHostName.IsUnknown()) {
		vyosData["host-name"] = o.LeafContainerNameHostName.ValueString()
	}
	if !(o.LeafContainerNameImage.IsNull() || o.LeafContainerNameImage.IsUnknown()) {
		vyosData["image"] = o.LeafContainerNameImage.ValueString()
	}
	if !(o.LeafContainerNameCommand.IsNull() || o.LeafContainerNameCommand.IsUnknown()) {
		vyosData["command"] = o.LeafContainerNameCommand.ValueString()
	}
	if !(o.LeafContainerNameArguments.IsNull() || o.LeafContainerNameArguments.IsUnknown()) {
		vyosData["arguments"] = o.LeafContainerNameArguments.ValueString()
	}
	if !(o.LeafContainerNameMemory.IsNull() || o.LeafContainerNameMemory.IsUnknown()) {
		vyosData["memory"] = o.LeafContainerNameMemory.ValueString()
	}
	if !(o.LeafContainerNameSharedMemory.IsNull() || o.LeafContainerNameSharedMemory.IsUnknown()) {
		vyosData["shared-memory"] = o.LeafContainerNameSharedMemory.ValueString()
	}
	if !(o.LeafContainerNameRestart.IsNull() || o.LeafContainerNameRestart.IsUnknown()) {
		vyosData["restart"] = o.LeafContainerNameRestart.ValueString()
	}

	// Tags
	if !(o.TagContainerNameDevice.IsNull() || o.TagContainerNameDevice.IsUnknown()) {
		subModel := make(map[string]ContainerNameDevice)
		diags.Append(o.TagContainerNameDevice.ElementsAs(ctx, &subModel, false)...)

		subData := make(map[string]interface{})
		for k, v := range subModel {
			subData[k] = v.TerraformToVyos(ctx, diags)
		}
		vyosData["device"] = subData
	}
	if !(o.TagContainerNameEnvironment.IsNull() || o.TagContainerNameEnvironment.IsUnknown()) {
		subModel := make(map[string]ContainerNameEnvironment)
		diags.Append(o.TagContainerNameEnvironment.ElementsAs(ctx, &subModel, false)...)

		subData := make(map[string]interface{})
		for k, v := range subModel {
			subData[k] = v.TerraformToVyos(ctx, diags)
		}
		vyosData["environment"] = subData
	}
	if !(o.TagContainerNameNetwork.IsNull() || o.TagContainerNameNetwork.IsUnknown()) {
		subModel := make(map[string]ContainerNameNetwork)
		diags.Append(o.TagContainerNameNetwork.ElementsAs(ctx, &subModel, false)...)

		subData := make(map[string]interface{})
		for k, v := range subModel {
			subData[k] = v.TerraformToVyos(ctx, diags)
		}
		vyosData["network"] = subData
	}
	if !(o.TagContainerNamePort.IsNull() || o.TagContainerNamePort.IsUnknown()) {
		subModel := make(map[string]ContainerNamePort)
		diags.Append(o.TagContainerNamePort.ElementsAs(ctx, &subModel, false)...)

		subData := make(map[string]interface{})
		for k, v := range subModel {
			subData[k] = v.TerraformToVyos(ctx, diags)
		}
		vyosData["port"] = subData
	}
	if !(o.TagContainerNameVolume.IsNull() || o.TagContainerNameVolume.IsUnknown()) {
		subModel := make(map[string]ContainerNameVolume)
		diags.Append(o.TagContainerNameVolume.ElementsAs(ctx, &subModel, false)...)

		subData := make(map[string]interface{})
		for k, v := range subModel {
			subData[k] = v.TerraformToVyos(ctx, diags)
		}
		vyosData["volume"] = subData
	}

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ContainerName) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"container", "name"}})

	// Leafs
	if value, ok := vyosData["allow-host-networks"]; ok {
		o.LeafContainerNameAllowHostNetworks = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafContainerNameAllowHostNetworks = basetypes.NewStringNull()
	}
	if value, ok := vyosData["cap-add"]; ok {
		o.LeafContainerNameCapAdd = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafContainerNameCapAdd = basetypes.NewStringNull()
	}
	if value, ok := vyosData["description"]; ok {
		o.LeafContainerNameDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafContainerNameDescrIPtion = basetypes.NewStringNull()
	}
	if value, ok := vyosData["disable"]; ok {
		o.LeafContainerNameDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafContainerNameDisable = basetypes.NewStringNull()
	}
	if value, ok := vyosData["entrypoint"]; ok {
		o.LeafContainerNameEntrypoint = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafContainerNameEntrypoint = basetypes.NewStringNull()
	}
	if value, ok := vyosData["host-name"]; ok {
		o.LeafContainerNameHostName = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafContainerNameHostName = basetypes.NewStringNull()
	}
	if value, ok := vyosData["image"]; ok {
		o.LeafContainerNameImage = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafContainerNameImage = basetypes.NewStringNull()
	}
	if value, ok := vyosData["command"]; ok {
		o.LeafContainerNameCommand = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafContainerNameCommand = basetypes.NewStringNull()
	}
	if value, ok := vyosData["arguments"]; ok {
		o.LeafContainerNameArguments = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafContainerNameArguments = basetypes.NewStringNull()
	}
	if value, ok := vyosData["memory"]; ok {
		o.LeafContainerNameMemory = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafContainerNameMemory = basetypes.NewStringNull()
	}
	if value, ok := vyosData["shared-memory"]; ok {
		o.LeafContainerNameSharedMemory = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafContainerNameSharedMemory = basetypes.NewStringNull()
	}
	if value, ok := vyosData["restart"]; ok {
		o.LeafContainerNameRestart = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafContainerNameRestart = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := vyosData["device"]; ok {
		data, d := types.MapValueFrom(ctx, types.ObjectType{AttrTypes: ContainerNameDevice{}.AttributeTypes()}, value.(map[string]interface{}))
		diags.Append(d...)
		o.TagContainerNameDevice = data
	} else {
		o.TagContainerNameDevice = basetypes.NewMapNull(types.ObjectType{})
	}
	if value, ok := vyosData["environment"]; ok {
		data, d := types.MapValueFrom(ctx, types.ObjectType{AttrTypes: ContainerNameEnvironment{}.AttributeTypes()}, value.(map[string]interface{}))
		diags.Append(d...)
		o.TagContainerNameEnvironment = data
	} else {
		o.TagContainerNameEnvironment = basetypes.NewMapNull(types.ObjectType{})
	}
	if value, ok := vyosData["network"]; ok {
		data, d := types.MapValueFrom(ctx, types.ObjectType{AttrTypes: ContainerNameNetwork{}.AttributeTypes()}, value.(map[string]interface{}))
		diags.Append(d...)
		o.TagContainerNameNetwork = data
	} else {
		o.TagContainerNameNetwork = basetypes.NewMapNull(types.ObjectType{})
	}
	if value, ok := vyosData["port"]; ok {
		data, d := types.MapValueFrom(ctx, types.ObjectType{AttrTypes: ContainerNamePort{}.AttributeTypes()}, value.(map[string]interface{}))
		diags.Append(d...)
		o.TagContainerNamePort = data
	} else {
		o.TagContainerNamePort = basetypes.NewMapNull(types.ObjectType{})
	}
	if value, ok := vyosData["volume"]; ok {
		data, d := types.MapValueFrom(ctx, types.ObjectType{AttrTypes: ContainerNameVolume{}.AttributeTypes()}, value.(map[string]interface{}))
		diags.Append(d...)
		o.TagContainerNameVolume = data
	} else {
		o.TagContainerNameVolume = basetypes.NewMapNull(types.ObjectType{})
	}

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"container", "name"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ContainerName) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"allow_host_networks": types.StringType,
		"cap_add":             types.StringType,
		"description":         types.StringType,
		"disable":             types.StringType,
		"entrypoint":          types.StringType,
		"host_name":           types.StringType,
		"image":               types.StringType,
		"command":             types.StringType,
		"arguments":           types.StringType,
		"memory":              types.StringType,
		"shared_memory":       types.StringType,
		"restart":             types.StringType,

		// Tags
		"device":      types.MapType{ElemType: types.ObjectType{AttrTypes: ContainerNameDevice{}.AttributeTypes()}},
		"environment": types.MapType{ElemType: types.ObjectType{AttrTypes: ContainerNameEnvironment{}.AttributeTypes()}},
		"network":     types.MapType{ElemType: types.ObjectType{AttrTypes: ContainerNameNetwork{}.AttributeTypes()}},
		"port":        types.MapType{ElemType: types.ObjectType{AttrTypes: ContainerNamePort{}.AttributeTypes()}},
		"volume":      types.MapType{ElemType: types.ObjectType{AttrTypes: ContainerNameVolume{}.AttributeTypes()}},

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ContainerName) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Container name

`,
		},

		// LeafNodes

		"allow_host_networks": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Allow host networks in container

`,
		},

		"cap_add": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Container capabilities/permissions

|  Format  |  Description  |
|----------|---------------|
|  net-admin  |  Network operations (interface, firewall, routing tables)  |
|  net-bind-service  |  Bind a socket to privileged ports (port numbers less than 1024)  |
|  net-raw  |  Permission to create raw network sockets  |
|  setpcap  |  Capability sets (from bounded or inherited set)  |
|  sys-admin  |  Administation operations (quotactl, mount, sethostname, setdomainame)  |
|  sys-time  |  Permission to set system clock  |

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

		"disable": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
		},

		"entrypoint": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Override the default ENTRYPOINT from the image

`,
		},

		"host_name": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Container host name

`,
		},

		"image": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Image name in the hub-registry

`,
		},

		"command": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Override the default CMD from the image

`,
		},

		"arguments": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `The command's arguments for this container

`,
		},

		"memory": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Memory (RAM) available to this container

|  Format  |  Description  |
|----------|---------------|
|  u32:0  |  Unlimited  |
|  u32:1-16384  |  Container memory in megabytes (MB)  |

`,

			// Default:          stringdefault.StaticString(`512`),
			Computed: true,
		},

		"shared_memory": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Shared memory available to this container

|  Format  |  Description  |
|----------|---------------|
|  u32:0  |  Unlimited  |
|  u32:1-8192  |  Container memory in megabytes (MB)  |

`,

			// Default:          stringdefault.StaticString(`64`),
			Computed: true,
		},

		"restart": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Restart options for container

|  Format  |  Description  |
|----------|---------------|
|  no  |  Do not restart containers on exit  |
|  on-failure  |  Restart containers when they exit with a non-zero exit code, retrying indefinitely  |
|  always  |  Restart containers when they exit, regardless of status, retrying indefinitely  |

`,

			// Default:          stringdefault.StaticString(`on-failure`),
			Computed: true,
		},

		// TagNodes

		"device": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ContainerNameDevice{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Add a host device to the container

`,
		},

		"environment": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ContainerNameEnvironment{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Add custom environment variables

`,
		},

		"network": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ContainerNameNetwork{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Attach user defined network to container

`,
		},

		"port": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ContainerNamePort{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Publish port to the container

`,
		},

		"volume": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ContainerNameVolume{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Mount a volume into the container

`,
		},

		// Nodes

	}
}