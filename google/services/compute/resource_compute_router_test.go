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
//	Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/compute/resource_compute_router_test.go.tmpl
//
//	DO NOT EDIT this file directly. Any changes made to this file will be
//	overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
package compute_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
)

func TestAccComputeRouter_basic(t *testing.T) {
	t.Parallel()

	testId := acctest.RandString(t, 10)
	routerName := fmt.Sprintf("tf-test-router-%s", testId)
	resourceRegion := "europe-west1"
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRouterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRouterBasic(routerName, resourceRegion),
			},
			{
				ResourceName:      "google_compute_router.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRouter_noRegion(t *testing.T) {
	t.Parallel()

	testId := acctest.RandString(t, 10)
	routerName := fmt.Sprintf("tf-test-router-%s", testId)
	providerRegion := "us-central1"
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRouterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRouterNoRegion(routerName, providerRegion),
			},
			{
				ResourceName:      "google_compute_router.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRouter_full(t *testing.T) {
	t.Parallel()

	testId := acctest.RandString(t, 10)
	routerName := fmt.Sprintf("tf-test-router-%s", testId)
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRouterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRouterFull(routerName),
			},
			{
				ResourceName:      "google_compute_router.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRouter_advertisedIpRangesOrder(t *testing.T) {
	t.Parallel()

	testId := acctest.RandString(t, 10)
	routerName := fmt.Sprintf("tf-test-router-%s", testId)
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRouterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRouterAdvertisedIpRangesOrder(routerName),
			},
			{
				ResourceName:            "google_compute_router.foobar",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"bgp.0.advertised_ip_ranges.0.range", "bgp.0.advertised_ip_ranges.1.range"},
			},
		},
	})
}

func TestAccComputeRouter_update(t *testing.T) {
	t.Parallel()

	testId := acctest.RandString(t, 10)
	routerName := fmt.Sprintf("tf-test-router-%s", testId)
	region := envvar.GetTestRegionFromEnv()
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRouterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRouterBasic(routerName, region),
			},
			{
				ResourceName:      "google_compute_router.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouterFull(routerName),
			},
			{
				ResourceName:      "google_compute_router.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouterBasic(routerName, region),
			},
			{
				ResourceName:      "google_compute_router.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRouter_updateAddRemoveBGP(t *testing.T) {
	t.Parallel()

	testId := acctest.RandString(t, 10)
	routerName := fmt.Sprintf("tf-test-router-%s", testId)
	region := envvar.GetTestRegionFromEnv()
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRouterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRouterBasic(routerName, region),
			},
			{
				ResourceName:      "google_compute_router.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouter_noBGP(routerName, region),
			},
			{
				ResourceName:      "google_compute_router.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouterBasic(routerName, region),
			},
			{
				ResourceName:      "google_compute_router.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRouter_addAndUpdateIdentifierRangeBgp(t *testing.T) {
	t.Parallel()

	testId := acctest.RandString(t, 10)
	routerName := fmt.Sprintf("tf-test-router-%s", testId)
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRouterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRouter_addIdentifierRangeBgp(routerName),
			},
			{
				ResourceName:      "google_compute_router.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouter_updateIdentifierRangeBgp(routerName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("google_compute_router.foobar", "bgp.0.identifier_range", "169.254.8.8/30"),
				),
			},
			{
				ResourceName:      "google_compute_router.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeRouterBasic(routerName, resourceRegion string) string {
	return fmt.Sprintf(`
resource "google_compute_network" "foobar" {
  name                    = "%s-net"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "foobar" {
  name          = "%s-subnet"
  network       = google_compute_network.foobar.self_link
  ip_cidr_range = "10.0.0.0/16"
  region        = "%s"
}

resource "google_compute_router" "foobar" {
  name    = "%s"
  region  = google_compute_subnetwork.foobar.region
  network = google_compute_network.foobar.name
  bgp {
    asn = 4294967294
  }
}
`, routerName, routerName, resourceRegion, routerName)
}

func testAccComputeRouterNoRegion(routerName, providerRegion string) string {
	return fmt.Sprintf(`
resource "google_compute_network" "foobar" {
  name                    = "%s-net"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "foobar" {
  name          = "%s-subnet"
  network       = google_compute_network.foobar.self_link
  ip_cidr_range = "10.0.0.0/16"
  region        = "%s"
}

resource "google_compute_router" "foobar" {
  name    = "%s"
  network = google_compute_network.foobar.name
  bgp {
    asn = 64514
  }
}
`, routerName, routerName, providerRegion, routerName)
}

func testAccComputeRouterFull(routerName string) string {
	return fmt.Sprintf(`
resource "google_compute_network" "foobar" {
  name                    = "%s-net"
  auto_create_subnetworks = false
}

resource "google_compute_router" "foobar" {
  name    = "%s"
  network = google_compute_network.foobar.name
  bgp {
    asn               = 64514
    advertise_mode    = "CUSTOM"
    advertised_groups = ["ALL_SUBNETS"]
    advertised_ip_ranges {
      range = "1.2.3.4"
    }
    advertised_ip_ranges {
      range = "6.7.0.0/16"
    }
    keepalive_interval = 25
  }
}
`, routerName, routerName)
}

func testAccComputeRouterAdvertisedIpRangesOrder(routerName string) string {
	return fmt.Sprintf(`
resource "google_compute_network" "foobar" {
  name                    = "%s-net"
  auto_create_subnetworks = false
}

resource "google_compute_router" "foobar" {
  name    = "%s"
  network = google_compute_network.foobar.name
  bgp {
    asn               = 64514
    advertise_mode    = "CUSTOM"
    advertised_groups = ["ALL_SUBNETS"]
	advertised_ip_ranges {
      range = "6.7.0.0/16"
    }
    advertised_ip_ranges {
      range = "1.2.3.4"
    }
    keepalive_interval = 25
  }
}
`, routerName, routerName)
}

func testAccComputeRouter_noBGP(routerName, resourceRegion string) string {
	return fmt.Sprintf(`
resource "google_compute_network" "foobar" {
  name                    = "%s-net"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "foobar" {
  name          = "%s-subnet"
  network       = google_compute_network.foobar.self_link
  ip_cidr_range = "10.0.0.0/16"
  region        = "%s"
}

resource "google_compute_router" "foobar" {
  name    = "%s"
  region  = google_compute_subnetwork.foobar.region
  network = google_compute_network.foobar.name
}
`, routerName, routerName, resourceRegion, routerName)
}

func testAccComputeRouter_addIdentifierRangeBgp(routerName string) string {
	return fmt.Sprintf(`
resource "google_compute_network" "foobar" {
  name                    = "%s-net"
  auto_create_subnetworks = false
}

resource "google_compute_router" "foobar" {
  name    = "%s"
  network = google_compute_network.foobar.name
  bgp {
    asn               = 64514
    advertise_mode    = "CUSTOM"
    advertised_groups = ["ALL_SUBNETS"]
    advertised_ip_ranges {
      range = "1.2.3.4"
    }
    advertised_ip_ranges {
      range = "6.7.0.0/16"
    }
	identifier_range = "169.254.8.8/29"
    keepalive_interval = 25
  }
}
`, routerName, routerName)
}

func testAccComputeRouter_updateIdentifierRangeBgp(routerName string) string {
	return fmt.Sprintf(`
resource "google_compute_network" "foobar" {
  name                    = "%s-net"
  auto_create_subnetworks = false
}

resource "google_compute_router" "foobar" {
  name    = "%s"
  network = google_compute_network.foobar.name
  bgp {
    asn               = 64514
    advertise_mode    = "CUSTOM"
    advertised_groups = ["ALL_SUBNETS"]
    advertised_ip_ranges {
      range = "1.2.3.4"
    }
    advertised_ip_ranges {
      range = "6.7.0.0/16"
    }
	identifier_range = "169.254.8.8/30"
    keepalive_interval = 25
  }
}
`, routerName, routerName)
}
