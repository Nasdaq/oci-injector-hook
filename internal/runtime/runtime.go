package runtime

import (
	"github.com/gclawes/oci-injector-hook/internal/config"
	specs "github.com/opencontainers/runtime-spec/specs-go"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
)

func CopyFile(src, dst string) {
	log.Infof("copying src=%s -> dst=%s", src, dst)
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcstat os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		log.Fatal(err)
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		log.Fatal(err)
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		log.Fatal(err)
	}

	if srcstat, err = os.Stat(src); err != nil {
		log.Fatal(err)
	}

	if err = os.Chmod(dst, srcstat.Mode()); err != nil {
		log.Fatal(err)
	}
}

func SetupDevices(config *config.InjectorConfig, state *specs.Spec) {
	log.Debugf("setting up devices '%s' under '%s'", config.Devices, state.Root.Path)
	log.Warn("SetupDevices not implemented!")
}

func CreateDirectories(config *config.InjectorConfig, state *specs.Spec) {
	log.Debugf("creating directories '%s' in '%s'", config.Directories, state.Root.Path)
	for _, dir := range config.Directories {
		dst_dir := filepath.Join(state.Root.Path, dir)
		log.Debugf("creating directory: %s", dst_dir)
		err := os.MkdirAll(dst_dir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func CopyBinaries(config *config.InjectorConfig, state *specs.Spec) {
	log.Debugf("copying binaries '%s' to '%s'", config.Binaries, state.Root.Path)
	log.Warn("CopyBinaries not implemented!")
}

func CopyLibraries(config *config.InjectorConfig, state *specs.Spec) {
	log.Debugf("copying libraries '%s' to '%s'", config.Libraries, state.Root.Path)
	log.Warn("CopyLibraries not implemented!")
}

func CopyMisc(config *config.InjectorConfig, state *specs.Spec) {
	log.Debugf("copying misc files '%s' to '%s'", config.Misc, state.Root.Path)
	log.Warn("CopyMisc not implemented!")
}
