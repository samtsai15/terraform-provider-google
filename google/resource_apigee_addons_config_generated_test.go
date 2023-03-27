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

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccApigeeAddonsConfig_apigeeAddonsTestExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":          GetTestOrgFromEnv(t),
		"billing_account": GetTestBillingAccountFromEnv(t),
		"random_suffix":   RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckApigeeAddonsConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccApigeeAddonsConfig_apigeeAddonsTestExample(context),
			},
			{
				ResourceName:            "google_apigee_addons_config.apigee_org_addons",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"org"},
			},
		},
	})
}

func testAccApigeeAddonsConfig_apigeeAddonsTestExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  project_id      = "tf-test-%{random_suffix}"
  name            = "tf-test-%{random_suffix}"
  org_id          = "%{org_id}"
  billing_account = "%{billing_account}"
}

resource "google_project_service" "apigee" {
  project = google_project.project.project_id
  service = "apigee.googleapis.com"
}

resource "google_project_service" "compute" {
  project = google_project.project.project_id
  service = "compute.googleapis.com"
  depends_on = [ google_project_service.servicenetworking ]
}

resource "google_project_service" "servicenetworking" {
  project = google_project.project.project_id
  service = "servicenetworking.googleapis.com"
  depends_on = [ google_project_service.apigee ]
}

resource "google_compute_network" "apigee_network" {
  name       = "apigee-network"
  project    = google_project.project.project_id
  depends_on = [ google_project_service.compute ]
}

resource "google_compute_global_address" "apigee_range" {
  name          = "apigee-range"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = google_compute_network.apigee_network.id
  project       = google_project.project.project_id
}

resource "google_service_networking_connection" "apigee_vpc_connection" {
  network                 = google_compute_network.apigee_network.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.apigee_range.name]
}

resource "google_apigee_organization" "org" {
  analytics_region   = "us-central1"
  project_id         = google_project.project.project_id
  authorized_network = google_compute_network.apigee_network.id
  billing_type       = "EVALUATION"
  depends_on         = [
    google_service_networking_connection.apigee_vpc_connection,
    google_project_service.apigee
  ]
}

resource "google_apigee_addons_config" "apigee_org_addons" {
  org = google_apigee_organization.org.name

  addons_config {
    integration_config {
      enabled = true
    }
    api_security_config {
      enabled = true
    }
    connectors_platform_config {
      enabled = true
    }
    monetization_config {
      enabled = true
    }
    advanced_api_ops_config {
      enabled = true
    }
  }
}
`, context)
}

func testAccCheckApigeeAddonsConfigDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_apigee_addons_config" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{ApigeeBasePath}}organizations/{{org}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			res, err := SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)

			if err != nil {
				// If the Apigee org doesn't exist, then a 403 can also be returned.
				if IsGoogleApiErrorWithCode(err, 403) || IsGoogleApiErrorWithCode(err, 404) {
					return nil
				} else {
					return err
				}
			}

			v, ok := res["addonsConfig"]

			if ok || v != nil {
				return fmt.Errorf("ApigeeAddonConfig still exists at %s", url)
			}
		}

		return nil
	}
}
