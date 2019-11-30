package main

import (
	"github.com/gclawes/oci-injector-hook/internal/config"
	"log"
)

func main() {
	log.Printf("oci-injector-hook: starting")

	configs := config.GetConfigs()
	for _, config := range configs {
		log.Printf("config.Name=%s", config.Name)
		log.Printf("config.Version=%s", config.Version)
		log.Printf("config.ActivationFlag=%s", config.ActivationFlag)
		log.Printf("config.DriverFeature=%s", config.DriverFeature)

		for name, inv := range config.Inventory {
			log.Printf("config.Inventory[%s].Devices=%s", name, inv.Devices)
			log.Printf("config.Inventory[%s].Binaries=%s", name, inv.Binaries)
			log.Printf("config.Inventory[%s].Libraries=%s", name, inv.Libraries)
			log.Printf("config.Inventory[%s].Directories=%s", name, inv.Directories)
			log.Printf("config.Inventory[%s].Misc=%s", name, inv.Misc)
		}
	}
}
