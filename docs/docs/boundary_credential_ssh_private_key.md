# Table: boundary_credential_ssh_private_key

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| description | string | X | √ | The description of the credential. | 
| id | string | X | √ | The ID of the credential. | 
| private_key_hmac | string | X | √ | The private key hmac. | 
| private_key_passphrase | string | X | √ | The passphrase of the private key associated with the credential. | 
| private_key_passphrase_hmac | string | X | √ | The private key passphrase hmac. | 
| credential_store_id | string | X | √ | ID of the credential store this credential belongs to. | 
| name | string | X | √ | The name of the credential. Defaults to the resource name. | 
| private_key | string | X | √ | The private key associated with the credential. | 
| username | string | X | √ | The username associated with the credential. | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


