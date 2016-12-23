package client

import (
	"fmt"
	"github.com/yamamoto-febc/sakuraio-api/gen/client/service"
	"github.com/yamamoto-febc/sakuraio-api/gen/models"
)

var (
	// ServiceSortIDAsc is sort target column `id` that used by [GET /v1/services]
	ServiceSortIDAsc = "id"
	// ServiceSortNameAsc is sort target column `name` that used by [GET /v1/services]
	ServiceSortNameAsc = "name"
	// ServiceSortTokenAsc is sort target column `token` that used by [GET /v1/services]
	ServiceSortTokenAsc = "token"
	// ServiceSortIDDesc is sort target column `-id` that used by [GET /v1/services]
	ServiceSortIDDesc = "-id"
	// ServiceSortNameDesc is sort target column `-name` that used by [GET /v1/services]
	ServiceSortNameDesc = "-name"
	// ServiceSortTokenDesc is sort target column `-token` that used by [GET /v1/services]
	ServiceSortTokenDesc = "-token"

	// ServiceSortCols is all sort target columns that used by [GET /v1/services]
	ServiceSortCols = []string{
		ServiceSortIDAsc, ServiceSortNameAsc, ServiceSortTokenAsc,
		ServiceSortIDDesc, ServiceSortNameDesc, ServiceSortTokenDesc,
	}
)

// SakuraAPIFuncService is interface of the Service functions
type SakuraAPIFuncService interface {
	// NewFindServicesParam return new FindServiceParam
	NewFindServicesParam() FindServiceParam

	// NewCreateWebSocketServiceParam returns new CreateWebSocketServiceParam
	NewCreateWebSocketServiceParam() CreateWebSocketServiceParam

	// NewCreateOutgoingWebhookServiceParam returns new CreateOutgoingWebhookServiceParam
	NewCreateOutgoingWebhookServiceParam() CreateOutgoingWebhookServiceParam

	// NewCreateIncomingWebhookServiceParam returns new CreateIncomingWebhookServiceParam
	NewCreateIncomingWebhookServiceParam() CreateIncomingWebhookServiceParam

	// NewReadWebSocketServiceParam returns new ReadWebSocketServiceParam
	NewReadWebSocketServiceParam(int64) ReadWebSocketServiceParam

	// NewReadOutgoingWebhookServiceParam returns new ReadOutgoingWebhookServiceParam
	NewReadOutgoingWebhookServiceParam(int64) ReadOutgoingWebhookServiceParam

	// NewReadIncomingWebhookServiceParam returns new ReadIncomingWebhookServiceParam
	NewReadIncomingWebhookServiceParam(int64) ReadIncomingWebhookServiceParam

	// NewUpdateWebSocketServiceParam returns new UpdateWebSocketServiceParam
	NewUpdateWebSocketServiceParam(int64) UpdateWebSocketServiceParam

	// NewUpdateOutgoingWebhookServiceParam returns new UpdateOutgoingWebhookServiceParam
	NewUpdateOutgoingWebhookServiceParam(int64) UpdateOutgoingWebhookServiceParam

	// NewUpdateIncomingWebhookServiceParam returns new UpdateIncomingWebhookServiceParam
	NewUpdateIncomingWebhookServiceParam(int64) UpdateIncomingWebhookServiceParam

	// NewDeleteServiceParam returns new DeleteServiceParam
	NewDeleteServiceParam(int64) DeleteServiceParam

	// FindServices calls [GET /v1/services/] API , returns services([]ServiceBase)
	FindServices(FindServiceParam) ([]*models.ServiceBase, error)

	// CreateWebSocketService calls [POST /v1/services/?type=websocket] API , returns created WebSocketService
	CreateWebSocketService(CreateWebSocketServiceParam) (*models.WebSocketService, error)

	// CreateOutgoingWebhookService calls [POST /v1/services/?type=outgoing-webhook] API , returns created OutgoingWebhookService
	CreateOutgoingWebhookService(CreateOutgoingWebhookServiceParam) (*models.OutgoingWebhookService, error)

	// CreateIncomingWebhookService calls [POST /v1/services/?type=incoming-webhook] API , returns created IncomingWebhookService
	CreateIncomingWebhookService(CreateIncomingWebhookServiceParam) (*models.IncomingWebhookService, error)

	// ReadWebSocketService calls [GET /v1/services/{serviceId}/] API , returns readed WebSocketService
	ReadWebSocketService(ReadWebSocketServiceParam) (*models.WebSocketService, error)

	// ReadOutgoingWebhookService calls [GET /v1/services/{serviceId}/] API , returns readed OutgoingWebhookService
	ReadOutgoingWebhookService(ReadOutgoingWebhookServiceParam) (*models.OutgoingWebhookService, error)

	// ReadIncomingWebhookService calls [GET /v1/services/{serviceId}/] API , returns readed IncomingWebhookService
	ReadIncomingWebhookService(ReadIncomingWebhookServiceParam) (*models.IncomingWebhookService, error)

	// UpdateWebSocketService calls [PUT /v1/services/{serviceId}] API , returns updated WebSocketService
	UpdateWebSocketService(UpdateWebSocketServiceParam) (*models.WebSocketService, error)

	// UpdateOutgoingWebhookService calls [PUT /v1/services/{serviceId}] API , returns updated OutgoingWebhookService
	UpdateOutgoingWebhookService(UpdateOutgoingWebhookServiceParam) (*models.OutgoingWebhookService, error)

	// UpdateIncomingWebhookService calls [PUT /v1/services/{serviceId}] API , returns updated IncomingWebhookService
	UpdateIncomingWebhookService(UpdateIncomingWebhookServiceParam) (*models.IncomingWebhookService, error)

	// DeleteService calls [DELETE /v1/services/{serviceId}] API
	DeleteService(DeleteServiceParam) error
}

// FindServiceParam is a parameter interface for calling the find services API
type FindServiceParam interface {
	SakuraAPIFuncParamBase
	SetProject(project *int64)
	SetSort(sort *string)
	SetType(typeVar *string)
}

// CreateWebSocketServiceParam is a parameter interface for calling the create WebSocketService API
type CreateWebSocketServiceParam interface {
	SakuraAPIFuncParamBase
	SetWebSocketService(*models.WebSocketServiceUpdate)
}

// CreateOutgoingWebhookServiceParam is a parameter interface for calling the create OutgoingWebhookService API
type CreateOutgoingWebhookServiceParam interface {
	SakuraAPIFuncParamBase
	SetOutgoingWebhookService(*models.OutgoingWebhookServiceUpdate)
}

// CreateIncomingWebhookServiceParam is a parameter interface for calling the create IncomingWebhookService API
type CreateIncomingWebhookServiceParam interface {
	SakuraAPIFuncParamBase
	SetIncomingWebhookService(*models.IncomingWebhookServiceUpdate)
}

// UpdateWebSocketServiceParam is a parameter interface for calling the update WebSocketService API
type UpdateWebSocketServiceParam interface {
	SakuraAPIFuncParamBase
	SetServiceID(serviceID int64)
	SetWebSocketService(webSocketService *models.WebSocketServiceUpdate)
}

// UpdateOutgoingWebhookServiceParam is a parameter interface for calling the update OutgoingWebhookService API
type UpdateOutgoingWebhookServiceParam interface {
	SakuraAPIFuncParamBase
	SetServiceID(serviceID int64)
	SetOutgoingWebhookService(outgoingWebhookService *models.OutgoingWebhookServiceUpdate)
}

// UpdateIncomingWebhookServiceParam is a parameter interface for calling the update IncomingWebhookService API
type UpdateIncomingWebhookServiceParam interface {
	SakuraAPIFuncParamBase
	SetServiceID(serviceID int64)
	SetIncomingWebhookService(incomingWebhookService *models.IncomingWebhookServiceUpdate)
}

// ReadWebSocketServiceParam is a parameter interface for calling the read WebSocketService API
type ReadWebSocketServiceParam interface {
	SakuraAPIFuncParamBase
	SetServiceID(serviceID int64)
}

// ReadOutgoingWebhookServiceParam is a parameter interface for calling the read OutgoingWebhookService API
type ReadOutgoingWebhookServiceParam interface {
	SakuraAPIFuncParamBase
	SetServiceID(serviceID int64)
}

// ReadIncomingWebhookServiceParam is a parameter interface for calling the read IncomingWebhookService API
type ReadIncomingWebhookServiceParam interface {
	SakuraAPIFuncParamBase
	SetServiceID(serviceID int64)
}

// DeleteServiceParam is a parameter interface for calling the delete service API
type DeleteServiceParam interface {
	SakuraAPIFuncParamBase
	SetServiceID(serviceID int64)
}

// NewFindServicesParam return new FindServiceParam
func (c *SakuraAPIClient) NewFindServicesParam() FindServiceParam {
	return service.NewGetV1ServicesParams()
}

// NewCreateWebSocketServiceParam returns new CreateWebSocketServiceParam
func (c *SakuraAPIClient) NewCreateWebSocketServiceParam() CreateWebSocketServiceParam {
	return service.NewPostV1ServicesTypeWebsocketParams()
}

// NewCreateOutgoingWebhookServiceParam returns new CreateOutgoingWebhookServiceParam
func (c *SakuraAPIClient) NewCreateOutgoingWebhookServiceParam() CreateOutgoingWebhookServiceParam {
	return service.NewPostV1ServicesTypeOutgoingWebhookParams()
}

// NewCreateIncomingWebhookServiceParam returns new CreateIncomingWebhookServiceParam
func (c *SakuraAPIClient) NewCreateIncomingWebhookServiceParam() CreateIncomingWebhookServiceParam {
	return service.NewPostV1ServicesTypeIncomingWebhookParams()
}

// NewReadWebSocketServiceParam returns new ReadWebSocketServiceParam
func (c *SakuraAPIClient) NewReadWebSocketServiceParam(id int64) ReadWebSocketServiceParam {
	p := service.NewGetV1ServicesServiceIDTypeWebsocketParams()
	p.SetServiceID(id)
	return p
}

// NewReadOutgoingWebhookServiceParam returns new ReadOutgoingWebhookServiceParam
func (c *SakuraAPIClient) NewReadOutgoingWebhookServiceParam(id int64) ReadOutgoingWebhookServiceParam {
	p := service.NewGetV1ServicesServiceIDTypeOutgoingWebhookParams()
	p.SetServiceID(id)
	return p
}

// NewReadIncomingWebhookServiceParam returns new ReadIncomingWebhookServiceParam
func (c *SakuraAPIClient) NewReadIncomingWebhookServiceParam(id int64) ReadIncomingWebhookServiceParam {
	p := service.NewGetV1ServicesServiceIDTypeIncomingWebhookParams()
	p.SetServiceID(id)
	return p
}

// NewUpdateWebSocketServiceParam returns new UpdateWebSocketServiceParam
func (c *SakuraAPIClient) NewUpdateWebSocketServiceParam(id int64) UpdateWebSocketServiceParam {
	p := service.NewPutV1ServicesServiceIDTypeWebsocketParams()
	p.SetServiceID(id)
	return p
}

// NewUpdateOutgoingWebhookServiceParam returns new UpdateOutgoingWebhookServiceParam
func (c *SakuraAPIClient) NewUpdateOutgoingWebhookServiceParam(id int64) UpdateOutgoingWebhookServiceParam {
	p := service.NewPutV1ServicesServiceIDTypeOutgoingWebhookParams()
	p.SetServiceID(id)
	return p
}

// NewUpdateIncomingWebhookServiceParam returns new UpdateIncomingWebhookServiceParam
func (c *SakuraAPIClient) NewUpdateIncomingWebhookServiceParam(id int64) UpdateIncomingWebhookServiceParam {
	p := service.NewPutV1ServicesServiceIDTypeIncomingWebhookParams()
	p.SetServiceID(id)
	return p
}

// NewDeleteServiceParam returns new DeleteServiceParam
func (c *SakuraAPIClient) NewDeleteServiceParam(id int64) DeleteServiceParam {
	p := service.NewDeleteV1ServicesServiceIDParams()
	p.SetServiceID(id)
	return p
}

// FindServices calls [GET /v1/services/] API , returns services([]ServiceBase)
func (c *SakuraAPIClient) FindServices(param FindServiceParam) ([]*models.ServiceBase, error) {

	if _, err := c.isAPIKeyEmpty(); err != nil {
		return nil, fmt.Errorf("Failed on FindServices(): %s", err)
	}

	if param == nil {
		c.NewFindServicesParam()
	}

	res, err := c.client.Service.GetV1Services(param.(*service.GetV1ServicesParams), *c.basicAuthInfoWriter)
	if err != nil {
		return nil, fmt.Errorf("Failed on FindServices(): %s", err)
	}
	return res.Payload, nil
}

// CreateWebSocketService calls [POST /v1/services/?type=websocket] API , returns created WebSocketService
func (c *SakuraAPIClient) CreateWebSocketService(param CreateWebSocketServiceParam) (*models.WebSocketService, error) {
	if _, err := c.isAPIKeyEmpty(); err != nil {
		return nil, fmt.Errorf("Failed on CreateWebSocketService(): %s", err)
	}

	if param == nil {
		return nil, fmt.Errorf("Failed on CreateWebSocketService(): %s", fmt.Errorf("param is required"))
	}

	webSocketParam := param.(*service.PostV1ServicesTypeWebsocketParams)
	if webSocketParam.WebSocketService == nil {
		return nil, fmt.Errorf("Failed on CreateWebSocketService(): %s", fmt.Errorf("WebSocketService is required"))
	}
	serviceType := "websocket"
	webSocketParam.WebSocketService.Type = &serviceType

	res, err := c.client.Service.PostV1ServicesTypeWebsocket(webSocketParam, *c.basicAuthInfoWriter)
	if err != nil {
		return nil, fmt.Errorf("Failed on CreateWebSocketService(): %s", err)
	}
	return res.Payload, nil
}

// CreateOutgoingWebhookService calls [POST /v1/services/?type=outgoing-webhook] API , returns created OutgoingWebhookService
func (c *SakuraAPIClient) CreateOutgoingWebhookService(param CreateOutgoingWebhookServiceParam) (*models.OutgoingWebhookService, error) {
	if _, err := c.isAPIKeyEmpty(); err != nil {
		return nil, fmt.Errorf("Failed on CreateOutgoingWebhookService(): %s", err)
	}

	if param == nil {
		return nil, fmt.Errorf("Failed on CreateOutgoingWebhookService(): %s", fmt.Errorf("param is required"))
	}

	webhookParam := param.(*service.PostV1ServicesTypeOutgoingWebhookParams)
	if webhookParam.OutgoingWebhookService == nil {
		return nil, fmt.Errorf("Failed on CreateOutgoingWebhookService(): %s", fmt.Errorf("OutgoingWebhookService is required"))
	}
	serviceType := "outgoing-webhook"
	webhookParam.OutgoingWebhookService.Type = &serviceType

	res, err := c.client.Service.PostV1ServicesTypeOutgoingWebhook(webhookParam, *c.basicAuthInfoWriter)
	if err != nil {
		return nil, fmt.Errorf("Failed on CreateOutgoingWebhookService(): %s", err)
	}
	return res.Payload, nil
}

// CreateIncomingWebhookService calls [POST /v1/services/?type=incoming-webhook] API , returns created IncomingWebhookService
func (c *SakuraAPIClient) CreateIncomingWebhookService(param CreateIncomingWebhookServiceParam) (*models.IncomingWebhookService, error) {
	if _, err := c.isAPIKeyEmpty(); err != nil {
		return nil, fmt.Errorf("Failed on CreateIncomingWebhookService(): %s", err)
	}

	if param == nil {
		return nil, fmt.Errorf("Failed on CreateIncomingWebhookService(): %s", fmt.Errorf("param is required"))
	}

	webhookParam := param.(*service.PostV1ServicesTypeIncomingWebhookParams)
	if webhookParam.IncomingWebhookService == nil {
		return nil, fmt.Errorf("Failed on CreateIncomingWebhookService(): %s", fmt.Errorf("CreateIncomingWebhookService is required"))
	}
	serviceType := "incoming-webhook"
	webhookParam.IncomingWebhookService.Type = &serviceType

	res, err := c.client.Service.PostV1ServicesTypeIncomingWebhook(webhookParam, *c.basicAuthInfoWriter)
	if err != nil {
		return nil, fmt.Errorf("Failed on CreateIncomingWebhookService(): %s", err)
	}
	return res.Payload, nil
}

// ReadWebSocketService calls [GET /v1/services/{serviceId}/] API , returns readed WebSocketService
func (c *SakuraAPIClient) ReadWebSocketService(param ReadWebSocketServiceParam) (*models.WebSocketService, error) {
	if _, err := c.isAPIKeyEmpty(); err != nil {
		return nil, fmt.Errorf("Failed on ReadWebSocketService(): %s", err)
	}

	if param == nil {
		return nil, fmt.Errorf("Failed on ReadWebSocketService(): %s", fmt.Errorf("param is required"))
	}

	res, err := c.client.Service.GetV1ServicesServiceIDTypeWebsocket(param.(*service.GetV1ServicesServiceIDTypeWebsocketParams), *c.basicAuthInfoWriter)
	if err != nil {
		return nil, fmt.Errorf("Failed on ReadWebSocketService(): %s", err)
	}
	return res.Payload, nil

}

// ReadOutgoingWebhookService calls [GET /v1/services/{serviceId}/] API , returns readed OutgoingWebhookService
func (c *SakuraAPIClient) ReadOutgoingWebhookService(param ReadOutgoingWebhookServiceParam) (*models.OutgoingWebhookService, error) {
	if _, err := c.isAPIKeyEmpty(); err != nil {
		return nil, fmt.Errorf("Failed on ReadOutgoingWebhookService(): %s", err)
	}

	if param == nil {
		return nil, fmt.Errorf("Failed on ReadOutgoingWebhookService(): %s", fmt.Errorf("param is required"))
	}

	res, err := c.client.Service.GetV1ServicesServiceIDTypeOutgoingWebhook(param.(*service.GetV1ServicesServiceIDTypeOutgoingWebhookParams), *c.basicAuthInfoWriter)
	if err != nil {
		return nil, fmt.Errorf("Failed on ReadOutgoingWebhookService(): %s", err)
	}
	return res.Payload, nil

}

// ReadIncomingWebhookService calls [GET /v1/services/{serviceId}/] API , returns readed IncomingWebhookService
func (c *SakuraAPIClient) ReadIncomingWebhookService(param ReadIncomingWebhookServiceParam) (*models.IncomingWebhookService, error) {
	if _, err := c.isAPIKeyEmpty(); err != nil {
		return nil, fmt.Errorf("Failed on ReadIncomingWebhookService(): %s", err)
	}

	if param == nil {
		return nil, fmt.Errorf("Failed on ReadIncomingWebhookService(): %s", fmt.Errorf("param is required"))
	}

	res, err := c.client.Service.GetV1ServicesServiceIDTypeIncomingWebhook(param.(*service.GetV1ServicesServiceIDTypeIncomingWebhookParams), *c.basicAuthInfoWriter)
	if err != nil {
		return nil, fmt.Errorf("Failed on ReadIncomingWebhookService(): %s", err)
	}
	return res.Payload, nil

}

// UpdateWebSocketService calls [PUT /v1/services/{serviceId}] API , returns updated WebSocketService
func (c *SakuraAPIClient) UpdateWebSocketService(param UpdateWebSocketServiceParam) (*models.WebSocketService, error) {
	if _, err := c.isAPIKeyEmpty(); err != nil {
		return nil, fmt.Errorf("Failed on UpdateWebSocketService(): %s", err)
	}

	if param == nil {
		return nil, fmt.Errorf("Failed on UpdateWebSocketService(): %s", fmt.Errorf("param is required"))
	}

	res, err := c.client.Service.PutV1ServicesServiceIDTypeWebsocket(param.(*service.PutV1ServicesServiceIDTypeWebsocketParams), *c.basicAuthInfoWriter)
	if err != nil {
		return nil, fmt.Errorf("Failed on UpdateWebSocketService(): %s", err)
	}
	return res.Payload, nil

}

// UpdateOutgoingWebhookService calls [PUT /v1/services/{serviceId}] API , returns updated OutgoingWebhookService
func (c *SakuraAPIClient) UpdateOutgoingWebhookService(param UpdateOutgoingWebhookServiceParam) (*models.OutgoingWebhookService, error) {
	if _, err := c.isAPIKeyEmpty(); err != nil {
		return nil, fmt.Errorf("Failed on UpdateOutgoingWebhookService(): %s", err)
	}

	if param == nil {
		return nil, fmt.Errorf("Failed on UpdateOutgoingWebhookService(): %s", fmt.Errorf("param is required"))
	}

	res, err := c.client.Service.PutV1ServicesServiceIDTypeOutgoingWebhook(param.(*service.PutV1ServicesServiceIDTypeOutgoingWebhookParams), *c.basicAuthInfoWriter)
	if err != nil {
		return nil, fmt.Errorf("Failed on UpdateOutgoingWebhookService(): %s", err)
	}
	return res.Payload, nil

}

// UpdateIncomingWebhookService calls [PUT /v1/services/{serviceId}] API , returns updated IncomingWebhookService
func (c *SakuraAPIClient) UpdateIncomingWebhookService(param UpdateIncomingWebhookServiceParam) (*models.IncomingWebhookService, error) {
	if _, err := c.isAPIKeyEmpty(); err != nil {
		return nil, fmt.Errorf("Failed on UpdateIncomingWebhookService(): %s", err)
	}

	if param == nil {
		return nil, fmt.Errorf("Failed on UpdateIncomingWebhookService(): %s", fmt.Errorf("param is required"))
	}

	res, err := c.client.Service.PutV1ServicesServiceIDTypeIncomingWebhook(param.(*service.PutV1ServicesServiceIDTypeIncomingWebhookParams), *c.basicAuthInfoWriter)
	if err != nil {
		return nil, fmt.Errorf("Failed on UpdateIncomingWebhookService(): %s", err)
	}
	return res.Payload, nil

}

// DeleteService calls [DELETE /v1/services/{serviceId}] API
func (c *SakuraAPIClient) DeleteService(param DeleteServiceParam) error {
	if _, err := c.isAPIKeyEmpty(); err != nil {
		return fmt.Errorf("Failed on DeleteService(): %s", err)
	}

	if param == nil {
		return fmt.Errorf("Failed on DeleteService(): %s", fmt.Errorf("param is required"))
	}

	_, _, err := c.client.Service.DeleteV1ServicesServiceID(param.(*service.DeleteV1ServicesServiceIDParams), *c.basicAuthInfoWriter)
	if err != nil {
		return fmt.Errorf("Failed on DeleteService(): %s", err)
	}
	return nil
}
