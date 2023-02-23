# Table: boundary_role

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| principal_ids | json | X | √ | A list of principal (user or group) IDs to add as principals on the role. | 
| scope_id | string | X | √ | The scope ID in which the resource is created. Defaults to the provider's `default_scope` if unset. | 
| description | string | X | √ | The role description. | 
| grant_scope_id | string | X | √ |  | 
| grant_strings | json | X | √ |  A list of stringified grants for the role. | 
| id | string | X | √ | The ID of the role. | 
| name | string | X | √ | The role name. Defaults to the resource name. | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


