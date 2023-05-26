package config

import "auto/pkg/conf"

func New() conf.Config {
	return conf.Config{
		Port: 8080,
	}
}
