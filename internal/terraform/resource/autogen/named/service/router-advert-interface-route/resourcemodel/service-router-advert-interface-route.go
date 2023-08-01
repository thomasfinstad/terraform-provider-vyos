// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ServiceRouterAdvertInterfaceRoute describes the resource data model.
type ServiceRouterAdvertInterfaceRoute struct {
	SelfIdentifier types.String `tfsdk:"route_id" vyos:",self-id"`

	ParentIDServiceRouterAdvertInterface types.String `tfsdk:"interface" vyos:"interface,parent-id"`

	// LeafNodes
	LeafServiceRouterAdvertInterfaceRouteValIDLifetime   types.String `tfsdk:"valid_lifetime" vyos:"valid-lifetime,omitempty"`
	LeafServiceRouterAdvertInterfaceRouteRoutePreference types.String `tfsdk:"route_preference" vyos:"route-preference,omitempty"`
	LeafServiceRouterAdvertInterfaceRouteNoRemoveRoute   types.Bool   `tfsdk:"no_remove_route" vyos:"no-remove-route,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceRouterAdvertInterfaceRoute) GetVyosPath() []string {
	return []string{
		"service",

		"router-advert",

		"interface",
		o.ParentIDServiceRouterAdvertInterface.ValueString(),

		"route",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceRouterAdvertInterfaceRoute) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, an amalgamation of the `route_id` and the parents `*_id` fields seperated by dunder `__` starting with top level ancestor.",
		},
		"route_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `IPv6 route to be advertised in Router Advertisements (RAs)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv6net  &emsp; |  IPv6 route to be advertized  |

`,
		},

		"interface_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Interface to send RA on

`,
		},

		// LeafNodes

		"valid_lifetime": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Time in seconds that the route will remain valid

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-4294967295  &emsp; |  Time in seconds that the route will remain valid  |
    |  infinity  &emsp; |  Route will remain preferred forever  |

`,

			// Default:          stringdefault.StaticString(`1800`),
			Computed: true,
		},

		"route_preference": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Preference associated with the route,

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  low  &emsp; |  Route has low preference  |
    |  medium  &emsp; |  Route has medium preference  |
    |  high  &emsp; |  Route has high preference  |

`,

			// Default:          stringdefault.StaticString(`medium`),
			Computed: true,
		},

		"no_remove_route": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not announce this route with a zero second lifetime upon shutdown

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServiceRouterAdvertInterfaceRoute) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ServiceRouterAdvertInterfaceRoute) UnmarshalJSON(_ []byte) error {
	return nil
}
