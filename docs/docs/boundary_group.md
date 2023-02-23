# Table: boundary_group

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| description | string | X | √ | The group description. | 
| id | string | X | √ | The ID of the group. | 
| member_ids | json | X | √ | Resource IDs for group members, these are most likely boundary users. | 
| name | string | X | √ | The group name. Defaults to the resource name. | 
| scope_id | string | X | √ | The scope ID in which the resource is created. Defaults to the provider's `default_scope` if unset. | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


