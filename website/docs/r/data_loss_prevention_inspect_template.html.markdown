---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/dlp/InspectTemplate.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Data Loss Prevention"
description: |-
  An inspect job template.
---

# google_data_loss_prevention_inspect_template

An inspect job template.


To get more information about InspectTemplate, see:

* [API documentation](https://cloud.google.com/dlp/docs/reference/rest/v2/projects.inspectTemplates)
* How-to Guides
    * [Official Documentation](https://cloud.google.com/dlp/docs/creating-templates-inspect)

## Example Usage - Dlp Inspect Template Basic


```hcl
resource "google_data_loss_prevention_inspect_template" "basic" {
	parent = "projects/my-project-name"
	description = "My description"
	display_name = "display_name"

	inspect_config {
		info_types {
			name = "EMAIL_ADDRESS"
		}
		info_types {
			name = "PERSON_NAME"
		}
		info_types {
			name = "LAST_NAME"
		}
		info_types {
			name = "DOMAIN_NAME"
		}
		info_types {
			name = "PHONE_NUMBER"
		}
		info_types {
			name = "FIRST_NAME"
		}

		min_likelihood = "UNLIKELY"
		rule_set {
			info_types {
				name = "EMAIL_ADDRESS"
			}
			rules {
				exclusion_rule {
					regex {
						pattern = ".+@example.com"
					}
					matching_type = "MATCHING_TYPE_FULL_MATCH"
				}
			}
		}
		rule_set {
			info_types {
				name = "EMAIL_ADDRESS"
			}
			info_types {
				name = "DOMAIN_NAME"
			}
			info_types {
				name = "PHONE_NUMBER"
			}
			info_types {
				name = "PERSON_NAME"
			}
			info_types {
				name = "FIRST_NAME"
			}
			rules {
				exclusion_rule {
					dictionary {
						word_list {
							words = ["TEST"]
						}
					}
					matching_type = "MATCHING_TYPE_PARTIAL_MATCH"
				}
			}
		}

		rule_set {
			info_types {
				name = "PERSON_NAME"
			}
			rules {
				hotword_rule {
					hotword_regex {
						pattern = "patient"
					}
					proximity {
						window_before = 50
					}
					likelihood_adjustment {
						fixed_likelihood = "VERY_LIKELY"
					}
				}
			}
		}

		limits {
			max_findings_per_item    = 10
			max_findings_per_request = 50
			max_findings_per_info_type {
				max_findings = "75"
				info_type {
					name = "PERSON_NAME"
				}
			}
			max_findings_per_info_type {
				max_findings = "80"
				info_type {
					name = "LAST_NAME"
				}
			}
		}
	}
}
```
## Example Usage - Dlp Inspect Template Custom Type


```hcl
resource "google_data_loss_prevention_inspect_template" "custom" {
	parent = "projects/my-project-name"
	description = "My description"
	display_name = "display_name"

	inspect_config {
		custom_info_types {
			info_type {
				name = "MY_CUSTOM_TYPE"
			}

			likelihood = "UNLIKELY"

			regex {
				pattern = "test*"
			}
		}

		info_types {
			name = "EMAIL_ADDRESS"
		}

		min_likelihood = "UNLIKELY"
		rule_set {
			info_types {
				name = "EMAIL_ADDRESS"
			}
			rules {
				exclusion_rule {
					regex {
						pattern = ".+@example.com"
					}
					matching_type = "MATCHING_TYPE_FULL_MATCH"
				}
			}
		}

		rule_set {
			info_types {
				name = "MY_CUSTOM_TYPE"
			}
			rules {
				hotword_rule {
					hotword_regex {
						pattern = "example*"
					}
					proximity {
						window_before = 50
					}
					likelihood_adjustment {
						fixed_likelihood = "VERY_LIKELY"
					}
				}
			}
		}

		limits {
			max_findings_per_item    = 10
			max_findings_per_request = 50
		}
	}
}
```
## Example Usage - Dlp Inspect Template Custom Type Surrogate


```hcl
resource "google_data_loss_prevention_inspect_template" "custom_type_surrogate" {
  parent = "projects/my-project-name"
  description = "My description"
  display_name = "display_name"

  inspect_config {
    custom_info_types {
      info_type {
        name = "MY_CUSTOM_TYPE"
      }

      likelihood = "UNLIKELY"

      surrogate_type {}
    }

    info_types {
      name = "EMAIL_ADDRESS"
    }

    min_likelihood = "UNLIKELY"
    rule_set {
      info_types {
        name = "EMAIL_ADDRESS"
      }
      rules {
        exclusion_rule {
          regex {
            pattern = ".+@example.com"
          }
          matching_type = "MATCHING_TYPE_FULL_MATCH"
        }
      }
    }

    rule_set {
      info_types {
        name = "MY_CUSTOM_TYPE"
      }
      rules {
        hotword_rule {
          hotword_regex {
            pattern = "example*"
          }
          proximity {
            window_before = 50
          }
          likelihood_adjustment {
            fixed_likelihood = "VERY_LIKELY"
          }
        }
      }
    }

    limits {
      max_findings_per_item    = 10
      max_findings_per_request = 50
    }
  }
}
```
## Example Usage - Dlp Inspect Template Max Infotype Per Finding Default


```hcl
resource "google_data_loss_prevention_inspect_template" "max_infotype_per_finding_default" {
  parent = "projects/my-project-name"

  inspect_config {
    info_types {
      name = "EMAIL_ADDRESS"
    }
    info_types {
      name = "PERSON_NAME"
    }

    min_likelihood = "UNLIKELY"
    limits {
        max_findings_per_request = 333
        max_findings_per_item = 222
        max_findings_per_info_type {
          # Entry with no info_type specifies the default value used by all info_types that don't specify their own limit
          max_findings = 111
        }
    }
  }
}
```

## Argument Reference

The following arguments are supported:


* `parent` -
  (Required)
  The parent of the inspect template in any of the following formats:
  * `projects/{{project}}`
  * `projects/{{project}}/locations/{{location}}`
  * `organizations/{{organization_id}}`
  * `organizations/{{organization_id}}/locations/{{location}}`


* `description` -
  (Optional)
  A description of the inspect template.

* `display_name` -
  (Optional)
  User set display name of the inspect template.

* `template_id` -
  (Optional)
  The template id can contain uppercase and lowercase letters, numbers, and hyphens;
  that is, it must match the regular expression: [a-zA-Z\d-_]+. The maximum length is
  100 characters. Can be empty to allow the system to generate one.

* `inspect_config` -
  (Optional)
  The core content of the template.
  Structure is [documented below](#nested_inspect_config).



<a name="nested_inspect_config"></a>The `inspect_config` block supports:

* `exclude_info_types` -
  (Optional)
  When true, excludes type information of the findings.

* `include_quote` -
  (Optional)
  When true, a contextual quote from the data that triggered a finding is included in the response.

* `min_likelihood` -
  (Optional)
  Only returns findings equal or above this threshold. See https://cloud.google.com/dlp/docs/likelihood for more info
  Default value is `POSSIBLE`.
  Possible values are: `VERY_UNLIKELY`, `UNLIKELY`, `POSSIBLE`, `LIKELY`, `VERY_LIKELY`.

* `limits` -
  (Optional)
  Configuration to control the number of findings returned.
  Structure is [documented below](#nested_inspect_config_limits).

* `info_types` -
  (Optional)
  Restricts what infoTypes to look for. The values must correspond to InfoType values returned by infoTypes.list
  or listed at https://cloud.google.com/dlp/docs/infotypes-reference.
  When no InfoTypes or CustomInfoTypes are specified in a request, the system may automatically choose what detectors to run.
  By default this may be all types, but may change over time as detectors are updated.
  Structure is [documented below](#nested_inspect_config_info_types).

* `content_options` -
  (Optional)
  List of options defining data content to scan. If empty, text, images, and other content will be included.
  Each value may be one of: `CONTENT_TEXT`, `CONTENT_IMAGE`.

* `rule_set` -
  (Optional)
  Set of rules to apply to the findings for this InspectConfig. Exclusion rules, contained in the set are executed in the end,
  other rules are executed in the order they are specified for each info type.
  Structure is [documented below](#nested_inspect_config_rule_set).

* `custom_info_types` -
  (Optional)
  Custom info types to be used. See https://cloud.google.com/dlp/docs/creating-custom-infotypes to learn more.
  Structure is [documented below](#nested_inspect_config_custom_info_types).


<a name="nested_inspect_config_limits"></a>The `limits` block supports:

* `max_findings_per_item` -
  (Required)
  Max number of findings that will be returned for each item scanned. The maximum returned is 2000.

* `max_findings_per_request` -
  (Required)
  Max number of findings that will be returned per request/job. The maximum returned is 2000.

* `max_findings_per_info_type` -
  (Optional)
  Configuration of findings limit given for specified infoTypes.
  Structure is [documented below](#nested_inspect_config_limits_max_findings_per_info_type).


<a name="nested_inspect_config_limits_max_findings_per_info_type"></a>The `max_findings_per_info_type` block supports:

* `info_type` -
  (Optional)
  Type of information the findings limit applies to. Only one limit per infoType should be provided. If InfoTypeLimit does
  not have an infoType, the DLP API applies the limit against all infoTypes that are found but not
  specified in another InfoTypeLimit.
  Structure is [documented below](#nested_inspect_config_limits_max_findings_per_info_type_max_findings_per_info_type_info_type).

* `max_findings` -
  (Required)
  Max findings limit for the given infoType.


<a name="nested_inspect_config_limits_max_findings_per_info_type_max_findings_per_info_type_info_type"></a>The `info_type` block supports:

* `name` -
  (Required)
  Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
  at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.

* `version` -
  (Optional)
  Version name for this InfoType.

* `sensitivity_score` -
  (Optional)
  Optional custom sensitivity for this InfoType. This only applies to data profiling.
  Structure is [documented below](#nested_inspect_config_limits_max_findings_per_info_type_max_findings_per_info_type_info_type_sensitivity_score).


<a name="nested_inspect_config_limits_max_findings_per_info_type_max_findings_per_info_type_info_type_sensitivity_score"></a>The `sensitivity_score` block supports:

* `score` -
  (Required)
  The sensitivity score applied to the resource.
  Possible values are: `SENSITIVITY_LOW`, `SENSITIVITY_MODERATE`, `SENSITIVITY_HIGH`.

<a name="nested_inspect_config_info_types"></a>The `info_types` block supports:

* `name` -
  (Required)
  Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
  at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.

* `version` -
  (Optional)
  Version of the information type to use. By default, the version is set to stable

* `sensitivity_score` -
  (Optional)
  Optional custom sensitivity for this InfoType. This only applies to data profiling.
  Structure is [documented below](#nested_inspect_config_info_types_info_types_sensitivity_score).


<a name="nested_inspect_config_info_types_info_types_sensitivity_score"></a>The `sensitivity_score` block supports:

* `score` -
  (Required)
  The sensitivity score applied to the resource.
  Possible values are: `SENSITIVITY_LOW`, `SENSITIVITY_MODERATE`, `SENSITIVITY_HIGH`.

<a name="nested_inspect_config_rule_set"></a>The `rule_set` block supports:

* `info_types` -
  (Required)
  List of infoTypes this rule set is applied to.
  Structure is [documented below](#nested_inspect_config_rule_set_rule_set_info_types).

* `rules` -
  (Required)
  Set of rules to be applied to infoTypes. The rules are applied in order.
  Structure is [documented below](#nested_inspect_config_rule_set_rule_set_rules).


<a name="nested_inspect_config_rule_set_rule_set_info_types"></a>The `info_types` block supports:

* `name` -
  (Required)
  Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
  at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.

* `version` -
  (Optional)
  Version name for this InfoType.

* `sensitivity_score` -
  (Optional)
  Optional custom sensitivity for this InfoType. This only applies to data profiling.
  Structure is [documented below](#nested_inspect_config_rule_set_rule_set_info_types_info_types_sensitivity_score).


<a name="nested_inspect_config_rule_set_rule_set_info_types_info_types_sensitivity_score"></a>The `sensitivity_score` block supports:

* `score` -
  (Required)
  The sensitivity score applied to the resource.
  Possible values are: `SENSITIVITY_LOW`, `SENSITIVITY_MODERATE`, `SENSITIVITY_HIGH`.

<a name="nested_inspect_config_rule_set_rule_set_rules"></a>The `rules` block supports:

* `hotword_rule` -
  (Optional)
  Hotword-based detection rule.
  Structure is [documented below](#nested_inspect_config_rule_set_rule_set_rules_rules_hotword_rule).

* `exclusion_rule` -
  (Optional)
  The rule that specifies conditions when findings of infoTypes specified in InspectionRuleSet are removed from results.
  Structure is [documented below](#nested_inspect_config_rule_set_rule_set_rules_rules_exclusion_rule).


<a name="nested_inspect_config_rule_set_rule_set_rules_rules_hotword_rule"></a>The `hotword_rule` block supports:

* `hotword_regex` -
  (Required)
  Regular expression pattern defining what qualifies as a hotword.
  Structure is [documented below](#nested_inspect_config_rule_set_rule_set_rules_rules_hotword_rule_hotword_regex).

* `proximity` -
  (Required)
  Proximity of the finding within which the entire hotword must reside. The total length of the window cannot
  exceed 1000 characters. Note that the finding itself will be included in the window, so that hotwords may be
  used to match substrings of the finding itself. For example, the certainty of a phone number regex
  `(\d{3}) \d{3}-\d{4}` could be adjusted upwards if the area code is known to be the local area code of a company
  office using the hotword regex `(xxx)`, where `xxx` is the area code in question.
  Structure is [documented below](#nested_inspect_config_rule_set_rule_set_rules_rules_hotword_rule_proximity).

* `likelihood_adjustment` -
  (Required)
  Likelihood adjustment to apply to all matching findings.
  Structure is [documented below](#nested_inspect_config_rule_set_rule_set_rules_rules_hotword_rule_likelihood_adjustment).


<a name="nested_inspect_config_rule_set_rule_set_rules_rules_hotword_rule_hotword_regex"></a>The `hotword_regex` block supports:

* `pattern` -
  (Required)
  Pattern defining the regular expression. Its syntax
  (https://github.com/google/re2/wiki/Syntax) can be found under the google/re2 repository on GitHub.

* `group_indexes` -
  (Optional)
  The index of the submatch to extract as findings. When not specified,
  the entire match is returned. No more than 3 may be included.

<a name="nested_inspect_config_rule_set_rule_set_rules_rules_hotword_rule_proximity"></a>The `proximity` block supports:

* `window_before` -
  (Optional)
  Number of characters before the finding to consider. Either this or window_after must be specified

* `window_after` -
  (Optional)
  Number of characters after the finding to consider. Either this or window_before must be specified

<a name="nested_inspect_config_rule_set_rule_set_rules_rules_hotword_rule_likelihood_adjustment"></a>The `likelihood_adjustment` block supports:

* `fixed_likelihood` -
  (Optional)
  Set the likelihood of a finding to a fixed value. Either this or relative_likelihood can be set.
  Possible values are: `VERY_UNLIKELY`, `UNLIKELY`, `POSSIBLE`, `LIKELY`, `VERY_LIKELY`.

* `relative_likelihood` -
  (Optional)
  Increase or decrease the likelihood by the specified number of levels. For example,
  if a finding would be POSSIBLE without the detection rule and relativeLikelihood is 1,
  then it is upgraded to LIKELY, while a value of -1 would downgrade it to UNLIKELY.
  Likelihood may never drop below VERY_UNLIKELY or exceed VERY_LIKELY, so applying an
  adjustment of 1 followed by an adjustment of -1 when base likelihood is VERY_LIKELY
  will result in a final likelihood of LIKELY. Either this or fixed_likelihood can be set.

<a name="nested_inspect_config_rule_set_rule_set_rules_rules_exclusion_rule"></a>The `exclusion_rule` block supports:

* `matching_type` -
  (Required)
  How the rule is applied. See the documentation for more information: https://cloud.google.com/dlp/docs/reference/rest/v2/InspectConfig#MatchingType
  Possible values are: `MATCHING_TYPE_FULL_MATCH`, `MATCHING_TYPE_PARTIAL_MATCH`, `MATCHING_TYPE_INVERSE_MATCH`.

* `dictionary` -
  (Optional)
  Dictionary which defines the rule.
  Structure is [documented below](#nested_inspect_config_rule_set_rule_set_rules_rules_exclusion_rule_dictionary).

* `regex` -
  (Optional)
  Regular expression which defines the rule.
  Structure is [documented below](#nested_inspect_config_rule_set_rule_set_rules_rules_exclusion_rule_regex).

* `exclude_info_types` -
  (Optional)
  Set of infoTypes for which findings would affect this rule.
  Structure is [documented below](#nested_inspect_config_rule_set_rule_set_rules_rules_exclusion_rule_exclude_info_types).

* `exclude_by_hotword` -
  (Optional)
  Drop if the hotword rule is contained in the proximate context.
  For tabular data, the context includes the column name.
  Structure is [documented below](#nested_inspect_config_rule_set_rule_set_rules_rules_exclusion_rule_exclude_by_hotword).


<a name="nested_inspect_config_rule_set_rule_set_rules_rules_exclusion_rule_dictionary"></a>The `dictionary` block supports:

* `word_list` -
  (Optional)
  List of words or phrases to search for.
  Structure is [documented below](#nested_inspect_config_rule_set_rule_set_rules_rules_exclusion_rule_dictionary_word_list).

* `cloud_storage_path` -
  (Optional)
  Newline-delimited file of words in Cloud Storage. Only a single file is accepted.
  Structure is [documented below](#nested_inspect_config_rule_set_rule_set_rules_rules_exclusion_rule_dictionary_cloud_storage_path).


<a name="nested_inspect_config_rule_set_rule_set_rules_rules_exclusion_rule_dictionary_word_list"></a>The `word_list` block supports:

* `words` -
  (Required)
  Words or phrases defining the dictionary. The dictionary must contain at least one
  phrase and every phrase must contain at least 2 characters that are letters or digits.

<a name="nested_inspect_config_rule_set_rule_set_rules_rules_exclusion_rule_dictionary_cloud_storage_path"></a>The `cloud_storage_path` block supports:

* `path` -
  (Required)
  A url representing a file or path (no wildcards) in Cloud Storage. Example: `gs://[BUCKET_NAME]/dictionary.txt`

<a name="nested_inspect_config_rule_set_rule_set_rules_rules_exclusion_rule_regex"></a>The `regex` block supports:

* `pattern` -
  (Required)
  Pattern defining the regular expression.
  Its syntax (https://github.com/google/re2/wiki/Syntax) can be found under the google/re2 repository on GitHub.

* `group_indexes` -
  (Optional)
  The index of the submatch to extract as findings. When not specified, the entire match is returned. No more than 3 may be included.

<a name="nested_inspect_config_rule_set_rule_set_rules_rules_exclusion_rule_exclude_info_types"></a>The `exclude_info_types` block supports:

* `info_types` -
  (Required)
  If a finding is matched by any of the infoType detectors listed here, the finding will be excluded from the scan results.
  Structure is [documented below](#nested_inspect_config_rule_set_rule_set_rules_rules_exclusion_rule_exclude_info_types_info_types).


<a name="nested_inspect_config_rule_set_rule_set_rules_rules_exclusion_rule_exclude_info_types_info_types"></a>The `info_types` block supports:

* `name` -
  (Required)
  Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
  at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.

* `version` -
  (Optional)
  Version name for this InfoType.

* `sensitivity_score` -
  (Optional)
  Optional custom sensitivity for this InfoType. This only applies to data profiling.
  Structure is [documented below](#nested_inspect_config_rule_set_rule_set_rules_rules_exclusion_rule_exclude_info_types_info_types_info_types_sensitivity_score).


<a name="nested_inspect_config_rule_set_rule_set_rules_rules_exclusion_rule_exclude_info_types_info_types_info_types_sensitivity_score"></a>The `sensitivity_score` block supports:

* `score` -
  (Required)
  The sensitivity score applied to the resource.
  Possible values are: `SENSITIVITY_LOW`, `SENSITIVITY_MODERATE`, `SENSITIVITY_HIGH`.

<a name="nested_inspect_config_rule_set_rule_set_rules_rules_exclusion_rule_exclude_by_hotword"></a>The `exclude_by_hotword` block supports:

* `hotword_regex` -
  (Required)
  Regular expression pattern defining what qualifies as a hotword.
  Structure is [documented below](#nested_inspect_config_rule_set_rule_set_rules_rules_exclusion_rule_exclude_by_hotword_hotword_regex).

* `proximity` -
  (Required)
  Proximity of the finding within which the entire hotword must reside. The total length of the window cannot
  exceed 1000 characters. Note that the finding itself will be included in the window, so that hotwords may be
  used to match substrings of the finding itself. For example, the certainty of a phone number regex
  `(\d{3}) \d{3}-\d{4}` could be adjusted upwards if the area code is known to be the local area code of a company
  office using the hotword regex `(xxx)`, where `xxx` is the area code in question.
  Structure is [documented below](#nested_inspect_config_rule_set_rule_set_rules_rules_exclusion_rule_exclude_by_hotword_proximity).


<a name="nested_inspect_config_rule_set_rule_set_rules_rules_exclusion_rule_exclude_by_hotword_hotword_regex"></a>The `hotword_regex` block supports:

* `pattern` -
  (Required)
  Pattern defining the regular expression. Its syntax
  (https://github.com/google/re2/wiki/Syntax) can be found under the google/re2 repository on GitHub.

* `group_indexes` -
  (Optional)
  The index of the submatch to extract as findings. When not specified,
  the entire match is returned. No more than 3 may be included.

<a name="nested_inspect_config_rule_set_rule_set_rules_rules_exclusion_rule_exclude_by_hotword_proximity"></a>The `proximity` block supports:

* `window_before` -
  (Optional)
  Number of characters before the finding to consider.

* `window_after` -
  (Optional)
  Number of characters after the finding to consider.

<a name="nested_inspect_config_custom_info_types"></a>The `custom_info_types` block supports:

* `info_type` -
  (Required)
  CustomInfoType can either be a new infoType, or an extension of built-in infoType, when the name matches one of existing
  infoTypes and that infoType is specified in `info_types` field. Specifying the latter adds findings to the
  one detected by the system. If built-in info type is not specified in `info_types` list then the name is
  treated as a custom info type.
  Structure is [documented below](#nested_inspect_config_custom_info_types_custom_info_types_info_type).

* `likelihood` -
  (Optional)
  Likelihood to return for this CustomInfoType. This base value can be altered by a detection rule if the finding meets the criteria
  specified by the rule.
  Default value is `VERY_LIKELY`.
  Possible values are: `VERY_UNLIKELY`, `UNLIKELY`, `POSSIBLE`, `LIKELY`, `VERY_LIKELY`.

* `exclusion_type` -
  (Optional)
  If set to EXCLUSION_TYPE_EXCLUDE this infoType will not cause a finding to be returned. It still can be used for rules matching.
  Possible values are: `EXCLUSION_TYPE_EXCLUDE`.

* `sensitivity_score` -
  (Optional)
  Optional custom sensitivity for this InfoType. This only applies to data profiling.
  Structure is [documented below](#nested_inspect_config_custom_info_types_custom_info_types_sensitivity_score).

* `regex` -
  (Optional)
  Regular expression which defines the rule.
  Structure is [documented below](#nested_inspect_config_custom_info_types_custom_info_types_regex).

* `dictionary` -
  (Optional)
  Dictionary which defines the rule.
  Structure is [documented below](#nested_inspect_config_custom_info_types_custom_info_types_dictionary).

* `surrogate_type` -
  (Optional)
  Message for detecting output from deidentification transformations that support reversing.

* `stored_type` -
  (Optional)
  A reference to a StoredInfoType to use with scanning.
  Structure is [documented below](#nested_inspect_config_custom_info_types_custom_info_types_stored_type).


<a name="nested_inspect_config_custom_info_types_custom_info_types_info_type"></a>The `info_type` block supports:

* `name` -
  (Required)
  Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names
  listed at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.

* `version` -
  (Optional)
  Version name for this InfoType.

* `sensitivity_score` -
  (Optional)
  Optional custom sensitivity for this InfoType. This only applies to data profiling.
  Structure is [documented below](#nested_inspect_config_custom_info_types_custom_info_types_info_type_sensitivity_score).


<a name="nested_inspect_config_custom_info_types_custom_info_types_info_type_sensitivity_score"></a>The `sensitivity_score` block supports:

* `score` -
  (Required)
  The sensitivity score applied to the resource.
  Possible values are: `SENSITIVITY_LOW`, `SENSITIVITY_MODERATE`, `SENSITIVITY_HIGH`.

<a name="nested_inspect_config_custom_info_types_custom_info_types_sensitivity_score"></a>The `sensitivity_score` block supports:

* `score` -
  (Required)
  The sensitivity score applied to the resource.
  Possible values are: `SENSITIVITY_LOW`, `SENSITIVITY_MODERATE`, `SENSITIVITY_HIGH`.

<a name="nested_inspect_config_custom_info_types_custom_info_types_regex"></a>The `regex` block supports:

* `pattern` -
  (Required)
  Pattern defining the regular expression.
  Its syntax (https://github.com/google/re2/wiki/Syntax) can be found under the google/re2 repository on GitHub.

* `group_indexes` -
  (Optional)
  The index of the submatch to extract as findings. When not specified, the entire match is returned. No more than 3 may be included.

<a name="nested_inspect_config_custom_info_types_custom_info_types_dictionary"></a>The `dictionary` block supports:

* `word_list` -
  (Optional)
  List of words or phrases to search for.
  Structure is [documented below](#nested_inspect_config_custom_info_types_custom_info_types_dictionary_word_list).

* `cloud_storage_path` -
  (Optional)
  Newline-delimited file of words in Cloud Storage. Only a single file is accepted.
  Structure is [documented below](#nested_inspect_config_custom_info_types_custom_info_types_dictionary_cloud_storage_path).


<a name="nested_inspect_config_custom_info_types_custom_info_types_dictionary_word_list"></a>The `word_list` block supports:

* `words` -
  (Required)
  Words or phrases defining the dictionary. The dictionary must contain at least one
  phrase and every phrase must contain at least 2 characters that are letters or digits.

<a name="nested_inspect_config_custom_info_types_custom_info_types_dictionary_cloud_storage_path"></a>The `cloud_storage_path` block supports:

* `path` -
  (Required)
  A url representing a file or path (no wildcards) in Cloud Storage. Example: `gs://[BUCKET_NAME]/dictionary.txt`

<a name="nested_inspect_config_custom_info_types_custom_info_types_stored_type"></a>The `stored_type` block supports:

* `name` -
  (Required)
  Resource name of the requested StoredInfoType, for example `organizations/433245324/storedInfoTypes/432452342`
  or `projects/project-id/storedInfoTypes/432452342`.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `{{parent}}/inspectTemplates/{{name}}`

* `name` -
  The resource name of the inspect template. Set by the server.


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


InspectTemplate can be imported using any of these accepted formats:

* `{{parent}}/inspectTemplates/{{name}}`
* `{{parent}}/{{name}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import InspectTemplate using one of the formats above. For example:

```tf
import {
  id = "{{parent}}/inspectTemplates/{{name}}"
  to = google_data_loss_prevention_inspect_template.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), InspectTemplate can be imported using one of the formats above. For example:

```
$ terraform import google_data_loss_prevention_inspect_template.default {{parent}}/inspectTemplates/{{name}}
$ terraform import google_data_loss_prevention_inspect_template.default {{parent}}/{{name}}
```
