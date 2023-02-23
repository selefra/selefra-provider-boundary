# Table: boundary_account_oidc

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| subject | string | X | √ | The OIDC subject. | 
| auth_method_id | string | X | √ | The resource ID for the auth method. | 
| description | string | X | √ | The account description. | 
| id | string | X | √ | The ID of the account. | 
| issuer | string | X | √ | The OIDC issuer. | 
| name | string | X | √ | The account name. Defaults to the resource name. | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


