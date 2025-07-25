---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/securitycenterv2/ProjectMuteConfig.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Security Command Center (SCC) v2 API"
description: |-
  Mute Findings is a volume management feature in Security Command Center
  that lets you manually or programmatically hide irrelevant findings,
  and create filters to automatically silence existing and future
  findings based on criteria you specify.
---

# google_scc_v2_project_mute_config

Mute Findings is a volume management feature in Security Command Center
that lets you manually or programmatically hide irrelevant findings,
and create filters to automatically silence existing and future
findings based on criteria you specify.


To get more information about ProjectMuteConfig, see:

* [API documentation](https://cloud.google.com/security-command-center/docs/reference/rest/v2/projects.muteConfigs)

## Example Usage - Scc V2 Project Mute Config Basic


```hcl
resource "google_scc_v2_project_mute_config" "default" {
  mute_config_id    = "my-config"
  project = "my-project-name"
  location     = "global"
  description  = "My custom Cloud Security Command Center Finding Project mute Configuration"
  filter = "severity = \"HIGH\""
  type = "STATIC"
}
```

## Argument Reference

The following arguments are supported:


* `filter` -
  (Required)
  An expression that defines the filter to apply across create/update
  events of findings. While creating a filter string, be mindful of
  the scope in which the mute configuration is being created. E.g.,
  If a filter contains project = X but is created under the
  project = Y scope, it might not match any findings.

* `type` -
  (Required)
  The type of the mute config.

* `mute_config_id` -
  (Required)
  Unique identifier provided by the client within the parent scope.


* `description` -
  (Optional)
  A description of the mute config.

* `location` -
  (Optional)
  location Id is provided by project. If not provided, Use global as default.

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.



## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{location}}/muteConfigs/{{mute_config_id}}`

* `name` -
  Name of the mute config. Its format is
  projects/{project}/locations/global/muteConfigs/{configId},
  folders/{folder}/locations/global/muteConfigs/{configId},
  or organizations/{organization}/locations/global/muteConfigs/{configId}

* `create_time` -
  The time at which the mute config was created. This field is set by
  the server and will be ignored if provided on config creation.

* `update_time` -
  Output only. The most recent time at which the mute config was
  updated. This field is set by the server and will be ignored if
  provided on config creation or update.

* `most_recent_editor` -
  Email address of the user who last edited the mute config. This
  field is set by the server and will be ignored if provided on
  config creation or update.


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


ProjectMuteConfig can be imported using any of these accepted formats:

* `projects/{{project}}/locations/{{location}}/muteConfigs/{{mute_config_id}}`
* `{{project}}/{{location}}/{{mute_config_id}}`
* `{{location}}/{{mute_config_id}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import ProjectMuteConfig using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/locations/{{location}}/muteConfigs/{{mute_config_id}}"
  to = google_scc_v2_project_mute_config.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), ProjectMuteConfig can be imported using one of the formats above. For example:

```
$ terraform import google_scc_v2_project_mute_config.default projects/{{project}}/locations/{{location}}/muteConfigs/{{mute_config_id}}
$ terraform import google_scc_v2_project_mute_config.default {{project}}/{{location}}/{{mute_config_id}}
$ terraform import google_scc_v2_project_mute_config.default {{location}}/{{mute_config_id}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
