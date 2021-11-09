package config

import (
"github.com/spf13/viper"
"time"
)

type (
	Config struct {
		HTTP        HTTPConfig
		Postgres    PostgresConfig
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
)

// Init Инициализация структуры конфига
func Init(configsDir string) (*Config, error) {
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
	//if env == EnvLocal {
	//	return nil
	//}
	//viper.SetConfigName(env)

	return viper.MergeInConfig()
}

// unmarshal заполение структуры конфигурации с помощью библиотеки viper
func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("postgres", &cfg.Postgres); err != nil {
		return err
	}
	return nil
}