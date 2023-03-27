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
)

func TestAccStorageBucketIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/storage.objectViewer",
		"admin_role":              "roles/storage.admin",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccStorageBucketIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_storage_bucket_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("b/%s roles/storage.objectViewer", fmt.Sprintf("tf-test-my-bucket%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccStorageBucketIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_storage_bucket_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("b/%s roles/storage.objectViewer", fmt.Sprintf("tf-test-my-bucket%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccStorageBucketIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/storage.objectViewer",
		"admin_role":              "roles/storage.admin",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccStorageBucketIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_storage_bucket_iam_member.foo",
				ImportStateId:     fmt.Sprintf("b/%s roles/storage.objectViewer user:admin@hashicorptest.com", fmt.Sprintf("tf-test-my-bucket%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccStorageBucketIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	// This may skip test, so do it first
	sa := GetTestServiceAccountFromEnv(t)
	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/storage.objectViewer",
		"admin_role":              "roles/storage.admin",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}
	context["service_account"] = sa

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccStorageBucketIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_storage_bucket_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("b/%s", fmt.Sprintf("tf-test-my-bucket%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccStorageBucketIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_storage_bucket_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("b/%s", fmt.Sprintf("tf-test-my-bucket%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccStorageBucketIamBindingGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/storage.objectViewer",
		"admin_role":              "roles/storage.admin",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccStorageBucketIamBinding_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_storage_bucket_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("b/%s roles/storage.objectViewer %s", fmt.Sprintf("tf-test-my-bucket%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccStorageBucketIamBindingGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/storage.objectViewer",
		"admin_role":              "roles/storage.admin",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccStorageBucketIamBinding_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_storage_bucket_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("b/%s roles/storage.objectViewer", fmt.Sprintf("tf-test-my-bucket%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_storage_bucket_iam_binding.foo2",
				ImportStateId:     fmt.Sprintf("b/%s roles/storage.objectViewer %s", fmt.Sprintf("tf-test-my-bucket%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_storage_bucket_iam_binding.foo3",
				ImportStateId:     fmt.Sprintf("b/%s roles/storage.objectViewer %s", fmt.Sprintf("tf-test-my-bucket%s", context["random_suffix"]), context["condition_title_no_desc"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccStorageBucketIamMemberGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/storage.objectViewer",
		"admin_role":              "roles/storage.admin",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccStorageBucketIamMember_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_storage_bucket_iam_member.foo",
				ImportStateId:     fmt.Sprintf("b/%s roles/storage.objectViewer user:admin@hashicorptest.com %s", fmt.Sprintf("tf-test-my-bucket%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccStorageBucketIamMemberGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/storage.objectViewer",
		"admin_role":              "roles/storage.admin",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccStorageBucketIamMember_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_storage_bucket_iam_member.foo",
				ImportStateId:     fmt.Sprintf("b/%s roles/storage.objectViewer user:admin@hashicorptest.com", fmt.Sprintf("tf-test-my-bucket%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_storage_bucket_iam_member.foo2",
				ImportStateId:     fmt.Sprintf("b/%s roles/storage.objectViewer user:admin@hashicorptest.com %s", fmt.Sprintf("tf-test-my-bucket%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_storage_bucket_iam_member.foo3",
				ImportStateId:     fmt.Sprintf("b/%s roles/storage.objectViewer user:admin@hashicorptest.com %s", fmt.Sprintf("tf-test-my-bucket%s", context["random_suffix"]), context["condition_title_no_desc"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccStorageBucketIamPolicyGenerated_withCondition(t *testing.T) {
	t.Parallel()

	// This may skip test, so do it first
	sa := GetTestServiceAccountFromEnv(t)
	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/storage.objectViewer",
		"admin_role":              "roles/storage.admin",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}
	context["service_account"] = sa

	// Test should have 3 bindings: one with a description and one without, and a third for an admin role. Any < chars are converted to a unicode character by the API.
	expectedPolicyData := Nprintf(`{"bindings":[{"members":["serviceAccount:%{service_account}"],"role":"%{admin_role}"},{"condition":{"description":"%{condition_desc}","expression":"%{condition_expr}","title":"%{condition_title}"},"members":["user:admin@hashicorptest.com"],"role":"%{role}"},{"condition":{"expression":"%{condition_expr}","title":"%{condition_title}-no-description"},"members":["user:admin@hashicorptest.com"],"role":"%{role}"}]}`, context)
	expectedPolicyData = strings.Replace(expectedPolicyData, "<", "\\u003c", -1)

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccStorageBucketIamPolicy_withConditionGenerated(context),
				Check: resource.ComposeAggregateTestCheckFunc(
					// TODO(SarahFrench) - uncomment once https://github.com/GoogleCloudPlatform/magic-modules/pull/6466 merged
					// resource.TestCheckResourceAttr("data.google_iam_policy.foo", "policy_data", expectedPolicyData),
					resource.TestCheckResourceAttr("google_storage_bucket_iam_policy.foo", "policy_data", expectedPolicyData),
					resource.TestCheckResourceAttrWith("data.google_iam_policy.foo", "policy_data", checkGoogleIamPolicy),
				),
			},
			{
				ResourceName:      "google_storage_bucket_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("b/%s", fmt.Sprintf("tf-test-my-bucket%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccStorageBucketIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "default" {
  name                        = "tf-test-my-bucket%{random_suffix}"
  location                    = "US"
  uniform_bucket_level_access = true
}

resource "google_storage_bucket_iam_member" "foo" {
  bucket = google_storage_bucket.default.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccStorageBucketIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "default" {
  name                        = "tf-test-my-bucket%{random_suffix}"
  location                    = "US"
  uniform_bucket_level_access = true
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
  binding {
    role = "%{admin_role}"
    members = ["serviceAccount:%{service_account}"]
  }
}

resource "google_storage_bucket_iam_policy" "foo" {
  bucket = google_storage_bucket.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccStorageBucketIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "default" {
  name                        = "tf-test-my-bucket%{random_suffix}"
  location                    = "US"
  uniform_bucket_level_access = true
}

data "google_iam_policy" "foo" {
}

resource "google_storage_bucket_iam_policy" "foo" {
  bucket = google_storage_bucket.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccStorageBucketIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "default" {
  name                        = "tf-test-my-bucket%{random_suffix}"
  location                    = "US"
  uniform_bucket_level_access = true
}

resource "google_storage_bucket_iam_binding" "foo" {
  bucket = google_storage_bucket.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccStorageBucketIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "default" {
  name                        = "tf-test-my-bucket%{random_suffix}"
  location                    = "US"
  uniform_bucket_level_access = true
}

resource "google_storage_bucket_iam_binding" "foo" {
  bucket = google_storage_bucket.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}

func testAccStorageBucketIamBinding_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "default" {
  name                        = "tf-test-my-bucket%{random_suffix}"
  location                    = "US"
  uniform_bucket_level_access = true
}

resource "google_storage_bucket_iam_binding" "foo" {
  bucket = google_storage_bucket.default.name
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

func testAccStorageBucketIamBinding_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "default" {
  name                        = "tf-test-my-bucket%{random_suffix}"
  location                    = "US"
  uniform_bucket_level_access = true
}

resource "google_storage_bucket_iam_binding" "foo" {
  bucket = google_storage_bucket.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}

resource "google_storage_bucket_iam_binding" "foo2" {
  bucket = google_storage_bucket.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}

resource "google_storage_bucket_iam_binding" "foo3" {
  bucket = google_storage_bucket.default.name
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

func testAccStorageBucketIamMember_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "default" {
  name                        = "tf-test-my-bucket%{random_suffix}"
  location                    = "US"
  uniform_bucket_level_access = true
}

resource "google_storage_bucket_iam_member" "foo" {
  bucket = google_storage_bucket.default.name
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

func testAccStorageBucketIamMember_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "default" {
  name                        = "tf-test-my-bucket%{random_suffix}"
  location                    = "US"
  uniform_bucket_level_access = true
}

resource "google_storage_bucket_iam_member" "foo" {
  bucket = google_storage_bucket.default.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}

resource "google_storage_bucket_iam_member" "foo2" {
  bucket = google_storage_bucket.default.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}

resource "google_storage_bucket_iam_member" "foo3" {
  bucket = google_storage_bucket.default.name
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

func testAccStorageBucketIamPolicy_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "default" {
  name                        = "tf-test-my-bucket%{random_suffix}"
  location                    = "US"
  uniform_bucket_level_access = true
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
  binding {
    role = "%{admin_role}"
    members = ["serviceAccount:%{service_account}"]
  }
}

resource "google_storage_bucket_iam_policy" "foo" {
  bucket = google_storage_bucket.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}
