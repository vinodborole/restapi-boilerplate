package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

var configFile Config

//MyAppConfig config
type MyAppConfig struct {
	Database    DatabaseConfig `yaml:"database"`
	Log         LogConfig      `yaml:"log"`
	Application Application    `yaml:"application"`
}

//Application application config
type Application struct {
	Port        string `yaml:"http-port"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

//DatabaseConfig database config
type DatabaseConfig struct {
	DBName   string `yaml:"dbname"`
	Username string `yaml:"user"`
	Password string `yaml:"password"`
	Location string `yaml:"location"`
	Port     string `yaml:"port"`
}

//LogConfig log config
type LogConfig struct {
	Location   string `yaml:"location"`
	Level      string `yaml:"level"`
	MaxBackups int    `yaml:"maxbackups"`
	MaxAge     int    `yaml:"maxage"`
}

//Config config
type Config struct {
	MyAppConfig MyAppConfig `yaml:"myappconfig"`
}

//GetConfig get config
func GetConfig() Config {
	return configFile
}

//ReadYamlConfigFile Initial Function
func ReadYamlConfigFile() error {
	var config = MyAppConfig{}
	// myapp Config
	yamlFile, err := ioutil.ReadFile(getConfigPath())
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return err
	}
	configFile = Config{MyAppConfig: config}
	return nil
}

func getConfigPath() string {
	//return GetConfigPath() + "/yaml/config.yaml"
	return "/Users/vinodborole/Documents/personal/projects/bredec/restapi-boilerplate/bin/yaml/config.yaml"
}

//GetConfigPath get config path
func GetConfigPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}
