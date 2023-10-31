package config

type Application struct {
	Host          string
	Port          int64
	Name          string
	Version       string
	Mode          string
	EnableDP      bool
	FileRootPath  string
	AmpKey        string
	IsSingleLogin bool
	Author        string
}

var ApplicationConfig = new(Application)
