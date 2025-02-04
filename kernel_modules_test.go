package main_test

import (
	a "github.com/gogunit/gunit/hammy"
	"github.com/mlmon/surveyor"
	"github.com/mlmon/surveyor/source"
	"testing"
)

func Test_kernel_modules_failure_on_invalid_procPath(t *testing.T) {
	assert := a.New(t)
	_, err := main.KernelModules("testdata/proc-modules-absent", "testdata/modules")()
	assert.Is(a.Error(err))
}

func Test_kernel_modules_success_on_invalid_moduleBasePath(t *testing.T) {
	assert := a.New(t)
	_, err := main.KernelModules("testdata/proc-modules", "testdata/modules-absent")()
	assert.Is(a.Error(err))
}

func Test_kernel_modules_success(t *testing.T) {
	defer stubUname(unameStub)()
	
	assert := a.New(t)
	records, _ := main.KernelModules("testdata/proc-modules", "testdata/modules")()
	assert.Is(a.Struct(records).EqualTo(&source.Records{
		Source: "kernel-modules",
		Entries: source.Entries{
			{Key: "overlay"},
			{Key: "rpcsec_gss_krb5"},
			{Key: "auth_rpcgss"},
			{Key: "nfsv4"},
			{Key: "nfs"},
			{Key: "netfs"},
			{Key: "tls"},
			{Key: "gdrdrv"},
			{Key: "sch_fq_codel"},
			{Key: "nvidia_uvm"},
			{Key: "binfmt_misc"},
			{Key: "nf_conntrack"},
			{Key: "nf_defrag_ipv6"},
			{Key: "nf_defrag_ipv4"},
			{Key: "libcrc32c"},
			{Key: "lockd"},
			{Key: "grace"},
			{Key: "nls_iso8859_1"},
			{Key: "lnet"},
			{Key: "libcfs"},
			{Key: "sunrpc"},
			{Key: "crct10dif_pclmul"},
			{Key: "crc32_pclmul"},
			{Key: "polyval_clmulni"},
			{Key: "polyval_generic"},
			{Key: "ghash_clmulni_intel"},
			{Key: "sha256_ssse3"},
			{Key: "sha1_ssse3"},
			{Key: "aesni_intel"},
			{Key: "crypto_simd"},
			{Key: "cryptd"},
			{Key: "psmouse"},
			{Key: "input_leds"},
			{Key: "serio_raw"},
			{Key: "msr"},
			{Key: "nvidia_drm"},
			{Key: "nvidia_modeset"},
			{Key: "video"},
			{Key: "wmi"},
			{Key: "nvidia"},
			{Key: "ecc"},
			{Key: "ena"},
			{Key: "ib_iser"},
			{Key: "libiscsi"},
			{Key: "scsi_transport_iscsi"},
			{Key: "rdma_cm"},
			{Key: "iw_cm"},
			{Key: "ib_cm"},
			{Key: "dm_multipath"},
			{Key: "scsi_dh_rdac"},
			{Key: "scsi_dh_emc"},
			{Key: "scsi_dh_alua"},
			{Key: "efa"},
			{Key: "ib_uverbs"},
			{Key: "ib_core"},
			{Key: "parport_pc"},
			{Key: "ppdev"},
			{Key: "lp"},
			{Key: "parport"},
			{Key: "efi_pstore"},
			{Key: "ip_tables"},
			{Key: "x_tables"},
			{Key: "autofs4"},
		},
	}))
}
