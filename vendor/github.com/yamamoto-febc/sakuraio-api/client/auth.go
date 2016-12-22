package client

import (
	"fmt"
	"github.com/yamamoto-febc/sakuraio-api/gen/models"
)

// SakuraAPIFuncAuth is interface of Auth function
type SakuraAPIFuncAuth interface {
	// Auth calls [GET /v1/auth/] API, returns *models.Auth object obtained using API token and secret
	Auth() (*models.Auth, error)
}

// Auth returns *models.Auth object obtained using API token and secret
func (c *SakuraAPIClient) Auth() (*models.Auth, error) {
	if _, err := c.isAPIKeyEmpty(); err != nil {
		return nil, fmt.Errorf("Failed on Auth(): %s", err)
	}

	v1AuthOK, err := c.client.Auth.GetV1Auth(nil, *c.basicAuthInfoWriter)
	if err != nil {
		return nil, fmt.Errorf("Failed on Auth(): %s", err)
	}
	return v1AuthOK.Payload, nil
}
