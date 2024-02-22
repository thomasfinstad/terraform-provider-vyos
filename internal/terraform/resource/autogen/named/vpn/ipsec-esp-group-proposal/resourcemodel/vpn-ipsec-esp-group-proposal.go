// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VpnIPsecEspGroupProposal describes the resource data model.
type VpnIPsecEspGroupProposal struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.Number `tfsdk:"proposal_id" vyos:",self-id"`

	ParentIDVpnIPsecEspGroup types.String `tfsdk:"esp_group" vyos:"esp-group,parent-id"`

	// LeafNodes
	LeafVpnIPsecEspGroupProposalEncryption types.String `tfsdk:"encryption" vyos:"encryption,omitempty"`
	LeafVpnIPsecEspGroupProposalHash       types.String `tfsdk:"hash" vyos:"hash,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// SetID configures the resource ID
func (o *VpnIPsecEspGroupProposal) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnIPsecEspGroupProposal) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"vpn",

		"ipsec",

		"esp-group",
		o.ParentIDVpnIPsecEspGroup.ValueString(),

		"proposal",
		o.SelfIdentifier.ValueBigFloat().String(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnIPsecEspGroupProposal) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"proposal_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `ESP group proposal

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  ESP group proposal number  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		"esp_group_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Encapsulating Security Payload (ESP) group name

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"encryption": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Encryption algorithm

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  null  &emsp; |  Null encryption  |
    |  aes128  &emsp; |  128 bit AES-CBC  |
    |  aes192  &emsp; |  192 bit AES-CBC  |
    |  aes256  &emsp; |  256 bit AES-CBC  |
    |  aes128ctr  &emsp; |  128 bit AES-COUNTER  |
    |  aes192ctr  &emsp; |  192 bit AES-COUNTER  |
    |  aes256ctr  &emsp; |  256 bit AES-COUNTER  |
    |  aes128ccm64  &emsp; |  128 bit AES-CCM with 64 bit ICV  |
    |  aes192ccm64  &emsp; |  192 bit AES-CCM with 64 bit ICV  |
    |  aes256ccm64  &emsp; |  256 bit AES-CCM with 64 bit ICV  |
    |  aes128ccm96  &emsp; |  128 bit AES-CCM with 96 bit ICV  |
    |  aes192ccm96  &emsp; |  192 bit AES-CCM with 96 bit ICV  |
    |  aes256ccm96  &emsp; |  256 bit AES-CCM with 96 bit ICV  |
    |  aes128ccm128  &emsp; |  128 bit AES-CCM with 128 bit ICV  |
    |  aes192ccm128  &emsp; |  192 bit AES-CCM with 128 bit IC  |
    |  aes256ccm128  &emsp; |  256 bit AES-CCM with 128 bit ICV  |
    |  aes128gcm64  &emsp; |  128 bit AES-GCM with 64 bit ICV  |
    |  aes192gcm64  &emsp; |  192 bit AES-GCM with 64 bit ICV  |
    |  aes256gcm64  &emsp; |  256 bit AES-GCM with 64 bit ICV  |
    |  aes128gcm96  &emsp; |  128 bit AES-GCM with 96 bit ICV  |
    |  aes192gcm96  &emsp; |  192 bit AES-GCM with 96 bit ICV  |
    |  aes256gcm96  &emsp; |  256 bit AES-GCM with 96 bit ICV  |
    |  aes128gcm128  &emsp; |  128 bit AES-GCM with 128 bit ICV  |
    |  aes192gcm128  &emsp; |  192 bit AES-GCM with 128 bit ICV  |
    |  aes256gcm128  &emsp; |  256 bit AES-GCM with 128 bit ICV  |
    |  aes128gmac  &emsp; |  Null encryption with 128 bit AES-GMAC  |
    |  aes192gmac  &emsp; |  Null encryption with 192 bit AES-GMAC  |
    |  aes256gmac  &emsp; |  Null encryption with 256 bit AES-GMAC  |
    |  3des  &emsp; |  168 bit 3DES-EDE-CBC  |
    |  blowfish128  &emsp; |  128 bit Blowfish-CBC  |
    |  blowfish192  &emsp; |  192 bit Blowfish-CBC  |
    |  blowfish256  &emsp; |  256 bit Blowfish-CBC  |
    |  camellia128  &emsp; |  128 bit Camellia-CBC  |
    |  camellia192  &emsp; |  192 bit Camellia-CBC  |
    |  camellia256  &emsp; |  256 bit Camellia-CBC  |
    |  camellia128ctr  &emsp; |  128 bit Camellia-COUNTER  |
    |  camellia192ctr  &emsp; |  192 bit Camellia-COUNTER  |
    |  camellia256ctr  &emsp; |  256 bit Camellia-COUNTER  |
    |  camellia128ccm64  &emsp; |  128 bit Camellia-CCM with 64 bit ICV  |
    |  camellia192ccm64  &emsp; |  192 bit Camellia-CCM with 64 bit ICV  |
    |  camellia256ccm64  &emsp; |  256 bit Camellia-CCM with 64 bit ICV  |
    |  camellia128ccm96  &emsp; |  128 bit Camellia-CCM with 96 bit ICV  |
    |  camellia192ccm96  &emsp; |  192 bit Camellia-CCM with 96 bit ICV  |
    |  camellia256ccm96  &emsp; |  256 bit Camellia-CCM with 96 bit ICV  |
    |  camellia128ccm128  &emsp; |  128 bit Camellia-CCM with 128 bit ICV  |
    |  camellia192ccm128  &emsp; |  192 bit Camellia-CCM with 128 bit ICV  |
    |  camellia256ccm128  &emsp; |  256 bit Camellia-CCM with 128 bit ICV  |
    |  serpent128  &emsp; |  128 bit Serpent-CBC  |
    |  serpent192  &emsp; |  192 bit Serpent-CBC  |
    |  serpent256  &emsp; |  256 bit Serpent-CBC  |
    |  twofish128  &emsp; |  128 bit Twofish-CBC  |
    |  twofish192  &emsp; |  192 bit Twofish-CBC  |
    |  twofish256  &emsp; |  256 bit Twofish-CBC  |
    |  cast128  &emsp; |  128 bit CAST-CBC  |
    |  chacha20poly1305  &emsp; |  256 bit ChaCha20/Poly1305 with 128 bit ICV  |

`,

			// Default:          stringdefault.StaticString(`aes128`),
			Computed: true,
		},

		"hash": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Hash algorithm

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  md5  &emsp; |  MD5 HMAC  |
    |  md5_128  &emsp; |  MD5_128 HMAC  |
    |  sha1  &emsp; |  SHA1 HMAC  |
    |  sha1_160  &emsp; |  SHA1_160 HMAC  |
    |  sha256  &emsp; |  SHA2_256_128 HMAC  |
    |  sha256_96  &emsp; |  SHA2_256_96 HMAC  |
    |  sha384  &emsp; |  SHA2_384_192 HMAC  |
    |  sha512  &emsp; |  SHA2_512_256 HMAC  |
    |  aesxcbc  &emsp; |  AES XCBC  |
    |  aescmac  &emsp; |  AES CMAC  |
    |  aes128gmac  &emsp; |  128-bit AES-GMAC  |
    |  aes192gmac  &emsp; |  192-bit AES-GMAC  |
    |  aes256gmac  &emsp; |  256-bit AES-GMAC  |

`,

			// Default:          stringdefault.StaticString(`sha1`),
			Computed: true,
		},

		// Nodes

	}
}
