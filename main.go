package main

import (
	"github.com/gclawes/oci-injector-hook/internal/config"
	"log"
	"os"
)

func main() {
	log.Printf("oci-injector-hook: starting")

	log.Printf("oci-injector-hook: getting container state from stdin")
	state, err := config.GetState(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("state.Version=%s", state.Version)
	log.Printf("state.ID=%s", state.ID)
	log.Printf("state.Status=%s", state.Status)
	log.Printf("state.Pid=%d", state.Pid)
	log.Printf("state.Bundle=%s", state.Bundle)
	log.Printf("state.Annotations=%s", state.Annotations)

	log.Printf("oci-injector-hook: getting configs")
	configs := config.GetConfigs()
	for _, config := range configs {
		log.Printf("configs[%s].ActivationFlag=%s", config.Name, config.ActivationFlag)
		log.Printf("configs[%s].Devices=%s", config.Name, config.Devices)
		log.Printf("configs[%s].Binaries=%s", config.Name, config.Binaries)
		log.Printf("configs[%s].Libraries=%s", config.Name, config.Libraries)
		log.Printf("configs[%s].Directories=%s", config.Name, config.Directories)
		log.Printf("configs[%s].Misc=%s", config.Name, config.Misc)
	}
}
