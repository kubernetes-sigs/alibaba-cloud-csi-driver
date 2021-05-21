package main

import "testing"
import "os"

func TestCheckCsiCmd(t *testing.T) {

	dir := "/var/lib/kubelet/pods/29125c96-1d6e-11ea-842e-00163e062fe1/volumes/kubernetes.io~csi/oss-csi-pv/mount"
	dir = "."

	// expect pass cases
	optionsList := []string{"-o max_stat_cache_size=0 -o allow_other",
		"",
		"-d",
		"--debug",
		"--debug -s -d",
		"-o max_stat_cache_size=0 -o allow_other -d",
		"  -o   max_stat_cache_size=0   -o   allow_other  ",
		"-o max_stat_cache_size=0 -o allow_other --debug",
		"-o max_stat_cache_size=0 --debug -o allow_other",
		"-o max_stat_cache_size=0 -s -o allow_other",
		"-o max_stat_cache_size=0 -o allow_other -o ro",
		"-omax_stat_cache_size=0 -oallow_other -o ro",
		"-omax_stat_cache_size=0 -oallow_other -oro",
		"-omax_stat_cache_size=0 -o ro,noresport -oallow_other ",
		"-o max_stat_cache_size=0 -oallow_other -o ro"}

	for _, options := range optionsList {
		cmdLine := "systemd-run --scope -- /usr/local/bin/ossfs gene-shenzhen:/fastq " + dir + " -ourl=oss-cn-shenzhen-internal.aliyuncs.com " + options
		os.MkdirAll(dir, 644)

		if err := checkOssfsCmd(cmdLine); err != nil {
			t.Errorf("Test with No Expect Error: %s, %s", err.Error(), options)
		}
	}

	// expect fail cases
	optionsList = []string{"-o -d max_stat_cache_size=0 -o allow_other",
		"-o max_stat_cache_size=0 debug -o allow_other",
		"-o -d max_stat_cache_size=0 -o allow_other -s",
		"-o -omax_stat_cache_size=0 -o allow_other -s",
		"--omax_stat_cache_size=0 -o allow_other -s",
		"-o max_stat_cache_size=0 -o -o allow_other"}
	for _, options := range optionsList {
		cmdLine := "systemd-run --scope -- /usr/local/bin/ossfs gene-shenzhen:/fastq " + dir + " -ourl=oss-cn-shenzhen-internal.aliyuncs.com " + options
		os.MkdirAll(dir, 644)

		if err := checkOssfsCmd(cmdLine); err != nil {
			t.Log("Test Success with Output: ", err.Error(), options)
		} else {
			t.Errorf("Test Fail with no error: %s", options)
		}
	}
}
