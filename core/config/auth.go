package config

type Auth struct {
	Use               string
	Timeout           int
	Secret            string
	DeviceCheck       bool
	TokenBlacklist    bool
	AllowMultiDevices bool
	MaxDevicesPerUser int
}

var AuthConfig = new(Auth)
