package config

import (
	"bufio"
	"encoding/json"
	specs "github.com/opencontainers/runtime-spec/specs-go"
	"github.com/spf13/viper"
	"io"
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

// GetVipers returns a map of config name -> *viper.Viper config objects
func GetConfigVipers() map[string]*viper.Viper {
	log.Printf("oci-injector-hook: getting configs")

	configDir, ok := os.LookupEnv("OCI_INJECTOR_CONFIG_DIR")
	if !ok {
		configDir = defaultConfigDir
	}

	// get config files in configDir
	configFiles, err := filepath.Glob(configDir + "/*.json")
	if err != nil {
		log.Fatal(err)
	}

	vipers := make(map[string]*viper.Viper)

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

func GetState(stdin io.Reader) (*specs.State, error) {
	var state specs.State
	//json.Unmarshal([]byte(input), &state
	err := json.NewDecoder(bufio.NewReader(stdin)).Decode(&state)

	return &state, err
}
