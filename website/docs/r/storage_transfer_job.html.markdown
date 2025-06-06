---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: Handwritten     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/website/docs/r/storage_transfer_job.html.markdown
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Storage Transfer Service"
description: |-
  Creates a new Transfer Job in Google Cloud Storage Transfer.
---

# google_storage_transfer_job

Creates a new Transfer Job in Google Cloud Storage Transfer.

To get more information about Google Cloud Storage Transfer, see:

* [Overview](https://cloud.google.com/storage-transfer/docs/overview)
* [API documentation](https://cloud.google.com/storage-transfer/docs/reference/rest/v1/transferJobs)
* How-to Guides
    * [Configuring Access to Data Sources and Sinks](https://cloud.google.com/storage-transfer/docs/configure-access)

## Example Usage

Example creating a nightly Transfer Job from an AWS S3 Bucket to a GCS bucket.

```hcl
data "google_storage_transfer_project_service_account" "default" {
  project = var.project
}

resource "google_storage_bucket" "s3-backup-bucket" {
  name          = "${var.aws_s3_bucket}-backup"
  storage_class = "NEARLINE"
  project       = var.project
  location      = "US"
}

resource "google_storage_bucket_iam_member" "s3-backup-bucket" {
  bucket     = google_storage_bucket.s3-backup-bucket.name
  role       = "roles/storage.admin"
  member     = "serviceAccount:${data.google_storage_transfer_project_service_account.default.email}"
  depends_on = [google_storage_bucket.s3-backup-bucket]
}

resource "google_pubsub_topic" "topic" {
  name = "${var.pubsub_topic_name}"
}

resource "google_pubsub_topic_iam_member" "notification_config" {
  topic = google_pubsub_topic.topic.id
  role = "roles/pubsub.publisher"
  member = "serviceAccount:${data.google_storage_transfer_project_service_account.default.email}"
}

resource "google_storage_transfer_job" "s3-bucket-nightly-backup" {
  description = "Nightly backup of S3 bucket"
  project     = var.project

  transfer_spec {
    object_conditions {
      max_time_elapsed_since_last_modification = "600s"
      exclude_prefixes = [
        "requests.gz",
      ]
    }
    transfer_options {
      delete_objects_unique_in_sink = false
    }
    aws_s3_data_source {
      bucket_name = var.aws_s3_bucket
      aws_access_key {
        access_key_id     = var.aws_access_key
        secret_access_key = var.aws_secret_key
      }
    }
    gcs_data_sink {
      bucket_name = google_storage_bucket.s3-backup-bucket.name
      path        = "foo/bar/"
    }
  }

  schedule {
    schedule_start_date {
      year  = 2018
      month = 10
      day   = 1
    }
    schedule_end_date {
      year  = 2019
      month = 1
      day   = 15
    }
    start_time_of_day {
      hours   = 23
      minutes = 30
      seconds = 0
      nanos   = 0
    }
    repeat_interval = "604800s"
  }

  notification_config {
    pubsub_topic  = google_pubsub_topic.topic.id
    event_types   = [
      "TRANSFER_OPERATION_SUCCESS",
      "TRANSFER_OPERATION_FAILED"
    ]
    payload_format = "JSON"
  }

   logging_config {
    log_actions       = [
      "COPY",
      "DELETE"
    ]
    log_action_states = [
      "SUCCEEDED",
      "FAILED"
    ]
  }

  depends_on = [google_storage_bucket_iam_member.s3-backup-bucket, google_pubsub_topic_iam_member.notification_config]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) The name of the Transfer Job. This name must start with "transferJobs/" prefix and end with a letter or a number, and should be no more than 128 characters ( `transferJobs/^(?!OPI)[A-Za-z0-9-._~]*[A-Za-z0-9]$` ). For transfers involving PosixFilesystem, this name must start with transferJobs/OPI specifically ( `transferJobs/OPI^[A-Za-z0-9-._~]*[A-Za-z0-9]$` ). For all other transfer types, this name must not start with transferJobs/OPI. Default the provider will assign a random unique name with `transferJobs/{{name}}` format, where `name` is a numeric value.

* `description` - (Required) Unique description to identify the Transfer Job.

* `transfer_spec` - (Optional) Transfer specification. Structure [documented below](#nested_transfer_spec). One of `transfer_spec`, or `replication_spec` can be specified.

* `replication_spec` - (Optional) Replication specification. Structure [documented below](#nested_replication_spec). User should not configure `schedule`, `event_stream` with this argument. One of `transfer_spec`, or `replication_spec` must be specified.

- - -

* `schedule` - (Optional) Schedule specification defining when the Transfer Job should be scheduled to start, end and what time to run. Structure [documented below](#nested_schedule). Either `schedule` or `event_stream` must be set.

* `event_stream` - (Optional) Specifies the Event-driven transfer options. Event-driven transfers listen to an event stream to transfer updated files. Structure [documented below](#nested_event_stream) Either `event_stream` or `schedule` must be set.

* `project` - (Optional) The project in which the resource belongs. If it
	is not provided, the provider project is used.

* `status` - (Optional) Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**

* `notification_config` - (Optional) Notification configuration. This is not supported for transfers involving PosixFilesystem. Structure [documented below](#nested_notification_config).

* `logging_config` - (Optional) Logging configuration. Structure [documented below](#nested_logging_config).

<a name="nested_transfer_spec"></a>The `transfer_spec` block supports:

* `source_agent_pool_name` - (Optional) Specifies the agent pool name associated with the posix data source. When unspecified, the default name is used.

* `sink_agent_pool_name` - (Optional) Specifies the agent pool name associated with the posix data sink. When unspecified, the default name is used.

* `gcs_data_sink` - (Optional) A Google Cloud Storage data sink. Structure [documented below](#nested_gcs_data_sink).

* `posix_data_sink` - (Optional) A POSIX data sink. Structure [documented below](#nested_posix_data_sink).

* `object_conditions` - (Optional) Only objects that satisfy these object conditions are included in the set of data source and data sink objects. Object conditions based on objects' `last_modification_time` do not exclude objects in a data sink. Structure [documented below](#nested_object_conditions).

* `transfer_options` - (Optional) Characteristics of how to treat files from datasource and sink during job. If the option `delete_objects_unique_in_sink` is true, object conditions based on objects' `last_modification_time` are ignored and do not exclude objects in a data source or a data sink. Structure [documented below](#nested_transfer_options).

* `gcs_data_source` - (Optional) A Google Cloud Storage data source. Structure [documented below](#nested_gcs_data_source).

* `posix_data_source` - (Optional) A POSIX filesystem data source. Structure [documented below](#nested_posix_data_source).

* `aws_s3_data_source` - (Optional) An AWS S3 data source. Structure [documented below](#nested_aws_s3_data_source).

* `http_data_source` - (Optional) A HTTP URL data source. Structure [documented below](#nested_http_data_source).

* `azure_blob_storage_data_source` - (Optional) An Azure Blob Storage data source. Structure [documented below](#nested_azure_blob_storage_data_source).

* `hdfs_data_source` - (Optional) An HDFS data source. Structure [documented below](#nested_hdfs_data_source).

<a name="nested_replication_spec"></a>The `replication_spec` block supports:

* `gcs_data_sink` - (Optional) A Google Cloud Storage data sink. Structure [documented below](#nested_gcs_data_sink).

* `gcs_data_source` - (Optional) A Google Cloud Storage data source. Structure [documented below](#nested_gcs_data_source).

* `object_conditions` - (Optional) Only objects that satisfy these object conditions are included in the set of data source and data sink objects. Object conditions based on objects' `last_modification_time` do not exclude objects in a data sink. Structure [documented below](#nested_object_conditions).

* `transfer_options` - (Optional) Characteristics of how to treat files from datasource and sink during job. If the option `delete_objects_unique_in_sink` is true, object conditions based on objects' `last_modification_time` are ignored and do not exclude objects in a data source or a data sink. Structure [documented below](#nested_transfer_options).

<a name="nested_schedule"></a>The `schedule` block supports:

* `schedule_start_date` - (Required) The first day the recurring transfer is scheduled to run. If `schedule_start_date` is in the past, the transfer will run for the first time on the following day. Structure [documented below](#nested_schedule_start_end_date).

* `schedule_end_date` - (Optional) The last day the recurring transfer will be run. If `schedule_end_date` is the same as `schedule_start_date`, the transfer will be executed only once. Structure [documented below](#nested_schedule_start_end_date).

* `start_time_of_day` - (Optional) The time in UTC at which the transfer will be scheduled to start in a day. Transfers may start later than this time. If not specified, recurring and one-time transfers that are scheduled to run today will run immediately; recurring transfers that are scheduled to run on a future date will start at approximately midnight UTC on that date. Note that when configuring a transfer with the Cloud Platform Console, the transfer's start time in a day is specified in your local timezone. Structure [documented below](#nested_start_time_of_day).

* `repeat_interval` - (Optional) Interval between the start of each scheduled transfer. If unspecified, the default value is 24 hours. This value may not be less than 1 hour. A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".

<a name="nested_event_stream"></a>The `event_stream` block supports:

* `name` - (Required) Specifies a unique name of the resource such as AWS SQS ARN in the form 'arn:aws:sqs:region:account_id:queue_name', or Pub/Sub subscription resource name in the form 'projects/{project}/subscriptions/{sub}'.

* `event_stream_start_time` - (Optional) Specifies the date and time that Storage Transfer Service starts listening for events from this stream. If no start time is specified or start time is in the past, Storage Transfer Service starts listening immediately. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".

* `event_stream_expiration_time` - (Optional) Specifies the data and time at which Storage Transfer Service stops listening for events from this stream. After this time, any transfers in progress will complete, but no new transfers are initiated.A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".

<a name="nested_object_conditions"></a>The `object_conditions` block supports:

* `max_time_elapsed_since_last_modification` - (Optional) A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".

* `min_time_elapsed_since_last_modification` - (Optional)
A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".

* `include_prefixes` - (Optional) If `include_prefixes` is specified, objects that satisfy the object conditions must have names that start with one of the `include_prefixes` and that do not start with any of the `exclude_prefixes`. If `include_prefixes` is not specified, all objects except those that have names starting with one of the `exclude_prefixes` must satisfy the object conditions. See [Requirements](https://cloud.google.com/storage-transfer/docs/reference/rest/v1/TransferSpec#ObjectConditions).

* `exclude_prefixes` - (Optional) `exclude_prefixes` must follow the requirements described for `include_prefixes`. See [Requirements](https://cloud.google.com/storage-transfer/docs/reference/rest/v1/TransferSpec#ObjectConditions).

* `last_modified_since` - (Optional) If specified, only objects with a "last modification time" on or after this timestamp and objects that don't have a "last modification time" are transferred. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".

* `last_modified_before` - (Optional) If specified, only objects with a "last modification time" before this timestamp and objects that don't have a "last modification time" are transferred. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".

<a name="nested_transfer_options"></a>The `transfer_options` block supports:

* `overwrite_objects_already_existing_in_sink` - (Optional) Whether overwriting objects that already exist in the sink is allowed.

* `delete_objects_unique_in_sink` - (Optional) Whether objects that exist only in the sink should be deleted. Note that this option and
`delete_objects_from_source_after_transfer` are mutually exclusive.

* `delete_objects_from_source_after_transfer` - (Optional) Whether objects should be deleted from the source after they are transferred to the sink. Note that this option and `delete_objects_unique_in_sink` are mutually exclusive.

* `overwrite_when` - (Optional) When to overwrite objects that already exist in the sink. If not set, overwrite behavior is determined by `overwrite_objects_already_existing_in_sink`. Possible values: ALWAYS, DIFFERENT, NEVER.

<a name="nested_gcs_data_sink"></a>The `gcs_data_sink` block supports:

* `bucket_name` - (Required) Google Cloud Storage bucket name.

* `path` - (Optional) Root path to transfer objects. Must be an empty string or full path name that ends with a '/'. This field is treated as an object prefix. As such, it should generally not begin with a '/'.

<a name="nested_gcs_data_source"></a>The `gcs_data_source` block supports:

* `bucket_name` - (Required) Google Cloud Storage bucket name.

* `path` - (Optional) Root path to transfer objects. Must be an empty string or full path name that ends with a '/'. This field is treated as an object prefix. As such, it should generally not begin with a '/'.

<a name="nested_posix_data_sink"></a>The `posix_data_sink` block supports:

* `root_directory` - (Required) Root directory path to the filesystem.

<a name="nested_posix_data_source"></a>The `posix_data_source` block supports:

* `root_directory` - (Required) Root directory path to the filesystem.

<a name="nested_hdfs_data_source"></a>The `hdfs_data_source` block supports:

* `path` - (Required) Root directory path to the filesystem.

<a name="nested_aws_s3_data_source"></a>The `aws_s3_data_source` block supports:

* `bucket_name` - (Required) S3 Bucket name.

* `path` - (Optional) Root path to transfer objects. Must be an empty string or full path name that ends with a '/'. This field is treated as an object prefix. As such, it should generally not begin with a '/'.

* `aws_access_key` - (Optional) AWS credentials block.

* `role_arn` - (Optional) The Amazon Resource Name (ARN) of the role to support temporary credentials via 'AssumeRoleWithWebIdentity'. For more information about ARNs, see [IAM ARNs](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-arns). When a role ARN is provided, Transfer Service fetches temporary credentials for the session using a 'AssumeRoleWithWebIdentity' call for the provided role using the [GoogleServiceAccount][] for this project.

* `managed_private_network` - (Optional) Egress bytes over a Google-managed private network. This network is shared between other users of Storage Transfer Service.

The `aws_access_key` block supports:

* `access_key_id` - (Required) AWS Key ID.

* `secret_access_key` - (Required) AWS Secret Access Key.

<a name="nested_http_data_source"></a>The `http_data_source` block supports:

* `list_url` - (Required) The URL that points to the file that stores the object list entries. This file must allow public access. Currently, only URLs with HTTP and HTTPS schemes are supported.

<a name="nested_azure_blob_storage_data_source"></a>The `azure_blob_storage_data_source` block supports:

* `storage_account` - (Required) The name of the Azure Storage account.

* `container` - (Required) The container to transfer from the Azure Storage account.`

* `path` - (Required) Root path to transfer objects. Must be an empty string or full path name that ends with a '/'. This field is treated as an object prefix. As such, it should generally not begin with a '/'.

* `credentials_secret` - (Optional, [Beta](https://terraform.io/docs/providers/google/guides/provider_versions.html)) Full Resource name of a secret in Secret Manager containing [SAS Credentials in JSON form](https://cloud.google.com/storage-transfer/docs/reference/rest/v1/TransferSpec#azureblobstoragedata:~:text=begin%20with%20a%20%27/%27.-,credentialsSecret,-string). Service Agent for Storage Transfer must have permissions to access secret. If credentials_secret is specified, do not specify azure_credentials.`,

* `azure_credentials` - (Required in GA, Optional in [Beta](https://terraform.io/docs/providers/google/guides/provider_versions.html)) Credentials used to authenticate API requests to Azure block.

The `azure_credentials` block supports:

* `sas_token` - (Required) Azure shared access signature. See [Grant limited access to Azure Storage resources using shared access signatures (SAS)](https://docs.microsoft.com/en-us/azure/storage/common/storage-sas-overview).

<a name="nested_schedule_start_end_date"></a>The `schedule_start_date` and `schedule_end_date` blocks support:

* `year` - (Required) Year of date. Must be from 1 to 9999.

* `month` - (Required) Month of year. Must be from 1 to 12.

* `day` - (Required) Day of month. Must be from 1 to 31 and valid for the year and month.

<a name="nested_start_time_of_day"></a>The `start_time_of_day` blocks support:

* `hours` - (Required) Hours of day in 24 hour format. Should be from 0 to 23

* `minutes` - (Required) Minutes of hour of day. Must be from 0 to 59.

* `seconds` - (Required) Seconds of minutes of the time. Must normally be from 0 to 59.

* `nanos` - (Required) Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.

<a name="nested_notification_config"></a>The `notification_config` block supports:

* `pubsub_topic` - (Required) The Topic.name of the Pub/Sub topic to which to publish notifications. Must be of the format: projects/{project}/topics/{topic}. Not matching this format results in an INVALID_ARGUMENT error.

* `event_types` - (Optional) Event types for which a notification is desired. If empty, send notifications for all event types. The valid types are "TRANSFER_OPERATION_SUCCESS", "TRANSFER_OPERATION_FAILED", "TRANSFER_OPERATION_ABORTED".

* `payload_format` - (Required) The desired format of the notification message payloads. One of "NONE" or "JSON".

<a name="nested_logging_config"></a>The `loggin_config` block supports:

* `log_actions` - (Optional) A list of actions to be logged. If empty, no logs are generated. Not supported for transfers with PosixFilesystem data sources; use enableOnpremGcsTransferLogs instead. 
Each action may be one of `FIND`, `DELETE`, and `COPY`.

* `log_action_states` - (Optional) A list of loggable action states. If empty, no logs are generated. Not supported for transfers with PosixFilesystem data sources; use enableOnpremGcsTransferLogs instead.
Each action state may be one of `SUCCEEDED`, and `FAILED`.

* `enable_on_prem_gcs_transfer` - (Optional) For transfers with a PosixFilesystem source, this option enables the Cloud Storage transfer logs for this transfer. 
Defaults to false.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are
exported:

* `name` - The name of the Transfer Job.

* `creation_time` - When the Transfer Job was created.

* `last_modification_time` - When the Transfer Job was last modified.

* `deletion_time` - When the Transfer Job was deleted.

## Import

Storage Transfer Jobs can be imported using the Transfer Job's `project` and `name` (without the `transferJob/` prefix), e.g.

* `{{project_id}}/{{name}}`, where `name` is a numeric value.

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Storage Transfer Jobs using one of the formats above. For example:

```tf
import {
  id = "{{project_id}}/{{name}}"
  to = google_storage_transfer_job.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), Storage Transfer Jobs can be imported using one of the formats above. For example:

```
$ terraform import google_storage_transfer_job.default {{project_id}}/123456789
```
