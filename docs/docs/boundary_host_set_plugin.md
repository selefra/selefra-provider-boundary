# Table: boundary_host_set_plugin

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | X | √ | The ID of the host set. | 
| name | string | X | √ | The host set name. Defaults to the resource name. | 
| preferred_endpoints | json | X | √ | The ordered list of preferred endpoints. | 
| sync_interval_seconds | float | X | √ | The value to set for the sync interval seconds. | 
| type | string | X | √ | The type of host set | 
| attributes_json | string | X | √ | The attributes for the host set. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the host set. | 
| description | string | X | √ | The host set description. | 
| host_catalog_id | string | X | √ | The catalog for the host set. | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


