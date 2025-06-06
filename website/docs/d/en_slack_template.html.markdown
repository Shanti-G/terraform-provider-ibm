---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_slack_template'
description: |-
  Get information about a Slack Template
---

# ibm_en_slack_template

Provides a read-only data source for Slack template. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example usage

```terraform
data "ibm_en_slack_template" "slack_template" {
  instance_guid  = ibm_resource_instance.en_terraform_test_resource.guid
  template_id = ibm_en_slack_template.en_slack_template.template_id
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `template_id` - (Required, String) Unique identifier for Template.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the `slack_template`.

- `name` - (String) The Template name.

- `description` - (String) The template description.

- `subscription_count` - (Integer) Number of subscriptions.

- `subscription_names` - (List) List of subscriptions.

- `type` - (String) Template type slack.notification.

- `updated_at` - (String) Last updated time.
