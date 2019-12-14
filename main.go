package main

import (
	"github.com/gclawes/oci-injector-hook/internal/config"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	debug, ok := os.LookupEnv("DEBUG")
	if ok && debug == "true" {
		log.SetLevel(log.DebugLevel)
	}
}

func main() {
	log.Debugf("oci-injector-hook: starting")

	log.Debugf("oci-injector-hook: getting container state from stdin")
	state, err := config.GetState(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	log.Debugf("state.Version=%s", state.Version)
	log.Debugf("state.ID=%s", state.ID)
	log.Debugf("state.Status=%s", state.Status)
	log.Debugf("state.Pid=%d", state.Pid)
	log.Debugf("state.Bundle=%s", state.Bundle)
	log.Debugf("state.Annotations=%s", state.Annotations)

	log.Debugf("oci-injector-hook: getting configs")
	configs := config.GetConfigs()
	for _, config := range configs {
		log.Debugf("configs[%s].ActivationFlag=%s", config.Name, config.ActivationFlag)
		log.Debugf("configs[%s].Devices=%s", config.Name, config.Devices)
		log.Debugf("configs[%s].Binaries=%s", config.Name, config.Binaries)
		log.Debugf("configs[%s].Libraries=%s", config.Name, config.Libraries)
		log.Debugf("configs[%s].Directories=%s", config.Name, config.Directories)
		log.Debugf("configs[%s].Misc=%s", config.Name, config.Misc)
	}
}
