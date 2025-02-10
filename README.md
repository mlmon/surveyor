# Surveyor

[![codecov](https://codecov.io/gh/mlmon/surveyor/graph/badge.svg?token=QGGV4Y1EMX)](https://codecov.io/gh/mlmon/surveyor)
[![test](https://github.com/mlmon/surveyor/actions/workflows/go.yml/badge.svg)](https://github.com/mlmon/surveyor/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/mlmon/surveyor)](https://goreportcard.com/report/github.com/mlmon/surveyor)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/mlmon/surveyor)

Collects common configuration elements from a linux host that are important to ML workloads including:

* kernel modules.
* nvidia SMI details.
* OS release.
* package list.
* uname.
* sysctl.

# Data Format

Currently the default is to output a CycloneDX JSON compatible file.

## Example Data

| Source          | Name         | Value                                |
|-----------------|--------------|--------------------------------------|
| /proc/cmdline   | nokaslr      | false                                |
| sysctl          | nofiles      | 4096                                 |
| dkms            | nvidia       | 2                                    |
| nvidia-smi      | model        | H100 80GB HBM3                       |
| ethtool         | driver       |                                      |
| env             | hosting      | aws                                  |
| /proc/cpuinfo   | cpu.model    | AMD EPYC 7R13 Processor              |
| /etc/os-release | os.release   | Ubuntu 22.04.3 LTS                   |
| uname           | kernel       | 6.2.0-1018-aws                       |
| date            | capture.date | 2025-01-28 11h33 UTC                 |
| identifier      | capture.uid  | f3f2e850-b5d4-11ef-ac7e-96584d5248b2 |

# Use cases

1. Collection and export.
2. Comparison to other systems.
3. Verification of all systems within the same environment.
