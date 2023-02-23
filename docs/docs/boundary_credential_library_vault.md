# Table: boundary_credential_library_vault

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| credential_type | string | X | √ | The type of credential the library generates. | 
| http_request_body | string | X | √ | The body of the HTTP request the library sends to Vault when requesting credentials. Only valid if `http_method` is set to `POST`. | 
| credential_mapping_overrides | json | X | √ | The credential mapping override. | 
| credential_store_id | string | X | √ | The ID of the credential store that this library belongs to. | 
| description | string | X | √ | The Vault credential library description. | 
| http_method | string | X | √ | The HTTP method the library uses when requesting credentials from Vault. Defaults to 'GET' | 
| id | string | X | √ | The ID of the Vault credential library. | 
| name | string | X | √ | The Vault credential library name. Defaults to the resource name. | 
| path | string | X | √ | The path in Vault to request credentials from. | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


