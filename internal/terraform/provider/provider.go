package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers/tools"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/provider/data"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen"
)

// Ensure ScaffoldingProvider satisfies various provider interfaces.
var _ provider.Provider = &VyosProvider{}

// VyosProvider defines the provider implementation.
type VyosProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	Version string
}

// VyosProviderModel Contains all global configurations for the provider
type VyosProviderModel struct {
	Endpoint               types.String `tfsdk:"endpoint"`
	Key                    types.String `tfsdk:"api_key"`
	Certificate            types.Object `tfsdk:"certificate"`
	OverwriteExistingRes   types.Bool   `tfsdk:"overwrite_existing_resources_on_create"`
	IgnoreMissingParentRes types.Bool   `tfsdk:"ignore_missing_parent_resource_on_create"`
	IgnoreChildResOnDelete types.Bool   `tfsdk:"ignore_child_resource_on_delete"`
	DefaultTimeouts        types.Number `tfsdk:"default_timeouts"`
}

// Metadata method to define the provider type name for inclusion in each data source and resource type name.
func (p *VyosProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "vyos"
	resp.Version = p.Version
}

// Schema method to define the schema for provider-level configuration.
func (p *VyosProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "VyOS API Endpoint",
				Required:            true,
			},
			"api_key": schema.StringAttribute{
				MarkdownDescription: "VyOS API Key",
				Required:            true,
			},
			"certificate": schema.SingleNestedAttribute{

				Optional: true,
				Attributes: map[string]schema.Attribute{
					"disable_verify": schema.BoolAttribute{
						MarkdownDescription: "Disable remote certificate verification, useful for selfsigned certs.",
						Optional:            true,
					},
				},
			},
			"default_timeouts": schema.NumberAttribute{
				MarkdownDescription: "Default Create/Read/Update/Destroy timeouts in minutes, can be overridden on a per resource basis. If not configured, defaults to 15.",
				Optional:            true,
			},
			"overwrite_existing_resources_on_create": schema.BoolAttribute{
				MarkdownDescription: `Disables the check to see if the resource already exists on the target machine, resulting in possibly overwriting configs without notice.
This can be helpful when trying to avoid and change many resources at once.`,
				Optional: true,
			},
			"ignore_missing_parent_resource_on_create": schema.BoolAttribute{
				MarkdownDescription: `Disables the check to see if the required parent resource exists on the target machine.
This can be helpful when encountering a bug with the provider.`,
				Optional: true,
			},
			"ignore_child_resource_on_delete": schema.BoolAttribute{
				MarkdownDescription: ` !> **WARNING:** This is extremly destructive and will delete everything below the destroyed resource.
Disables the check to see if the resouce has any child resources.
This can be useful when only a parent resource is configured via terraform.
This has no effect on global resources.`,
				Optional: true,
			},
		},
	}
}

// Configure method to configure shared clients for data source and resource implementations.
func (p *VyosProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tools.Info(ctx, "Configuring vyos provider")

	var providerModel VyosProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &providerModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	endpoint := providerModel.Endpoint.ValueString()
	apiKey := providerModel.Key.ValueString()

	// Configuration values validation
	if endpoint == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("endpoint"),
			"VyOS Endpoint",
			"A valid http(s) endpoint is required",
		)
	}

	if apiKey == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"VyOS API Key",
			"API Key required to connect",
		)
	}

	// SSL settings
	disableVerifyAttr := providerModel.Certificate.Attributes()["disable_verify"]

	var disableVerify bool
	if disableVerifyAttr != nil {
		disableVerify = disableVerifyAttr.(types.Bool).ValueBool()
	} else {
		disableVerify = false
	}

	ctxMutilators := data.CtxMutilators(endpoint, apiKey)

	// Run ctx mutilators for client
	for _, fn := range ctxMutilators {
		ctx = fn(ctx)
	}

	// Client configuration for data sources and resources
	config := data.NewProviderData(
		client.NewClient(ctx, endpoint, apiKey, "TODO: add useragent with provider version", disableVerify),
	)

	// Add ctx mutilators to provider data
	config.CtxMutilatorAdd(ctxMutilators...)

	// Default timeout
	var defaultTimeout float64
	if providerModel.DefaultTimeouts.IsNull() || providerModel.DefaultTimeouts.IsUnknown() {
		defaultTimeout = 15
	} else {
		defaultTimeout, _ = providerModel.DefaultTimeouts.ValueBigFloat().Float64()
	}

	if defaultTimeout == 0.0 {
		config.Config.CrudDefaultTimeouts = 15
	} else {
		config.Config.CrudDefaultTimeouts = defaultTimeout
	}

	// CRUD checks
	config.Config.CrudSkipExistingResourceCheck = providerModel.OverwriteExistingRes.ValueBool()
	config.Config.CrudSkipCheckParentBeforeCreate = providerModel.IgnoreMissingParentRes.ValueBool()
	config.Config.CrudSkipCheckChildBeforeDelete = providerModel.IgnoreChildResOnDelete.ValueBool()

	// Send provider data
	resp.DataSourceData = config
	resp.ResourceData = config
}

// Resources method to define the provider's resources.
func (p *VyosProvider) Resources(ctx context.Context) []func() resource.Resource {
	return autogen.GetResources()
}

// DataSources method to define the provider's data sources.
func (p *VyosProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		//NewDataSource,
	}
}

// New method to return the provider generator function
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &VyosProvider{
			Version: version,
		}
	}
}
