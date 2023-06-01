package config

import "auto/pkg/conf"

func New() conf.Config {
	return conf.Config{
		Port:     8080,
		Mode:     "dev",
		LogLevel: "debug",
		Mongo: conf.MongoConfig{
			DB: "yuqi",
			// Url: "mongodb://root:xy2089@mongo:27017",
			Url: "mongodb://root:xy2089@localhost:27017",
		},
		Redis: conf.CacheConfig{
			// Url: "redis:6379",
			Url: "localhost:6379",
			// Todo 暂时不用密码
			Password: "",
		},
	}
}
