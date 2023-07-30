// Package namedinterfacesethernetvifsdhcpvsixoptionspdinterface code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesethernetvifsdhcpvsixoptionspdinterface

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r interfacesEthernetVifSDhcpvsixOptionsPdInterface) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Network interfaces

Ethernet Interface

    |  Format  |  Description  |
    |----------|---------------|
    |  ethN  |  Ethernet interface name  |

QinQ TAG-S Virtual Local Area Network (VLAN) ID

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:0-4094  |  QinQ Virtual Local Area Network (VLAN) ID  |

DHCPv6 client settings/options

DHCPv6 prefix delegation interface statement

    |  Format  |  Description  |
    |----------|---------------|
    |  instance number  |  Prefix delegation instance (>= 0)  |

Delegate IPv6 prefix from provider to this interface

`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
