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
//	Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/beyondcorp/data_source_google_beyondcorp_app_gateway_test.go
//
//	DO NOT EDIT this file directly. Any changes made to this file will be
//	overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
package beyondcorp_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google/google/acctest"
)

func TestAccDataSourceGoogleBeyondcorpAppGateway_basic(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckBeyondcorpAppGatewayDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGoogleBeyondcorpAppGateway_basic(context),
				Check: resource.ComposeTestCheckFunc(
					acctest.CheckDataSourceStateMatchesResourceState("data.google_beyondcorp_app_gateway.foo", "google_beyondcorp_app_gateway.foo"),
				),
			},
		},
	})
}

func TestAccDataSourceGoogleBeyondcorpAppGateway_optionalProject(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckBeyondcorpAppGatewayDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGoogleBeyondcorpAppGateway_optionalProject(context),
				Check: resource.ComposeTestCheckFunc(
					acctest.CheckDataSourceStateMatchesResourceState("data.google_beyondcorp_app_gateway.foo", "google_beyondcorp_app_gateway.foo"),
				),
			},
		},
	})
}

func TestAccDataSourceGoogleBeyondcorpAppGateway_optionalRegion(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckBeyondcorpAppGatewayDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGoogleBeyondcorpAppGateway_optionalRegion(context),
				Check: resource.ComposeTestCheckFunc(
					acctest.CheckDataSourceStateMatchesResourceState("data.google_beyondcorp_app_gateway.foo", "google_beyondcorp_app_gateway.foo"),
				),
			},
		},
	})
}

func TestAccDataSourceGoogleBeyondcorpAppGateway_optionalProjectRegion(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckBeyondcorpAppGatewayDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGoogleBeyondcorpAppGateway_optionalProjectRegion(context),
				Check: resource.ComposeTestCheckFunc(
					acctest.CheckDataSourceStateMatchesResourceState("data.google_beyondcorp_app_gateway.foo", "google_beyondcorp_app_gateway.foo"),
				),
			},
		},
	})
}

func testAccDataSourceGoogleBeyondcorpAppGateway_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_beyondcorp_app_gateway" "foo" {
	name      = "tf-test-appgateway-%{random_suffix}"
	type      = "TCP_PROXY"
	host_type = "GCP_REGIONAL_MIG"
	labels = {
		my-label = "my-label-value"
	}
}

data "google_beyondcorp_app_gateway" "foo" {
	name    = google_beyondcorp_app_gateway.foo.name
	project = google_beyondcorp_app_gateway.foo.project
	region  = google_beyondcorp_app_gateway.foo.region
}
`, context)
}

func testAccDataSourceGoogleBeyondcorpAppGateway_optionalProject(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_beyondcorp_app_gateway" "foo" {
	name      = "tf-test-appgateway-%{random_suffix}"
	type      = "TCP_PROXY"
	host_type = "GCP_REGIONAL_MIG"
}

data "google_beyondcorp_app_gateway" "foo" {
	name   = google_beyondcorp_app_gateway.foo.name
	region = google_beyondcorp_app_gateway.foo.region
}
`, context)
}

func testAccDataSourceGoogleBeyondcorpAppGateway_optionalRegion(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_beyondcorp_app_gateway" "foo" {
	name      = "tf-test-appgateway-%{random_suffix}"
	type      = "TCP_PROXY"
	host_type = "GCP_REGIONAL_MIG"
}

data "google_beyondcorp_app_gateway" "foo" {
	name    = google_beyondcorp_app_gateway.foo.name
	project = google_beyondcorp_app_gateway.foo.project
}
`, context)
}

func testAccDataSourceGoogleBeyondcorpAppGateway_optionalProjectRegion(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_beyondcorp_app_gateway" "foo" {
	name      = "tf-test-appgateway-%{random_suffix}"
	type      = "TCP_PROXY"
	host_type = "GCP_REGIONAL_MIG"
}

data "google_beyondcorp_app_gateway" "foo" {
	name = google_beyondcorp_app_gateway.foo.name
}
`, context)
}
