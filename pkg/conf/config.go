package conf

type Config struct {
	Port     int
	Mode     string // dev or prod
	LogLevel string // debug, info, warn, error
	Mongo    MongoConfig
	Redis    CacheConfig
}

type MongoConfig struct {
	DB  string
	Url string
}

type CacheConfig struct {
	Url      string
	Password string
}
