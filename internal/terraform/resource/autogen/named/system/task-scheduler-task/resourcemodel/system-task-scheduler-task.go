// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SystemTaskSchedulerTask describes the resource data model.
type SystemTaskSchedulerTask struct {
	SelfIdentifier types.String `tfsdk:"task_id" vyos:",self-id"`

	// LeafNodes
	LeafSystemTaskSchedulerTaskCrontabSpec types.String `tfsdk:"crontab_spec" vyos:"crontab-spec,omitempty"`
	LeafSystemTaskSchedulerTaskInterval    types.String `tfsdk:"interval" vyos:"interval,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeSystemTaskSchedulerTaskExecutable *SystemTaskSchedulerTaskExecutable `tfsdk:"executable" vyos:"executable,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *SystemTaskSchedulerTask) GetVyosPath() []string {
	return []string{
		"system",

		"task-scheduler",

		"task",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemTaskSchedulerTask) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, an amalgamation of the `task_id` and the parents `*_id` fields seperated by dunder `__` starting with top level ancestor.",
		},
		"task_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Scheduled task

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Task name  |

`,
		},

		// LeafNodes

		"crontab_spec": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `UNIX crontab time specification string

`,
		},

		"interval": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Execution interval

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  <minutes>  &emsp; |  Execution interval in minutes  |
    |  <minutes>m  &emsp; |  Execution interval in minutes  |
    |  <hours>h  &emsp; |  Execution interval in hours  |
    |  <days>d  &emsp; |  Execution interval in days  |

`,
		},

		// Nodes

		"executable": schema.SingleNestedAttribute{
			Attributes: SystemTaskSchedulerTaskExecutable{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Executable path and arguments

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *SystemTaskSchedulerTask) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *SystemTaskSchedulerTask) UnmarshalJSON(_ []byte) error {
	return nil
}
