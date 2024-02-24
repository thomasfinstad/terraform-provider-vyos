// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ContainerName describes the resource data model.
type ContainerName struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"name_id" vyos:"-,self-id"`

	// LeafNodes
	LeafContainerNameAllowHostNetworks types.Bool   `tfsdk:"allow_host_networks" vyos:"allow-host-networks,omitempty"`
	LeafContainerNameCapAdd            types.List   `tfsdk:"cap_add" vyos:"cap-add,omitempty"`
	LeafContainerNameDescrIPtion       types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafContainerNameDisable           types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafContainerNameEntrypoint        types.String `tfsdk:"entrypoint" vyos:"entrypoint,omitempty"`
	LeafContainerNameHostName          types.String `tfsdk:"host_name" vyos:"host-name,omitempty"`
	LeafContainerNameImage             types.String `tfsdk:"image" vyos:"image,omitempty"`
	LeafContainerNameCommand           types.String `tfsdk:"command" vyos:"command,omitempty"`
	LeafContainerNameArguments         types.String `tfsdk:"arguments" vyos:"arguments,omitempty"`
	LeafContainerNameMemory            types.Number `tfsdk:"memory" vyos:"memory,omitempty"`
	LeafContainerNameSharedMemory      types.Number `tfsdk:"shared_memory" vyos:"shared-memory,omitempty"`
	LeafContainerNameRestart           types.String `tfsdk:"restart" vyos:"restart,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagContainerNameDevice      bool `tfsdk:"-" vyos:"device,ignore,child"`
	ExistsTagContainerNameEnvironment bool `tfsdk:"-" vyos:"environment,ignore,child"`
	ExistsTagContainerNameNetwork     bool `tfsdk:"-" vyos:"network,ignore,child"`
	ExistsTagContainerNamePort        bool `tfsdk:"-" vyos:"port,ignore,child"`
	ExistsTagContainerNameVolume      bool `tfsdk:"-" vyos:"volume,ignore,child"`

	// Nodes
}

// SetID configures the resource ID
func (o *ContainerName) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ContainerName) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"container",

		"name",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ContainerName) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"name_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Container name

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"allow_host_networks": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Allow host networks in container

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"cap_add": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Container capabilities/permissions

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  net-admin  &emsp; |  Network operations (interface, firewall, routing tables)  |
    |  net-bind-service  &emsp; |  Bind a socket to privileged ports (port numbers less than 1024)  |
    |  net-raw  &emsp; |  Permission to create raw network sockets  |
    |  setpcap  &emsp; |  Capability sets (from bounded or inherited set)  |
    |  sys-admin  &emsp; |  Administation operations (quotactl, mount, sethostname, setdomainame)  |
    |  sys-time  &emsp; |  Permission to set system clock  |

`,
		},

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Description  |

`,
		},

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
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

		"memory": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Memory (RAM) available to this container

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0  &emsp; |  Unlimited  |
    |  number: 1-16384  &emsp; |  Container memory in megabytes (MB)  |

`,

			// Default:          stringdefault.StaticString(`512`),
			Computed: true,
		},

		"shared_memory": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Shared memory available to this container

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0  &emsp; |  Unlimited  |
    |  number: 1-8192  &emsp; |  Container memory in megabytes (MB)  |

`,

			// Default:          stringdefault.StaticString(`64`),
			Computed: true,
		},

		"restart": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Restart options for container

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  no  &emsp; |  Do not restart containers on exit  |
    |  on-failure  &emsp; |  Restart containers when they exit with a non-zero exit code, retrying indefinitely  |
    |  always  &emsp; |  Restart containers when they exit, regardless of status, retrying indefinitely  |

`,

			// Default:          stringdefault.StaticString(`on-failure`),
			Computed: true,
		},

		// Nodes

	}
}
