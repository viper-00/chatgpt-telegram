package config

import (
	"errors"
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	v *viper.Viper

	// OpenAISession string
	OpenAiAuthorization string
}

// LoadOrCreatePersistentConfig uses the default config directory for the current OS
// to load or create a config file named "chatgpt.json"
func LoadOrCreatePersistentConfig(path ...string) (*Config, error) {
	// currentFilePath, err := filepath.Abs(os.Args[0])
	// if err != nil {
	// 	return nil, errors.New(err.Error())
	// }
	// configPath := filepath.Join(filepath.Dir(currentFilePath), "chatgpt.json")
	// if err != nil {
	// 	return nil, errors.New(fmt.Sprintf("Couldn't get user config dir: %v", err))
	// }

	var config string

	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose a config file.")
		flag.Parse()
		if config == "" {
			panic(fmt.Errorf("could not find config file, please use `-c` command pointing to the specified file"))
		}
	} else {
		config = path[0]
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("json")
	// v.AddConfigPath(currentFilePath)

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if err := v.SafeWriteConfig(); err != nil {
				return nil, errors.New(fmt.Sprintf("Couldn't create config file: %v", err))
			}
		} else {
			return nil, errors.New(fmt.Sprintf("Couldn't read config file: %v", err))
		}
	}

	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error parsing config: %v", err))
	}
	cfg.v = v

	return &cfg, nil
}

// func (cfg *Config) SetSessionToken(token string) error {
// 	// key must match the struct field name
// 	cfg.v.Set("OpenAISession", token)
// 	cfg.OpenAISession = token
// 	return cfg.v.WriteConfig()
// }

// func (cfg *Config) SetAuthorization(apiKey string) error {
// 	cfg.v.Set("Authorization", apiKey)
// 	cfg.OpenAiAuthorization = apiKey
// 	return cfg.v.WriteConfig()
// }
