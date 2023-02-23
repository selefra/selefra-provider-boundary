# Table: boundary_host_set_static

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| description | string | X | √ | The host set description. | 
| host_catalog_id | string | X | √ | The catalog for the host set. | 
| host_ids | json | X | √ | The list of host IDs contained in this set. | 
| id | string | X | √ | The ID of the host set. | 
| name | string | X | √ | The host set name. Defaults to the resource name. | 
| type | string | X | √ | The type of host set | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


