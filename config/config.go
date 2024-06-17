package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Config структура для хранения конфигурации
type Config struct {
	TargetSite struct {
		URL string `yaml:"url"`
	} `yaml:"target_site"`
	Output struct {
		File string `yaml:"file"`
	} `yaml:"output"`
}

// LoadConfig загружает конфигурацию из файла
func LoadConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// GetConfig загружает конфигурацию и возвращает ее
func GetConfig() *Config {
	config, err := LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}
	return config
}
