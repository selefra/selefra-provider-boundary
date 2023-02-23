# Table: boundary_managed_group

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| description | string | X | √ | The managed group description. | 
| filter | string | X | √ | Boolean expression to filter the workers for this managed group. | 
| id | string | X | √ | The ID of the group. | 
| name | string | X | √ | The managed group name. Defaults to the resource name. | 
| auth_method_id | string | X | √ | The resource ID for the auth method. | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


