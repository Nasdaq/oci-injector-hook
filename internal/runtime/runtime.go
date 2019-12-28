package runtime

import (
	"bufio"
	"fmt"
	"github.com/gclawes/oci-injector-hook/internal/config"
	specs "github.com/opencontainers/runtime-spec/specs-go"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"syscall"
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

/*
FIXME: replace with cgroups.ParseCgroupFile from github.com/opencontainers/runc/libcontainer/cgroups
This is currently broken due to some go dependency bug with github.com/Sirupsen/logrus,
which has a case-sensitivie dependency conflict with github.com/sirupsen/logrus
	cgroup_procfile := filepath.Join("/proc", strconv.Itoa(state.Pid), "cgroup")
	cgroup, err := cgroups.ParseCgroupFile(cgroup_procfile)
*/
func GetDevicesCgroup(pid int) (string, error) {
	var cgroup string
	filepath := filepath.Join("/proc", strconv.Itoa(pid), "cgroup")

	procfile, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer procfile.Close()

	scanner := bufio.NewScanner(procfile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if match, _ := regexp.MatchString("^[0-9]+:devices", line); match {
			return strings.Split(line, ":")[2], nil
		}
	}

	return cgroup, fmt.Errorf("unable to determine devices cgroup for pid=%d", pid)
}

func GetDeviceType(stat syscall.Stat_t) (string, error) {
	if stat.Mode&syscall.DT_CHR == syscall.DT_CHR {
		return "c", nil
	} else if stat.Mode&syscall.DT_BLK == syscall.DT_BLK {
		return "b", nil
	} else {
		return "", fmt.Errorf("cant determine device type for stat.Mode=%d", stat.Mode)
	}
}

func SetupDevices(config *config.InjectorConfig, containerConfig *specs.Spec, state *specs.State) {
	log.Debugf("setting up devices '%s' under '%s'", config.Devices, containerConfig.Root.Path)
	for _, devname := range config.Devices {
		var dev string
		var err error
		if match, _ := regexp.MatchString(`^/dev/`, devname); match {
			dev = devname
		} else {
			dev = filepath.Join("/dev", devname)
		}

		host_stat := syscall.Stat_t{}
		if err = syscall.Stat(dev, &host_stat); err != nil {
			log.Fatal(err)
		}

		major := unix.Major(uint64(host_stat.Rdev))
		minor := unix.Minor(uint64(host_stat.Rdev))

		// create device file in rootfs
		log.Infof("creating device %s", filepath.Join(containerConfig.Root.Path, dev))
		if err = syscall.Mknod(filepath.Join(containerConfig.Root.Path, dev), uint32(host_stat.Mode), int(unix.Mkdev(major, minor))); err != nil {
			log.Fatal(err)
		}

		// get cgroup name
		cgroup, err := GetDevicesCgroup(state.Pid)
		if err != nil {
			log.Fatal(err)
		}

		devtype, err := GetDeviceType(host_stat)
		log.Infof("type=%s cgroup=%s", devtype, cgroup)
		// run cgset to allow the device in the container cgroup
		allow_str := fmt.Sprintf("devices=%s %d:%d rwm")
		if _, err := exec.Command("cgset", "-r", allow_str, cgroup).Output(); err != nil {
			log.Fatal(err)
		}
	}
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
