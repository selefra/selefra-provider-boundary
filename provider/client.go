package provider

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

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

func (c *Config) isVaild() bool {
	if c.Addr == "" || c.AuthMethodId == "" || c.LoginName == "" || c.PassWord == "" {
		return false
	}
	return true
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
	fmt.Printf("----------------->\n")

	// Must be four param
	// find param in ~/.terraformrc
	if !config.isVaild() {
		homedir, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("get param failed: %v", err)
		}

		rcfContent, err := os.ReadFile(filepath.Join(homedir, ".terraformrc"))
		if err != nil {
			return nil, fmt.Errorf("get param failed: %v", err)
		}
		config.Addr, err = getBoundaryParam(rcfContent, "addr")
		if err != nil {
			return nil, fmt.Errorf("get param addr failed: %v", err)
		}
		config.AuthMethodId, err = getBoundaryParam(rcfContent, "auth_method_id")
		if err != nil {
			return nil, fmt.Errorf("get param addr failed: %v", err)
		}
		config.LoginName, err = getBoundaryParam(rcfContent, "password_auth_method_login_name")
		if err != nil {
			return nil, fmt.Errorf("get param addr failed: %v", err)
		}
		config.PassWord, err = getBoundaryParam(rcfContent, "password_auth_method_password")
		if err != nil {
			return nil, fmt.Errorf("get param addr failed: %v", err)
		}
	}

	// Env var
	if !config.isVaild() {
		config.Addr = os.Getenv("BOUNDARY_ADDR")
		config.AuthMethodId = os.Getenv("AUTH_METHOD_ID")
		config.LoginName = os.Getenv("PASSWORD_AUTH_METHOD_LOGIN_NAME")
		config.PassWord = os.Getenv("PASSWORD_AUTH_METHOD_PASSWORD")
	}

	if !config.isVaild() {
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

// input addr = xx12345
// getBoundaryParam(rcfContent, addr) -> return xx12345
func getBoundaryParam(rcfContent []byte, str string) (string, error) {
	exp, err := regexp.Compile(fmt.Sprintf(`%s\s?=\s?"?\w+\.\w+.\w+"?`, str))
	if err != nil {
		return "", fmt.Errorf("get %s failed: %v", str, err)
	}

	strExp := exp.Find(rcfContent)

	rawToken := strings.Split(string(strExp), "=")
	if len(rawToken) < 1 {
		return "", fmt.Errorf("failed to get boundary str, please set your boundary param correct.")
	}
	result := strings.TrimSpace(strings.Replace(rawToken[1], "\"", "", -1))
	return result, nil
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
