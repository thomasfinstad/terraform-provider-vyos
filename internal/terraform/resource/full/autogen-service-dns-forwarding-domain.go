package resourcefull

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &service_dns_forwarding_domain{}

// var _ resource.ResourceWithImportState = &service_dns_forwarding_domain{}

// service_dns_forwarding_domain defines the resource implementation.
type service_dns_forwarding_domain struct {
	client *http.Client
}

// service dns forwarding domainModel describes the resource data model.
type service_dns_forwarding_domainModel struct {
	ConfigurableAttribute types.String `tfsdk:"configurable_attribute"`
	Defaulted             types.String `tfsdk:"defaulted"`
	ID                    types.String `tfsdk:"id"`

	Server types.String `tfsdk:server`

	Addnta types.String `tfsdk:addnta`

	Recursion_desired types.String `tfsdk:recursion_desired`
}

// Metadata method to define the resource type name.
func (r *service_dns_forwarding_domain) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service dns forwarding domain"
}

// service_dns_forwarding_domainResource method to return the example resource reference
func service_dns_forwarding_domainResource() resource.Resource {
	return &service_dns_forwarding_domain{}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r *service_dns_forwarding_domain) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example resource",

		Attributes: map[string]schema.Attribute{
			// TODO handle non-string types
			"server": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Domain Name Server (DNS) to forward queries to

Format: ipv4
Domain Name Server (DNS) IPv4 address
Format: ipv6
Domain Name Server (DNS) IPv6 address
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"addnta": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Add NTA (negative trust anchor) for this domain (must be set if the domain does not support DNSSEC)

`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"recursion_desired": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Set the "recursion desired" bit in requests to the upstream nameserver

`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},
		},
	}
}

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func (r *service_dns_forwarding_domain) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *service_dns_forwarding_domainModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create example, got error: %s", err))
	//     return
	// }

	// For the purposes of this example code, hardcoding a response value to
	// save into the Terraform state.
	data.ID = types.StringValue("example-id")

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "created a resource")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read method to define the logic which refreshes the Terraform state for the resource.
func (r *service_dns_forwarding_domain) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *service_dns_forwarding_domainModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read example, got error: %s", err))
	//     return
	// }

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update method to define the logic which updates the resource and sets the updated Terraform state on success.
func (r *service_dns_forwarding_domain) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *service_dns_forwarding_domainModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update example, got error: %s", err))
	//     return
	// }

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete method to define the logic which deletes the resource and removes the Terraform state on success.
func (r *service_dns_forwarding_domain) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *service_dns_forwarding_domainModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete example, got error: %s", err))
	//     return
	// }
}