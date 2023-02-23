# Table: boundary_target

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| default_port | float | X | √ | The default port for this target. | 
| ingress_worker_filter | string | X | √ | HCP Only. Boolean expression to filter the workers a user will connect to when initiating a session against this target | 
| session_max_seconds | float | X | √ |  | 
| type | string | X | √ | The target resource type. | 
| brokered_credential_source_ids | json | X | √ | A list of brokered credential source ID's. | 
| egress_worker_filter | string | X | √ | Boolean expression to filter the workers used to access this target | 
| id | string | X | √ | The ID of the target. | 
| injected_application_credential_source_ids | json | X | √ | A list of injected application credential source ID's. | 
| name | string | X | √ | The target name. Defaults to the resource name. | 
| address | string | X | √ | Optionally, a valid network address to connect to for this target. Cannot be used alongside host_source_ids. | 
| host_source_ids | json | X | √ | A list of host source ID's. Cannot be used alongside address. | 
| scope_id | string | X | √ | The scope ID in which the resource is created. Defaults to the provider's `default_scope` if unset. | 
| worker_filter | string | X | √ | Boolean expression to filter the workers for this target | 
| description | string | X | √ | The target description. | 
| session_connection_limit | float | X | √ |  | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


