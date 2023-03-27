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

func TestAccArtifactRegistryRepository_artifactRegistryRepositoryBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckArtifactRegistryRepositoryDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccArtifactRegistryRepository_artifactRegistryRepositoryBasicExample(context),
			},
			{
				ResourceName:            "google_artifact_registry_repository.my-repo",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"repository_id", "location"},
			},
		},
	})
}

func testAccArtifactRegistryRepository_artifactRegistryRepositoryBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_artifact_registry_repository" "my-repo" {
  location      = "us-central1"
  repository_id = "tf-test-my-repository%{random_suffix}"
  description   = "example docker repository%{random_suffix}"
  format        = "DOCKER"
}
`, context)
}

func TestAccArtifactRegistryRepository_artifactRegistryRepositoryCmekExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"kms_key_name":  BootstrapKMSKeyInLocation(t, "us-central1").CryptoKey.Name,
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckArtifactRegistryRepositoryDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccArtifactRegistryRepository_artifactRegistryRepositoryCmekExample(context),
			},
			{
				ResourceName:            "google_artifact_registry_repository.my-repo",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"repository_id", "location"},
			},
		},
	})
}

func testAccArtifactRegistryRepository_artifactRegistryRepositoryCmekExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_artifact_registry_repository" "my-repo" {
  location      = "us-central1"
  repository_id = "tf-test-my-repository%{random_suffix}"
  description   = "example docker repository with cmek"
  format        = "DOCKER"
  kms_key_name  = "%{kms_key_name}"
  depends_on = [
    google_kms_crypto_key_iam_member.crypto_key
  ]
}

resource "google_kms_crypto_key_iam_member" "crypto_key" {
  crypto_key_id = "%{kms_key_name}"
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member        = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-artifactregistry.iam.gserviceaccount.com"
}

data "google_project" "project" {}
`, context)
}

func testAccCheckArtifactRegistryRepositoryDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_artifact_registry_repository" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{ArtifactRegistryBasePath}}projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("ArtifactRegistryRepository still exists at %s", url)
			}
		}

		return nil
	}
}
