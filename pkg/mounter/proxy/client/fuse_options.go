package client

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"k8s.io/apimachinery/pkg/util/sets"
)

const NULLVAL string = "&NULLVAL"

var defaultFuseOptionsMap = map[string]string{
	"nodev":    NULLVAL,
	"nosuid":   NULLVAL,
	"rootmode": "40000",
	// Note: user_id and group_id are the owner of the FUSE filesystem. Together with
	// allow_other, they enable other users to access the FUSE filesystem. The uid and
	// gid parameters of ossfs are application-level parameters used to simulate file permissions.
	"allow_other":         NULLVAL,
	"default_permissions": NULLVAL,
	"user_id":             strconv.Itoa(os.Getuid()),
	"group_id":            strconv.Itoa(os.Getgid()),
}

type FuseOptionType int

const (
	SupportedOption FuseOptionType = iota
	UnsupportedOption
	IgnoredOption
)

// Refer: https://man7.org/linux/man-pages/man8/mount.fuse3.8.html
// Note: Disable dev and suid options for security reasons. Disable rw option
// because it's the default mode, and we only need to set ro to disable write
// access, avoiding conflicts with ro settings from other mount points.
var fuseOptionKeys = map[string]FuseOptionType{
	"exec":         SupportedOption,
	"noexec":       SupportedOption,
	"atime":        SupportedOption,
	"noatime":      SupportedOption,
	"sync":         SupportedOption,
	"async":        SupportedOption,
	"dirsync":      SupportedOption,
	"debug":        SupportedOption,
	"ro":           SupportedOption,
	"dev":          UnsupportedOption,
	"suid":         UnsupportedOption,
	"auto_unmount": UnsupportedOption,
	"rw":           IgnoredOption,
}

func splitFuseOptions(options []string) (fuseOptions, daemonOptions []string, err error) {
	fuseOptionsMap := make(map[string]string)
	optionSet := sets.New(options...)

	// inital with default fuse options
	for key, val := range defaultFuseOptionsMap {
		fuseOptionsMap[key] = val
	}

	for _, o := range options {
		kv := strings.SplitN(o, "=", 2)
		// empty kv, invalid
		if len(kv) == 0 {
			optionSet.Delete(o)
			continue
		}
		key := kv[0]
		if key == "" {
			optionSet.Delete(o)
			continue
		}
		v, ok := defaultFuseOptionsMap[key]
		// if the key is default, compare the value
		// if not equal, update the defualt value
		if ok {
			if v != NULLVAL && len(kv) == 2 && v != kv[1] {
				fuseOptionsMap[key] = kv[1]
			}
			continue
		}
		// if the key is not for fuse, see as daemon options
		t, ok := fuseOptionKeys[key]
		if !ok {
			continue
		}
		switch t {
		case SupportedOption:
			fuseOptionsMap[key] = NULLVAL
		case IgnoredOption:
			optionSet.Delete(o)
		case UnsupportedOption:
			return nil, nil, fmt.Errorf("Unsupported option: %s", o)
		}
	}
	daemonOptions = optionSet.UnsortedList()
	for k, v := range fuseOptionsMap {
		if v == NULLVAL {
			fuseOptions = append(fuseOptions, k)
		} else {
			fuseOptions = append(fuseOptions, fmt.Sprintf("%s=%s", k, v))
		}
	}

	return
}
