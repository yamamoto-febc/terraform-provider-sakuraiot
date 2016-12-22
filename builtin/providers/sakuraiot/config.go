package sakuraiot

import (
	"fmt"
	"github.com/yamamoto-febc/sakuraio-api/client"
	"github.com/yamamoto-febc/sakuraio-api/gen/models"
)

const (
	APITokenParamName  = "SAKURAIOT_AUTH_TOKEN"
	APISecretParamName = "SAKURAIOT_AUTH_SECRET"
)

type Config struct {
	Token  string
	Secret string
}

func (c *Config) NewClient() (*SakuraClient, error) {

	client := client.SakuraAPI

	client.SetAPIKey(c.Token, c.Secret)

	auth, err := client.Auth()
	if err != nil {
		return nil, fmt.Errorf("Failed on NewClient#Auth():%s", err)
	}

	return &SakuraClient{
		SakuraAPIFuncs: client,
		AuthInfo:       auth,
	}, nil
}

type SakuraClient struct {
	client.SakuraAPIFuncs
	AuthInfo *models.Auth
}
