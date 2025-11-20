package goconfig

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

var yamlConfigFilePath string
var envConfigFilePath string
var envPrefix string

func InitOnce() {
	flagYamlConfigPath := flag.String("yaml-config", "", "yaml config file path")
	flagEnvConfigPath := flag.String("env-config", "", ".env config file path")
	flagEnvPrefix := flag.String("env-prefix", "", "environment settings prefix")
	flag.Parse()

	yamlConfigFilePath = *flagYamlConfigPath
	envConfigFilePath = *flagEnvConfigPath
	envPrefix = *flagEnvPrefix
}

// Fill fills config structure.
func Fill(config any) error {
	validate := validator.New()

	if err := fillFromFile(config); err != nil {
		return err
	}
	if err := fillFromEnv(config); err != nil {
		return err
	}

	return validate.Struct(config)
}

func fillFromFile(cfg any) error {
	if yamlConfigFilePath == "" {
		fmt.Println("fdvd")
		return nil
	}

	file, err := os.Open(yamlConfigFilePath)
	if err == nil {
		defer file.Close()
		err = yaml.NewDecoder(file).Decode(cfg)
	}

	return err
}

func fillFromEnv(config any) error {
	if envConfigFilePath != "" {
		godotenv.Load(strings.Split(envConfigFilePath, ",")...)
	} else {
		godotenv.Load()
	}

	return envconfig.Process(envPrefix, config)
}
