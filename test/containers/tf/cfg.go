package tf

import (
	"fmt"

	"github.com/bigmikesolutions/wingman/test/containers"
)

type Config struct {
	Env     string
	Region  string
	Cognito CognitoConfig
}

type CognitoConfig struct {
	Env           string
	AppClientName string
	UserPoolName  string
	IdEndpoint    string
}

func NewCfg(cfg containers.LocalstackCfg) Config {
	return Config{
		Env:    "localstack",
		Region: cfg.Region,
		Cognito: CognitoConfig{
			IdEndpoint:    fmt.Sprintf("http://localhost:%d/v1/", cfg.Port),
			UserPoolName:  "wingman",
			AppClientName: "localstack",
		},
	}
}

func (c *Config) Vars() map[string]any {
	return map[string]any{
		"env":                     c.Env,
		"aws_region":              c.Region,
		"aws_access_key":          "", // not needed for localstack
		"aws_secret_key":          "", // not needed for localstack
		"cognito_idp_endpoint":    c.Cognito.IdEndpoint,
		"cognito_user_pool_name":  c.Cognito.UserPoolName,
		"cognito_app_client_name": c.Cognito.AppClientName,
	}
}
