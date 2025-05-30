// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package identityplatform_test

import (
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
)

func TestAccIdentityPlatformConfig_identityPlatformConfigBasicExample(t *testing.T) {
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"billing_acct":     envvar.GetTestBillingAccountFromEnv(t),
		"org_id":           envvar.GetTestOrgFromEnv(t),
		"quota_start_time": time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
		"random_suffix":    acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIdentityPlatformConfig_identityPlatformConfigBasicExample(context),
			},
			{
				ResourceName:            "google_identity_platform_config.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"client.0.api_key", "client.0.firebase_subdomain"},
			},
		},
	})
}

func testAccIdentityPlatformConfig_identityPlatformConfigBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_project" "default" {
  project_id = "tf-test-my-project%{random_suffix}"
  name       = "tf-test-my-project%{random_suffix}"
  org_id     = "%{org_id}"
  billing_account =  "%{billing_acct}"
  deletion_policy = "DELETE"
  labels = {
    firebase = "enabled"
  }
}

resource "google_project_service" "identitytoolkit" {
  project = google_project.default.project_id
  service = "identitytoolkit.googleapis.com"
}

resource "google_identity_platform_config" "default" {
  project = google_project.default.project_id
  autodelete_anonymous_users = true
  sign_in {
    allow_duplicate_emails = true
   
    anonymous {
        enabled = true
    }
    email {
        enabled = true
        password_required = false
    }
    phone_number {
        enabled = true
        test_phone_numbers = {
            "+11231231234" = "000000"
        }
    }
  }
  sms_region_config {
    allowlist_only {
      allowed_regions = [
        "US",
        "CA",
      ]
    }
  }
  blocking_functions {
    triggers {
      event_type = "beforeSignIn"
      function_uri = "https://us-east1-tf-test-my-project%{random_suffix}.cloudfunctions.net/before-sign-in"
    }
    forward_inbound_credentials {
      refresh_token = true
      access_token = true
      id_token = true
    }
  }
  quota {
    sign_up_quota_config {
      quota = 1000
      start_time = "%{quota_start_time}"
      quota_duration = "7200s"
    }
  }
  authorized_domains = [
    "localhost",
    "tf-test-my-project%{random_suffix}.firebaseapp.com",
    "tf-test-my-project%{random_suffix}.web.app",
  ]
}
`, context)
}

func TestAccIdentityPlatformConfig_identityPlatformConfigMinimalExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"billing_acct":  envvar.GetTestBillingAccountFromEnv(t),
		"org_id":        envvar.GetTestOrgFromEnv(t),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIdentityPlatformConfig_identityPlatformConfigMinimalExample(context),
			},
			{
				ResourceName:            "google_identity_platform_config.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"client.0.api_key", "client.0.firebase_subdomain"},
			},
		},
	})
}

func testAccIdentityPlatformConfig_identityPlatformConfigMinimalExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_project" "default" {
  project_id = "tf-test-my-project-1%{random_suffix}"
  name       = "tf-test-my-project-1%{random_suffix}"
  org_id     = "%{org_id}"
  billing_account =  "%{billing_acct}"
  deletion_policy = "DELETE"
  labels = {
    firebase = "enabled"
  }
}

resource "google_project_service" "identitytoolkit" {
  project = google_project.default.project_id
  service = "identitytoolkit.googleapis.com"
}


resource "google_identity_platform_config" "default" {
  project = google_project.default.project_id
  client {
    permissions {
      disabled_user_deletion = false
      disabled_user_signup   = true
    }
  }

  mfa {
    enabled_providers = ["PHONE_SMS"]
    provider_configs {
      state = "ENABLED"
      totp_provider_config {
        adjacent_intervals = 3
      }
    }
    state = "ENABLED"
  }
  monitoring {
    request_logging {
      enabled = true
    }
  }
  multi_tenant {
    allow_tenants           = true
    default_tenant_location = "organizations/%{org_id}"
  }

  depends_on = [
    google_project_service.identitytoolkit
  ]
}
`, context)
}

func TestAccIdentityPlatformConfig_identityPlatformConfigWithFalseValuesExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"billing_acct":  envvar.GetTestBillingAccountFromEnv(t),
		"org_id":        envvar.GetTestOrgFromEnv(t),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIdentityPlatformConfig_identityPlatformConfigWithFalseValuesExample(context),
			},
			{
				ResourceName:            "google_identity_platform_config.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"client.0.api_key", "client.0.firebase_subdomain"},
			},
		},
	})
}

func testAccIdentityPlatformConfig_identityPlatformConfigWithFalseValuesExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_project" "default" {
  project_id = "tf-test-my-project-2%{random_suffix}"
  name       = "tf-test-my-project-2%{random_suffix}"
  org_id     = "%{org_id}"
  billing_account =  "%{billing_acct}"
  deletion_policy = "DELETE"
  labels = {
    firebase = "enabled"
  }
}

resource "google_project_service" "identitytoolkit" {
  project = google_project.default.project_id
  service = "identitytoolkit.googleapis.com"
}

resource "google_identity_platform_config" "default" {
  project = google_project.default.project_id
  autodelete_anonymous_users = false
  sign_in {
    allow_duplicate_emails = false
   
    anonymous {
        enabled = false
    }
    email {
        enabled = false
    }
    phone_number {
        enabled = false
    }
  }
  blocking_functions {
    triggers {
      event_type   = "beforeSignIn"
      function_uri = "https://us-east1-tf-test-my-project-2%{random_suffix}.cloudfunctions.net/before-sign-in"
    }
    forward_inbound_credentials {
      refresh_token = false
      access_token  = false
      id_token      = false
    }
  }
  client {
    permissions {
      disabled_user_signup = false
      disabled_user_deletion = false
    }
  }
  multi_tenant {
    allow_tenants = false
  }
  monitoring {
    request_logging {
      enabled = false
    }
  }
}
`, context)
}
