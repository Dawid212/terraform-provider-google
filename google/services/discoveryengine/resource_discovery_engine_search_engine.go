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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/discoveryengine/SearchEngine.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package discoveryengine

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
	"github.com/hashicorp/terraform-provider-google/google/verify"
)

func ResourceDiscoveryEngineSearchEngine() *schema.Resource {
	return &schema.Resource{
		Create: resourceDiscoveryEngineSearchEngineCreate,
		Read:   resourceDiscoveryEngineSearchEngineRead,
		Update: resourceDiscoveryEngineSearchEngineUpdate,
		Delete: resourceDiscoveryEngineSearchEngineDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDiscoveryEngineSearchEngineImport,
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
			"collection_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The collection ID.`,
			},
			"data_store_ids": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `The data stores associated with this engine. For SOLUTION_TYPE_SEARCH type of engines, they can only associate with at most one data store.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Required. The display name of the engine. Should be human readable. UTF-8 encoded string with limit of 1024 characters.`,
			},
			"engine_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Unique ID to use for Search Engine App.`,
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Location.`,
			},
			"search_engine_config": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `Configurations for a Search Engine.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"search_add_ons": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `The add-on that this search engine enables. Possible values: ["SEARCH_ADD_ON_LLM"]`,
							Elem: &schema.Schema{
								Type:         schema.TypeString,
								ValidateFunc: verify.ValidateEnum([]string{"SEARCH_ADD_ON_LLM"}),
							},
						},
						"search_tier": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: verify.ValidateEnum([]string{"SEARCH_TIER_STANDARD", "SEARCH_TIER_ENTERPRISE", ""}),
							Description:  `The search feature tier of this engine. Defaults to SearchTier.SEARCH_TIER_STANDARD if not specified. Default value: "SEARCH_TIER_STANDARD" Possible values: ["SEARCH_TIER_STANDARD", "SEARCH_TIER_ENTERPRISE"]`,
							Default:      "SEARCH_TIER_STANDARD",
						},
					},
				},
			},
			"common_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `Common config spec that specifies the metadata of the engine.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"company_name": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: `The name of the company, business or entity that is associated with the engine. Setting this may help improve LLM related features.cd`,
						},
					},
				},
			},
			"industry_vertical": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"GENERIC", "MEDIA", "HEALTHCARE_FHIR", ""}),
				Description:  `The industry vertical that the engine registers. The restriction of the Engine industry vertical is based on DataStore: If unspecified, default to GENERIC. Vertical on Engine has to match vertical of the DataStore liniked to the engine. Default value: "GENERIC" Possible values: ["GENERIC", "MEDIA", "HEALTHCARE_FHIR"]`,
				Default:      "GENERIC",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Timestamp the Engine was created at.`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The unique full resource name of the search engine. Values are of the format
'projects/{project}/locations/{location}/collections/{collection_id}/engines/{engine_id}'.
This field must be a UTF-8 encoded string with a length limit of 1024
characters.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Timestamp the Engine was last updated.`,
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

func resourceDiscoveryEngineSearchEngineCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	industryVerticalProp, err := expandDiscoveryEngineSearchEngineIndustryVertical(d.Get("industry_vertical"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("industry_vertical"); !tpgresource.IsEmptyValue(reflect.ValueOf(industryVerticalProp)) && (ok || !reflect.DeepEqual(v, industryVerticalProp)) {
		obj["industryVertical"] = industryVerticalProp
	}
	displayNameProp, err := expandDiscoveryEngineSearchEngineDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	dataStoreIdsProp, err := expandDiscoveryEngineSearchEngineDataStoreIds(d.Get("data_store_ids"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("data_store_ids"); !tpgresource.IsEmptyValue(reflect.ValueOf(dataStoreIdsProp)) && (ok || !reflect.DeepEqual(v, dataStoreIdsProp)) {
		obj["dataStoreIds"] = dataStoreIdsProp
	}
	searchEngineConfigProp, err := expandDiscoveryEngineSearchEngineSearchEngineConfig(d.Get("search_engine_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("search_engine_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(searchEngineConfigProp)) && (ok || !reflect.DeepEqual(v, searchEngineConfigProp)) {
		obj["searchEngineConfig"] = searchEngineConfigProp
	}
	commonConfigProp, err := expandDiscoveryEngineSearchEngineCommonConfig(d.Get("common_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("common_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(commonConfigProp)) && (ok || !reflect.DeepEqual(v, commonConfigProp)) {
		obj["commonConfig"] = commonConfigProp
	}

	obj, err = resourceDiscoveryEngineSearchEngineEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DiscoveryEngineBasePath}}projects/{{project}}/locations/{{location}}/collections/{{collection_id}}/engines?engineId={{engine_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new SearchEngine: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for SearchEngine: %s", err)
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
		return fmt.Errorf("Error creating SearchEngine: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/collections/{{collection_id}}/engines/{{engine_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = DiscoveryEngineOperationWaitTime(
		config, res, project, "Creating SearchEngine", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create SearchEngine: %s", err)
	}

	log.Printf("[DEBUG] Finished creating SearchEngine %q: %#v", d.Id(), res)

	return resourceDiscoveryEngineSearchEngineRead(d, meta)
}

func resourceDiscoveryEngineSearchEngineRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DiscoveryEngineBasePath}}projects/{{project}}/locations/{{location}}/collections/{{collection_id}}/engines/{{engine_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for SearchEngine: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("DiscoveryEngineSearchEngine %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading SearchEngine: %s", err)
	}

	if err := d.Set("name", flattenDiscoveryEngineSearchEngineName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading SearchEngine: %s", err)
	}
	if err := d.Set("industry_vertical", flattenDiscoveryEngineSearchEngineIndustryVertical(res["industryVertical"], d, config)); err != nil {
		return fmt.Errorf("Error reading SearchEngine: %s", err)
	}
	if err := d.Set("display_name", flattenDiscoveryEngineSearchEngineDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading SearchEngine: %s", err)
	}
	if err := d.Set("data_store_ids", flattenDiscoveryEngineSearchEngineDataStoreIds(res["dataStoreIds"], d, config)); err != nil {
		return fmt.Errorf("Error reading SearchEngine: %s", err)
	}
	if err := d.Set("create_time", flattenDiscoveryEngineSearchEngineCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading SearchEngine: %s", err)
	}
	if err := d.Set("update_time", flattenDiscoveryEngineSearchEngineUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading SearchEngine: %s", err)
	}
	if err := d.Set("search_engine_config", flattenDiscoveryEngineSearchEngineSearchEngineConfig(res["searchEngineConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading SearchEngine: %s", err)
	}
	if err := d.Set("common_config", flattenDiscoveryEngineSearchEngineCommonConfig(res["commonConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading SearchEngine: %s", err)
	}

	return nil
}

func resourceDiscoveryEngineSearchEngineUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for SearchEngine: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandDiscoveryEngineSearchEngineDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	dataStoreIdsProp, err := expandDiscoveryEngineSearchEngineDataStoreIds(d.Get("data_store_ids"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("data_store_ids"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, dataStoreIdsProp)) {
		obj["dataStoreIds"] = dataStoreIdsProp
	}
	searchEngineConfigProp, err := expandDiscoveryEngineSearchEngineSearchEngineConfig(d.Get("search_engine_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("search_engine_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, searchEngineConfigProp)) {
		obj["searchEngineConfig"] = searchEngineConfigProp
	}

	obj, err = resourceDiscoveryEngineSearchEngineEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DiscoveryEngineBasePath}}projects/{{project}}/locations/{{location}}/collections/{{collection_id}}/engines/{{engine_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating SearchEngine %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("data_store_ids") {
		updateMask = append(updateMask, "dataStoreIds")
	}

	if d.HasChange("search_engine_config") {
		updateMask = append(updateMask, "searchEngineConfig")
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
			return fmt.Errorf("Error updating SearchEngine %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating SearchEngine %q: %#v", d.Id(), res)
		}

	}

	return resourceDiscoveryEngineSearchEngineRead(d, meta)
}

func resourceDiscoveryEngineSearchEngineDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for SearchEngine: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{DiscoveryEngineBasePath}}projects/{{project}}/locations/{{location}}/collections/{{collection_id}}/engines/{{engine_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting SearchEngine %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "SearchEngine")
	}

	err = DiscoveryEngineOperationWaitTime(
		config, res, project, "Deleting SearchEngine", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting SearchEngine %q: %#v", d.Id(), res)
	return nil
}

func resourceDiscoveryEngineSearchEngineImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/collections/(?P<collection_id>[^/]+)/engines/(?P<engine_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<collection_id>[^/]+)/(?P<engine_id>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<collection_id>[^/]+)/(?P<engine_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/collections/{{collection_id}}/engines/{{engine_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenDiscoveryEngineSearchEngineName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDiscoveryEngineSearchEngineIndustryVertical(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDiscoveryEngineSearchEngineDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDiscoveryEngineSearchEngineDataStoreIds(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDiscoveryEngineSearchEngineCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDiscoveryEngineSearchEngineUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDiscoveryEngineSearchEngineSearchEngineConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["search_tier"] =
		flattenDiscoveryEngineSearchEngineSearchEngineConfigSearchTier(original["searchTier"], d, config)
	transformed["search_add_ons"] =
		flattenDiscoveryEngineSearchEngineSearchEngineConfigSearchAddOns(original["searchAddOns"], d, config)
	return []interface{}{transformed}
}
func flattenDiscoveryEngineSearchEngineSearchEngineConfigSearchTier(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDiscoveryEngineSearchEngineSearchEngineConfigSearchAddOns(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDiscoveryEngineSearchEngineCommonConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["company_name"] =
		flattenDiscoveryEngineSearchEngineCommonConfigCompanyName(original["companyName"], d, config)
	return []interface{}{transformed}
}
func flattenDiscoveryEngineSearchEngineCommonConfigCompanyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandDiscoveryEngineSearchEngineIndustryVertical(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDiscoveryEngineSearchEngineDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDiscoveryEngineSearchEngineDataStoreIds(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDiscoveryEngineSearchEngineSearchEngineConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSearchTier, err := expandDiscoveryEngineSearchEngineSearchEngineConfigSearchTier(original["search_tier"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSearchTier); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["searchTier"] = transformedSearchTier
	}

	transformedSearchAddOns, err := expandDiscoveryEngineSearchEngineSearchEngineConfigSearchAddOns(original["search_add_ons"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSearchAddOns); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["searchAddOns"] = transformedSearchAddOns
	}

	return transformed, nil
}

func expandDiscoveryEngineSearchEngineSearchEngineConfigSearchTier(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDiscoveryEngineSearchEngineSearchEngineConfigSearchAddOns(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDiscoveryEngineSearchEngineCommonConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCompanyName, err := expandDiscoveryEngineSearchEngineCommonConfigCompanyName(original["company_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCompanyName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["companyName"] = transformedCompanyName
	}

	return transformed, nil
}

func expandDiscoveryEngineSearchEngineCommonConfigCompanyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceDiscoveryEngineSearchEngineEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// hard code solutionType to "SOLUTION_TYPE_SEARCH" for search engine resource
	obj["solutionType"] = "SOLUTION_TYPE_SEARCH"
	return obj, nil
}
