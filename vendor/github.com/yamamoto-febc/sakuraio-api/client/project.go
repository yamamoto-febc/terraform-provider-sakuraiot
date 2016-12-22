package client

import (
	"fmt"
	"github.com/yamamoto-febc/sakuraio-api/gen/client/project"
	"github.com/yamamoto-febc/sakuraio-api/gen/models"
)

var (
	// ProjectSortNameAsc is sort target column `name` that used by [GET /v1/projects]
	ProjectSortNameAsc = "name"

	// ProjectSortIDAsc is sort target column `id` that used by [GET /v1/projects]
	ProjectSortIDAsc = "id"

	// ProjectSortNameDesc is sort target column `-name` that used by [GET /v1/projects]
	ProjectSortNameDesc = "-name"

	// ProjectSortIDDesc is sort target column `-id` that used by [GET /v1/projects]
	ProjectSortIDDesc = "-id"

	// ProjectSortCols is all sort target columns that used by [GET /v1/projects]
	ProjectSortCols = []string{ProjectSortNameAsc, ProjectSortIDAsc, ProjectSortNameDesc, ProjectSortIDDesc}
)

// SakuraAPIFuncProject is interface of the Project functions
type SakuraAPIFuncProject interface {
	// NewFindProjectsParam returns new FindProjectParam
	NewFindProjectsParam() FindProjectParam

	// NewCreateProjectParam returns new CreateProjectParam
	NewCreateProjectParam() CreateProjectParam

	// NewReadProjectParam returns new ReadPorjectParam
	NewReadProjectParam(int64) ReadProjectParam

	// NewUpdateProjectParam returns new UpdatePorjectParam
	NewUpdateProjectParam(int64) UpdateProjectParam

	// NewDeleteProjectParam returns new DeleteProjectParam
	NewDeleteProjectParam(int64) DeleteProjectParam

	// FindProjects calls [GET /v1/projects/] API , returns projects
	FindProjects(FindProjectParam) ([]*models.Project, error)

	// CreateProject calls [POST /v1/projects/] API , returns created project
	CreateProject(CreateProjectParam) (*models.Project, error)

	// ReadProject calls [GET /v1/projects/{projectId}/] API , returns readed project
	ReadProject(ReadProjectParam) (*models.Project, error)

	// UpdateProject calls [PUT /v1/projects/{projectId}] API , returns updated project
	UpdateProject(UpdateProjectParam) (*models.Project, error)

	// DeleteProject calls [DELETE /v1/projects/{projectId}] API
	DeleteProject(DeleteProjectParam) error
}

// FindProjectParam is a parameter interface for calling the find projects API
type FindProjectParam interface {
	SakuraAPIFuncParamBase
	SetName(name *string)
	SetSort(sort *string)
}

// CreateProjectParam is a parameter interface for calling the create project API
type CreateProjectParam interface {
	SakuraAPIFuncParamBase
	SetProject(project *models.ProjectUpdate)
}

// UpdateProjectParam is a parameter interface for calling the update project API
type UpdateProjectParam interface {
	SakuraAPIFuncParamBase
	SetProjectID(projectID int64)
	SetProject(project *models.ProjectUpdate)
}

// ReadProjectParam is a parameter interface for calling the read project API
type ReadProjectParam interface {
	SakuraAPIFuncParamBase
	SetProjectID(projectID int64)
}

// DeleteProjectParam is a parameter interface for calling the delete project API
type DeleteProjectParam interface {
	SakuraAPIFuncParamBase
	SetProjectID(projectID int64)
}

// NewFindProjectsParam returns new FindProjectParam
func (c *SakuraAPIClient) NewFindProjectsParam() FindProjectParam {
	return project.NewGetV1ProjectsParams()
}

// NewCreateProjectParam returns new CreateProjectParam
func (c *SakuraAPIClient) NewCreateProjectParam() CreateProjectParam {
	return project.NewPostV1ProjectsParams()
}

// NewReadProjectParam returns new ReadPorjectParam
func (c *SakuraAPIClient) NewReadProjectParam(id int64) ReadProjectParam {
	p := project.NewGetV1ProjectsProjectIDParams()
	p.SetProjectID(id)
	return p
}

// NewUpdateProjectParam returns new UpdatePorjectParam
func (c *SakuraAPIClient) NewUpdateProjectParam(id int64) UpdateProjectParam {
	p := project.NewPutV1ProjectsProjectIDParams()
	p.SetProjectID(id)
	return p
}

// NewDeleteProjectParam returns new DeleteProjectParam
func (c *SakuraAPIClient) NewDeleteProjectParam(id int64) DeleteProjectParam {
	p := project.NewDeleteV1ProjectsProjectIDParams()
	p.SetProjectID(id)
	return p
}

// FindProjects calls [GET /v1/projects/] API , returns projects
func (c *SakuraAPIClient) FindProjects(param FindProjectParam) ([]*models.Project, error) {

	if _, err := c.isAPIKeyEmpty(); err != nil {
		return nil, fmt.Errorf("Failed on FindProjects(): %s", err)
	}

	if param == nil {
		c.NewFindProjectsParam()
	}

	res, err := c.client.Project.GetV1Projects(param.(*project.GetV1ProjectsParams), *c.basicAuthInfoWriter)
	if err != nil {
		return nil, fmt.Errorf("Failed on FindProjects(): %s", err)
	}
	return res.Payload, nil
}

// CreateProject calls [POST /v1/projects/] API , returns created project
func (c *SakuraAPIClient) CreateProject(param CreateProjectParam) (*models.Project, error) {
	if _, err := c.isAPIKeyEmpty(); err != nil {
		return nil, fmt.Errorf("Failed on CreateProject(): %s", err)
	}

	if param == nil {
		return nil, fmt.Errorf("Failed on CreateProject(): %s", fmt.Errorf("param is required"))
	}

	res, err := c.client.Project.PostV1Projects(param.(*project.PostV1ProjectsParams), *c.basicAuthInfoWriter)
	if err != nil {
		return nil, fmt.Errorf("Failed on CreateProject(): %s", err)
	}
	return res.Payload, nil
}

// ReadProject calls [GET /v1/projects/{projectId}/] API , returns readed project
func (c *SakuraAPIClient) ReadProject(param ReadProjectParam) (*models.Project, error) {
	if _, err := c.isAPIKeyEmpty(); err != nil {
		return nil, fmt.Errorf("Failed on ReadProject(): %s", err)
	}

	if param == nil {
		return nil, fmt.Errorf("Failed on ReadProject(): %s", fmt.Errorf("param is required"))
	}

	res, err := c.client.Project.GetV1ProjectsProjectID(param.(*project.GetV1ProjectsProjectIDParams), *c.basicAuthInfoWriter)
	if err != nil {
		return nil, fmt.Errorf("Failed on ReadProject(): %s", err)
	}
	return res.Payload, nil

}

// UpdateProject calls [PUT /v1/projects/{projectId}] API , returns updated project
func (c *SakuraAPIClient) UpdateProject(param UpdateProjectParam) (*models.Project, error) {
	if _, err := c.isAPIKeyEmpty(); err != nil {
		return nil, fmt.Errorf("Failed on UpdateProject(): %s", err)
	}

	if param == nil {
		return nil, fmt.Errorf("Failed on UpdateProject(): %s", fmt.Errorf("param is required"))
	}

	res, err := c.client.Project.PutV1ProjectsProjectID(param.(*project.PutV1ProjectsProjectIDParams), *c.basicAuthInfoWriter)
	if err != nil {
		return nil, fmt.Errorf("Failed on UpdateProject(): %s", err)
	}
	return res.Payload, nil

}

// DeleteProject calls [DELETE /v1/projects/{projectId}] API
func (c *SakuraAPIClient) DeleteProject(param DeleteProjectParam) error {
	if _, err := c.isAPIKeyEmpty(); err != nil {
		return fmt.Errorf("Failed on DeleteProject(): %s", err)
	}

	if param == nil {
		return fmt.Errorf("Failed on DeleteProject(): %s", fmt.Errorf("param is required"))
	}

	_, _, err := c.client.Project.DeleteV1ProjectsProjectID(param.(*project.DeleteV1ProjectsProjectIDParams), *c.basicAuthInfoWriter)
	if err != nil {
		return fmt.Errorf("Failed on DeleteProject(): %s", err)
	}
	return nil
}
