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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/beyondcorp/SecurityGateway.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package beyondcorp

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
	"github.com/hashicorp/terraform-provider-google/google/verify"
)

func beyondcorpSecurityGatewayHubsHash(v interface{}) int {
	if v == nil {
		return 0
	}

	var buf bytes.Buffer
	m := v.(map[string]interface{})

	buf.WriteString(fmt.Sprintf("%s-", m["region"].(string)))

	return tpgresource.Hashcode(buf.String())
}

func ResourceBeyondcorpSecurityGateway() *schema.Resource {
	return &schema.Resource{
		Create: resourceBeyondcorpSecurityGatewayCreate,
		Read:   resourceBeyondcorpSecurityGatewayRead,
		Update: resourceBeyondcorpSecurityGatewayUpdate,
		Delete: resourceBeyondcorpSecurityGatewayDelete,

		Importer: &schema.ResourceImporter{
			State: resourceBeyondcorpSecurityGatewayImport,
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
			"security_gateway_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Optional. User-settable SecurityGateway resource ID.
* Must start with a letter.
* Must contain between 4-63 characters from '/a-z-/'.
* Must end with a number or letter.`,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `Optional. An arbitrary user-provided name for the SecurityGateway.
Cannot exceed 64 characters.`,
			},
			"hubs": {
				Type:     schema.TypeSet,
				Optional: true,
				Description: `Optional. Map of Hubs that represents regional data path deployment with GCP region
as a key.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"region": {
							Type:     schema.TypeString,
							Required: true,
						},
						"internet_gateway": {
							Type:        schema.TypeList,
							Computed:    true,
							Optional:    true,
							Description: `Internet Gateway configuration.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"assigned_ips": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: `Output only. List of IP addresses assigned to the Cloud NAT.`,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
				Set: beyondcorpSecurityGatewayHubsHash,
			},
			"location": {
				Type:         schema.TypeString,
				Optional:     true,
				Deprecated:   "`location` is deprecated and will be removed in a future major release.",
				ForceNew:     true,
				ValidateFunc: verify.ValidateRegexp(`^global$`),
				Description:  `Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122. Must be omitted or set to 'global'.`,
				Default:      "global",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. Timestamp when the resource was created.`,
			},
			"delegating_service_account": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Service account used for operations that involve resources in consumer projects.`,
			},
			"external_ips": {
				Type:     schema.TypeList,
				Computed: true,
				Description: `Output only. IP addresses that will be used for establishing
connection to the endpoints.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Identifier. Name of the resource.`,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Output only. The operational state of the SecurityGateway.
Possible values:
STATE_UNSPECIFIED
CREATING
UPDATING
DELETING
RUNNING
DOWN
ERROR`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. Timestamp when the resource was last modified.`,
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

func resourceBeyondcorpSecurityGatewayCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	hubsProp, err := expandBeyondcorpSecurityGatewayHubs(d.Get("hubs"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("hubs"); !tpgresource.IsEmptyValue(reflect.ValueOf(hubsProp)) && (ok || !reflect.DeepEqual(v, hubsProp)) {
		obj["hubs"] = hubsProp
	}
	displayNameProp, err := expandBeyondcorpSecurityGatewayDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BeyondcorpBasePath}}projects/{{project}}/locations/{{location}}/securityGateways?securityGatewayId={{security_gateway_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new SecurityGateway: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for SecurityGateway: %s", err)
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
		return fmt.Errorf("Error creating SecurityGateway: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/securityGateways/{{security_gateway_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = BeyondcorpOperationWaitTime(
		config, res, project, "Creating SecurityGateway", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create SecurityGateway: %s", err)
	}

	log.Printf("[DEBUG] Finished creating SecurityGateway %q: %#v", d.Id(), res)

	return resourceBeyondcorpSecurityGatewayRead(d, meta)
}

func resourceBeyondcorpSecurityGatewayRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BeyondcorpBasePath}}projects/{{project}}/locations/{{location}}/securityGateways/{{security_gateway_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for SecurityGateway: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("BeyondcorpSecurityGateway %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading SecurityGateway: %s", err)
	}

	if err := d.Set("state", flattenBeyondcorpSecurityGatewayState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecurityGateway: %s", err)
	}
	if err := d.Set("update_time", flattenBeyondcorpSecurityGatewayUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecurityGateway: %s", err)
	}
	if err := d.Set("create_time", flattenBeyondcorpSecurityGatewayCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecurityGateway: %s", err)
	}
	if err := d.Set("hubs", flattenBeyondcorpSecurityGatewayHubs(res["hubs"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecurityGateway: %s", err)
	}
	if err := d.Set("display_name", flattenBeyondcorpSecurityGatewayDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecurityGateway: %s", err)
	}
	if err := d.Set("external_ips", flattenBeyondcorpSecurityGatewayExternalIps(res["externalIps"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecurityGateway: %s", err)
	}
	if err := d.Set("name", flattenBeyondcorpSecurityGatewayName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecurityGateway: %s", err)
	}
	if err := d.Set("delegating_service_account", flattenBeyondcorpSecurityGatewayDelegatingServiceAccount(res["delegatingServiceAccount"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecurityGateway: %s", err)
	}

	return nil
}

func resourceBeyondcorpSecurityGatewayUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for SecurityGateway: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	hubsProp, err := expandBeyondcorpSecurityGatewayHubs(d.Get("hubs"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("hubs"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, hubsProp)) {
		obj["hubs"] = hubsProp
	}
	displayNameProp, err := expandBeyondcorpSecurityGatewayDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BeyondcorpBasePath}}projects/{{project}}/locations/{{location}}/securityGateways/{{security_gateway_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating SecurityGateway %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("hubs") {
		updateMask = append(updateMask, "hubs")
	}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// if updateMask is empty we are not updating anything so skip the post
	if len(updateMask) > 0 {
		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "PATCH",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
			Body:      obj,
			Timeout:   d.Timeout(schema.TimeoutUpdate),
			Headers:   headers,
		})

		if err != nil {
			return fmt.Errorf("Error updating SecurityGateway %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating SecurityGateway %q: %#v", d.Id(), res)
		}

		err = BeyondcorpOperationWaitTime(
			config, res, project, "Updating SecurityGateway", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceBeyondcorpSecurityGatewayRead(d, meta)
}

func resourceBeyondcorpSecurityGatewayDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for SecurityGateway: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{BeyondcorpBasePath}}projects/{{project}}/locations/{{location}}/securityGateways/{{security_gateway_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting SecurityGateway %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "SecurityGateway")
	}

	err = BeyondcorpOperationWaitTime(
		config, res, project, "Deleting SecurityGateway", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting SecurityGateway %q: %#v", d.Id(), res)
	return nil
}

func resourceBeyondcorpSecurityGatewayImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/securityGateways/(?P<security_gateway_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<security_gateway_id>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<security_gateway_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/securityGateways/{{security_gateway_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenBeyondcorpSecurityGatewayState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBeyondcorpSecurityGatewayUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBeyondcorpSecurityGatewayCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBeyondcorpSecurityGatewayHubs(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.(map[string]interface{})
	transformed := make([]interface{}, 0, len(l))
	for k, raw := range l {
		original := raw.(map[string]interface{})
		transformed = append(transformed, map[string]interface{}{
			"region":           k,
			"internet_gateway": flattenBeyondcorpSecurityGatewayHubsInternetGateway(original["internetGateway"], d, config),
		})
	}
	return transformed
}
func flattenBeyondcorpSecurityGatewayHubsInternetGateway(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["assigned_ips"] =
		flattenBeyondcorpSecurityGatewayHubsInternetGatewayAssignedIps(original["assignedIps"], d, config)
	return []interface{}{transformed}
}
func flattenBeyondcorpSecurityGatewayHubsInternetGatewayAssignedIps(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBeyondcorpSecurityGatewayDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBeyondcorpSecurityGatewayExternalIps(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBeyondcorpSecurityGatewayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBeyondcorpSecurityGatewayDelegatingServiceAccount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandBeyondcorpSecurityGatewayHubs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	if v == nil {
		return map[string]interface{}{}, nil
	}
	m := make(map[string]interface{})
	for _, raw := range v.(*schema.Set).List() {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedInternetGateway, err := expandBeyondcorpSecurityGatewayHubsInternetGateway(original["internet_gateway"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedInternetGateway); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["internetGateway"] = transformedInternetGateway
		}

		transformedRegion, err := tpgresource.ExpandString(original["region"], d, config)
		if err != nil {
			return nil, err
		}
		m[transformedRegion] = transformed
	}
	return m, nil
}

func expandBeyondcorpSecurityGatewayHubsInternetGateway(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAssignedIps, err := expandBeyondcorpSecurityGatewayHubsInternetGatewayAssignedIps(original["assigned_ips"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAssignedIps); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["assignedIps"] = transformedAssignedIps
	}

	return transformed, nil
}

func expandBeyondcorpSecurityGatewayHubsInternetGatewayAssignedIps(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBeyondcorpSecurityGatewayDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
