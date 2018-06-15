package models

type Template struct {
	ControllerPath string
	ServerTemplate ServerTemplate
	// only can be assigned the value automatically
	// Deprivated and replaced with Resources
	fileList []string
	// key is API name and the value indicates
	// that what methods does this API have
	APIs map[string][]string
	// being generated based initialized setting
	Resources []string
}
