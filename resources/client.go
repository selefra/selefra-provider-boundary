package resources

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/hashicorp/boundary/api"
	"github.com/hashicorp/boundary/api/authmethods"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/terraform/bridge"
)

type Client struct {
	TerraformBridge	*bridge.TerraformBridge

	// TODO You can continue to refine your client
	ApiClient	*api.Client
}

/* *
* New Client, with    param: BOUNDARY_ADDR PASSWORD_AUTH_METHOD_PASSWORD  PASSWORD_AUTH_METHOD_PASSWORD
* Terrform URL: https://registry.terraform.io/providers/bo
 */
func newClient(clientMeta *schema.ClientMeta) (*Client, error) {
	cfg := &api.Config{
		Addr: os.Getenv("BOUNDARY_ADDR"),
	}

	// The default address points to the default dev mode address
	client, err := api.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	token, err := getToken(client)
	if err != nil {
		ErrorF(clientMeta, "获取Token信息失败: %s", err.Error())
		return nil, err
	}

	client.SetToken(token)

	return &Client{
		ApiClient: client,
	}, nil
}

func getToken(client *api.Client) (string, error) {
	credentials := map[string]interface{}{
		"login_name":	os.Getenv("PASSWORD_AUTH_METHOD_LOGIN_NAME"),
		"password":	os.Getenv("PASSWORD_AUTH_METHOD_PASSWORD"),
	}

	amClient := authmethods.NewClient(client)
	authenticationResult, err := amClient.Authenticate(context.Background(), os.Getenv("AUTH_METHOD_ID"), "login", credentials)
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
