package client

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"golang.org/x/net/context"
	"time"
)

// SakuraAPIFuncs is interface of the SakuraIoT APIs
type SakuraAPIFuncs interface {
	SakuraAPIFuncAPIKey
	SakuraAPIFuncAuth
	SakuraAPIFuncProject
	SakuraAPIFuncModule
	SakuraAPIFuncService
}

// SakuraAPIFuncParamBase is interface of API parameter
type SakuraAPIFuncParamBase interface {
	// SetTimeout adds the timeout to the request parameter
	SetTimeout(timeout time.Duration)
	// SetTimeout adds the context to the request parameter
	SetContext(ctx context.Context)
	// WriteToRequest writes these params to a swagger request
	WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error
}
