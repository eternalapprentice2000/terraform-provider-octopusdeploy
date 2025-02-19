---
page_title: "octopusdeploy_aws_account Resource - terraform-provider-octopusdeploy"
subcategory: "Accounts"
description: |-
  This resource manages AWS accounts in Octopus Deploy.
---

# octopusdeploy_aws_account: Resource

This resource manages AWS accounts in Octopus Deploy.

## Example Usage

```terraform
resource "octopusdeploy_aws_account" "example" {
  access_key   = "access-key"
  name         = "AWS Account (OK to Delete)"
  secret_key   = "###########" # required; get from secure environment/store
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `access_key` (String) The access key associated with this AWS account.
- `name` (String) The name of this AWS account.
- `secret_key` (String, Sensitive) The secret key associated with this resource.

### Optional

- `description` (String) A user-friendly description of this AWS account.
- `environments` (List of String) A list of environment IDs associated with this resource.
- `space_id` (String) The space ID associated with this resource.
- `tenant_tags` (List of String) A list of tenant tags associated with this resource.
- `tenanted_deployment_participation` (String) The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
- `tenants` (List of String) A list of tenant IDs associated with this resource.

### Read-Only

- `id` (String) The ID of this resource.


