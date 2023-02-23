# Table: boundary_credential_store_static

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| description | string | X | √ | The static credential store description. | 
| id | string | X | √ | The ID of the static credential store. | 
| name | string | X | √ | The static credential store name. Defaults to the resource name. | 
| scope_id | string | X | √ | The scope for this credential store. | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


