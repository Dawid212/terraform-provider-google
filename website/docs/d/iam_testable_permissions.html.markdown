---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: Handwritten     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/website/docs/d/iam_testable_permissions.html.markdown
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Cloud Platform"
description: |-
  Retrieve a list of testable permissions for a resource. Testable permissions mean the permissions that user can add or remove in a role at a given resource. The resource can be referenced either via the full resource name or via a URI.
---

# google_iam_testable_permissions

Retrieve a list of testable permissions for a resource. Testable permissions mean the permissions that user can add or remove in a role at a given resource. The resource can be referenced either via the full resource name or via a URI.

## Example Usage

Retrieve all the supported permissions able to be set on `my-project` that are in either GA or BETA. This is useful for dynamically constructing custom roles.

```hcl
data "google_iam_testable_permissions" "perms" {
	full_resource_name = "//cloudresourcemanager.googleapis.com/projects/my-project"
	stages             = ["GA", "BETA"]
}
```

## Argument Reference

The following arguments are supported:

* `full_resource_name` - (Required) See [full resource name documentation](https://cloud.google.com/apis/design/resource_names#full_resource_name) for more detail.
* `stages` - (Optional) The acceptable release stages of the permission in the output. Note that `BETA` does not include permissions in `GA`, but you can specify both with `["GA", "BETA"]` for example. Can be a list of `"ALPHA"`, `"BETA"`, `"GA"`, `"DEPRECATED"`. Default is `["GA"]`.
* `custom_support_level` - (Optional) The level of support for custom roles. Can be one of `"NOT_SUPPORTED"`, `"SUPPORTED"`, `"TESTING"`. Default is `"SUPPORTED"`

## Attributes Reference

The following attributes are exported:

* `permissions` - A list of permissions matching the provided input. Structure is [defined below](#nested_permissions).

<a name="nested_permissions"></a>The `permissions` block supports:

* `name` - Name of the permission.
* `title` - Human readable title of the permission.
* `stage` - Release stage of the permission.
* `custom_support_level` - The the support level of this permission for custom roles.
* `api_disabled` - Whether the corresponding API has been enabled for the resource.

