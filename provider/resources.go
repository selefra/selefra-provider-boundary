package provider

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/terraform/selefra_terraform_schema"
)

// terraform resource: boundary_host_static
func GetResource_boundary_host_static() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "boundary_host_static",
		TerraformResourceName: "boundary_host_static",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO
			return nil, nil
		},
	}
}

// terraform resource: boundary_credential_ssh_private_key
func GetResource_boundary_credential_ssh_private_key() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "boundary_credential_ssh_private_key",
		TerraformResourceName: "boundary_credential_ssh_private_key",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO
			return nil, nil
		},
	}
}

// terraform resource: boundary_credential_json
func GetResource_boundary_credential_json() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "boundary_credential_json",
		TerraformResourceName: "boundary_credential_json",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO
			return nil, nil
		},
	}
}

// terraform resource: boundary_account_password
func GetResource_boundary_account_password() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "boundary_account_password",
		TerraformResourceName: "boundary_account_password",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO
			return nil, nil
		},
	}
}

// terraform resource: boundary_credential_username_password
func GetResource_boundary_credential_username_password() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "boundary_credential_username_password",
		TerraformResourceName: "boundary_credential_username_password",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO
			return nil, nil
		},
	}
}

// terraform resource: boundary_role
func GetResource_boundary_role() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "boundary_role",
		TerraformResourceName: "boundary_role",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO
			return nil, nil
		},
	}
}

// terraform resource: boundary_host_set_plugin
func GetResource_boundary_host_set_plugin() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "boundary_host_set_plugin",
		TerraformResourceName: "boundary_host_set_plugin",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO
			return nil, nil
		},
	}
}

// terraform resource: boundary_credential_library_vault_ssh_certificate
func GetResource_boundary_credential_library_vault_ssh_certificate() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "boundary_credential_library_vault_ssh_certificate",
		TerraformResourceName: "boundary_credential_library_vault_ssh_certificate",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO
			return nil, nil
		},
	}
}

// terraform resource: boundary_auth_method_oidc
func GetResource_boundary_auth_method_oidc() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "boundary_auth_method_oidc",
		TerraformResourceName: "boundary_auth_method_oidc",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO
			return nil, nil
		},
	}
}

type ListScopesReq struct {
	ScopeId   string `json:"scope_id"`
	Recursive bool   `json:"recursive"`
	Filter    string `json:"filter"`
}

func (r *ListScopesReq) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListScopesResp(data []byte) (ListScopesResp, error) {
	var r ListScopesResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type ListScopesResp struct {
	ListScopesRespItems []ListScopesRespItems `json:"items"`
}

type ListScopesRespItems struct {
	ID      string `json:"id"`
	ScopeID string `json:"scope_id"`
}

func sendBoundaryScopeReq(client *Client) (ListScopesResp, error) {
	reqBody := ListScopesReq{
		ScopeId:   "global",
		Recursive: true,
		Filter:    "",
	}
	reqBodyByte, err := reqBody.Marshal()
	if err != nil {
		return ListScopesResp{}, err
	}

	retryReq, err := client.ApiClient.NewRequest(context.Background(), "GET", "scopes", reqBodyByte)
	if err != nil {
		return ListScopesResp{}, err

	}

	res, err := client.ApiClient.Do(retryReq)
	if err != nil {
		return ListScopesResp{}, err
	}

	resBodyByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ListScopesResp{}, err
	}

	scopes, err := UnmarshalListScopesResp(resBodyByte)
	if err != nil {
		return ListScopesResp{}, err
	}
	return scopes, err
}

// terraform resource: boundary_scope
func GetResource_boundary_scope() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "boundary_scope",
		TerraformResourceName: "boundary_scope",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO
			client := taskClient.(*Client)

			scopes, err := sendBoundaryScopeReq(client)
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, scope := range scopes.ListScopesRespItems {
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ID: scope.ID,
					ArgumentMap: map[string]any{
						"scope_id": scope.ScopeID,
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
		SelefraTableName:      "boundary_target",
		TerraformResourceName: "boundary_target",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO
			return nil, nil
		},
	}
}

// terraform resource: boundary_host_catalog_static
func GetResource_boundary_host_catalog_static() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "boundary_host_catalog_static",
		TerraformResourceName: "boundary_host_catalog_static",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO
			return nil, nil
		},
	}
}

// terraform resource: boundary_auth_method
func GetResource_boundary_auth_method() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "boundary_auth_method",
		TerraformResourceName: "boundary_auth_method",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO
			return nil, nil
		},
	}
}

// terraform resource: boundary_host_catalog_plugin
func GetResource_boundary_host_catalog_plugin() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "boundary_host_catalog_plugin",
		TerraformResourceName: "boundary_host_catalog_plugin",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO
			return nil, nil
		},
	}
}

// terraform resource: boundary_group
func GetResource_boundary_group() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "boundary_group",
		TerraformResourceName: "boundary_group",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO
			return nil, nil
		},
	}
}

// terraform resource: boundary_host_set
func GetResource_boundary_host_set() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "boundary_host_set",
		TerraformResourceName: "boundary_host_set",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO
			return nil, nil
		},
	}
}

// terraform resource: boundary_credential_library_vault
func GetResource_boundary_credential_library_vault() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "boundary_credential_library_vault",
		TerraformResourceName: "boundary_credential_library_vault",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO
			return nil, nil
		},
	}
}

// terraform resource: boundary_credential_store_vault
func GetResource_boundary_credential_store_vault() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "boundary_credential_store_vault",
		TerraformResourceName: "boundary_credential_store_vault",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO
			return nil, nil
		},
	}
}

// terraform resource: boundary_user
func GetResource_boundary_user() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "boundary_user",
		TerraformResourceName: "boundary_user",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO
			return nil, nil
		},
	}
}

// terraform resource: boundary_account_oidc
func GetResource_boundary_account_oidc() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "boundary_account_oidc",
		TerraformResourceName: "boundary_account_oidc",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO
			return nil, nil
		},
	}
}

// terraform resource: boundary_host_set_static
func GetResource_boundary_host_set_static() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "boundary_host_set_static",
		TerraformResourceName: "boundary_host_set_static",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO
			return nil, nil
		},
	}
}

// terraform resource: boundary_host
func GetResource_boundary_host() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "boundary_host",
		TerraformResourceName: "boundary_host",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO
			return nil, nil
		},
	}
}

// terraform resource: boundary_auth_method_password
func GetResource_boundary_auth_method_password() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "boundary_auth_method_password",
		TerraformResourceName: "boundary_auth_method_password",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO
			return nil, nil
		},
	}
}

// terraform resource: boundary_credential_store_static
func GetResource_boundary_credential_store_static() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "boundary_credential_store_static",
		TerraformResourceName: "boundary_credential_store_static",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO
			return nil, nil
		},
	}
}

// terraform resource: boundary_host_catalog
func GetResource_boundary_host_catalog() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "boundary_host_catalog",
		TerraformResourceName: "boundary_host_catalog",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO
			return nil, nil
		},
	}
}

// terraform resource: boundary_managed_group
func GetResource_boundary_managed_group() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "boundary_managed_group",
		TerraformResourceName: "boundary_managed_group",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO
			return nil, nil
		},
	}
}

// terraform resource: boundary_worker
func GetResource_boundary_worker() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "boundary_worker",
		TerraformResourceName: "boundary_worker",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO
			return nil, nil
		},
	}
}

// terraform resource: boundary_account
func GetResource_boundary_account() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "boundary_account",
		TerraformResourceName: "boundary_account",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO
			return nil, nil
		},
	}
}
