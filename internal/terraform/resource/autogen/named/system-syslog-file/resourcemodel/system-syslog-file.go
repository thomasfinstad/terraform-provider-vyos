// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SystemSyslogFile describes the resource data model.
type SystemSyslogFile struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes

	// TagNodes
	TagSystemSyslogFileFacility *map[string]SystemSyslogFileFacility `tfsdk:"facility" json:"facility,omitempty"`

	// Nodes
	NodeSystemSyslogFileArchive *SystemSyslogFileArchive `tfsdk:"archive" json:"archive,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *SystemSyslogFile) GetVyosPath() []string {
	return []string{
		"system",
		"syslog",
		"file",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemSyslogFile) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Logging to a file

`,
		},

		// LeafNodes

		// TagNodes

		"facility": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: SystemSyslogFileFacility{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Facility for logging

|  Format  |  Description  |
|----------|---------------|
|  all  |  All facilities excluding "mark"  |
|  auth  |  Authentication and authorization  |
|  authpriv  |  Non-system authorization  |
|  cron  |  Cron daemon  |
|  daemon  |  System daemons  |
|  kern  |  Kernel  |
|  lpr  |  Line printer spooler  |
|  mail  |  Mail subsystem  |
|  mark  |  Timestamp  |
|  news  |  USENET subsystem  |
|  protocols  |  depricated will be set to local7  |
|  security  |  depricated will be set to auth  |
|  syslog  |  Authentication and authorization  |
|  user  |  Application processes  |
|  uucp  |  UUCP subsystem  |
|  local0  |  Local facility 0  |
|  local1  |  Local facility 1  |
|  local2  |  Local facility 2  |
|  local3  |  Local facility 3  |
|  local4  |  Local facility 4  |
|  local5  |  Local facility 5  |
|  local6  |  Local facility 6  |
|  local7  |  Local facility 7  |

`,
		},

		// Nodes

		"archive": schema.SingleNestedAttribute{
			Attributes: SystemSyslogFileArchive{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Log file size and rotation characteristics

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *SystemSyslogFile) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	// Tags

	if !reflect.ValueOf(o.TagSystemSyslogFileFacility).IsZero() {
		subJSONStr, err := json.Marshal(o.TagSystemSyslogFileFacility)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["facility"] = subData
	}

	// Nodes

	if !reflect.ValueOf(o.NodeSystemSyslogFileArchive).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeSystemSyslogFileArchive)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["archive"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *SystemSyslogFile) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	// Tags
	if value, ok := jsonData["facility"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.TagSystemSyslogFileFacility = &map[string]SystemSyslogFileFacility{}

		err = json.Unmarshal(subJSONStr, o.TagSystemSyslogFileFacility)
		if err != nil {
			return err
		}
	}

	// Nodes
	if value, ok := jsonData["archive"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeSystemSyslogFileArchive = &SystemSyslogFileArchive{}

		err = json.Unmarshal(subJSONStr, o.NodeSystemSyslogFileArchive)
		if err != nil {
			return err
		}
	}

	return nil
}
