# Table: boundary_credential_json

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| credential_store_id | string | X | √ | The credential store in which to save this json credential. | 
| description | string | X | √ | The description of this json credential. | 
| id | string | X | √ | The ID of this json credential. | 
| name | string | X | √ | The name of this json credential. Defaults to the resource name. | 
| object | string | X | √ | The object for the this json credential. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file | 
| object_hmac | string | X | √ | The object hmac. | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


