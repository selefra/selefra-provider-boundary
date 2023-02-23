package resources

import (
	"context"
	"log"

	"github.com/selefra/selefra-provider-sdk/terraform/bridge"
	terraform_providers "github.com/selefra/selefra-provider-sdk/terraform/provider"
	"github.com/selefra/selefra-provider-sdk/terraform/selefra_terraform_schema"

	"github.com/joho/godotenv"
	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/spf13/viper"
)

const Version = "v0.0.1"

func GetSelefraTerraformProvider() *selefra_terraform_schema.SelefraTerraformProvider {
	return &selefra_terraform_schema.SelefraTerraformProvider{
		Name:		"selefra-terraform-provider-boundary",
		Version:	Version,
		ResourceList:	getResources(),
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {
				err := godotenv.Load()
				if err != nil {
					log.Fatal(err)
				}
				diagnostics := schema.NewDiagnostics()

				client, err := newClient(clientMeta)
				if err != nil {
					log.Fatalf("初始化client错误: %s", err.Error())
				}

				// run terraform providers
				if clientMeta.Runtime().Workspace != "" {
					providerSaveDirectory := clientMeta.Runtime().Workspace + "/" + clientMeta.Runtime().ProviderName + "/" + clientMeta.Runtime().ProviderVersion
					providerFileSlice := getTerraformProviderExecuteFileSlice()
					providerExecFilePath, err := terraform_providers.NewProviderDownloader(providerFileSlice).Download(providerSaveDirectory)
					if err != nil {
						return nil, diagnostics.AddError(err)
					}
					bridge := bridge.NewTerraformBridge(providerExecFilePath)

					// read terraform config from selefra provider's config file
					terraformProviderConfig := make(map[string]any, 0)
					if config != nil {
						err := config.Unmarshal(&terraformProviderConfig)
						if err != nil {
							return nil, schema.NewDiagnostics().AddError(err)
						}
					}

					err = bridge.StartBridge(context.Background(), terraformProviderConfig)
					if err != nil {
						return nil, diagnostics.AddError(err)
					}
					client.TerraformBridge = bridge
				}

				return []any{client}, nil
			},
		},
		ConfigMeta: provider.ConfigMeta{
			GetDefaultConfigTemplate: func(ctx context.Context) string {
				// TODO
				return ``
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {
				// TODO
				return nil
			},
		},
		TransformerMeta: schema.TransformerMeta{
			DefaultColumnValueConvertorBlackList:	[]string{},
			DataSourcePullResultAutoExpand:		true,
		},
		ErrorsHandlerMeta: schema.ErrorsHandlerMeta{
			IgnoredErrors: []schema.IgnoredError{schema.IgnoredErrorOnSaveResult},
		},
	}
}

func getTerraformProviderExecuteFileSlice() []*terraform_providers.TerraformProviderFile {
	providerFileSlice := make([]*terraform_providers.TerraformProviderFile, 0)

	providerFileSlice = append(providerFileSlice, &terraform_providers.TerraformProviderFile{
		ProviderName:		"",
		ProviderVersion:	"1.1.4",
		DownloadUrl:		"https://releases.hashicorp.com/terraform-provider-boundary/1.1.4/terraform-provider-boundary_1.1.4_darwin_amd64.zip",
		Arch:			"amd64",
		OS:			"darwin",
	})

	providerFileSlice = append(providerFileSlice, &terraform_providers.TerraformProviderFile{
		ProviderName:		"",
		ProviderVersion:	"1.1.4",
		DownloadUrl:		"https://releases.hashicorp.com/terraform-provider-boundary/1.1.4/terraform-provider-boundary_1.1.4_darwin_arm64.zip",
		Arch:			"arm64",
		OS:			"darwin",
	})

	providerFileSlice = append(providerFileSlice, &terraform_providers.TerraformProviderFile{
		ProviderName:		"",
		ProviderVersion:	"1.1.4",
		DownloadUrl:		"https://releases.hashicorp.com/terraform-provider-boundary/1.1.4/terraform-provider-boundary_1.1.4_freebsd_386.zip",
		Arch:			"386",
		OS:			"freebsd",
	})

	providerFileSlice = append(providerFileSlice, &terraform_providers.TerraformProviderFile{
		ProviderName:		"",
		ProviderVersion:	"1.1.4",
		DownloadUrl:		"https://releases.hashicorp.com/terraform-provider-boundary/1.1.4/terraform-provider-boundary_1.1.4_freebsd_amd64.zip",
		Arch:			"amd64",
		OS:			"freebsd",
	})

	providerFileSlice = append(providerFileSlice, &terraform_providers.TerraformProviderFile{
		ProviderName:		"",
		ProviderVersion:	"1.1.4",
		DownloadUrl:		"https://releases.hashicorp.com/terraform-provider-boundary/1.1.4/terraform-provider-boundary_1.1.4_freebsd_arm.zip",
		Arch:			"arm",
		OS:			"freebsd",
	})

	providerFileSlice = append(providerFileSlice, &terraform_providers.TerraformProviderFile{
		ProviderName:		"",
		ProviderVersion:	"1.1.4",
		DownloadUrl:		"https://releases.hashicorp.com/terraform-provider-boundary/1.1.4/terraform-provider-boundary_1.1.4_linux_386.zip",
		Arch:			"386",
		OS:			"linux",
	})

	providerFileSlice = append(providerFileSlice, &terraform_providers.TerraformProviderFile{
		ProviderName:		"",
		ProviderVersion:	"1.1.4",
		DownloadUrl:		"https://releases.hashicorp.com/terraform-provider-boundary/1.1.4/terraform-provider-boundary_1.1.4_linux_amd64.zip",
		Arch:			"amd64",
		OS:			"linux",
	})

	providerFileSlice = append(providerFileSlice, &terraform_providers.TerraformProviderFile{
		ProviderName:		"",
		ProviderVersion:	"1.1.4",
		DownloadUrl:		"https://releases.hashicorp.com/terraform-provider-boundary/1.1.4/terraform-provider-boundary_1.1.4_linux_arm.zip",
		Arch:			"arm",
		OS:			"linux",
	})

	providerFileSlice = append(providerFileSlice, &terraform_providers.TerraformProviderFile{
		ProviderName:		"",
		ProviderVersion:	"1.1.4",
		DownloadUrl:		"https://releases.hashicorp.com/terraform-provider-boundary/1.1.4/terraform-provider-boundary_1.1.4_linux_arm64.zip",
		Arch:			"arm64",
		OS:			"linux",
	})

	providerFileSlice = append(providerFileSlice, &terraform_providers.TerraformProviderFile{
		ProviderName:		"",
		ProviderVersion:	"1.1.4",
		DownloadUrl:		"https://releases.hashicorp.com/terraform-provider-boundary/1.1.4/terraform-provider-boundary_1.1.4_windows_386.zip",
		Arch:			"386",
		OS:			"windows",
	})

	providerFileSlice = append(providerFileSlice, &terraform_providers.TerraformProviderFile{
		ProviderName:		"",
		ProviderVersion:	"1.1.4",
		DownloadUrl:		"https://releases.hashicorp.com/terraform-provider-boundary/1.1.4/terraform-provider-boundary_1.1.4_windows_amd64.zip",
		Arch:			"amd64",
		OS:			"windows",
	})

	return providerFileSlice
}

func getResources() []*selefra_terraform_schema.SelefraTerraformResource {
	return []*selefra_terraform_schema.SelefraTerraformResource{
		// GetResource_example(),
	}
}
