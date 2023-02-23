# Table: boundary_host_catalog_plugin

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| internal_force_update | string | X | √ | Internal only. Used to force update so that we can always check the value of secrets. | 
| internal_hmac_used_for_secrets_config_hmac | string | X | √ | Internal only. The Boundary-provided HMAC used to calculate the current value of the HMAC'd config. Used for drift detection. | 
| plugin_name | string | X | √ | The name of the plugin that should back the resource. This or plugin_id must be defined. | 
| scope_id | string | X | √ | The scope ID in which the resource is created. | 
| secrets_json | string | X | √ | The secrets for the host catalog. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" to clear any existing values. NOTE: Unlike "attributes_json", removing this block will NOT clear secrets from the host catalog; this allows injecting secrets for one call, then removing them for storage. | 
| attributes_json | string | X | √ | The attributes for the host catalog. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the host catalog. | 
| description | string | X | √ | The host catalog description. | 
| id | string | X | √ | The ID of the host catalog. | 
| internal_secrets_config_hmac | string | X | √ | Internal only. HMAC of (serverSecretsHmac + config secrets). Used for proper secrets handling. | 
| name | string | X | √ | The host catalog name. Defaults to the resource name. | 
| plugin_id | string | X | √ | The ID of the plugin that should back the resource. This or plugin_name must be defined. | 
| secrets_hmac | string | X | √ | The HMAC'd secrets value returned from the server. | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


