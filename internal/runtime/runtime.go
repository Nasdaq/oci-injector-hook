package runtime

import (
	"github.com/gclawes/oci-injector-hook/internal/config"
	specs "github.com/opencontainers/runtime-spec/specs-go"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

func CopyFile(src, dst string) {
	log.Debugf("copying src=%s -> dst=%s", src, dst)
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

func SetupDevices(config *config.InjectorConfig, containerConfig *specs.Spec) {
	log.Debugf("setting up devices '%s' under '%s'", config.Devices, containerConfig.Root.Path)
	log.Warn("SetupDevices not implemented!")
}

func CreateDirectories(config *config.InjectorConfig, containerConfig *specs.Spec) {
	log.Debugf("creating directories '%s' in '%s'", config.Directories, containerConfig.Root.Path)
	for _, dir := range config.Directories {
		dst_dir := filepath.Join(containerConfig.Root.Path, dir)
		log.Debugf("creating directory: %s", dst_dir)
		err := os.MkdirAll(dst_dir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func CopyBinaries(config *config.InjectorConfig, containerConfig *specs.Spec) {
	log.Debugf("copying binaries '%s' to '%s'", config.Binaries, containerConfig.Root.Path)
	for _, bin := range config.Binaries {
		dst := filepath.Join(containerConfig.Root.Path, bin)
		log.Debugf("copying binary: %s -> %s", bin, dst)
		CopyFile(bin, dst)
	}
}

func CopyLibraries(config *config.InjectorConfig, containerConfig *specs.Spec) {
	log.Debugf("copying libraries '%s' to '%s'", config.Libraries, containerConfig.Root.Path)
	for _, lib := range config.Libraries {
		dst := filepath.Join(containerConfig.Root.Path, lib)
		log.Debugf("copying library: %s -> %s", lib, dst)
		CopyFile(lib, dst)
	}

	if _, err := exec.Command("chroot", containerConfig.Root.Path, "ldconfig").Output(); err != nil {
		log.Fatal(err)
	}
}

func CopyMisc(config *config.InjectorConfig, containerConfig *specs.Spec) {
	log.Debugf("copying misc files '%s' to '%s'", config.Misc, containerConfig.Root.Path)
	for _, file := range config.Misc {
		dst := filepath.Join(containerConfig.Root.Path, file)
		log.Debugf("copying misc file: %s -> %s", file, dst)
		CopyFile(file, dst)
	}
}
