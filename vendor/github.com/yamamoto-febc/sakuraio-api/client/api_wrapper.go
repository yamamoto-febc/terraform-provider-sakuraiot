package client

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/yamamoto-febc/sakuraio-api/gen/client"
)

// SakuraAPI SakuraIoT HTTP client.
var SakuraAPI = NewSakuraClient(nil)

// NewSakuraClient creates a new SakuraIoT HTTP client.
func NewSakuraClient(formats strfmt.Registry) SakuraAPIFuncs {
	if formats == nil {
		formats = strfmt.Default
	}
	transport := httptransport.New("api.sakura.io", "/", []string{"https"})
	return NewSakuraAPI(transport, formats)
}

// NewSakuraAPI creates a new SakuraIoT client
func NewSakuraAPI(transport runtime.ClientTransport, formats strfmt.Registry) SakuraAPIFuncs {

	cli := client.New(transport, formats)
	sakura := &SakuraAPIClient{
		client: cli,
	}
	return sakura
}

// SakuraAPIClient is a client for sakura IoT
type SakuraAPIClient struct {
	client              *client.SakuraIoT
	basicAuthInfoWriter *runtime.ClientAuthInfoWriter
}
