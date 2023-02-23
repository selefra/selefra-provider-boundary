# Table: boundary_user

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_ids | json | X | √ | Account ID's to associate with this user resource. | 
| description | string | X | √ | The user description. | 
| id | string | X | √ | The ID of the user. | 
| name | string | X | √ | The username. Defaults to the resource name. | 
| scope_id | string | X | √ | The scope ID in which the resource is created. Defaults to the provider's `default_scope` if unset. | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


