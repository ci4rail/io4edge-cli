# io4edge-cli
Command line tool to manage io4edge devices.

## Description

Io4edge devices are intelligent I/O devices invented by Ci4Rail, connected to the host via network.

The `io4edge-cli` tool is intended to run on the host machine. Via the command line tool, you can:
    - Identify the currently running firmware
    - Load new firmware
    - Identify HW (name, revision, serial number)
    - Program HW identification
    - Restart device

This repo contains also a device simulator `io4edge-devsim` to test the `io4edge-cli`

## Examples

### Identify currently running firmware:
```bash
$ io4edge-cli -d 192.168.7.1:9999 fw
Firmware name: fw_esp_io4edge_default, Version 1f3f2a2-dirty
```

### Get hardware inventory information:
```bash
$ io4edge-cli -d 192.168.7.1:9999 hw
Hardware name: S101-IOU04, rev: 2, serial: 70a3b920-7eb7-434e-b20d-6d0a12618ffe
```

### Load firmware from a firmware package
A firmware package contains the firmware binary and a manifest file. The io4edge-cli checks if the firmware is suitable for the device before loading it.

```bash
$ io4edge-cli -d 192.168.7.1:9999 load-firmware fw-cpu01uc-tty_accdl-1.1.0.beta1.fwpkg
Reconnect to restarted device
...
```


### Load raw firmware
"Raw Firmware" means: Load a firmware binary that is not embedded in a firmware package file. In this case, the

```bash
$ io4edge-cli -d 192.168.7.1:9999 load-raw-firmware build/fw_esp_io4edge_default.bin
Reconnect to restarted device
Reading back firmware id
Firmware name: fw_esp_io4edge_default, Version 1f3f2a2-dirty
```

### Program HW inventory
```bash
$ io4edge-cli -d 192.168.7.1:9999 program-hwid S101-IOU04 2 70a3b920-7eb7-434e-b20d-6d0a12618ffe
Success. Read back programmed ID
Hardware name: S101-IOU04, rev: 2, serial: 70a3b920-7eb7-434e-b20d-6d0a12618ffe
```

## Building

### Local Builds

go compiler has to be installed.

```
make
```
