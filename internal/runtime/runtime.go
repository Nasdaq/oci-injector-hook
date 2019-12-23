package runtime

import (
	"github.com/gclawes/oci-injector-hook/internal/config"
	specs "github.com/opencontainers/runtime-spec/specs-go"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

func SetupDevices(config *config.InjectorConfig, state *specs.Spec) {
	log.Debugf("setting up devices '%s' under '%s'", config.Devices, state.Root.Path)
	log.Warn("SetupDevices not implemented!")
}

func CopyBinaries(config *config.InjectorConfig, state *specs.Spec) {
	log.Debugf("copying binaries '%s' to '%s'", config.Binaries, state.Root.Path)
	log.Warn("CopyBinaries not implemented!")
}

func CopyLibraries(config *config.InjectorConfig, state *specs.Spec) {
	log.Debugf("copying libraries '%s' to '%s'", config.Libraries, state.Root.Path)
	log.Warn("CopyLibraries not implemented!")
}

func CreateDirectories(config *config.InjectorConfig, state *specs.Spec) {
	log.Debugf("creating directories '%s' in '%s'", config.Directories, state.Root.Path)
	for _, dir := range config.Directories {
		dst_dir := filepath.Join(state.Root.Path, dir)
		log.Infof("creating directory: %s", dst_dir)
		err := os.MkdirAll(dst_dir, os.FileMode(int(0755)))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func CopyMisc(config *config.InjectorConfig, state *specs.Spec) {
	log.Debugf("copying misc files '%s' to '%s'", config.Misc, state.Root.Path)
	log.Warn("CopyMisc not implemented!")
}
