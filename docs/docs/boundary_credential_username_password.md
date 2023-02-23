# Table: boundary_credential_username_password

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| username | string | X | √ | The username of this username/password credential. | 
| credential_store_id | string | X | √ | The credential store in which to save this username/password credential. | 
| description | string | X | √ | The description of this username/password credential. | 
| id | string | X | √ | The ID of this username/password credential. | 
| name | string | X | √ | The name of this username/password credential. Defaults to the resource name. | 
| password | string | X | √ | The password of this username/password credential. | 
| password_hmac | string | X | √ | The password hmac. | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


