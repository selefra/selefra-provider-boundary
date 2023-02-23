# Table: boundary_host_catalog_static

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| description | string | X | √ | The host catalog description. | 
| id | string | X | √ | The ID of the host catalog. | 
| name | string | X | √ | The host catalog name. Defaults to the resource name. | 
| scope_id | string | X | √ | The scope ID in which the resource is created. | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


