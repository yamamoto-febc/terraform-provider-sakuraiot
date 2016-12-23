package client

import (
	"fmt"
	"github.com/yamamoto-febc/sakuraio-api/gen/client/module"
	"github.com/yamamoto-febc/sakuraio-api/gen/models"
)

var (
	// ModuleSortProjectAsc is sort target column `project` that used by [GET /v1/modules]
	ModuleSortProjectAsc = "project"
	// ModuleSortNameAsc is sort target column `name` that used by [GET /v1/modules]
	ModuleSortNameAsc = "name"
	// ModuleSortIDAsc is sort target column `id` that used by [GET /v1/modules]
	ModuleSortIDAsc = "id"
	// ModuleSortSerialNumberAsc is sort target column `serial_number` that used by [GET /v1/modules]
	ModuleSortSerialNumberAsc = "serial_number"
	// ModuleSortModelAsc is sort target column `model` that used by [GET /v1/modules]
	ModuleSortModelAsc = "model"
	// ModuleSortProjectDesc is sort target column `-project` that used by [GET /v1/modules]
	ModuleSortProjectDesc = "-project"
	// ModuleSortNameDesc is sort target column `-name` that used by [GET /v1/modules]
	ModuleSortNameDesc = "-name"
	// ModuleSortIDDesc is sort target column `-id` that used by [GET /v1/modules]
	ModuleSortIDDesc = "-id"
	// ModuleSortSerialNumberDesc is sort target column `-serial_number` that used by [GET /v1/modules]
	ModuleSortSerialNumberDesc = "-serial_number"
	// ModuleSortModelDesc is sort target column `-model` that used by [GET /v1/modules]
	ModuleSortModelDesc = "-model"

	// ModuleSortCols is all sort target columns that used by [GET /v1/modules]
	ModuleSortCols = []string{
		ModuleSortProjectAsc, ModuleSortNameAsc, ModuleSortIDAsc, ModuleSortSerialNumberAsc, ModuleSortModelAsc,
		ModuleSortProjectDesc, ModuleSortNameDesc, ModuleSortIDDesc, ModuleSortSerialNumberDesc, ModuleSortModelDesc,
	}
)

// SakuraAPIFuncModule is interface of the Module functions
type SakuraAPIFuncModule interface {
	// NewFindModulesParam returns new FindModuleParam
	NewFindModulesParam() FindModuleParam

	// NewCreateModuleParam return new CreateModuleParam
	NewCreateModuleParam() CreateModuleParam

	// NewReadModuleParam returns new ReadModuleParam
	NewReadModuleParam(string) ReadModuleParam

	// NewUpdateModuleParam returns new UpdateModuleParam
	NewUpdateModuleParam(string) UpdateModuleParam

	// NewDeleteModuleParam returns DeleteModuleParam
	NewDeleteModuleParam(string) DeleteModuleParam

	// FindModules calls [GET /v1/modules/] API , returns Modules
	FindModules(FindModuleParam) ([]*models.Module, error)

	// CreateModule calls [POST /v1/modules/] API, returns created Module
	CreateModule(CreateModuleParam) (*models.Module, error)

	// ReadModule calls [GET /v1/modules/{moduleId}/] API, returns readed Module
	ReadModule(ReadModuleParam) (*models.Module, error)

	// UpdateModule calls [PUT /v1/modules/{moduleId}/] API, returns updated Module
	UpdateModule(UpdateModuleParam) (*models.Module, error)

	// DeleteModule calls [DELETE /v1/modules/{moduleId}/] API
	DeleteModule(DeleteModuleParam) error
}

// FindModuleParam is a parameter interface for calling the find modules API
type FindModuleParam interface {
	SakuraAPIFuncParamBase
	SetModel(model *string)
	SetProject(project *int64)
	SetSerialNumber(serialNumber *string)
	SetSort(sort *string)
}

// CreateModuleParam is a parameter interface for calling the create module API
type CreateModuleParam interface {
	SakuraAPIFuncParamBase
	SetRegisterInformation(registerInformation *models.ModuleRegister)
}

// ReadModuleParam is a parameter interface for calling the read module API
type ReadModuleParam interface {
	SakuraAPIFuncParamBase
	SetModuleID(moduleID string)
}

// UpdateModuleParam is a parameter interface for calling the update module API
type UpdateModuleParam interface {
	SakuraAPIFuncParamBase
	SetModuleID(moduleID string)
	SetBody(body *models.ModuleUpdate)
}

// DeleteModuleParam is a parameter interface for calling the delete module API
type DeleteModuleParam interface {
	SakuraAPIFuncParamBase
	SetModuleID(moduleID string)
}

// NewFindModulesParam returns new FindModuleParam
func (c *SakuraAPIClient) NewFindModulesParam() FindModuleParam {
	return module.NewGetV1ModulesParams()
}

// NewCreateModuleParam return new CreateModuleParam
func (c *SakuraAPIClient) NewCreateModuleParam() CreateModuleParam {
	return module.NewPostV1ModulesParams()
}

// NewReadModuleParam returns new ReadModuleParam
func (c *SakuraAPIClient) NewReadModuleParam(id string) ReadModuleParam {
	p := module.NewGetV1ModulesModuleIDParams()
	p.SetModuleID(id)
	return p
}

// NewUpdateModuleParam returns new UpdateModuleParam
func (c *SakuraAPIClient) NewUpdateModuleParam(id string) UpdateModuleParam {
	p := module.NewPutV1ModulesModuleIDParams()
	p.SetModuleID(id)
	return p
}

// NewDeleteModuleParam returns DeleteModuleParam
func (c *SakuraAPIClient) NewDeleteModuleParam(id string) DeleteModuleParam {
	p := module.NewDeleteV1ModulesModuleIDParams()
	p.SetModuleID(id)
	return p
}

// FindModules calls [GET /v1/modules/] API , returns Modules
func (c *SakuraAPIClient) FindModules(param FindModuleParam) ([]*models.Module, error) {

	if _, err := c.isAPIKeyEmpty(); err != nil {
		return nil, fmt.Errorf("Failed on FindModules(): %s", err)
	}

	if param == nil {
		c.NewFindModulesParam()
	}

	res, err := c.client.Module.GetV1Modules(param.(*module.GetV1ModulesParams), *c.basicAuthInfoWriter)
	if err != nil {
		return nil, fmt.Errorf("Failed on FindModules(): %s", err)
	}
	return res.Payload, nil
}

// CreateModule calls [POST /v1/modules/] API, returns created Module
func (c *SakuraAPIClient) CreateModule(param CreateModuleParam) (*models.Module, error) {
	if _, err := c.isAPIKeyEmpty(); err != nil {
		return nil, fmt.Errorf("Failed on CreateModule(): %s", err)
	}

	if param == nil {
		return nil, fmt.Errorf("Failed on CreateModule(): %s", fmt.Errorf("param is required"))
	}

	res, err := c.client.Module.PostV1Modules(param.(*module.PostV1ModulesParams), *c.basicAuthInfoWriter)
	if err != nil {
		return nil, fmt.Errorf("Failed on CreateModule(): %s", err)
	}
	return res.Payload, nil
}

// ReadModule calls [GET /v1/modules/{moduleId}/] API, returns readed Module
func (c *SakuraAPIClient) ReadModule(param ReadModuleParam) (*models.Module, error) {
	if _, err := c.isAPIKeyEmpty(); err != nil {
		return nil, fmt.Errorf("Failed on ReadModule(): %s", err)
	}

	if param == nil {
		return nil, fmt.Errorf("Failed on ReadModule(): %s", fmt.Errorf("param is required"))
	}

	res, err := c.client.Module.GetV1ModulesModuleID(param.(*module.GetV1ModulesModuleIDParams), *c.basicAuthInfoWriter)
	if err != nil {
		return nil, fmt.Errorf("Failed on ReadModule(): %s", err)
	}
	return res.Payload, nil

}

// UpdateModule calls [PUT /v1/modules/{moduleId}/] API, returns updated Module
func (c *SakuraAPIClient) UpdateModule(param UpdateModuleParam) (*models.Module, error) {
	if _, err := c.isAPIKeyEmpty(); err != nil {
		return nil, fmt.Errorf("Failed on UpdateModule(): %s", err)
	}

	if param == nil {
		return nil, fmt.Errorf("Failed on UpdateModule(): %s", fmt.Errorf("param is required"))
	}

	res, err := c.client.Module.PutV1ModulesModuleID(param.(*module.PutV1ModulesModuleIDParams), *c.basicAuthInfoWriter)
	if err != nil {
		return nil, fmt.Errorf("Failed on UpdateModule(): %s", err)
	}
	return res.Payload, nil

}

// DeleteModule calls [DELETE /v1/modules/{moduleId}/] API
func (c *SakuraAPIClient) DeleteModule(param DeleteModuleParam) error {
	if _, err := c.isAPIKeyEmpty(); err != nil {
		return fmt.Errorf("Failed on DeleteModule(): %s", err)
	}

	if param == nil {
		return fmt.Errorf("Failed on DeleteModule(): %s", fmt.Errorf("param is required"))
	}

	_, _, err := c.client.Module.DeleteV1ModulesModuleID(param.(*module.DeleteV1ModulesModuleIDParams), *c.basicAuthInfoWriter)
	if err != nil {
		return fmt.Errorf("Failed on DeleteModule(): %s", err)
	}
	return nil
}
