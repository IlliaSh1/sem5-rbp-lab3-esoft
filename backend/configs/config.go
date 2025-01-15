package configs

import (
	"time"
)

type Config struct {
	// Env         string        `yaml:"env" env-default:"local"`
	// StoragePath string        `yaml:"storage_path" env-required:"true"`
	// TokenTTL    time.Duration `yaml:"token_ttl" env-required:"true"`
	// GRPC        ServerConfig  `yaml:"grpc" env-required:"true"`
	ServerConfig ServerConfig `yaml:"server" env-required:"true"`
	DBConfig     DBConfig     `yaml:"database" env-required:"true"`
	JWTConfig    JWTConfig    `yaml:"jwt" env-required:"true"`
}

type ServerConfig struct {
	Port int `yaml:"port" env-required:"true"`
}

type DBConfig struct {
	DBType     string `yaml:"db_type" env-required:"true"`
	Host       string `yaml:"host" env-required:"true"`
	Port       int    `yaml:"port" env-required:"true"`
	DBName     string `yaml:"db_name" env-required:"true"`
	AdminName  string `yaml:"admin_name" env-required:"true"`
	DBPassword string `yaml:"db_password" env-required:"true"`
	SSL        string `yaml:"ssl" env-required:"true"`
}

type JWTConfig struct {
	SecretKey            string        `yaml:"secret_key" env-required:"true"`
	AccessTokenDuration  time.Duration `yaml:"access_token_duration" env-required:"true"`
	RefreshTokenDuration time.Duration `yaml:"refresh_token_duration" env-required:"true"`
}
