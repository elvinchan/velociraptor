package services

type ExternalResourceType string

const (
	ExternalResourceTypeClient    ExternalResourceType = "client"
	ExternalResourceTypeFileStore ExternalResourceType = "file_store"
	ExternalResourceTypeHunt      ExternalResourceType = "hunt"
	ExternalResourceTypeNotebook  ExternalResourceType = "notebook"
)

type ExternalFilter interface {
	FilterResource(
		username string,
		resourceType ExternalResourceType, resourceId string) bool
	GetResource(
		username string,
		resourceType ExternalResourceType) []string
	RecordResource(username string,
		resourceType ExternalResourceType, resourceId string) error
}

var global_external_filter ExternalFilter

func RegisterExternalFilter(dispatcher ExternalFilter) {
	mu.Lock()
	defer mu.Unlock()

	global_external_filter = dispatcher
}

func GetExternalFilter() ExternalFilter {
	mu.Lock()
	defer mu.Unlock()

	return global_external_filter
}
