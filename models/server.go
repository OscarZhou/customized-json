package models

type CustomConfig map[string]interface{}

type Server struct {
	Method        string
	Path          string
	ProxyScheme   string
	ProxyPass     string
	ProxyPassPath string
	APIVersion    string
	CustomConfigs CustomConfig
}
