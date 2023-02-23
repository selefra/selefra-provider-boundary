# Table: boundary_scope

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| auto_create_admin_role | bool | X | √ | If set, when a new scope is created, the provider will not disable the functionality that automatically creates a role in the new scope and gives permissions to manage the scope to the provider's user. Marking this true makes for simpler HCL but results in role resources that are unmanaged by Terraform. | 
| auto_create_default_role | bool | X | √ | Only relevant when creating an org scope. If set, when a new scope is created, the provider will not disable the functionality that automatically creates a role in the new scope and gives listing of scopes and auth methods and the ability to authenticate to the anonymous user. Marking this true makes for simpler HCL but results in role resources that are unmanaged by Terraform. | 
| description | string | X | √ | The scope description. | 
| global_scope | bool | X | √ | Indicates that the scope containing this value is the global scope, which triggers some specialized behavior to allow it to be imported and managed. | 
| id | string | X | √ | The ID of the scope. | 
| name | string | X | √ | The scope name. Defaults to the resource name. | 
| scope_id | string | X | √ | The scope ID containing the sub scope resource. | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


