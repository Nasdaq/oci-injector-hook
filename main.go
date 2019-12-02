package main

import (
	"github.com/gclawes/oci-injector-hook/internal/config"
	"log"
)

func main() {
	log.Printf("oci-injector-hook: starting")

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
