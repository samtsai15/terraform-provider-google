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

func TestAccDocumentAIProcessor_documentaiProcessorExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDocumentAIProcessorDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDocumentAIProcessor_documentaiProcessorExample(context),
			},
			{
				ResourceName:            "google_document_ai_processor.processor",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location"},
			},
		},
	})
}

func testAccDocumentAIProcessor_documentaiProcessorExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_document_ai_processor" "processor" {
  location = "us"
  display_name = "tf-test-test-processor%{random_suffix}"
  type = "OCR_PROCESSOR"
}
`, context)
}

func TestAccDocumentAIProcessor_documentaiProcessorEuExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDocumentAIProcessorDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDocumentAIProcessor_documentaiProcessorEuExample(context),
			},
			{
				ResourceName:            "google_document_ai_processor.processor",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location"},
			},
		},
	})
}

func testAccDocumentAIProcessor_documentaiProcessorEuExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_document_ai_processor" "processor" {
  location = "eu"
  display_name = "tf-test-test-processor%{random_suffix}"
  type = "OCR_PROCESSOR"
}
`, context)
}

func testAccCheckDocumentAIProcessorDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_document_ai_processor" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{DocumentAIBasePath}}projects/{{project}}/locations/{{location}}/processors/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("DocumentAIProcessor still exists at %s", url)
			}
		}

		return nil
	}
}
