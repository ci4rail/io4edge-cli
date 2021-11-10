# io4edge-cli
Command line tool to manage io4edge devices.

## Description

Io4edge devices are intelligent I/O devices invented by Ci4Rail, connected to the host via network.

The `io4edge-cli` tool is intended to run on the host machine. Via the command line tool, you can:
* Identify the currently running firmware
* Load new firmware
* Identify HW (name, revision, serial number)
* Program HW identification
* Restart device

## Examples

io4edge-cli addresses io4edge devices via their mdns service names. To find out available services (and therefore io4edege devices), you can use a MDNS browser, such as `avahi-browse`. Search for services named `_io4edge-core._tcp`:

```shell
$ avahi-browse -t _io4edge-core._tcp
+ usb_io_ctrl IPv4 cpu01uc-usb_io_ctrl   _io4edge-core._tcp   local
```

Then concatenate the `instance` with the `service` name, i.e. `cpu01uc-usb_io_ctrl._io4edge-core._tcp` in our example. Pass this string via the `-s` parameter to `io4edge-cli`

### Identify currently running firmware:
```bash
$ io4edge-cli -s cpu01uc-usb_io_ctrl._io4edge-core._tcp fw
Firmware name: fw_esp_io4edge_default, Version 1f3f2a2-dirty
```

### Get hardware inventory information:
```bash
$ io4edge-cli -s cpu01uc-usb_io_ctrl._io4edge-core._tcp hw
Hardware name: S101-IOU04, rev: 2, serial: 70a3b920-7eb7-434e-b20d-6d0a12618ffe
```

### Load firmware from a firmware package
A firmware package contains the firmware binary and a manifest file. The io4edge-cli checks if the firmware is suitable for the device before loading it.

```bash
$ io4edge-cli -s cpu01uc-usb_io_ctrl._io4edge-core._tcp load-firmware fw-cpu01uc-tty_accdl-1.1.0.beta1.fwpkg
Reconnect to restarted device
...
```
### Load raw firmware
"Raw Firmware" means: Load a firmware binary that is not embedded in a firmware package file. In this case, the

```bash
$ io4edge-cli -s cpu01uc-usb_io_ctrl._io4edge-core._tcp load-raw-firmware build/fw_esp_io4edge_default.bin
Reconnect to restarted device
Reading back firmware id
Firmware name: fw_esp_io4edge_default, Version 1f3f2a2-dirty
```

### Program HW inventory
```bash
$ io4edge-cli -s cpu01uc-usb_io_ctrl._io4edge-core._tcp program-hwid S101-IOU04 2 70a3b920-7eb7-434e-b20d-6d0a12618ffe
Success. Read back programmed ID
Hardware name: S101-IOU04, rev: 2, serial: 70a3b920-7eb7-434e-b20d-6d0a12618ffe
```

## Building

### Local Builds

go compiler has to be installed.

```
make
```

# Simulator
This repo contains also a device simulator `io4edge-devsim` to test the `io4edge-cli` (Currently not usable as it doesn't support MDNS).