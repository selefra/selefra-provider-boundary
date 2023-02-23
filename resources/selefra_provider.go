package resources

import (
	"context"
	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/terraform/bridge"
)

func GetSelefraProvider() *provider.Provider {
	diagnostics := schema.NewDiagnostics()
	selefraProvider, d := GetSelefraTerraformProvider().ToSelefraProvider(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) *bridge.TerraformBridge {
		return taskClient.(*Client).TerraformBridge
	})
	if diagnostics.AddDiagnostics(d).HasError() {
		panic(diagnostics.ToString())
	}

    selefraProvider.TableList = GetSelefraTables()

	return selefraProvider
}

func GetSelefraTables() []*schema.Table {

    diagnostics := schema.NewDiagnostics()
    tables := make([]*schema.Table, 0)
    var table *schema.Table
    var d *schema.Diagnostics

    
    table, d = TableSchemaGenerator_boundary_host_static()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_boundary_credential_ssh_private_key()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_boundary_credential_json()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_boundary_account_password()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_boundary_credential_username_password()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_boundary_role()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_boundary_host_set_plugin()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_boundary_credential_library_vault_ssh_certificate()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_boundary_auth_method_oidc()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_boundary_scope()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_boundary_target()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_boundary_host_catalog_static()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_boundary_auth_method()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_boundary_host_catalog_plugin()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_boundary_group()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_boundary_host_set()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_boundary_credential_library_vault()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_boundary_credential_store_vault()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_boundary_user()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_boundary_account_oidc()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_boundary_host_set_static()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_boundary_host()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_boundary_auth_method_password()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_boundary_credential_store_static()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_boundary_host_catalog()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_boundary_managed_group()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_boundary_worker()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_boundary_account()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    

    if diagnostics.HasError() {
        panic(diagnostics.ToString())
    }

	return tables
}
