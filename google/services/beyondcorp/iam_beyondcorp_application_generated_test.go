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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/beyondcorp/Application.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/examples/base_configs/iam_test_file.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package beyondcorp_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
)

func TestAccBeyondcorpApplicationIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/beyondcorp.securityGatewayUser",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBeyondcorpApplicationIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_beyondcorp_application_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/global/securityGateways/%s/applications/%s roles/beyondcorp.securityGatewayUser", envvar.GetTestProjectFromEnv(), fmt.Sprintf("default%s", context["random_suffix"]), fmt.Sprintf("google%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccBeyondcorpApplicationIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_beyondcorp_application_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/global/securityGateways/%s/applications/%s roles/beyondcorp.securityGatewayUser", envvar.GetTestProjectFromEnv(), fmt.Sprintf("default%s", context["random_suffix"]), fmt.Sprintf("google%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBeyondcorpApplicationIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/beyondcorp.securityGatewayUser",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccBeyondcorpApplicationIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_beyondcorp_application_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/global/securityGateways/%s/applications/%s roles/beyondcorp.securityGatewayUser user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), fmt.Sprintf("default%s", context["random_suffix"]), fmt.Sprintf("google%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBeyondcorpApplicationIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/beyondcorp.securityGatewayUser",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBeyondcorpApplicationIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_beyondcorp_application_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_beyondcorp_application_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/global/securityGateways/%s/applications/%s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("default%s", context["random_suffix"]), fmt.Sprintf("google%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccBeyondcorpApplicationIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_beyondcorp_application_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/global/securityGateways/%s/applications/%s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("default%s", context["random_suffix"]), fmt.Sprintf("google%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBeyondcorpApplicationIamBindingGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/beyondcorp.securityGatewayUser",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBeyondcorpApplicationIamBinding_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_beyondcorp_application_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/global/securityGateways/%s/applications/%s roles/beyondcorp.securityGatewayUser %s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("default%s", context["random_suffix"]), fmt.Sprintf("google%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBeyondcorpApplicationIamBindingGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/beyondcorp.securityGatewayUser",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBeyondcorpApplicationIamBinding_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_beyondcorp_application_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/global/securityGateways/%s/applications/%s roles/beyondcorp.securityGatewayUser", envvar.GetTestProjectFromEnv(), fmt.Sprintf("default%s", context["random_suffix"]), fmt.Sprintf("google%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_beyondcorp_application_iam_binding.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/global/securityGateways/%s/applications/%s roles/beyondcorp.securityGatewayUser %s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("default%s", context["random_suffix"]), fmt.Sprintf("google%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_beyondcorp_application_iam_binding.foo3",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/global/securityGateways/%s/applications/%s roles/beyondcorp.securityGatewayUser %s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("default%s", context["random_suffix"]), fmt.Sprintf("google%s", context["random_suffix"]), context["condition_title_no_desc"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBeyondcorpApplicationIamMemberGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/beyondcorp.securityGatewayUser",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBeyondcorpApplicationIamMember_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_beyondcorp_application_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/global/securityGateways/%s/applications/%s roles/beyondcorp.securityGatewayUser user:admin@hashicorptest.com %s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("default%s", context["random_suffix"]), fmt.Sprintf("google%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBeyondcorpApplicationIamMemberGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/beyondcorp.securityGatewayUser",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBeyondcorpApplicationIamMember_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_beyondcorp_application_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/global/securityGateways/%s/applications/%s roles/beyondcorp.securityGatewayUser user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), fmt.Sprintf("default%s", context["random_suffix"]), fmt.Sprintf("google%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_beyondcorp_application_iam_member.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/global/securityGateways/%s/applications/%s roles/beyondcorp.securityGatewayUser user:admin@hashicorptest.com %s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("default%s", context["random_suffix"]), fmt.Sprintf("google%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_beyondcorp_application_iam_member.foo3",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/global/securityGateways/%s/applications/%s roles/beyondcorp.securityGatewayUser user:admin@hashicorptest.com %s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("default%s", context["random_suffix"]), fmt.Sprintf("google%s", context["random_suffix"]), context["condition_title_no_desc"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBeyondcorpApplicationIamPolicyGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/beyondcorp.securityGatewayUser",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	// Test should have 2 bindings: one with a description and one without. Any < chars are converted to a unicode character by the API.
	expectedPolicyData := acctest.Nprintf(`{"bindings":[{"condition":{"description":"%{condition_desc}","expression":"%{condition_expr}","title":"%{condition_title}"},"members":["user:admin@hashicorptest.com"],"role":"%{role}"},{"condition":{"expression":"%{condition_expr}","title":"%{condition_title}-no-description"},"members":["user:admin@hashicorptest.com"],"role":"%{role}"}]}`, context)
	expectedPolicyData = strings.Replace(expectedPolicyData, "<", "\\u003c", -1)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBeyondcorpApplicationIamPolicy_withConditionGenerated(context),
				Check: resource.ComposeAggregateTestCheckFunc(
					// TODO - uncomment once https://github.com/GoogleCloudPlatform/magic-modules/pull/6466 merged
					// resource.TestCheckResourceAttr("data.google_iam_policy.foo", "policy_data", expectedPolicyData),
					resource.TestCheckResourceAttr("google_beyondcorp_application_iam_policy.foo", "policy_data", expectedPolicyData),
					resource.TestCheckResourceAttrWith("data.google_iam_policy.foo", "policy_data", tpgresource.CheckGoogleIamPolicy),
				),
			},
			{
				ResourceName:      "google_beyondcorp_application_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/global/securityGateways/%s/applications/%s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("default%s", context["random_suffix"]), fmt.Sprintf("google%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccBeyondcorpApplicationIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_beyondcorp_security_gateway" "default" {
  security_gateway_id = "default%{random_suffix}"
  display_name = "My Security Gateway resource"
  hubs { region = "us-central1" }
}

resource "google_beyondcorp_application" "example" {
  security_gateways_id = google_beyondcorp_security_gateway.default.security_gateway_id
  application_id = "google%{random_suffix}"
  endpoint_matchers {
    hostname = "google.com"
  }
}

resource "google_beyondcorp_application_iam_member" "foo" {
  project = google_beyondcorp_application.example.project
  security_gateways_id = google_beyondcorp_application.example.security_gateways_id
  application_id = google_beyondcorp_application.example.application_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccBeyondcorpApplicationIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_beyondcorp_security_gateway" "default" {
  security_gateway_id = "default%{random_suffix}"
  display_name = "My Security Gateway resource"
  hubs { region = "us-central1" }
}

resource "google_beyondcorp_application" "example" {
  security_gateways_id = google_beyondcorp_security_gateway.default.security_gateway_id
  application_id = "google%{random_suffix}"
  endpoint_matchers {
    hostname = "google.com"
  }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_beyondcorp_application_iam_policy" "foo" {
  project = google_beyondcorp_application.example.project
  security_gateways_id = google_beyondcorp_application.example.security_gateways_id
  application_id = google_beyondcorp_application.example.application_id
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_beyondcorp_application_iam_policy" "foo" {
  project = google_beyondcorp_application.example.project
  security_gateways_id = google_beyondcorp_application.example.security_gateways_id
  application_id = google_beyondcorp_application.example.application_id
  depends_on = [
    google_beyondcorp_application_iam_policy.foo
  ]
}
`, context)
}

func testAccBeyondcorpApplicationIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_beyondcorp_security_gateway" "default" {
  security_gateway_id = "default%{random_suffix}"
  display_name = "My Security Gateway resource"
  hubs { region = "us-central1" }
}

resource "google_beyondcorp_application" "example" {
  security_gateways_id = google_beyondcorp_security_gateway.default.security_gateway_id
  application_id = "google%{random_suffix}"
  endpoint_matchers {
    hostname = "google.com"
  }
}

data "google_iam_policy" "foo" {
}

resource "google_beyondcorp_application_iam_policy" "foo" {
  project = google_beyondcorp_application.example.project
  security_gateways_id = google_beyondcorp_application.example.security_gateways_id
  application_id = google_beyondcorp_application.example.application_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccBeyondcorpApplicationIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_beyondcorp_security_gateway" "default" {
  security_gateway_id = "default%{random_suffix}"
  display_name = "My Security Gateway resource"
  hubs { region = "us-central1" }
}

resource "google_beyondcorp_application" "example" {
  security_gateways_id = google_beyondcorp_security_gateway.default.security_gateway_id
  application_id = "google%{random_suffix}"
  endpoint_matchers {
    hostname = "google.com"
  }
}

resource "google_beyondcorp_application_iam_binding" "foo" {
  project = google_beyondcorp_application.example.project
  security_gateways_id = google_beyondcorp_application.example.security_gateways_id
  application_id = google_beyondcorp_application.example.application_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccBeyondcorpApplicationIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_beyondcorp_security_gateway" "default" {
  security_gateway_id = "default%{random_suffix}"
  display_name = "My Security Gateway resource"
  hubs { region = "us-central1" }
}

resource "google_beyondcorp_application" "example" {
  security_gateways_id = google_beyondcorp_security_gateway.default.security_gateway_id
  application_id = "google%{random_suffix}"
  endpoint_matchers {
    hostname = "google.com"
  }
}

resource "google_beyondcorp_application_iam_binding" "foo" {
  project = google_beyondcorp_application.example.project
  security_gateways_id = google_beyondcorp_application.example.security_gateways_id
  application_id = google_beyondcorp_application.example.application_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}

func testAccBeyondcorpApplicationIamBinding_withConditionGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_beyondcorp_security_gateway" "default" {
  security_gateway_id = "default%{random_suffix}"
  display_name = "My Security Gateway resource"
  hubs { region = "us-central1" }
}

resource "google_beyondcorp_application" "example" {
  security_gateways_id = google_beyondcorp_security_gateway.default.security_gateway_id
  application_id = "google%{random_suffix}"
  endpoint_matchers {
    hostname = "google.com"
  }
}

resource "google_beyondcorp_application_iam_binding" "foo" {
  project = google_beyondcorp_application.example.project
  security_gateways_id = google_beyondcorp_application.example.security_gateways_id
  application_id = google_beyondcorp_application.example.application_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccBeyondcorpApplicationIamBinding_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_beyondcorp_security_gateway" "default" {
  security_gateway_id = "default%{random_suffix}"
  display_name = "My Security Gateway resource"
  hubs { region = "us-central1" }
}

resource "google_beyondcorp_application" "example" {
  security_gateways_id = google_beyondcorp_security_gateway.default.security_gateway_id
  application_id = "google%{random_suffix}"
  endpoint_matchers {
    hostname = "google.com"
  }
}

resource "google_beyondcorp_application_iam_binding" "foo" {
  project = google_beyondcorp_application.example.project
  security_gateways_id = google_beyondcorp_application.example.security_gateways_id
  application_id = google_beyondcorp_application.example.application_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}

resource "google_beyondcorp_application_iam_binding" "foo2" {
  project = google_beyondcorp_application.example.project
  security_gateways_id = google_beyondcorp_application.example.security_gateways_id
  application_id = google_beyondcorp_application.example.application_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}

resource "google_beyondcorp_application_iam_binding" "foo3" {
  project = google_beyondcorp_application.example.project
  security_gateways_id = google_beyondcorp_application.example.security_gateways_id
  application_id = google_beyondcorp_application.example.application_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    # Check that lack of description doesn't cause any issues
    # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
    title       = "%{condition_title_no_desc}"
    expression  = "%{condition_expr_no_desc}"
  }
}
`, context)
}

func testAccBeyondcorpApplicationIamMember_withConditionGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_beyondcorp_security_gateway" "default" {
  security_gateway_id = "default%{random_suffix}"
  display_name = "My Security Gateway resource"
  hubs { region = "us-central1" }
}

resource "google_beyondcorp_application" "example" {
  security_gateways_id = google_beyondcorp_security_gateway.default.security_gateway_id
  application_id = "google%{random_suffix}"
  endpoint_matchers {
    hostname = "google.com"
  }
}

resource "google_beyondcorp_application_iam_member" "foo" {
  project = google_beyondcorp_application.example.project
  security_gateways_id = google_beyondcorp_application.example.security_gateways_id
  application_id = google_beyondcorp_application.example.application_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccBeyondcorpApplicationIamMember_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_beyondcorp_security_gateway" "default" {
  security_gateway_id = "default%{random_suffix}"
  display_name = "My Security Gateway resource"
  hubs { region = "us-central1" }
}

resource "google_beyondcorp_application" "example" {
  security_gateways_id = google_beyondcorp_security_gateway.default.security_gateway_id
  application_id = "google%{random_suffix}"
  endpoint_matchers {
    hostname = "google.com"
  }
}

resource "google_beyondcorp_application_iam_member" "foo" {
  project = google_beyondcorp_application.example.project
  security_gateways_id = google_beyondcorp_application.example.security_gateways_id
  application_id = google_beyondcorp_application.example.application_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}

resource "google_beyondcorp_application_iam_member" "foo2" {
  project = google_beyondcorp_application.example.project
  security_gateways_id = google_beyondcorp_application.example.security_gateways_id
  application_id = google_beyondcorp_application.example.application_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}

resource "google_beyondcorp_application_iam_member" "foo3" {
  project = google_beyondcorp_application.example.project
  security_gateways_id = google_beyondcorp_application.example.security_gateways_id
  application_id = google_beyondcorp_application.example.application_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    # Check that lack of description doesn't cause any issues
    # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
    title       = "%{condition_title_no_desc}"
    expression  = "%{condition_expr_no_desc}"
  }
}
`, context)
}

func testAccBeyondcorpApplicationIamPolicy_withConditionGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_beyondcorp_security_gateway" "default" {
  security_gateway_id = "default%{random_suffix}"
  display_name = "My Security Gateway resource"
  hubs { region = "us-central1" }
}

resource "google_beyondcorp_application" "example" {
  security_gateways_id = google_beyondcorp_security_gateway.default.security_gateway_id
  application_id = "google%{random_suffix}"
  endpoint_matchers {
    hostname = "google.com"
  }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
    condition {
      # Check that lack of description doesn't cause any issues
      # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
      title       = "%{condition_title_no_desc}"
      expression  = "%{condition_expr_no_desc}"
    }
  }
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
    condition {
      title       = "%{condition_title}"
      description = "%{condition_desc}"
      expression  = "%{condition_expr}"
    }
  }
}

resource "google_beyondcorp_application_iam_policy" "foo" {
  project = google_beyondcorp_application.example.project
  security_gateways_id = google_beyondcorp_application.example.security_gateways_id
  application_id = google_beyondcorp_application.example.application_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}
