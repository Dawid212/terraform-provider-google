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
//	Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/dataproc/iam_dataproc_cluster_test.go
//
//	DO NOT EDIT this file directly. Any changes made to this file will be
//	overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
package dataproc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
)

func TestAccDataprocClusterIamBinding(t *testing.T) {
	t.Parallel()

	cluster := "tf-test-iam-" + acctest.RandString(t, 10)
	account := "tf-test-dpiam-" + acctest.RandString(t, 10)
	role := "roles/editor"

	networkName := acctest.BootstrapSharedTestNetwork(t, "dataproc-cluster")
	subnetworkName := acctest.BootstrapSubnet(t, "dataproc-cluster", networkName)
	acctest.BootstrapFirewallForDataprocSharedNetwork(t, "dataproc-cluster", networkName)

	importId := fmt.Sprintf("projects/%s/regions/%s/clusters/%s %s",
		envvar.GetTestProjectFromEnv(), "us-central1", cluster, role)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test IAM Binding creation
				Config: testAccDataprocClusterIamBinding_basic(cluster, subnetworkName, account, role),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"google_dataproc_cluster_iam_binding.binding", "role", role),
				),
			},
			{
				ResourceName:      "google_dataproc_cluster_iam_binding.binding",
				ImportStateId:     importId,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test IAM Binding update
				Config: testAccDataprocClusterIamBinding_update(cluster, subnetworkName, account, role),
			},
			{
				ResourceName:      "google_dataproc_cluster_iam_binding.binding",
				ImportStateId:     importId,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDataprocClusterIamMember(t *testing.T) {
	t.Parallel()

	cluster := "tf-test-iam-" + acctest.RandString(t, 10)
	account := "tf-test-dpiam-" + acctest.RandString(t, 10)
	role := "roles/editor"

	networkName := acctest.BootstrapSharedTestNetwork(t, "dataproc-cluster")
	subnetworkName := acctest.BootstrapSubnet(t, "dataproc-cluster", networkName)
	acctest.BootstrapFirewallForDataprocSharedNetwork(t, "dataproc-cluster", networkName)

	importId := fmt.Sprintf("projects/%s/regions/%s/clusters/%s %s serviceAccount:%s",
		envvar.GetTestProjectFromEnv(),
		"us-central1",
		cluster,
		role,
		envvar.ServiceAccountCanonicalEmail(account))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test IAM Binding creation
				Config: testAccDataprocClusterIamMember(cluster, subnetworkName, account, role),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"google_dataproc_cluster_iam_member.member", "role", role),
					resource.TestCheckResourceAttr(
						"google_dataproc_cluster_iam_member.member", "member", "serviceAccount:"+envvar.ServiceAccountCanonicalEmail(account)),
				),
			},
			{
				ResourceName:      "google_dataproc_cluster_iam_member.member",
				ImportStateId:     importId,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDataprocClusterIamPolicy(t *testing.T) {
	t.Parallel()

	cluster := "tf-test-iam-" + acctest.RandString(t, 10)
	account := "tf-test-dpiam-" + acctest.RandString(t, 10)
	role := "roles/editor"

	networkName := acctest.BootstrapSharedTestNetwork(t, "dataproc-cluster")
	subnetworkName := acctest.BootstrapSubnet(t, "dataproc-cluster", networkName)
	acctest.BootstrapFirewallForDataprocSharedNetwork(t, "dataproc-cluster", networkName)

	importId := fmt.Sprintf("projects/%s/regions/%s/clusters/%s",
		envvar.GetTestProjectFromEnv(), "us-central1", cluster)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test IAM Binding creation
				Config: testAccDataprocClusterIamPolicy(cluster, subnetworkName, account, role),
				Check:  resource.TestCheckResourceAttrSet("data.google_dataproc_cluster_iam_policy.policy", "policy_data"),
			},
			{
				ResourceName:      "google_dataproc_cluster_iam_policy.policy",
				ImportStateId:     importId,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDataprocClusterIamBinding_basic(cluster, subnetworkName, account, role string) string {
	return fmt.Sprintf(testDataprocIamSingleNodeCluster+`
resource "google_service_account" "test-account1" {
  account_id   = "%s-1"
  display_name = "Dataproc Cluster IAM Testing Account"
}

resource "google_service_account" "test-account2" {
  account_id   = "%s-2"
  display_name = "Dataproc Cluster Iam Testing Account"
}

resource "google_dataproc_cluster_iam_binding" "binding" {
  cluster = google_dataproc_cluster.cluster.name
  region  = "us-central1"
  role    = "%s"
  members = [
    "serviceAccount:${google_service_account.test-account1.email}",
  ]
}
`, cluster, subnetworkName, account, account, role)
}

func testAccDataprocClusterIamBinding_update(cluster, subnetworkName, account, role string) string {
	return fmt.Sprintf(testDataprocIamSingleNodeCluster+`
resource "google_service_account" "test-account1" {
  account_id   = "%s-1"
  display_name = "Dataproc Cluster IAM Testing Account"
}

resource "google_service_account" "test-account2" {
  account_id   = "%s-2"
  display_name = "Dataproc Cluster Iam Testing Account"
}

resource "google_dataproc_cluster_iam_binding" "binding" {
  cluster = google_dataproc_cluster.cluster.name
  region  = "us-central1"
  role    = "%s"
  members = [
    "serviceAccount:${google_service_account.test-account1.email}",
    "serviceAccount:${google_service_account.test-account2.email}",
  ]
}
`, cluster, subnetworkName, account, account, role)
}

func testAccDataprocClusterIamMember(cluster, subnetworkName, account, role string) string {
	return fmt.Sprintf(testDataprocIamSingleNodeCluster+`
resource "google_service_account" "test-account" {
  account_id   = "%s"
  display_name = "Dataproc Cluster IAM Testing Account"
}

resource "google_dataproc_cluster_iam_member" "member" {
  cluster = google_dataproc_cluster.cluster.name
  role    = "%s"
  member  = "serviceAccount:${google_service_account.test-account.email}"
}
`, cluster, subnetworkName, account, role)
}

func testAccDataprocClusterIamPolicy(cluster, subnetworkName, account, role string) string {
	return fmt.Sprintf(testDataprocIamSingleNodeCluster+`
resource "google_service_account" "test-account" {
  account_id   = "%s"
  display_name = "Dataproc Cluster IAM Testing Account"
}

data "google_iam_policy" "policy" {
  binding {
    role    = "%s"
    members = ["serviceAccount:${google_service_account.test-account.email}"]
  }
}

resource "google_dataproc_cluster_iam_policy" "policy" {
  cluster     = google_dataproc_cluster.cluster.name
  region      = "us-central1"
  policy_data = data.google_iam_policy.policy.policy_data
}

data "google_dataproc_cluster_iam_policy" "policy" {
  cluster     = google_dataproc_cluster.cluster.name
  region      = "us-central1"
}

`, cluster, subnetworkName, account, role)
}

// Smallest cluster possible for testing
var testDataprocIamSingleNodeCluster = `
resource "google_dataproc_cluster" "cluster" {
  name   = "%s"
  region = "us-central1"

  cluster_config {
    gce_cluster_config {
      subnetwork = "%s"
    }
      
    # Keep the costs down with smallest config we can get away with
    software_config {
      override_properties = {
        "dataproc:dataproc.allow.zero.workers" = "true"
      }
    }

    master_config {
      num_instances = 1
      machine_type  = "e2-standard-2"
      disk_config {
        boot_disk_size_gb = 35
      }
    }
  }
}`
