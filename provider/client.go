package provider

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/boundary/api"
	"github.com/hashicorp/boundary/api/authmethods"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/terraform/bridge"
)

type Config struct {
	Addr         string `yaml:"addr" json:"addr" mapstructure:"addr"`
	AuthMethodId string `yaml:"auth_mathod_id" json:"auth_method_id" mapstructure:"auth_method_id"`
	LoginName    string `yaml:"password_auth_method_login_name" json:"password_auth_method_login_name" mapstructure:"password_auth_method_login_name"`
	PassWord     string `yaml:"password_auth_method_password" json:"password_auth_method_password" mapstructure:"password_auth_method_password"`
}

type Client struct {
	TerraformBridge *bridge.TerraformBridge
	Config

	// TODO You can continue to refine your client
	ApiClient *api.Client
}

/* *
* New Client, with    param: BOUNDARY_ADDR PASSWORD_AUTH_METHOD_PASSWORD  PASSWORD_AUTH_METHOD_PASSWORD
* Terrform URL: https://registry.terraform.io/providers/bo
 */
func newClient(clientMeta *schema.ClientMeta, config *Config) (*Client, error) {
	if config.Addr == "" || config.AuthMethodId == "" || config.LoginName == "" || config.PassWord == "" {
		ErrorF(clientMeta, "Config Error!")
		return nil, errors.New("Get Config Error!")
	}

	cfg := &api.Config{
		Addr: config.Addr,
	}

	// The default address points to the default dev mode address
	client, err := api.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	token, err := getToken(client, config)
	if err != nil {
		ErrorF(clientMeta, "Get token error: %s", err.Error())
		return nil, err
	}

	client.SetToken(token)

	return &Client{
		ApiClient: client,
		Config:    *config,
	}, nil
}

func getToken(client *api.Client, config *Config) (string, error) {
	credentials := map[string]interface{}{
		"login_name": config.LoginName,
		"password":   config.PassWord,
	}

	amClient := authmethods.NewClient(client)
	authenticationResult, err := amClient.Authenticate(context.Background(), config.AuthMethodId, "login", credentials)
	if err != nil {
		return "", err
	}
	if authenticationResult.Attributes["token"] == "" {
		return "", errors.New("Token Not Found!")
	}
	return fmt.Sprint(authenticationResult.Attributes["token"]), nil
}

func DebugF(clientMeta *schema.ClientMeta, msg string, args ...any) {
	if clientMeta != nil {
		clientMeta.DebugF(msg, args...)
	}
}

func ErrorF(clientMeta *schema.ClientMeta, msg string, args ...any) {
	if clientMeta != nil {
		clientMeta.ErrorF(msg, args...)
	}
}
