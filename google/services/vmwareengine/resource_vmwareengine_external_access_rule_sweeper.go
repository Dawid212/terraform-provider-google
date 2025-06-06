// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
// ----------------------------------------------------------------------------
//
//	***     AUTO GENERATED CODE    ***    Type: Handwritten     ***
//
// ----------------------------------------------------------------------------
//
//	This code is generated by Magic Modules using the following:
//
//	Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/vmwareengine/resource_vmwareengine_external_access_rule_sweeper.go
//
//	DO NOT EDIT this file directly. Any changes made to this file will be
//	overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
package vmwareengine

import (
	"context"
	"log"
	"testing"

	"github.com/hashicorp/terraform-provider-google/google/envvar"
	"github.com/hashicorp/terraform-provider-google/google/sweeper"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func init() {
	sweeper.AddTestSweepersLegacy("VmwareengineExternalAccessRule", testSweepVmwareengineExternalAccessRule)
}

// At the time of writing, the CI only passes us-central1 as the region
func testSweepVmwareengineExternalAccessRule(region string) error {
	resourceName := "VmwareengineExternalAccessRule"
	log.Printf("[INFO][SWEEPER_LOG] Starting sweeper for %s", resourceName)

	config, err := sweeper.SharedConfigForRegion(region)
	if err != nil {
		log.Printf("[INFO][SWEEPER_LOG] error getting shared config for region: %s", err)
		return err
	}

	err = config.LoadAndValidate(context.Background())
	if err != nil {
		log.Printf("[INFO][SWEEPER_LOG] error loading: %s", err)
		return err
	}

	t := &testing.T{}
	billingId := envvar.GetTestBillingAccountFromEnv(t)

	// List of location values includes:
	//   * zones used for this resource type's acc tests in the past
	//   * the 'region' passed to the sweeper
	locations := []string{region, "us-central1", "southamerica-west1", "me-west1"}
	log.Printf("[INFO][SWEEPER_LOG] Sweeping will include these locations: %v.", locations)
	for _, location := range locations {

		// Setup variables to replace in list template
		d := &tpgresource.ResourceDataMock{
			FieldsInSchema: map[string]interface{}{
				"project":         config.Project,
				"region":          location,
				"location":        location,
				"zone":            location,
				"billing_account": billingId,
			},
		}

		log.Printf("[INFO][SWEEPER_LOG] looking for parent resources in location '%s'.", location)

		parentResponseField := "networkPolicies"
		parentListUrlTemplate := "https://vmwareengine.googleapis.com/v1/projects/{{project}}/locations/{{location}}/networkPolicies"
		parentNames, err := sweeper.ListParentResourcesInLocation(d, config, parentListUrlTemplate, parentResponseField)
		if err != nil {
			log.Printf("[INFO][SWEEPER_LOG] error finding parental resources in location %s: %s", location, err)
			continue
		}

		for _, parent := range parentNames {

			// `parent` will be string of form projects/my-project/locations/us-central1-a/privateClouds/my-cloud
			// Change on each loop, so new value used in tpgresource.ReplaceVars
			d.Set("parent", parent)

			listTemplate := "https://vmwareengine.googleapis.com/v1/{{parent}}/externalAccessRules"
			listUrl, err := tpgresource.ReplaceVars(d, config, listTemplate)
			if err != nil {
				log.Printf("[INFO][SWEEPER_LOG] error preparing sweeper list url: %s", err)
				return nil
			}

			res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   config.Project,
				RawURL:    listUrl,
				UserAgent: config.UserAgent,
			})
			if err != nil {
				log.Printf("[INFO][SWEEPER_LOG] Error in response from request %s: %s", listUrl, err)
				return nil
			}

			resourceList, ok := res["externalAccessRules"]
			if !ok {
				log.Printf("[INFO][SWEEPER_LOG] Nothing found in response.")
				return nil
			}

			rl := resourceList.([]interface{})

			log.Printf("[INFO][SWEEPER_LOG] Found %d items in %s list response.", len(rl), resourceName)
			// Keep count of items that aren't sweepable for logging.
			nonPrefixCount := 0
			for _, ri := range rl {
				obj := ri.(map[string]interface{})
				if obj["name"] == nil {
					log.Printf("[INFO][SWEEPER_LOG] %s resource name was nil", resourceName)
					return nil
				}

				name := tpgresource.GetResourceNameFromSelfLink(obj["name"].(string))
				// Skip resources that shouldn't be sweeped
				if !sweeper.IsSweepableTestResource(name) {
					nonPrefixCount++
					continue
				}

				deleteTemplate := "https://vmwareengine.googleapis.com/v1/{{parent}}/externalAccessRules/{{name}}"
				deleteUrl, err := tpgresource.ReplaceVars(d, config, deleteTemplate)
				if err != nil {
					log.Printf("[INFO][SWEEPER_LOG] error preparing delete url: %s", err)
					return nil
				}
				deleteUrl = deleteUrl + name

				// Don't wait on operations as we may have a lot to delete
				_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
					Config:    config,
					Method:    "DELETE",
					Project:   config.Project,
					RawURL:    deleteUrl,
					UserAgent: config.UserAgent,
				})
				if err != nil {
					log.Printf("[INFO][SWEEPER_LOG] Error deleting for url %s : %s", deleteUrl, err)
				} else {
					log.Printf("[INFO][SWEEPER_LOG] Sent delete request for %s resource: %s", resourceName, name)
				}
			}

			if nonPrefixCount > 0 {
				log.Printf("[INFO][SWEEPER_LOG] %d items were non-sweepable and skipped.", nonPrefixCount)
			}
		}
	}
	return nil
}
