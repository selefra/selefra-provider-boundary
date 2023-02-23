# Table: boundary_account_password

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| login_name | string | X | √ | The login name for this account. | 
| name | string | X | √ | The account name. Defaults to the resource name. | 
| password | string | X | √ | The account password. Only set on create, changes will not be reflected when updating account. | 
| type | string | X | √ | The resource type. | 
| auth_method_id | string | X | √ | The resource ID for the auth method. | 
| description | string | X | √ | The account description. | 
| id | string | X | √ | The ID of the account. | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


