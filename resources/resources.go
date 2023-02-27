package resources

import (
	"context"

	"github.com/hashicorp/boundary/api/authmethods"
	"github.com/hashicorp/boundary/api/groups"
	"github.com/hashicorp/boundary/api/roles"
	"github.com/hashicorp/boundary/api/scopes"
	"github.com/hashicorp/boundary/api/targets"
	"github.com/hashicorp/boundary/api/users"
	"github.com/hashicorp/boundary/api/workers"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/terraform/selefra_terraform_schema"
)

// // terraform resource: boundary_host_static
// func GetResource_boundary_host_static() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "boundary_host_static",
// 		TerraformResourceName: "boundary_host_static",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			// TODO
// 			return nil, nil
// 		},
// 	}
// }

// // terraform resource: boundary_credential_ssh_private_key
// func GetResource_boundary_credential_ssh_private_key() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "boundary_credential_ssh_private_key",
// 		TerraformResourceName: "boundary_credential_ssh_private_key",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			// TODO
// 			return nil, nil
// 		},
// 	}
// }

// // terraform resource: boundary_credential_json
// func GetResource_boundary_credential_json() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "boundary_credential_json",
// 		TerraformResourceName: "boundary_credential_json",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			// TODO
// 			return nil, nil
// 		},
// 	}
// }

// // terraform resource: boundary_account_password
// func GetResource_boundary_account_password() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "boundary_account_password",
// 		TerraformResourceName: "boundary_account_password",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			// TODO
// 			return nil, nil
// 		},
// 	}
// }

// terraform resource: boundary_credential_username_password
// func GetResource_boundary_credential_username_password() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "boundary_credential_username_password",
// 		TerraformResourceName: "boundary_credential_username_password",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			client := taskClient.(*Client)
// 			credentialsClient := credentials.NewClient(client.ApiClient)
// 			credentialsListRes, err := credentialsClient.List(ctx, "global")
// 			if err != nil {
// 				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
// 			}

// 			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
// 			for _, item := range credentialsListRes.Items {
// 				userInfo, err := item.GetUsernamePasswordAttributes()
// 				if err != nil {
// 					return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
// 				}
// 				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
// 					ID: item.Id,
// 					ArgumentMap: map[string]any{
// 						"credential_store_id": item.CredentialStoreId,
// 						"username":            userInfo.Username,
// 						"password":            userInfo.Password,
// 					},
// 				})
// 			}
// 			return resourceRequestParamSlice, nil
// 		},
// 	}
// }

// terraform resource: boundary_role
func GetResource_boundary_role() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:	"boundary_role",
		TerraformResourceName:	"boundary_role",
		Description:		"",
		SubTables:		nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)
			roleClient := roles.NewClient(client.ApiClient)

			rolesListRes, err := roleClient.List(ctx, "global", roles.WithRecursive(true))
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, roles := range rolesListRes.Items {
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ID:	roles.Id,
					ArgumentMap: map[string]any{
						"scope_id": roles.ScopeId,
					},
				})
			}

			return resourceRequestParamSlice, nil
		},
	}
}

// // terraform resource: boundary_host_set_plugin
// func GetResource_boundary_host_set_plugin() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "boundary_host_set_plugin",
// 		TerraformResourceName: "boundary_host_set_plugin",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			// TODO
// 			return nil, nil
// 		},
// 	}
// }

// // terraform resource: boundary_credential_library_vault_ssh_certificate
// func GetResource_boundary_credential_library_vault_ssh_certificate() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "boundary_credential_library_vault_ssh_certificate",
// 		TerraformResourceName: "boundary_credential_library_vault_ssh_certificate",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			// TODO
// 			return nil, nil
// 		},
// 	}
// }

// terraform resource: boundary_auth_method_oidc
func GetResource_boundary_auth_method_oidc() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:	"boundary_auth_method_oidc",
		TerraformResourceName:	"boundary_auth_method_oidc",
		Description:		"",
		SubTables:		nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)
			amClient := authmethods.NewClient(client.ApiClient)

			authMetuodsListResult, err := amClient.List(context.Background(), "global", authmethods.DefaultAttributes())
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, authMethod := range authMetuodsListResult.Items {
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ID:	authMethod.Id,
					ArgumentMap: map[string]any{
						"scope_id": authMethod.ScopeId,
					},
				})
			}

			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: boundary_scope
func GetResource_boundary_scope() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:	"boundary_scope",
		TerraformResourceName:	"boundary_scope",
		Description:		"",
		SubTables:		nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)
			scopeClient := scopes.NewClient(client.ApiClient)
			scopesListRes, err := scopeClient.List(ctx, "global", scopes.WithRecursive(true))
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, scope := range scopesListRes.Items {
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ID:	scope.Id,
					ArgumentMap: map[string]any{
						"scope_id": scope.ScopeId,
					},
				})
			}
			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: boundary_target
func GetResource_boundary_target() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:	"boundary_target",
		TerraformResourceName:	"boundary_target",
		Description:		"",
		SubTables:		nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)
			targetsClient := targets.NewClient(client.ApiClient)

			targetsListRes, err := targetsClient.List(context.Background(), "global", targets.WithRecursive(true))
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, target := range targetsListRes.Items {
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ID:	target.Id,
					ArgumentMap: map[string]any{
						"scope_id":	target.ScopeId,
						"type":		target.Type,
					},
				})
			}

			return resourceRequestParamSlice, nil
		},
	}
}

// // terraform resource: boundary_host_catalog_static
// func GetResource_boundary_host_catalog_static() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "boundary_host_catalog_static",
// 		TerraformResourceName: "boundary_host_catalog_static",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			// TODO
// 			return nil, nil
// 		},
// 	}
// }

// terraform resource: boundary_auth_method
func GetResource_boundary_auth_method() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:	"boundary_auth_method",
		TerraformResourceName:	"boundary_auth_method",
		Description:		"",
		SubTables:		nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)
			amClient := authmethods.NewClient(client.ApiClient)

			authMetuodsListResult, err := amClient.List(context.Background(), "global", authmethods.DefaultAttributes())
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, authMethod := range authMetuodsListResult.Items {
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ID:	authMethod.Id,
					ArgumentMap: map[string]any{
						"scope_id":	authMethod.ScopeId,
						"type":		authMethod.Type,
					},
				})
			}

			return resourceRequestParamSlice, nil
		},
	}
}

// // terraform resource: boundary_host_catalog_plugin
// func GetResource_boundary_host_catalog_plugin() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "boundary_host_catalog_plugin",
// 		TerraformResourceName: "boundary_host_catalog_plugin",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			// TODO
// 			return nil, nil
// 		},
// 	}
// }

// terraform resource: boundary_group
func GetResource_boundary_group() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:	"boundary_group",
		TerraformResourceName:	"boundary_group",
		Description:		"",
		SubTables:		nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)
			groupsClient := groups.NewClient(client.ApiClient)

			groupsListRes, err := groupsClient.List(context.Background(), "global", groups.WithRecursive(true))
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, authMethod := range groupsListRes.Items {
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ID:	authMethod.Id,
					ArgumentMap: map[string]any{
						"scope_id": authMethod.ScopeId,
					},
				})
			}

			return resourceRequestParamSlice, nil
		},
	}
}

// // terraform resource: boundary_host_set
// func GetResource_boundary_host_set() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "boundary_host_set",
// 		TerraformResourceName: "boundary_host_set",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			// TODO
// 			return nil, nil
// 		},
// 	}
// }

// // terraform resource: boundary_credential_library_vault
// func GetResource_boundary_credential_library_vault() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "boundary_credential_library_vault",
// 		TerraformResourceName: "boundary_credential_library_vault",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			// TODO
// 			return nil, nil
// 		},
// 	}
// }

// // terraform resource: boundary_credential_store_vault
// func GetResource_boundary_credential_store_vault() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "boundary_credential_store_vault",
// 		TerraformResourceName: "boundary_credential_store_vault",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			// TODO
// 			return nil, nil
// 		},
// 	}
// }

// terraform resource: boundary_user
func GetResource_boundary_user() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:	"boundary_user",
		TerraformResourceName:	"boundary_user",
		Description:		"",
		SubTables:		nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)
			usersClient := users.NewClient(client.ApiClient)
			usersListRes, err := usersClient.List(context.Background(), "global", users.WithRecursive(true))
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, user := range usersListRes.Items {
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ID:	user.Id,
					ArgumentMap: map[string]any{
						"scope_id": user.ScopeId,
					},
				})
			}
			return resourceRequestParamSlice, nil
		},
	}
}

// // terraform resource: boundary_account_oidc
// func GetResource_boundary_account_oidc() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "boundary_account_oidc",
// 		TerraformResourceName: "boundary_account_oidc",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			// TODO
// 			return nil, nil
// 		},
// 	}
// }

// // terraform resource: boundary_host_set_static
// func GetResource_boundary_host_set_static() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "boundary_host_set_static",
// 		TerraformResourceName: "boundary_host_set_static",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			// TODO
// 			return nil, nil
// 		},
// 	}
// }

// // terraform resource: boundary_host
// func GetResource_boundary_host() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "boundary_host",
// 		TerraformResourceName: "boundary_host",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			// TODO
// 			return nil, nil
// 		},
// 	}
// }

// terraform resource: boundary_auth_method_password
func GetResource_boundary_auth_method_password() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:	"boundary_auth_method_password",
		TerraformResourceName:	"boundary_auth_method_password",
		Description:		"",
		SubTables:		nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)
			authmethodsClient := authmethods.NewClient(client.ApiClient)
			authmethodsListRes, err := authmethodsClient.List(context.Background(), "global", authmethods.WithRecursive(true))
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, authmetohd := range authmethodsListRes.Items {
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ID:	authmetohd.Id,
					ArgumentMap: map[string]any{
						"scope_id":	authmetohd.ScopeId,
						"type":		authmetohd.Type,
					},
				})
			}
			return resourceRequestParamSlice, nil
		},
	}
}

// // terraform resource: boundary_credential_store_static
// func GetResource_boundary_credential_store_static() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "boundary_credential_store_static",
// 		TerraformResourceName: "boundary_credential_store_static",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			// TODO
// 			return nil, nil
// 		},
// 	}
// }

// // terraform resource: boundary_host_catalog
// func GetResource_boundary_host_catalog() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "boundary_host_catalog",
// 		TerraformResourceName: "boundary_host_catalog",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			// TODO
// 			return nil, nil
// 		},
// 	}
// }

// // terraform resource: boundary_managed_group
// func GetResource_boundary_managed_group() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "boundary_managed_group",
// 		TerraformResourceName: "boundary_managed_group",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			// TODO
// 			return nil, nil
// 		},
// 	}
// }

// terraform resource: boundary_worker
func GetResource_boundary_worker() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:	"boundary_worker",
		TerraformResourceName:	"boundary_worker",
		Description:		"",
		SubTables:		nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)
			workersClient := workers.NewClient(client.ApiClient)
			workersListRes, err := workersClient.List(context.Background(), "global", workers.WithRecursive(true))
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, user := range workersListRes.Items {
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ID: user.Id,
				})
			}
			return resourceRequestParamSlice, nil
		},
	}
}

// // terraform resource: boundary_account
// func GetResource_boundary_account() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "boundary_account",
// 		TerraformResourceName: "boundary_account",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			// TODO
// 			return nil, nil
// 		},
// 	}
// }
