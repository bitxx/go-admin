package config

type Auth struct {
	Timeout           int
	Secret            string
	EnableDeviceCheck bool
	EnableBlacklist   bool
	MaxDeviceCount    int
}

var AuthConfig = new(Auth)
