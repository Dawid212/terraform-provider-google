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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/networksecurity/InterceptEndpointGroup.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package networksecurity

import (
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
)

func ResourceNetworkSecurityInterceptEndpointGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkSecurityInterceptEndpointGroupCreate,
		Read:   resourceNetworkSecurityInterceptEndpointGroupRead,
		Update: resourceNetworkSecurityInterceptEndpointGroupUpdate,
		Delete: resourceNetworkSecurityInterceptEndpointGroupDelete,

		Importer: &schema.ResourceImporter{
			State: resourceNetworkSecurityInterceptEndpointGroupImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"intercept_deployment_group": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The deployment group that this endpoint group is connected to, for example:
'projects/123456789/locations/global/interceptDeploymentGroups/my-dg'.
See https://google.aip.dev/124.`,
			},
			"intercept_endpoint_group_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The ID to use for the endpoint group, which will become the final component
of the endpoint group's resource name.`,
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The cloud location of the endpoint group, currently restricted to 'global'.`,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `User-provided description of the endpoint group.
Used as additional context for the endpoint group.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Labels are key/value pairs that help to organize and filter resources.

**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"associations": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: `List of associations to this endpoint group.`,
				Elem:        networksecurityInterceptEndpointGroupAssociationsSchema(),
				// Default schema.HashSchema is used.
			},
			"connected_deployment_group": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `The endpoint group's view of a connected deployment group.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"locations": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: `The list of locations where the deployment group is present.`,
							Elem:        networksecurityInterceptEndpointGroupConnectedDeploymentGroupLocationsSchema(),
							// Default schema.HashSchema is used.
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
							Description: `The connected deployment group's resource name, for example:
'projects/123456789/locations/global/interceptDeploymentGroups/my-dg'.
See https://google.aip.dev/124.`,
						},
					},
				},
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The timestamp when the resource was created.
See https://google.aip.dev/148#timestamps.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The resource name of this endpoint group, for example:
'projects/123456789/locations/global/interceptEndpointGroups/my-eg'.
See https://google.aip.dev/122 for more details.`,
			},
			"reconciling": {
				Type:     schema.TypeBool,
				Computed: true,
				Description: `The current state of the resource does not match the user's intended state,
and the system is working to reconcile them. This is part of the normal
operation (e.g. adding a new association to the group).
See https://google.aip.dev/128.`,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The current state of the endpoint group.
See https://google.aip.dev/216.
Possible values:
STATE_UNSPECIFIED
ACTIVE
CLOSED
CREATING
DELETING
OUT_OF_SYNC
DELETE_FAILED`,
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"update_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The timestamp when the resource was most recently updated.
See https://google.aip.dev/148#timestamps.`,
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

func networksecurityInterceptEndpointGroupAssociationsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The connected association's resource name, for example:
'projects/123456789/locations/global/interceptEndpointGroupAssociations/my-ega'.
See https://google.aip.dev/124.`,
			},
			"network": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The associated network, for example:
projects/123456789/global/networks/my-network.
See https://google.aip.dev/124.`,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Most recent known state of the association.
Possible values:
STATE_UNSPECIFIED
ACTIVE
CREATING
DELETING
CLOSED
OUT_OF_SYNC
DELETE_FAILED`,
			},
		},
	}
}

func networksecurityInterceptEndpointGroupConnectedDeploymentGroupLocationsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The cloud location, e.g. 'us-central1-a' or 'asia-south1-b'.`,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The current state of the association in this location.
Possible values:
STATE_UNSPECIFIED
ACTIVE
OUT_OF_SYNC`,
			},
		},
	}
}

func resourceNetworkSecurityInterceptEndpointGroupCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	interceptDeploymentGroupProp, err := expandNetworkSecurityInterceptEndpointGroupInterceptDeploymentGroup(d.Get("intercept_deployment_group"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("intercept_deployment_group"); !tpgresource.IsEmptyValue(reflect.ValueOf(interceptDeploymentGroupProp)) && (ok || !reflect.DeepEqual(v, interceptDeploymentGroupProp)) {
		obj["interceptDeploymentGroup"] = interceptDeploymentGroupProp
	}
	descriptionProp, err := expandNetworkSecurityInterceptEndpointGroupDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandNetworkSecurityInterceptEndpointGroupEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/interceptEndpointGroups?interceptEndpointGroupId={{intercept_endpoint_group_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new InterceptEndpointGroup: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for InterceptEndpointGroup: %s", err)
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
		return fmt.Errorf("Error creating InterceptEndpointGroup: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/interceptEndpointGroups/{{intercept_endpoint_group_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = NetworkSecurityOperationWaitTime(
		config, res, project, "Creating InterceptEndpointGroup", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create InterceptEndpointGroup: %s", err)
	}

	log.Printf("[DEBUG] Finished creating InterceptEndpointGroup %q: %#v", d.Id(), res)

	return resourceNetworkSecurityInterceptEndpointGroupRead(d, meta)
}

func resourceNetworkSecurityInterceptEndpointGroupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/interceptEndpointGroups/{{intercept_endpoint_group_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for InterceptEndpointGroup: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("NetworkSecurityInterceptEndpointGroup %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading InterceptEndpointGroup: %s", err)
	}

	if err := d.Set("name", flattenNetworkSecurityInterceptEndpointGroupName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterceptEndpointGroup: %s", err)
	}
	if err := d.Set("create_time", flattenNetworkSecurityInterceptEndpointGroupCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterceptEndpointGroup: %s", err)
	}
	if err := d.Set("update_time", flattenNetworkSecurityInterceptEndpointGroupUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterceptEndpointGroup: %s", err)
	}
	if err := d.Set("labels", flattenNetworkSecurityInterceptEndpointGroupLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterceptEndpointGroup: %s", err)
	}
	if err := d.Set("intercept_deployment_group", flattenNetworkSecurityInterceptEndpointGroupInterceptDeploymentGroup(res["interceptDeploymentGroup"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterceptEndpointGroup: %s", err)
	}
	if err := d.Set("state", flattenNetworkSecurityInterceptEndpointGroupState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterceptEndpointGroup: %s", err)
	}
	if err := d.Set("reconciling", flattenNetworkSecurityInterceptEndpointGroupReconciling(res["reconciling"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterceptEndpointGroup: %s", err)
	}
	if err := d.Set("description", flattenNetworkSecurityInterceptEndpointGroupDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterceptEndpointGroup: %s", err)
	}
	if err := d.Set("associations", flattenNetworkSecurityInterceptEndpointGroupAssociations(res["associations"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterceptEndpointGroup: %s", err)
	}
	if err := d.Set("connected_deployment_group", flattenNetworkSecurityInterceptEndpointGroupConnectedDeploymentGroup(res["connectedDeploymentGroup"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterceptEndpointGroup: %s", err)
	}
	if err := d.Set("terraform_labels", flattenNetworkSecurityInterceptEndpointGroupTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterceptEndpointGroup: %s", err)
	}
	if err := d.Set("effective_labels", flattenNetworkSecurityInterceptEndpointGroupEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterceptEndpointGroup: %s", err)
	}

	return nil
}

func resourceNetworkSecurityInterceptEndpointGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for InterceptEndpointGroup: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkSecurityInterceptEndpointGroupDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandNetworkSecurityInterceptEndpointGroupEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/interceptEndpointGroups/{{intercept_endpoint_group_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating InterceptEndpointGroup %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("effective_labels") {
		updateMask = append(updateMask, "labels")
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
			return fmt.Errorf("Error updating InterceptEndpointGroup %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating InterceptEndpointGroup %q: %#v", d.Id(), res)
		}

		err = NetworkSecurityOperationWaitTime(
			config, res, project, "Updating InterceptEndpointGroup", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceNetworkSecurityInterceptEndpointGroupRead(d, meta)
}

func resourceNetworkSecurityInterceptEndpointGroupDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for InterceptEndpointGroup: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/interceptEndpointGroups/{{intercept_endpoint_group_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting InterceptEndpointGroup %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "InterceptEndpointGroup")
	}

	err = NetworkSecurityOperationWaitTime(
		config, res, project, "Deleting InterceptEndpointGroup", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting InterceptEndpointGroup %q: %#v", d.Id(), res)
	return nil
}

func resourceNetworkSecurityInterceptEndpointGroupImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/interceptEndpointGroups/(?P<intercept_endpoint_group_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<intercept_endpoint_group_id>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<intercept_endpoint_group_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/interceptEndpointGroups/{{intercept_endpoint_group_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNetworkSecurityInterceptEndpointGroupName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityInterceptEndpointGroupCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityInterceptEndpointGroupUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityInterceptEndpointGroupLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenNetworkSecurityInterceptEndpointGroupInterceptDeploymentGroup(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityInterceptEndpointGroupState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityInterceptEndpointGroupReconciling(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityInterceptEndpointGroupDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityInterceptEndpointGroupAssociations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(schema.HashResource(networksecurityInterceptEndpointGroupAssociationsSchema()), []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"name":    flattenNetworkSecurityInterceptEndpointGroupAssociationsName(original["name"], d, config),
			"network": flattenNetworkSecurityInterceptEndpointGroupAssociationsNetwork(original["network"], d, config),
			"state":   flattenNetworkSecurityInterceptEndpointGroupAssociationsState(original["state"], d, config),
		})
	}
	return transformed
}
func flattenNetworkSecurityInterceptEndpointGroupAssociationsName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityInterceptEndpointGroupAssociationsNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityInterceptEndpointGroupAssociationsState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityInterceptEndpointGroupConnectedDeploymentGroup(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["name"] =
		flattenNetworkSecurityInterceptEndpointGroupConnectedDeploymentGroupName(original["name"], d, config)
	transformed["locations"] =
		flattenNetworkSecurityInterceptEndpointGroupConnectedDeploymentGroupLocations(original["locations"], d, config)
	return []interface{}{transformed}
}
func flattenNetworkSecurityInterceptEndpointGroupConnectedDeploymentGroupName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityInterceptEndpointGroupConnectedDeploymentGroupLocations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(schema.HashResource(networksecurityInterceptEndpointGroupConnectedDeploymentGroupLocationsSchema()), []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"location": flattenNetworkSecurityInterceptEndpointGroupConnectedDeploymentGroupLocationsLocation(original["location"], d, config),
			"state":    flattenNetworkSecurityInterceptEndpointGroupConnectedDeploymentGroupLocationsState(original["state"], d, config),
		})
	}
	return transformed
}
func flattenNetworkSecurityInterceptEndpointGroupConnectedDeploymentGroupLocationsLocation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityInterceptEndpointGroupConnectedDeploymentGroupLocationsState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityInterceptEndpointGroupTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("terraform_labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenNetworkSecurityInterceptEndpointGroupEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNetworkSecurityInterceptEndpointGroupInterceptDeploymentGroup(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityInterceptEndpointGroupDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityInterceptEndpointGroupEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
