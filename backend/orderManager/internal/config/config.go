package config

import (
	"github.com/spf13/viper"
	"time"
)

const (
	defaultHost					  = "localhost"
	defaultHTTPPort               = "8080"
	defaultHTTPRWTimeout          = 10 * time.Second
	defaultHTTPMaxHeaderMegabytes = 1
)

type (
	Config struct {
		HTTP        HTTPConfig
		Postgres    PostgresConfig
		Router 		RouterConfig
		Hubs		HubsConfig
	}

	HTTPConfig struct {
		Host               string
		Port               string
		ReadTimeout        time.Duration
		WriteTimeout       time.Duration
		MaxHeaderMegabytes int
	}

	PostgresConfig struct {
		Host     string
		Port     string
		Username string
		Password string
		DBName   string
		SSLMode  string
	}
	RouterConfig struct {
		Host 				string
		Port				string
		VCopter				float32
		MaxDistanceCopter 	float32
		TimeWaitOrder		float32
	}
	HubsConfig struct {
		Host		string
		BasePort	int
	}
)

// Init Инициализация структуры конфига
func Init(configsDir string) (*Config, error) {
	InitDefault()

	err := parseConfigFile(configsDir)
	if err != nil {
		return nil, err
	}

	var cfg Config

	err = unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

// parseConfigFile прикрепление библиотеки viper к файлу конфигурации
func parseConfigFile(folder string) error {
	viper.AddConfigPath(folder)
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.MergeInConfig()
}

// unmarshal заполение структуры конфигурации с помощью библиотеки viper
func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("HTTPConfig", &cfg.HTTP); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("PostgresConfig", &cfg.Postgres); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("RouterConfig", &cfg.Router); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("HubsConfig", &cfg.Hubs); err != nil {
		return err
	}

	return nil
}

// InitDefault инициализация структуры конфигурации стандарными значениями
func InitDefault() {
	viper.SetDefault("HTTP.Host", defaultHost)
	viper.SetDefault("HTTP.Port", defaultHTTPPort)
	viper.SetDefault("HTTP.MaxHeaderMegabytes", defaultHTTPMaxHeaderMegabytes)
	viper.SetDefault("HTTP.ReadTimeout", defaultHTTPRWTimeout)
	viper.SetDefault("HTTP.WriteTimeout", defaultHTTPRWTimeout)
}
