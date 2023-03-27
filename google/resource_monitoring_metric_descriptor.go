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
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceMonitoringMetricDescriptor() *schema.Resource {
	return &schema.Resource{
		Create: resourceMonitoringMetricDescriptorCreate,
		Read:   resourceMonitoringMetricDescriptorRead,
		Update: resourceMonitoringMetricDescriptorUpdate,
		Delete: resourceMonitoringMetricDescriptorDelete,

		Importer: &schema.ResourceImporter{
			State: resourceMonitoringMetricDescriptorImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `A detailed description of the metric, which can be used in documentation.`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `A concise name for the metric, which can be displayed in user interfaces. Use sentence case without an ending period, for example "Request count".`,
			},
			"metric_kind": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateEnum([]string{"METRIC_KIND_UNSPECIFIED", "GAUGE", "DELTA", "CUMULATIVE"}),
				Description:  `Whether the metric records instantaneous values, changes to a value, etc. Some combinations of metricKind and valueType might not be supported. Possible values: ["METRIC_KIND_UNSPECIFIED", "GAUGE", "DELTA", "CUMULATIVE"]`,
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The metric type, including its DNS name prefix. The type is not URL-encoded. All service defined metrics must be prefixed with the service name, in the format of {service name}/{relative metric name}, such as cloudsql.googleapis.com/database/cpu/utilization. The relative metric name must have only upper and lower-case letters, digits, '/' and underscores '_' are allowed. Additionally, the maximum number of characters allowed for the relative_metric_name is 100. All user-defined metric types have the DNS name custom.googleapis.com, external.googleapis.com, or logging.googleapis.com/user/.`,
			},
			"value_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateEnum([]string{"BOOL", "INT64", "DOUBLE", "STRING", "DISTRIBUTION"}),
				Description:  `Whether the measurement is an integer, a floating-point number, etc. Some combinations of metricKind and valueType might not be supported. Possible values: ["BOOL", "INT64", "DOUBLE", "STRING", "DISTRIBUTION"]`,
			},
			"labels": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: `The set of labels that can be used to describe a specific instance of this metric type. In order to delete a label, the entire resource must be deleted, then created with the desired labels.`,
				Elem:        monitoringMetricDescriptorLabelsSchema(),
				// Default schema.HashSchema is used.
			},
			"launch_stage": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validateEnum([]string{"LAUNCH_STAGE_UNSPECIFIED", "UNIMPLEMENTED", "PRELAUNCH", "EARLY_ACCESS", "ALPHA", "BETA", "GA", "DEPRECATED", ""}),
				Description:  `The launch stage of the metric definition. Possible values: ["LAUNCH_STAGE_UNSPECIFIED", "UNIMPLEMENTED", "PRELAUNCH", "EARLY_ACCESS", "ALPHA", "BETA", "GA", "DEPRECATED"]`,
			},
			"metadata": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `Metadata which can be used to guide usage of the metric.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ingest_delay": {
							Type:         schema.TypeString,
							Optional:     true,
							Description:  `The delay of data points caused by ingestion. Data points older than this age are guaranteed to be ingested and available to be read, excluding data loss due to errors. In '[duration format](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf?&_ga=2.264881487.1507873253.1593446723-935052455.1591817775#google.protobuf.Duration)'.`,
							AtLeastOneOf: []string{"metadata.0.sample_period", "metadata.0.ingest_delay"},
						},
						"sample_period": {
							Type:         schema.TypeString,
							Optional:     true,
							Description:  `The sampling period of metric data points. For metrics which are written periodically, consecutive data points are stored at this time interval, excluding data loss due to errors. Metrics with a higher granularity have a smaller sampling period. In '[duration format](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf?&_ga=2.264881487.1507873253.1593446723-935052455.1591817775#google.protobuf.Duration)'.`,
							AtLeastOneOf: []string{"metadata.0.sample_period", "metadata.0.ingest_delay"},
						},
					},
				},
			},
			"unit": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The units in which the metric value is reported. It is only applicable if the
valueType is INT64, DOUBLE, or DISTRIBUTION. The unit defines the representation of
the stored metric values.

Different systems may scale the values to be more easily displayed (so a value of
0.02KBy might be displayed as 20By, and a value of 3523KBy might be displayed as
3.5MBy). However, if the unit is KBy, then the value of the metric is always in
thousands of bytes, no matter how it may be displayed.

If you want a custom metric to record the exact number of CPU-seconds used by a job,
you can create an INT64 CUMULATIVE metric whose unit is s{CPU} (or equivalently
1s{CPU} or just s). If the job uses 12,005 CPU-seconds, then the value is written as
12005.

Alternatively, if you want a custom metric to record data in a more granular way, you
can create a DOUBLE CUMULATIVE metric whose unit is ks{CPU}, and then write the value
12.005 (which is 12005/1000), or use Kis{CPU} and write 11.723 (which is 12005/1024).
The supported units are a subset of The Unified Code for Units of Measure standard.
More info can be found in the API documentation
(https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.metricDescriptors).`,
			},
			"monitored_resource_types": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `If present, then a time series, which is identified partially by a metric type and a MonitoredResourceDescriptor, that is associated with this metric type can only be associated with one of the monitored resource types listed here. This field allows time series to be associated with the intersection of this metric type and the monitored resource types in this list.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The resource name of the metric descriptor.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func monitoringMetricDescriptorLabelsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The key for this label. The key must not exceed 100 characters. The first character of the key must be an upper- or lower-case letter, the remaining characters must be letters, digits or underscores, and the key must match the regular expression [a-zA-Z][a-zA-Z0-9_]*`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A human-readable description for the label.`,
			},
			"value_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateEnum([]string{"STRING", "BOOL", "INT64", ""}),
				Description:  `The type of data that can be assigned to the label. Default value: "STRING" Possible values: ["STRING", "BOOL", "INT64"]`,
				Default:      "STRING",
			},
		},
	}
}

func resourceMonitoringMetricDescriptorCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	typeProp, err := expandMonitoringMetricDescriptorType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !isEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	labelsProp, err := expandMonitoringMetricDescriptorLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	metricKindProp, err := expandMonitoringMetricDescriptorMetricKind(d.Get("metric_kind"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("metric_kind"); !isEmptyValue(reflect.ValueOf(metricKindProp)) && (ok || !reflect.DeepEqual(v, metricKindProp)) {
		obj["metricKind"] = metricKindProp
	}
	valueTypeProp, err := expandMonitoringMetricDescriptorValueType(d.Get("value_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("value_type"); !isEmptyValue(reflect.ValueOf(valueTypeProp)) && (ok || !reflect.DeepEqual(v, valueTypeProp)) {
		obj["valueType"] = valueTypeProp
	}
	unitProp, err := expandMonitoringMetricDescriptorUnit(d.Get("unit"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("unit"); !isEmptyValue(reflect.ValueOf(unitProp)) && (ok || !reflect.DeepEqual(v, unitProp)) {
		obj["unit"] = unitProp
	}
	descriptionProp, err := expandMonitoringMetricDescriptorDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	displayNameProp, err := expandMonitoringMetricDescriptorDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	metadataProp, err := expandMonitoringMetricDescriptorMetadata(d.Get("metadata"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("metadata"); !isEmptyValue(reflect.ValueOf(metadataProp)) && (ok || !reflect.DeepEqual(v, metadataProp)) {
		obj["metadata"] = metadataProp
	}
	launchStageProp, err := expandMonitoringMetricDescriptorLaunchStage(d.Get("launch_stage"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("launch_stage"); !isEmptyValue(reflect.ValueOf(launchStageProp)) && (ok || !reflect.DeepEqual(v, launchStageProp)) {
		obj["launchStage"] = launchStageProp
	}

	url, err := replaceVars(d, config, "{{MonitoringBasePath}}v3/projects/{{project}}/metricDescriptors")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new MetricDescriptor: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for MetricDescriptor: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate), isMonitoringConcurrentEditError)
	if err != nil {
		return fmt.Errorf("Error creating MetricDescriptor: %s", err)
	}
	if err := d.Set("name", flattenMonitoringMetricDescriptorName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = PollingWaitTime(resourceMonitoringMetricDescriptorPollRead(d, meta), PollCheckForExistence, "Creating MetricDescriptor", d.Timeout(schema.TimeoutCreate), 20)
	if err != nil {
		return fmt.Errorf("Error waiting to create MetricDescriptor: %s", err)
	}

	log.Printf("[DEBUG] Finished creating MetricDescriptor %q: %#v", d.Id(), res)

	return resourceMonitoringMetricDescriptorRead(d, meta)
}

func resourceMonitoringMetricDescriptorPollRead(d *schema.ResourceData, meta interface{}) PollReadFunc {
	return func() (map[string]interface{}, error) {
		config := meta.(*Config)

		url, err := replaceVars(d, config, "{{MonitoringBasePath}}v3/{{name}}")
		if err != nil {
			return nil, err
		}

		billingProject := ""

		project, err := getProject(d, config)
		if err != nil {
			return nil, fmt.Errorf("Error fetching project for MetricDescriptor: %s", err)
		}
		billingProject = project

		// err == nil indicates that the billing_project value was found
		if bp, err := getBillingProject(d, config); err == nil {
			billingProject = bp
		}

		userAgent, err := generateUserAgentString(d, config.UserAgent)
		if err != nil {
			return nil, err
		}

		res, err := SendRequest(config, "GET", billingProject, url, userAgent, nil, isMonitoringConcurrentEditError)
		if err != nil {
			return res, err
		}
		return res, nil
	}
}

func resourceMonitoringMetricDescriptorRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{MonitoringBasePath}}v3/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for MetricDescriptor: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequest(config, "GET", billingProject, url, userAgent, nil, isMonitoringConcurrentEditError)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("MonitoringMetricDescriptor %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading MetricDescriptor: %s", err)
	}

	if err := d.Set("name", flattenMonitoringMetricDescriptorName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading MetricDescriptor: %s", err)
	}
	if err := d.Set("type", flattenMonitoringMetricDescriptorType(res["type"], d, config)); err != nil {
		return fmt.Errorf("Error reading MetricDescriptor: %s", err)
	}
	if err := d.Set("labels", flattenMonitoringMetricDescriptorLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading MetricDescriptor: %s", err)
	}
	if err := d.Set("metric_kind", flattenMonitoringMetricDescriptorMetricKind(res["metricKind"], d, config)); err != nil {
		return fmt.Errorf("Error reading MetricDescriptor: %s", err)
	}
	if err := d.Set("value_type", flattenMonitoringMetricDescriptorValueType(res["valueType"], d, config)); err != nil {
		return fmt.Errorf("Error reading MetricDescriptor: %s", err)
	}
	if err := d.Set("unit", flattenMonitoringMetricDescriptorUnit(res["unit"], d, config)); err != nil {
		return fmt.Errorf("Error reading MetricDescriptor: %s", err)
	}
	if err := d.Set("description", flattenMonitoringMetricDescriptorDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading MetricDescriptor: %s", err)
	}
	if err := d.Set("display_name", flattenMonitoringMetricDescriptorDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading MetricDescriptor: %s", err)
	}
	if err := d.Set("monitored_resource_types", flattenMonitoringMetricDescriptorMonitoredResourceTypes(res["monitoredResourceTypes"], d, config)); err != nil {
		return fmt.Errorf("Error reading MetricDescriptor: %s", err)
	}

	return nil
}

func resourceMonitoringMetricDescriptorUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for MetricDescriptor: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	typeProp, err := expandMonitoringMetricDescriptorType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	labelsProp, err := expandMonitoringMetricDescriptorLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	metricKindProp, err := expandMonitoringMetricDescriptorMetricKind(d.Get("metric_kind"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("metric_kind"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, metricKindProp)) {
		obj["metricKind"] = metricKindProp
	}
	valueTypeProp, err := expandMonitoringMetricDescriptorValueType(d.Get("value_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("value_type"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, valueTypeProp)) {
		obj["valueType"] = valueTypeProp
	}
	unitProp, err := expandMonitoringMetricDescriptorUnit(d.Get("unit"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("unit"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, unitProp)) {
		obj["unit"] = unitProp
	}
	descriptionProp, err := expandMonitoringMetricDescriptorDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	displayNameProp, err := expandMonitoringMetricDescriptorDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	metadataProp, err := expandMonitoringMetricDescriptorMetadata(d.Get("metadata"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("metadata"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, metadataProp)) {
		obj["metadata"] = metadataProp
	}
	launchStageProp, err := expandMonitoringMetricDescriptorLaunchStage(d.Get("launch_stage"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("launch_stage"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, launchStageProp)) {
		obj["launchStage"] = launchStageProp
	}

	url, err := replaceVars(d, config, "{{MonitoringBasePath}}v3/projects/{{project}}/metricDescriptors")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating MetricDescriptor %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate), isMonitoringConcurrentEditError)

	if err != nil {
		return fmt.Errorf("Error updating MetricDescriptor %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating MetricDescriptor %q: %#v", d.Id(), res)
	}

	err = PollingWaitTime(resourceMonitoringMetricDescriptorPollRead(d, meta), PollCheckForExistence, "Updating MetricDescriptor", d.Timeout(schema.TimeoutUpdate), 20)
	if err != nil {
		return err
	}

	return resourceMonitoringMetricDescriptorRead(d, meta)
}

func resourceMonitoringMetricDescriptorDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for MetricDescriptor: %s", err)
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{MonitoringBasePath}}v3/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting MetricDescriptor %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete), isMonitoringConcurrentEditError)
	if err != nil {
		return handleNotFoundError(err, d, "MetricDescriptor")
	}

	err = PollingWaitTime(resourceMonitoringMetricDescriptorPollRead(d, meta), PollCheckForAbsence, "Deleting MetricDescriptor", d.Timeout(schema.TimeoutCreate), 20)
	if err != nil {
		return fmt.Errorf("Error waiting to delete MetricDescriptor: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting MetricDescriptor %q: %#v", d.Id(), res)
	return nil
}

func resourceMonitoringMetricDescriptorImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	config := meta.(*Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := parseImportId([]string{"(?P<project>[^ ]+) (?P<name>[^ ]+)", "(?P<name>[^ ]+)"}, d, config); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func flattenMonitoringMetricDescriptorName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMonitoringMetricDescriptorType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMonitoringMetricDescriptorLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(schema.HashResource(monitoringMetricDescriptorLabelsSchema()), []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"key":         flattenMonitoringMetricDescriptorLabelsKey(original["key"], d, config),
			"value_type":  flattenMonitoringMetricDescriptorLabelsValueType(original["valueType"], d, config),
			"description": flattenMonitoringMetricDescriptorLabelsDescription(original["description"], d, config),
		})
	}
	return transformed
}
func flattenMonitoringMetricDescriptorLabelsKey(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMonitoringMetricDescriptorLabelsValueType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil || isEmptyValue(reflect.ValueOf(v)) {
		return "STRING"
	}

	return v
}

func flattenMonitoringMetricDescriptorLabelsDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMonitoringMetricDescriptorMetricKind(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMonitoringMetricDescriptorValueType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMonitoringMetricDescriptorUnit(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMonitoringMetricDescriptorDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMonitoringMetricDescriptorDisplayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMonitoringMetricDescriptorMonitoredResourceTypes(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandMonitoringMetricDescriptorType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringMetricDescriptorLabels(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedKey, err := expandMonitoringMetricDescriptorLabelsKey(original["key"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedKey); val.IsValid() && !isEmptyValue(val) {
			transformed["key"] = transformedKey
		}

		transformedValueType, err := expandMonitoringMetricDescriptorLabelsValueType(original["value_type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValueType); val.IsValid() && !isEmptyValue(val) {
			transformed["valueType"] = transformedValueType
		}

		transformedDescription, err := expandMonitoringMetricDescriptorLabelsDescription(original["description"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !isEmptyValue(val) {
			transformed["description"] = transformedDescription
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandMonitoringMetricDescriptorLabelsKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringMetricDescriptorLabelsValueType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringMetricDescriptorLabelsDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringMetricDescriptorMetricKind(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringMetricDescriptorValueType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringMetricDescriptorUnit(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringMetricDescriptorDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringMetricDescriptorDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringMetricDescriptorMetadata(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSamplePeriod, err := expandMonitoringMetricDescriptorMetadataSamplePeriod(original["sample_period"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSamplePeriod); val.IsValid() && !isEmptyValue(val) {
		transformed["samplePeriod"] = transformedSamplePeriod
	}

	transformedIngestDelay, err := expandMonitoringMetricDescriptorMetadataIngestDelay(original["ingest_delay"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIngestDelay); val.IsValid() && !isEmptyValue(val) {
		transformed["ingestDelay"] = transformedIngestDelay
	}

	return transformed, nil
}

func expandMonitoringMetricDescriptorMetadataSamplePeriod(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringMetricDescriptorMetadataIngestDelay(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringMetricDescriptorLaunchStage(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
