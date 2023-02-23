# Table: boundary_credential_store_vault

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| description | string | X | √ | The Vault credential store description. | 
| tls_server_name | string | X | √ | Name to use as the SNI host when connecting to Vault via TLS. | 
| token_hmac | string | X | √ | The Vault token hmac. | 
| client_certificate | string | X | √ | A PEM-encoded client certificate to use for TLS authentication to the Vault server. | 
| client_certificate_key | string | X | √ | A PEM-encoded private key matching the client certificate from 'client_certificate'. | 
| client_certificate_key_hmac | string | X | √ | The Vault client certificate key hmac. | 
| scope_id | string | X | √ | The scope for this credential store. | 
| tls_skip_verify | bool | X | √ | Whether or not to skip TLS verification. | 
| ca_cert | string | X | √ | A PEM-encoded CA certificate to verify the Vault server's TLS certificate. | 
| name | string | X | √ | The Vault credential store name. Defaults to the resource name. | 
| token | string | X | √ | A token used for accessing Vault. | 
| address | string | X | √ | The address to Vault server. This should be a complete URL such as 'https://127.0.0.1:8200' | 
| id | string | X | √ | The ID of the Vault credential store. | 
| namespace | string | X | √ | The namespace within Vault to use. | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


