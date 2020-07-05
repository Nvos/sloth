package infra

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"time"
)

type Mode string

const (
	Production Mode = "production"
	Test       Mode = "test"
	Develop    Mode = "develop"
)

type LoggerConfig struct {
	Level string
}

type EnvironmentConfig struct {
	Mode Mode
}

type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Schema   string

	ShouldDrop bool
	ShouldSeed bool
}

type SecurityConfig struct {
	JwtSigningSecret        string
	JwtTokenLifetime        int
	JwtRefreshTokenLifetime int
}

type ServerConfig struct {
	Host string
	Port int
}

type Configuration struct {
	Environment EnvironmentConfig
	Logger      LoggerConfig
	Database    DatabaseConfig
	Server      ServerConfig
}

func NewConfiguration() (*Configuration, error) {
	cfg :=  Configuration{}
	err := cfg.ReadIn()

	return &cfg, err
}

func MustNewConfiguration() *Configuration {
	cfg := Configuration{}
	if err := cfg.ReadIn(); err != nil {
		panic(fmt.Errorf("failed to read in configuration, err=%w", err))
	}

	return &cfg
}

func applyDefaults(v *viper.Viper) {
	v.SetDefault("logger.level", "info")
	v.SetDefault("environment.mode", "develop")

	v.SetDefault("server.host", "http://localhost")
	v.SetDefault("server.port", 8080)

	v.SetDefault("security.jwtTokenLifetime", time.Minute * 30)
	v.SetDefault("security.jwtRefreshTokenLifetime", time.Minute * 60 * 24)
}

func (c *Configuration) ReadIn() error {
	v := viper.New()
	applyDefaults(v)

	v.AddConfigPath(".")
	v.SetConfigName("config")
	v.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			_, err := os.Create("config.toml")
			if err != nil {
				return err
			}

			err = v.WriteConfigAs("config.toml")
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	err := v.Unmarshal(&c)
	if err != nil {
		return err
	}

	return nil
}
