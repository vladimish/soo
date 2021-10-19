package configurator

import (
	"encoding/json"
	"io"
	"os"
)

type Level string

const (
	DEBUG   Level = "DEBUG"
	RELEASE Level = "RELEASE"
)

type Config struct {
	Logging struct {
		LogLevel Level `json:"log_level"`
	} `json:"logging"`
	NetworkConfiguration struct {
		Address string `json:"address"`
	} `json:"network_configuration"`
	Persistence struct {
		Auth struct {
			ConnectionString string `json:"connection_string"`
		} `json:"auth"`
	} `json:"persistence"`
}

var Cfg = Config{}

// LoadSettings loads data from .json config
// and parsing it to global object Cfg.
func LoadSettings(path string) error {
	configFile, err := os.Open(path)
	if err != nil {
		return err
	}
	fileBytes, err := io.ReadAll(configFile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(fileBytes, &Cfg)
	if err != nil {
		return err
	}

	return nil
}
