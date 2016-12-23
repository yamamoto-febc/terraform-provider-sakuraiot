package client

import (
	"fmt"
	httptransport "github.com/go-openapi/runtime/client"
)

// SakuraAPIFuncAPIKey is interface of handling APIKey function
type SakuraAPIFuncAPIKey interface {
	// SetAPIKey sets APIKeys(token and secret)
	SetAPIKey(token string, secret string)
	// ClearAPIKey clear APIKey
	ClearAPIKey()
}

// SetAPIKey sets APIKeys(token and secret)
func (c *SakuraAPIClient) SetAPIKey(token string, secret string) {
	writer := httptransport.BasicAuth(token, secret)
	c.basicAuthInfoWriter = &writer
}

// ClearAPIKey clear APIKey
func (c *SakuraAPIClient) ClearAPIKey() {
	c.basicAuthInfoWriter = nil
}

func (c *SakuraAPIClient) isAPIKeyEmpty() (bool, error) {
	if c.basicAuthInfoWriter == nil {
		return true, fmt.Errorf("APIKey is empty , Please set APIKey")
	}

	return false, nil
}
