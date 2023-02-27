# Table: boundary_auth_method_password

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| description | string | X | √ | The auth method description. | 
| id | string | X | √ | The ID of the account. | 
| min_login_name_length | float | X | √ | The minimum login name length. | 
| min_password_length | float | X | √ | The minimum password length. | 
| name | string | X | √ | The auth method name. Defaults to the resource name. | 
| scope_id | string | X | √ | The scope ID. | 
| type | string | X | √ | The resource type, hardcoded per resource | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


