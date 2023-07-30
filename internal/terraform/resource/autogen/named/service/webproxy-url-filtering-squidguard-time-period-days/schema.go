// Package namedservicewebproxyurlfilteringsquidguardtimeperioddays code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicewebproxyurlfilteringsquidguardtimeperioddays

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r serviceWebproxyURLFilteringSquIDguardTimePeriodDays) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Webproxy service settings

URL filtering settings

URL filtering via squidGuard redirector

Time period name

Time-period days

    |  Format  |  Description  |
    |----------|---------------|
    |  Sun  |  Sunday  |
    |  Mon  |  Monday  |
    |  Tue  |  Tuesday  |
    |  Wed  |  Wednesday  |
    |  Thu  |  Thursday  |
    |  Fri  |  Friday  |
    |  Sat  |  Saturday  |
    |  weekdays  |  Monday through Friday  |
    |  weekend  |  Saturday and Sunday  |
    |  all  |  All days of the week  |

`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
