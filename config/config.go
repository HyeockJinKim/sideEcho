package config

import (
	"io/ioutil"
	"os"

	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/pflag"
)

const (
	flagConfig = "config"
)

const (
	defaultConfigPath = "configs/config.toml"
)

type Config struct {
	Flags  Flags
	Config ServerConfig
}

type ServerConfig struct {
	Api Api
}

type Api struct {
	Port               int
	InternalPort       int
	TimeoutSec         int
	ShutDownTimeoutSec int
}

type Flags struct {
	ConfigPath string
	Auth       Auth
}

type Auth struct {
	Username []byte
	Password []byte
}

func ReadFlags() *Flags {
	flags := &Flags{}
	flagSet := pflag.NewFlagSet("sideEcho", pflag.ExitOnError)
	flagSet.StringVarP(&flags.ConfigPath, flagConfig, "c", defaultConfigPath, "configuration file path")

	flags.Auth.Username = []byte(os.Getenv("AUTH_USERNAME"))
	flags.Auth.Password = []byte(os.Getenv("AUTH_PASSWORD"))
	return flags
}

func ReadConfigFile(flags *Flags) (*Config, error) {
	b, err := ioutil.ReadFile(flags.ConfigPath)
	if err != nil {
		return nil, err
	}
	var configs ServerConfig
	err = toml.Unmarshal(b, &configs)
	if err != nil {
		return nil, err
	}

	return &Config{
		Flags:  *flags,
		Config: configs,
	}, nil
}
