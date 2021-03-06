package config

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/endpoints"
)

type Config struct {
	APIServer           string
	Interface           string
	ConfigDir           string
	BinaryDir           string
	BootstrapToken      string
	WireGuardBinary     string
	WireGuardGoBinary   string
	PrivateKeyPath      string
	WireGuardConfigPath string
	BootstrapConfigPath string
	LogLevel            string
	OAuth2Config        oauth2.Config
	Platform            string
	BootstrapAPI        string
	SessionInfoPath     string
}

func (c *Config) SetDefaults() {
	c.Platform = Platform
	c.SetPlatformDefaults()
	c.PrivateKeyPath = filepath.Join(c.ConfigDir, "private.key")
	c.WireGuardConfigPath = filepath.Join(c.ConfigDir, c.Interface+".conf")
	c.BootstrapConfigPath = filepath.Join(c.ConfigDir, "bootstrapconfig.json")
	c.SessionInfoPath = filepath.Join(c.ConfigDir, "sessioninfo.json")
}

func DefaultConfig() Config {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal("Getting user config dir: %w", err)
	}

	return Config{
		APIServer:    "http://10.255.240.1",
		BootstrapAPI: "https://bootstrap.device.nais.io",
		ConfigDir:    filepath.Join(userConfigDir, "naisdevice"),
		BinaryDir:    "/usr/local/bin",
		LogLevel:     "info",
		OAuth2Config: oauth2.Config{
			ClientID:    "8086d321-c6d3-4398-87da-0d54e3d93967",
			Scopes:      []string{"openid", "6e45010d-2637-4a40-b91d-d4cbb451fb57/.default", "offline_access"},
			Endpoint:    endpoints.AzureAD("62366534-1ec3-4962-8869-9b5535279d0b"),
			RedirectURL: "http://localhost:51800",
		},
	}
}
