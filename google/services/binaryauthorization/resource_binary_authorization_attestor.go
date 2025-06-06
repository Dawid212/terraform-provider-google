// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/binaryauthorization/Attestor.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package binaryauthorization

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func CompareSignatureAlgorithm(_, old, new string, _ *schema.ResourceData) bool {
	// See https://cloud.google.com/binary-authorization/docs/reference/rest/v1/projects.attestors#signaturealgorithm
	normalizedAlgorithms := map[string]string{
		"ECDSA_P256_SHA256":   "ECDSA_P256_SHA256",
		"EC_SIGN_P256_SHA256": "ECDSA_P256_SHA256",
		"ECDSA_P384_SHA384":   "ECDSA_P384_SHA384",
		"EC_SIGN_P384_SHA384": "ECDSA_P384_SHA384",
		"ECDSA_P521_SHA512":   "ECDSA_P521_SHA512",
		"EC_SIGN_P521_SHA512": "ECDSA_P521_SHA512",
	}

	normalizedOld := old
	normalizedNew := new

	if normalized, ok := normalizedAlgorithms[old]; ok {
		normalizedOld = normalized
	}
	if normalized, ok := normalizedAlgorithms[new]; ok {
		normalizedNew = normalized
	}

	if normalizedNew == normalizedOld {
		return true
	}

	return false
}

func ResourceBinaryAuthorizationAttestor() *schema.Resource {
	return &schema.Resource{
		Create: resourceBinaryAuthorizationAttestorCreate,
		Read:   resourceBinaryAuthorizationAttestorRead,
		Update: resourceBinaryAuthorizationAttestorUpdate,
		Delete: resourceBinaryAuthorizationAttestorDelete,

		Importer: &schema.ResourceImporter{
			State: resourceBinaryAuthorizationAttestorImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"attestation_authority_note": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `A Container Analysis ATTESTATION_AUTHORITY Note, created by the user.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"note_reference": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
							Description: `The resource name of a ATTESTATION_AUTHORITY Note, created by the
user. If the Note is in a different project from the Attestor, it
should be specified in the format 'projects/*/notes/*' (or the legacy
'providers/*/notes/*'). This field may not be updated.
An attestation by this attestor is stored as a Container Analysis
ATTESTATION_AUTHORITY Occurrence that names a container image
and that links to this Note.`,
						},
						"public_keys": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `Public keys that verify attestations signed by this attestor. This
field may be updated.
If this field is non-empty, one of the specified public keys must
verify that an attestation was signed by this attestor for the
image specified in the admission request.
If this field is empty, this attestor always returns that no valid
attestations exist.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ascii_armored_pgp_public_key": {
										Type:     schema.TypeString,
										Optional: true,
										Description: `ASCII-armored representation of a PGP public key, as the
entire output by the command
'gpg --export --armor foo@example.com' (either LF or CRLF
line endings). When using this field, id should be left
blank. The BinAuthz API handlers will calculate the ID
and fill it in automatically. BinAuthz computes this ID
as the OpenPGP RFC4880 V4 fingerprint, represented as
upper-case hex. If id is provided by the caller, it will
be overwritten by the API-calculated ID.`,
									},
									"comment": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `A descriptive comment. This field may be updated.`,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
										Optional: true,
										Description: `The ID of this public key. Signatures verified by BinAuthz
must include the ID of the public key that can be used to
verify them, and that ID must match the contents of this
field exactly. Additional restrictions on this field can
be imposed based on which public key type is encapsulated.
See the documentation on publicKey cases below for details.`,
									},
									"pkix_public_key": {
										Type:     schema.TypeList,
										Optional: true,
										Description: `A raw PKIX SubjectPublicKeyInfo format public key.

NOTE: id may be explicitly provided by the caller when using this
type of public key, but it MUST be a valid RFC3986 URI. If id is left
blank, a default one will be computed based on the digest of the DER
encoding of the public key.`,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"public_key_pem": {
													Type:     schema.TypeString,
													Optional: true,
													Description: `A PEM-encoded public key, as described in
'https://tools.ietf.org/html/rfc7468#section-13'`,
												},
												"signature_algorithm": {
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: CompareSignatureAlgorithm,
													Description: `The signature algorithm used to verify a message against
a signature using this key. These signature algorithm must
match the structure and any object identifiers encoded in
publicKeyPem (i.e. this algorithm must match that of the
public key).`,
												},
											},
										},
									},
								},
							},
						},
						"delegation_service_account_email": {
							Type:     schema.TypeString,
							Computed: true,
							Description: `This field will contain the service account email address that
this Attestor will use as the principal when querying Container
Analysis. Attestor administrators must grant this service account
the IAM role needed to read attestations from the noteReference in
Container Analysis (containeranalysis.notes.occurrences.viewer).
This email address is fixed for the lifetime of the Attestor, but
callers should not make any other assumptions about the service
account email; future versions may use an email based on a
different naming pattern.`,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The resource name.`,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `A descriptive comment. This field may be updated. The field may be
displayed in chooser dialogs.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceBinaryAuthorizationAttestorCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandBinaryAuthorizationAttestorName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandBinaryAuthorizationAttestorDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	userOwnedGrafeasNoteProp, err := expandBinaryAuthorizationAttestorAttestationAuthorityNote(d.Get("attestation_authority_note"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("attestation_authority_note"); !tpgresource.IsEmptyValue(reflect.ValueOf(userOwnedGrafeasNoteProp)) && (ok || !reflect.DeepEqual(v, userOwnedGrafeasNoteProp)) {
		obj["userOwnedGrafeasNote"] = userOwnedGrafeasNoteProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BinaryAuthorizationBasePath}}projects/{{project}}/attestors?attestorId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Attestor: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Attestor: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating Attestor: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/attestors/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Attestor %q: %#v", d.Id(), res)

	return resourceBinaryAuthorizationAttestorRead(d, meta)
}

func resourceBinaryAuthorizationAttestorRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BinaryAuthorizationBasePath}}projects/{{project}}/attestors/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Attestor: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("BinaryAuthorizationAttestor %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Attestor: %s", err)
	}

	if err := d.Set("name", flattenBinaryAuthorizationAttestorName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Attestor: %s", err)
	}
	if err := d.Set("description", flattenBinaryAuthorizationAttestorDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Attestor: %s", err)
	}
	if err := d.Set("attestation_authority_note", flattenBinaryAuthorizationAttestorAttestationAuthorityNote(res["userOwnedGrafeasNote"], d, config)); err != nil {
		return fmt.Errorf("Error reading Attestor: %s", err)
	}

	return nil
}

func resourceBinaryAuthorizationAttestorUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Attestor: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	nameProp, err := expandBinaryAuthorizationAttestorName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandBinaryAuthorizationAttestorDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	userOwnedGrafeasNoteProp, err := expandBinaryAuthorizationAttestorAttestationAuthorityNote(d.Get("attestation_authority_note"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("attestation_authority_note"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, userOwnedGrafeasNoteProp)) {
		obj["userOwnedGrafeasNote"] = userOwnedGrafeasNoteProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BinaryAuthorizationBasePath}}projects/{{project}}/attestors/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Attestor %q: %#v", d.Id(), obj)
	headers := make(http.Header)

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PUT",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
		Headers:   headers,
	})

	if err != nil {
		return fmt.Errorf("Error updating Attestor %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Attestor %q: %#v", d.Id(), res)
	}

	return resourceBinaryAuthorizationAttestorRead(d, meta)
}

func resourceBinaryAuthorizationAttestorDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Attestor: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{BinaryAuthorizationBasePath}}projects/{{project}}/attestors/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting Attestor %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Attestor")
	}

	log.Printf("[DEBUG] Finished deleting Attestor %q: %#v", d.Id(), res)
	return nil
}

func resourceBinaryAuthorizationAttestorImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/attestors/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/attestors/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenBinaryAuthorizationAttestorName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.GetResourceNameFromSelfLink(v.(string))
}

func flattenBinaryAuthorizationAttestorDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBinaryAuthorizationAttestorAttestationAuthorityNote(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["note_reference"] =
		flattenBinaryAuthorizationAttestorAttestationAuthorityNoteNoteReference(original["noteReference"], d, config)
	transformed["public_keys"] =
		flattenBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeys(original["publicKeys"], d, config)
	transformed["delegation_service_account_email"] =
		flattenBinaryAuthorizationAttestorAttestationAuthorityNoteDelegationServiceAccountEmail(original["delegationServiceAccountEmail"], d, config)
	return []interface{}{transformed}
}
func flattenBinaryAuthorizationAttestorAttestationAuthorityNoteNoteReference(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeys(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"comment":                      flattenBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysComment(original["comment"], d, config),
			"id":                           flattenBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysId(original["id"], d, config),
			"ascii_armored_pgp_public_key": flattenBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysAsciiArmoredPgpPublicKey(original["asciiArmoredPgpPublicKey"], d, config),
			"pkix_public_key":              flattenBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKey(original["pkixPublicKey"], d, config),
		})
	}
	return transformed
}
func flattenBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysComment(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysAsciiArmoredPgpPublicKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["public_key_pem"] =
		flattenBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKeyPublicKeyPem(original["publicKeyPem"], d, config)
	transformed["signature_algorithm"] =
		flattenBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKeySignatureAlgorithm(original["signatureAlgorithm"], d, config)
	return []interface{}{transformed}
}
func flattenBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKeyPublicKeyPem(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKeySignatureAlgorithm(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBinaryAuthorizationAttestorAttestationAuthorityNoteDelegationServiceAccountEmail(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandBinaryAuthorizationAttestorName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return tpgresource.ReplaceVars(d, config, "projects/{{project}}/attestors/{{name}}")
}

func expandBinaryAuthorizationAttestorDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNote(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNoteReference, err := expandBinaryAuthorizationAttestorAttestationAuthorityNoteNoteReference(original["note_reference"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNoteReference); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["noteReference"] = transformedNoteReference
	}

	transformedPublicKeys, err := expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeys(original["public_keys"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPublicKeys); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["publicKeys"] = transformedPublicKeys
	}

	transformedDelegationServiceAccountEmail, err := expandBinaryAuthorizationAttestorAttestationAuthorityNoteDelegationServiceAccountEmail(original["delegation_service_account_email"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDelegationServiceAccountEmail); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["delegationServiceAccountEmail"] = transformedDelegationServiceAccountEmail
	}

	return transformed, nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNoteNoteReference(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	r := regexp.MustCompile("projects/(.+)/notes/(.+)")
	if r.MatchString(v.(string)) {
		return v.(string), nil
	}

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return nil, err
	}

	return fmt.Sprintf("projects/%s/notes/%s", project, v.(string)), nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeys(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedComment, err := expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysComment(original["comment"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedComment); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["comment"] = transformedComment
		}

		transformedId, err := expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysId(original["id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["id"] = transformedId
		}

		transformedAsciiArmoredPgpPublicKey, err := expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysAsciiArmoredPgpPublicKey(original["ascii_armored_pgp_public_key"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAsciiArmoredPgpPublicKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["asciiArmoredPgpPublicKey"] = transformedAsciiArmoredPgpPublicKey
		}

		transformedPkixPublicKey, err := expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKey(original["pkix_public_key"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPkixPublicKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["pkixPublicKey"] = transformedPkixPublicKey
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysComment(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysAsciiArmoredPgpPublicKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPublicKeyPem, err := expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKeyPublicKeyPem(original["public_key_pem"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPublicKeyPem); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["publicKeyPem"] = transformedPublicKeyPem
	}

	transformedSignatureAlgorithm, err := expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKeySignatureAlgorithm(original["signature_algorithm"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSignatureAlgorithm); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["signatureAlgorithm"] = transformedSignatureAlgorithm
	}

	return transformed, nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKeyPublicKeyPem(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKeySignatureAlgorithm(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNoteDelegationServiceAccountEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
