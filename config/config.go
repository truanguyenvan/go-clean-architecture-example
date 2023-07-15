package config

import (
	"errors"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

var Set = wire.NewSet(NewConfig)

// Configuration
type Configuration struct {
	Server        ServerConfig
	Postgres      PostgresConfig
	Logger        Logger
	Redis         RedisConfig
	RedisCluster  RedisClusterConfig
	MongoDB       MongoDB
	Authorization Authorization
}

// ServerConfig struct
type ServerConfig struct {
	Name                string
	AppVersion          string
	Port                string
	BaseURI             string
	Mode                string
	ReadTimeout         time.Duration
	WriteTimeout        time.Duration
	SSL                 bool
	CtxDefaultTimeout   time.Duration
	CSRF                bool
	Debug               bool
	GrRunningThreshold  int //  threshold for goroutines are running (which could indicate a resource leak).
	GcPauseThreshold    int //  threshold threshold garbage collection pause exceeds. (Millisecond)
	CacheDeploymentType int
}

// Casbin struct
type Authorization struct {
	CasbinModelFilePath  string
	CasbinPolicyFilePath string
	JWTSecret            string
}

// Logger config
type Logger struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

// PostgresConfig config
type PostgresConfig struct {
	PostgresqlHost     string
	PostgresqlPort     string
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDbname   string
	PostgresqlSSLMode  bool
	PgDriver           string
}

// RedisConfig struct
type RedisConfig struct {
	Address     string
	Password    string
	DefaultDb   string
	MinIdleCons int
	PoolSize    int
	PoolTimeout int
	DB          int
}

// Redis config
type RedisClusterConfig struct {
	Delimiter   string
	ReadOnly    bool
	Address     string
	DefaultDb   string
	MinIdleCons int
	PoolSize    int
	PoolTimeout int
	Password    string
	DB          int
}

// MongoDB config
type MongoDB struct {
	MongoURI        string
	MongoUser       string
	MongoPassword   string
	connectTimeout  int
	maxConnIdleTime int
	minPoolSize     uint64
	maxPoolSize     uint64
}

// Get config path for local or docker
func getDefaultConfig() string {
	return "./config/config-local"
}

// Load config file from given path
func NewConfig() (*Configuration, error) {
	path := os.Getenv("cfgPath")
	if path == "" {
		path = getDefaultConfig()
	}

	v := viper.New()

	v.SetConfigName(path)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	err := v.Unmarshal(&DefaultConfig)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &DefaultConfig, nil
}
