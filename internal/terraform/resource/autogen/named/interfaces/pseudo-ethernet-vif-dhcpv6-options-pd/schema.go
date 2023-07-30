// Package namedinterfacespseudoethernetvifdhcpvsixoptionspd code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacespseudoethernetvifdhcpvsixoptionspd

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r interfacesPseudoEthernetVifDhcpvsixOptionsPd) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Pseudo Ethernet Interface (Macvlan)

    |  Format  |  Description  |
    |----------|---------------|
    |  pethN  |  Pseudo Ethernet interface name  |

Virtual Local Area Network (VLAN) ID

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:0-4094  |  Virtual Local Area Network (VLAN) ID  |

DHCPv6 client settings/options

DHCPv6 prefix delegation interface statement

    |  Format  |  Description  |
    |----------|---------------|
    |  instance number  |  Prefix delegation instance (>= 0)  |

`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
