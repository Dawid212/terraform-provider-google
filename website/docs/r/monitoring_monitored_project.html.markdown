---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/monitoring/MonitoredProject.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Cloud (Stackdriver) Monitoring"
description: |-
  A [project being monitored](https://cloud.
---

# google_monitoring_monitored_project

A [project being monitored](https://cloud.google.com/monitoring/settings/multiple-projects#create-multi) by a Metrics Scope.


To get more information about MonitoredProject, see:

* [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v1/locations.global.metricsScopes.projects)
* How-to Guides
    * [Official Documentation](https://cloud.google.com/monitoring/settings/manage-api)

## Example Usage - Monitoring Monitored Project Basic


```hcl
resource "google_monitoring_monitored_project" "primary" {
  metrics_scope = "my-project-name"
  name          = google_project.basic.project_id
}

resource "google_project" "basic" {
  project_id = "m-id"
  name       = "m-id-display"
  org_id     = "123456789"
  deletion_policy = "DELETE"
}
```

## Argument Reference

The following arguments are supported:


* `name` -
  (Required)
  Immutable. The resource name of the `MonitoredProject`. On input, the resource name includes the scoping project ID and monitored project ID. On output, it contains the equivalent project numbers. Example: `locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}/projects/{MONITORED_PROJECT_ID_OR_NUMBER}`

* `metrics_scope` -
  (Required)
  Required. The resource name of the existing Metrics Scope that will monitor this project. Example: locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}




## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `locations/global/metricsScopes/{{metrics_scope}}/projects/{{name}}`

* `create_time` -
  Output only. The time when this `MonitoredProject` was created.


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


MonitoredProject can be imported using any of these accepted formats:

* `v1/locations/global/metricsScopes/{{name}}`
* `{{name}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import MonitoredProject using one of the formats above. For example:

```tf
import {
  id = "v1/locations/global/metricsScopes/{{name}}"
  to = google_monitoring_monitored_project.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), MonitoredProject can be imported using one of the formats above. For example:

```
$ terraform import google_monitoring_monitored_project.default v1/locations/global/metricsScopes/{{name}}
$ terraform import google_monitoring_monitored_project.default {{name}}
```
