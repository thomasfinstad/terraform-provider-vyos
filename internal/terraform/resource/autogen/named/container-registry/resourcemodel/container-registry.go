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

// ContainerRegistry describes the resource data model.
type ContainerRegistry struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafContainerRegistryDisable types.String `tfsdk:"disable"`

	// TagNodes

	// Nodes
	NodeContainerRegistryAuthentication types.Object `tfsdk:"authentication"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ContainerRegistry) GetVyosPath() []string {
	return []string{
		"container",
		"registry",
		o.ID.ValueString(),
	}
}

// TerraformToVyos converts terraform data to vyos data
func (o *ContainerRegistry) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"container", "registry"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafContainerRegistryDisable.IsNull() || o.LeafContainerRegistryDisable.IsUnknown()) {
		vyosData["disable"] = o.LeafContainerRegistryDisable.ValueString()
	}

	// Tags

	// Nodes
	if !(o.NodeContainerRegistryAuthentication.IsNull() || o.NodeContainerRegistryAuthentication.IsUnknown()) {
		var subModel ContainerRegistryAuthentication
		diags.Append(o.NodeContainerRegistryAuthentication.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["authentication"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ContainerRegistry) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"container", "registry"}})

	// Leafs
	if value, ok := vyosData["disable"]; ok {
		o.LeafContainerRegistryDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafContainerRegistryDisable = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := vyosData["authentication"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, ContainerRegistryAuthentication{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeContainerRegistryAuthentication = data

	} else {
		o.NodeContainerRegistryAuthentication = basetypes.NewObjectNull(ContainerRegistryAuthentication{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"container", "registry"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ContainerRegistry) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"disable": types.StringType,

		// Tags

		// Nodes
		"authentication": types.ObjectType{AttrTypes: ContainerRegistryAuthentication{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ContainerRegistry) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Registry Name

`,
		},

		// LeafNodes

		"disable": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
		},

		// TagNodes

		// Nodes

		"authentication": schema.SingleNestedAttribute{
			Attributes: ContainerRegistryAuthentication{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Authentication settings

`,
		},
	}
}
