# Surveyor

Collects common configuration elements from a linux host taht are important to ML workloads.

# Architecture

* Collector - 
* Aggregator - aggregates

# Data Format

| Name   | Description               |
|--------|---------------------------|
| Source | Origin of the data        |
| Name   | Index key                 |
| Value  | Value assigned to the key |

## Example

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