---
page_title: "octopusdeploy_azure_subscription_account Resource - terraform-provider-octopusdeploy"
subcategory: "Accounts"
description: |-
  This resource manages Azure subscription accounts in Octopus Deploy.
---

# octopusdeploy_azure_subscription_account: Resource

This resource manages Azure subscription accounts in Octopus Deploy.

## Example Usage

```terraform
resource "octopusdeploy_azure_subscription_account" "example" {
  name            = "Azure Subscription Account (OK to Delete)"
  subscription_id = "00000000-0000-0000-0000-000000000000"
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `management_endpoint` (String)
- `name` (String) The name of this resource.
- `storage_endpoint_suffix` (String) The storage endpoint suffix associated with this Azure subscription account.
- `subscription_id` (String) The subscription ID of this resource.

### Optional

- `azure_environment` (String) The Azure environment associated with this resource. Valid Azure environments are `AzureCloud`, `AzureChinaCloud`, `AzureGermanCloud`, or `AzureUSGovernment`.
- `certificate` (String, Sensitive)
- `certificate_thumbprint` (String, Sensitive)
- `description` (String) The description of this Azure subscription account.
- `environments` (List of String) A list of environment IDs associated with this resource.
- `space_id` (String) The space ID associated with this resource.
- `tenant_tags` (List of String) A list of tenant tags associated with this resource.
- `tenanted_deployment_participation` (String) The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
- `tenants` (List of String) A list of tenant IDs associated with this resource.

### Read-Only

- `id` (String) The ID of this resource.


