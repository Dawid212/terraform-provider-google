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
//	Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/securitycenter/resource_scc_project_notification_config_test.go
//
//	DO NOT EDIT this file directly. Any changes made to this file will be
//	overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
package securitycenter_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
)

func TestAccSecurityCenterProjectNotificationConfig_updateStreamingConfigFilter(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSecurityCenterProjectNotificationConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityCenterProjectNotificationConfig_sccProjectNotificationConfigBasicExample(context),
			},
			{
				ResourceName:            "google_scc_project_notification_config.custom_notification_config",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"project", "config_id"},
			},
			{
				Config: testAccSecurityCenterProjectNotificationConfig_updateStreamingConfigFilter(context),
			},
			{
				ResourceName:            "google_scc_project_notification_config.custom_notification_config",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"project", "config_id"},
			},
			{
				Config: testAccSecurityCenterProjectNotificationConfig_emptyStreamingConfigFilter(context),
			},
			{
				ResourceName:            "google_scc_project_notification_config.custom_notification_config",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"project", "config_id"},
			},
		},
	})
}

func testAccSecurityCenterProjectNotificationConfig_updateStreamingConfigFilter(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_pubsub_topic" "scc_project_notification" {
  name = "tf-test-my-topic%{random_suffix}"
}

resource "google_scc_project_notification_config" "custom_notification_config" {
  config_id    = "tf-test-my-config%{random_suffix}"
  project      = "%{project}"
  description  = "My custom Cloud Security Command Center Finding Notification Configuration"
  pubsub_topic = google_pubsub_topic.scc_project_notification.id

  streaming_config {
    filter = "category = \"OPEN_FIREWALL\""
  }
}
`, context)
}

func testAccSecurityCenterProjectNotificationConfig_emptyStreamingConfigFilter(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_pubsub_topic" "scc_project_notification" {
  name = "tf-test-my-topic%{random_suffix}"
}

resource "google_scc_project_notification_config" "custom_notification_config" {
  config_id    = "tf-test-my-config%{random_suffix}"
  project      = "%{project}"
  description  = "My custom Cloud Security Command Center Finding Notification Configuration"
  pubsub_topic = google_pubsub_topic.scc_project_notification.id

  streaming_config {
    filter = ""
  }
}
`, context)
}
