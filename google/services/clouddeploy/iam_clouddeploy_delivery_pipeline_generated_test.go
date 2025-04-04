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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/clouddeploy/DeliveryPipeline.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/examples/base_configs/iam_test_file.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package clouddeploy_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
)

func TestAccClouddeployDeliveryPipelineIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccClouddeployDeliveryPipelineIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_clouddeploy_delivery_pipeline_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s roles/viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-cd-delivery-pipeline%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccClouddeployDeliveryPipelineIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_clouddeploy_delivery_pipeline_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s roles/viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-cd-delivery-pipeline%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccClouddeployDeliveryPipelineIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccClouddeployDeliveryPipelineIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_clouddeploy_delivery_pipeline_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s roles/viewer user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-cd-delivery-pipeline%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccClouddeployDeliveryPipelineIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccClouddeployDeliveryPipelineIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_clouddeploy_delivery_pipeline_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_clouddeploy_delivery_pipeline_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-cd-delivery-pipeline%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccClouddeployDeliveryPipelineIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_clouddeploy_delivery_pipeline_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-cd-delivery-pipeline%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccClouddeployDeliveryPipelineIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_clouddeploy_delivery_pipeline" "default" {
  name = "tf-test-cd-delivery-pipeline%{random_suffix}"
  location = "us-central1"
  serial_pipeline  {
    stages {
      target_id = "test"
      profiles = []
    }
  }
 }

resource "google_clouddeploy_delivery_pipeline_iam_member" "foo" {
  project = google_clouddeploy_delivery_pipeline.default.project
  location = google_clouddeploy_delivery_pipeline.default.location
  name = google_clouddeploy_delivery_pipeline.default.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccClouddeployDeliveryPipelineIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_clouddeploy_delivery_pipeline" "default" {
  name = "tf-test-cd-delivery-pipeline%{random_suffix}"
  location = "us-central1"
  serial_pipeline  {
    stages {
      target_id = "test"
      profiles = []
    }
  }
 }

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_clouddeploy_delivery_pipeline_iam_policy" "foo" {
  project = google_clouddeploy_delivery_pipeline.default.project
  location = google_clouddeploy_delivery_pipeline.default.location
  name = google_clouddeploy_delivery_pipeline.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_clouddeploy_delivery_pipeline_iam_policy" "foo" {
  project = google_clouddeploy_delivery_pipeline.default.project
  location = google_clouddeploy_delivery_pipeline.default.location
  name = google_clouddeploy_delivery_pipeline.default.name
  depends_on = [
    google_clouddeploy_delivery_pipeline_iam_policy.foo
  ]
}
`, context)
}

func testAccClouddeployDeliveryPipelineIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_clouddeploy_delivery_pipeline" "default" {
  name = "tf-test-cd-delivery-pipeline%{random_suffix}"
  location = "us-central1"
  serial_pipeline  {
    stages {
      target_id = "test"
      profiles = []
    }
  }
 }

data "google_iam_policy" "foo" {
}

resource "google_clouddeploy_delivery_pipeline_iam_policy" "foo" {
  project = google_clouddeploy_delivery_pipeline.default.project
  location = google_clouddeploy_delivery_pipeline.default.location
  name = google_clouddeploy_delivery_pipeline.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccClouddeployDeliveryPipelineIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_clouddeploy_delivery_pipeline" "default" {
  name = "tf-test-cd-delivery-pipeline%{random_suffix}"
  location = "us-central1"
  serial_pipeline  {
    stages {
      target_id = "test"
      profiles = []
    }
  }
 }

resource "google_clouddeploy_delivery_pipeline_iam_binding" "foo" {
  project = google_clouddeploy_delivery_pipeline.default.project
  location = google_clouddeploy_delivery_pipeline.default.location
  name = google_clouddeploy_delivery_pipeline.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccClouddeployDeliveryPipelineIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_clouddeploy_delivery_pipeline" "default" {
  name = "tf-test-cd-delivery-pipeline%{random_suffix}"
  location = "us-central1"
  serial_pipeline  {
    stages {
      target_id = "test"
      profiles = []
    }
  }
 }

resource "google_clouddeploy_delivery_pipeline_iam_binding" "foo" {
  project = google_clouddeploy_delivery_pipeline.default.project
  location = google_clouddeploy_delivery_pipeline.default.location
  name = google_clouddeploy_delivery_pipeline.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
