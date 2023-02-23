# Table: boundary_host

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ | The host name. Defaults to the resource name. | 
| type | string | X | √ | The type of host | 
| address | string | X | √ | The static address of the host resource as `<IP>` (note: port assignment occurs in the target resource definition, do not add :port here) or a domain name. | 
| description | string | X | √ | The host description. | 
| host_catalog_id | string | X | √ |  | 
| id | string | X | √ | The ID of the host. | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


