---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/datafusion/Instance.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/datasource_iam.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Cloud Data Fusion"
description: |-
  A datasource to retrieve the IAM policy state for Cloud Data Fusion Instance
---


# google_data_fusion_instance_iam_policy

Retrieves the current IAM policy data for instance


## Example Usage


```hcl
data "google_data_fusion_instance_iam_policy" "policy" {
  project = google_data_fusion_instance.basic_instance.project
  region = google_data_fusion_instance.basic_instance.region
  name = google_data_fusion_instance.basic_instance.name
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional) The region of the Data Fusion instance.
 Used to find the parent resource to bind the IAM policy to. If not specified,
  the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
  region is specified, it is taken from the provider configuration.
* `name` - (Required) Used to find the parent resource to bind the IAM policy to

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.

## Attributes Reference

The attributes are exported:

* `etag` - (Computed) The etag of the IAM policy.

* `policy_data` - (Required only by `google_data_fusion_instance_iam_policy`) The policy data generated by
  a `google_iam_policy` data source.
