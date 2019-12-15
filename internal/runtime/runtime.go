package runtime

import (
	"github.com/gclawes/oci-injector-hook/internal/config"
	specs "github.com/opencontainers/runtime-spec/specs-go"
	log "github.com/sirupsen/logrus"
	// "os"
)

func SetupDevices(config *config.InjectorConfig, state *specs.State) {
	log.Debugf("setting up devices '%s' under '%s'", config.Devices, state.Bundle+"/rootfs")
	log.Warn("SetupDevices not implemented!")
}

func CopyBinaries(config *config.InjectorConfig, state *specs.State) {
	log.Debugf("copying binaries '%s' to '%s'", config.Binaries, state.Bundle+"/rootfs")
	log.Warn("CopyBinaries not implemented!")
}

func CopyLibraries(config *config.InjectorConfig, state *specs.State) {
	log.Debugf("copying libraries '%s' to '%s'", config.Libraries, state.Bundle+"/rootfs")
	log.Warn("CopyLibraries not implemented!")
}

func CopyDirectories(config *config.InjectorConfig, state *specs.State) {
	log.Debugf("copying directories '%s' to '%s'", config.Directories, state.Bundle+"/rootfs")
	log.Warn("CopyDirectories not implemented!")
}

func CopyMisc(config *config.InjectorConfig, state *specs.State) {
	log.Debugf("copying misc files '%s' to '%s'", config.Misc, state.Bundle+"/rootfs")
	log.Warn("CopyMisc not implemented!")
}
