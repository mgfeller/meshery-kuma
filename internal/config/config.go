package config

import (
	"github.com/mgfeller/common-adapter-library/config"
)

// New returns the interface of the config handler
func New(name string) (config.Handler, error) {
	switch name {
	case "local":
		return NewLocal()
	case "viper":
		return NewViper()
	}
	return nil, config.ErrEmptyConfig
}
