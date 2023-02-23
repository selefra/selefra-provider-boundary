# Table: boundary_credential_library_vault_ssh_certificate

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| key_bits | float | X | √ | Specifies the number of bits to use for the generated keys. | 
| key_type | string | X | √ | Specifies the desired key type; must be ed25519, ecdsa, or rsa. | 
| path | string | X | √ | The path in Vault to request credentials from. | 
| ttl | string | X | √ | Specifies the requested time to live for a certificate returned from the library. | 
| username | string | X | √ | The username to use with the certificate returned by the library. | 
| extensions | json | X | √ | Specifies a map of the extensions that the certificate should be signed for. | 
| id | string | X | √ | The ID of the Vault credential library. | 
| description | string | X | √ | The Vault credential library description. | 
| key_id | string | X | √ | Specifies the key id a certificate should have. | 
| name | string | X | √ | The Vault credential library name. Defaults to the resource name. | 
| credential_store_id | string | X | √ | The ID of the credential store that this library belongs to. | 
| critical_options | json | X | √ | Specifies a map of the critical options that the certificate should be signed for. | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


