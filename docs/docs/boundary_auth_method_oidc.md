# Table: boundary_auth_method_oidc

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_claim_maps | json | X | √ | Account claim maps for the to_claim of sub. | 
| allowed_audiences | json | X | √ | Audiences for which the provider responses will be allowed | 
| api_url_prefix | string | X | √ | The API prefix to use when generating callback URLs for the provider. Should be set to an address at which the provider can reach back to the controller. | 
| callback_url | string | X | √ | The URL that should be provided to the IdP for callbacks. | 
| client_id | string | X | √ | The client ID assigned to this auth method from the provider. | 
| client_secret_hmac | string | X | √ | The HMAC of the client secret returned by the Boundary controller, which is used for comparison after initial setting of the value. | 
| idp_ca_certs | json | X | √ | A list of CA certificates to trust when validating the IdP's token signatures. | 
| claims_scopes | json | X | √ | Claims scopes. | 
| description | string | X | √ | The auth method description. | 
| disable_discovered_config_validation | bool | X | √ | Disables validation logic ensuring that the OIDC provider's information from its discovery endpoint matches the information here. The validation is only performed at create or update time. | 
| is_primary_for_scope | bool | X | √ | When true, makes this auth method the primary auth method for the scope in which it resides. The primary auth method for a scope means the the user will be automatically created when they login using an OIDC account. | 
| max_age | float | X | √ | The max age to provide to the provider, indicating how much time is allowed to have passed since the last authentication before the user is challenged again. | 
| name | string | X | √ | The auth method name. Defaults to the resource name. | 
| id | string | X | √ | The ID of the auth method. | 
| issuer | string | X | √ | The issuer corresponding to the provider, which must match the issuer field in generated tokens. | 
| state | string | X | √ | Can be one of 'inactive', 'active-private', or 'active-public'. Currently automatically set to active-public. | 
| type | string | X | √ | The type of auth method; hardcoded. | 
| client_secret | string | X | √ | The secret key assigned to this auth method from the provider. Once set, only the hash will be kept and the original value can be removed from configuration. | 
| scope_id | string | X | √ | The scope ID. | 
| signing_algorithms | json | X | √ | Allowed signing algorithms for the provider's issued tokens. | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


