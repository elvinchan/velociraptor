package external_filter

import (
	"www.velocidex.com/golang/velociraptor/services"
)

const ExternalPermissionURL = "http://localhost:5000/api/v1/"

type Resource struct {
	Type  string // 资源类型，eg：HuntId
	Value string // 资源值（资源名或资源ID）
}

type ExternalFilterRunner struct {

	// key: username, value: resource
	cache map[string]map[string]string
}

func NewExternalFilterService() *ExternalFilterRunner {
	return &ExternalFilterRunner{
		cache: make(map[string]map[string]string),
	}
}

func (ExternalFilterRunner) FilterResource(username string,
	resourceType services.ExternalResourceType, resourceId string) bool {
	return true
}

func (ExternalFilterRunner) GetResource(username string,
	resourceType services.ExternalResourceType) []string {
	return []string{}
}

func (ExternalFilterRunner) RecordResource(username string,
	resourceType services.ExternalResourceType, resourceId string) error {
	return nil
}

func init() {
	services.RegisterExternalFilter(NewExternalFilterService())
}
