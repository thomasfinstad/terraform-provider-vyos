// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// SystemSyslogGlobalFacility describes the resource data model.
type SystemSyslogGlobalFacility struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafSystemSyslogGlobalFacilityLevel types.String `tfsdk:"level" json:"level,omitempty"`

	// TagNodes

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *SystemSyslogGlobalFacility) GetVyosPath() []string {
	return []string{
		"system",
		"syslog",
		"global",
		"facility",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemSyslogGlobalFacility) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
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

		// LeafNodes

		"level": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Logging level

|  Format  |  Description  |
|----------|---------------|
|  emerg  |  Emergency messages  |
|  alert  |  Urgent messages  |
|  crit  |  Critical messages  |
|  err  |  Error messages  |
|  warning  |  Warning messages  |
|  notice  |  Messages for further investigation  |
|  info  |  Informational messages  |
|  debug  |  Debug messages  |
|  all  |  Log everything  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *SystemSyslogGlobalFacility) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafSystemSyslogGlobalFacilityLevel.IsNull() && !o.LeafSystemSyslogGlobalFacilityLevel.IsUnknown() {
		jsonData["level"] = o.LeafSystemSyslogGlobalFacilityLevel.ValueString()
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
func (o *SystemSyslogGlobalFacility) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["level"]; ok {
		o.LeafSystemSyslogGlobalFacilityLevel = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafSystemSyslogGlobalFacilityLevel = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
