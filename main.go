package main

import (
	"encoding/json"
	"github.com/gclawes/oci-injector-hook/internal/config"
	"github.com/gclawes/oci-injector-hook/internal/runtime"
	specs "github.com/opencontainers/runtime-spec/specs-go"
	log "github.com/sirupsen/logrus"
	logrus_syslog "github.com/sirupsen/logrus/hooks/syslog"
	"io/ioutil"
	"log/syslog"
	"os"
	"path/filepath"
)

func init() {
	debug, ok := os.LookupEnv("DEBUG")
	if ok && debug == "true" {
		log.SetLevel(log.DebugLevel)
	}

	hook, err := logrus_syslog.NewSyslogHook("", "", syslog.LOG_INFO, "")
	if err != nil {
		log.Fatal("Unable to connect to local syslog daemon")
	} else {
		log.AddHook(hook)
	}
}

func main() {
	state, err := config.GetState(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	log.Debugf("state=%+v", state)

	configs := config.GetConfigs()
	for _, config := range configs {
		configJson, err := ioutil.ReadFile(filepath.Join(state.Bundle, "config.json"))
		if err != nil {
			log.Fatal(err)
		}

		var containerConfig specs.Spec
		err = json.Unmarshal(configJson, &containerConfig)
		if err != nil {
			log.Fatal(err)
		}

		log.Debug("containerConfig.Process.Env=%s", containerConfig.Process.Env)

		if config.ActivationFlagPresent(containerConfig.Process.Env) {
			runtime.SetupDevices(config, state)
			runtime.CopyBinaries(config, state)
			runtime.CopyLibraries(config, state)
			runtime.CopyDirectories(config, state)
			runtime.CopyMisc(config, state)
		} else {
			log.Infof("activation flag %s not present!", config.ActivationFlag)
		}

	}
}
