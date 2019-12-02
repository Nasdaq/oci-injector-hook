package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var defaultConfigDir = "/etc/oci-injector-hook"
var configExt = ".json"

type InjectorConfig struct {
	Name           string
	ActivationFlag string   `mapstructure:"activation_flag"`
	Devices        []string `mapstructure:"devices"`
	Binaries       []string `mapstructure:"binaries"`
	Libraries      []string `mapstructure:"libraries"`
	Directories    []string `mapstructure:"directories"`
	Misc           []string `mapstructure:"miscellaneous"`
}

// visitConfigDir returns files in a directory with the extension configExt
// from https://flaviocopes.com/go-list-files/
func visitConfigDir(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		if filepath.Ext(path) == configExt {
			*files = append(*files, path)
		}
		return nil
	}
}

// GetVipers returns a map of config name -> *viper.Viper config objects
func GetConfigVipers() map[string]*viper.Viper {
	log.Printf("oci-injector-hook: getting configs")

	configDir, ok := os.LookupEnv("OCI_INJECTOR_CONFIG_DIR")
	if !ok {
		configDir = defaultConfigDir
	}

	var configFiles []string
	// var configs []*InjectorConfig
	vipers := make(map[string]*viper.Viper)

	// get config files in configDir
	err := filepath.Walk(configDir, visitConfigDir(&configFiles))
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range configFiles {
		configName := strings.TrimSuffix(file, configExt)

		v := viper.New()
		v.SetConfigName(configName)
		v.AddConfigPath(configDir)

		if err := v.ReadInConfig(); err != nil {
			log.Fatal("couldn't read config: %s", err)
		}

		vipers[configName] = v
	}

	return vipers
}

func GetConfigs() []*InjectorConfig {
	var configs []*InjectorConfig
	for name, v := range GetConfigVipers() {
		var config InjectorConfig

		if err := v.Unmarshal(&config); err != nil {
			log.Fatal("couldn't unmarshal config: %s", err)
		}

		config.Name = name

		configs = append(configs, &config)
	}

	return configs
}
