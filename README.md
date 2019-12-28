# oci-injector-hook

It is sometimes useful to inject platform-specific files and devices into an OCI container at runtime.  Common use cases include device drivers (GPUs, network adapters, FPGAs, etc).  This hook uses the [POSIX-platform Hooks](https://github.com/opencontainers/runtime-spec/blob/master/config.md#posix-platform-hooks) from the OCI Runtime Spec to inject these files into a container's rootfs before the container is started.

## Supported File Types
* Devices (not yet implemented) - Device files under /dev/
* Directories - create directories in the container rootfs
* Binaries - exectuable binaries
* Libraries - library files (updates ld.so.cache)
* Miscellaneous - ordinary files to copy in (chmod +x/ldconfig not required)


## Configuration
Configurations are definied in .json files placed in the `/etc/oci-injector-hook/` directory.  Each configuration has an `activation_flag`, which indicates an environment variable that must be present in the container's environment for the hook to execute.

An example configuration file:
`/etc/oci-injector-hook/foo.json`
```
{
  "activation_flag": "OCI_FOO",
  "devices": [
    "/dev/foo",
  ],
  "binaries": [
    "/usr/bin/runfoo",
  ],
  "libraries": [
    "/usr/lib64/libfoo.so",
  ],
  "directories": [
    "/etc/foo",
  ],
  "miscellaneous": [
    "/etc/foo/config.json",
  ]
}
```

## License
This project is licensed under the Apache-2.0 License

## Inspiration
This project was inspired by work done in the following projects to support SolarFlare Network Adapters:
https://github.com/zvonkok/oci-decorator
https://github.com/solarflarecommunications/sfc-k8s-prestart-hook
