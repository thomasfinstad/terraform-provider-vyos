// Package namedprotocolsospfneighbor code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsospfneighbor

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func (r protocolsOspfNeighbor) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
}

// Read method to define the logic which refreshes the Terraform state for the resource.
func (r protocolsOspfNeighbor) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

// Update method to define the logic which updates the resource and sets the updated Terraform state on success.
func (r protocolsOspfNeighbor) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	data := r.model

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
func (r protocolsOspfNeighbor) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	data := r.model

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
