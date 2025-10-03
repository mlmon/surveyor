package source_test

import (
	"testing"

	a "github.com/gogunit/gunit/hammy"
	"github.com/mlmon/surveyor/source"
)

func Test_procfs_successful(t *testing.T) {
	assert := a.New(t)
	records, _ := source.ProcSys("testdata/procfs")()
	assert.Is(a.Struct(records).EqualTo(&source.Records{
		Source: "procfs",
		Entries: source.Entries{
			{"fs.file-max", "9223372036854775807"},
			{"kernel.hostname", "nfisher-mbp"},
		}}))
}

func Test_proc_nvidia_params_successful(t *testing.T) {
	assert := a.New(t)
	records, _ := source.ProcNvidiaParams("testdata/procfs")()

	assert.Is(a.Struct(records).EqualTo(&source.Records{
		Source: "nvidia_driver_params",
		Entries: source.Entries{
			{Key: "ResmanDebugLevel", Value: "4294967295"},
			{Key: "RmLogonRC", Value: "1"},
			{Key: "ModifyDeviceFiles", Value: "1"},
			{Key: "DeviceFileUID", Value: "0"},
			{Key: "DeviceFileGID", Value: "0"},
			{Key: "DeviceFileMode", Value: "438"},
			{Key: "InitializeSystemMemoryAllocations", Value: "1"},
			{Key: "UsePageAttributeTable", Value: "4294967295"},
			{Key: "EnableMSI", Value: "1"},
			{Key: "EnablePCIeGen3", Value: "0"},
			{Key: "MemoryPoolSize", Value: "0"},
			{Key: "KMallocHeapMaxSize", Value: "0"},
			{Key: "VMallocHeapMaxSize", Value: "0"},
			{Key: "IgnoreMMIOCheck", Value: "0"},
			{Key: "EnableStreamMemOPs", Value: "0"},
			{Key: "EnableUserNUMAManagement", Value: "0"},
			{Key: "NvLinkDisable", Value: "0"},
			{Key: "RmProfilingAdminOnly", Value: "1"},
			{Key: "PreserveVideoMemoryAllocations", Value: "0"},
			{Key: "EnableS0ixPowerManagement", Value: "0"},
			{Key: "S0ixPowerManagementVideoMemoryThreshold", Value: "256"},
			{Key: "DynamicPowerManagement", Value: "3"},
			{Key: "DynamicPowerManagementVideoMemoryThreshold", Value: "200"},
			{Key: "RegisterPCIDriver", Value: "1"},
			{Key: "EnablePCIERelaxedOrderingMode", Value: "0"},
			{Key: "EnableResizableBar", Value: "0"},
			{Key: "EnableGpuFirmware", Value: "18"},
			{Key: "EnableGpuFirmwareLogs", Value: "2"},
			{Key: "RmNvlinkBandwidthLinkCount", Value: "0"},
			{Key: "EnableDbgBreakpoint", Value: "0"},
			{Key: "OpenRmEnableUnsupportedGpus", Value: "1"},
			{Key: "DmaRemapPeerMmio", Value: "1"},
			{Key: "ImexChannelCount", Value: "2048"},
			{Key: "CreateImexChannel0", Value: "0"},
			{Key: "GrdmaPciTopoCheckOverride", Value: "0"},
			{Key: "CoherentGPUMemoryMode", Value: `"driver"`},
			{Key: "RegistryDwords", Value: `""`},
			{Key: "RegistryDwordsPerDevice", Value: `""`},
			{Key: "RmMsg", Value: `""`},
			{Key: "GpuBlacklist", Value: `""`},
			{Key: "TemporaryFilePath", Value: `""`},
			{Key: "ExcludedGpus", Value: `""`},
		},
	}))
}
