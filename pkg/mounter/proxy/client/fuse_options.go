package client

import (
	"os"
	"strconv"
	"strings"

	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

const nullVal string = "&NULLVAL"

var defaultFuseOptionsMap = map[string]string{
	"nodev":    nullVal,
	"nosuid":   nullVal,
	"rootmode": "40000", // S_IFDIR: FUSE root is a directory
	// Note: user_id and group_id are the owner of the FUSE filesystem. Together with
	// allow_other, they enable other users to access the FUSE filesystem. The uid and
	// gid parameters of ossfs are application-level parameters used to simulate file permissions.
	"allow_other":         nullVal,
	"default_permissions": nullVal,
	"user_id":             strconv.Itoa(os.Getuid()),
	"group_id":            strconv.Itoa(os.Getgid()),
}

// splitFuseOptions splits mount options into FUSE kernel options and daemon options.
// FUSE kernel options are used for the mount(2) call; everything else is passed
// to the client as daemon options.
//
// Merge rules:
//   - defaultFuseOptionsMap and mountFlags are BOTH placed into fuseOptions in full.
//     When their keys conflict, mountFlags win.
//   - For each entry in options: if its key already exists in fuseOptions (i.e. came
//     from defaults or mountFlags), options win and a warning is logged. Otherwise
//     the entry becomes a daemon option.
//
// "rw" and empty/blank entries are silently ignored throughout (the default mode is
// rw; we drop it to avoid conflicts with explicit "ro").
//
// TODO: splitFuseOptions currently only handles ossfs2's fd-passing requirements.
// If future clients have significantly different FUSE option support, this should
// be abstracted into the FuseMounterType interface so each client can implement
// its own option routing logic.
func splitFuseOptions(options []string, mountFlags []string) (fuseOptions, daemonOptions []string) {
	fuseOptionsMap := make(map[string]string)
	daemonOptionSet := sets.NewString()

	// 1. Initialize with default fuse options.
	for key, val := range defaultFuseOptionsMap {
		fuseOptionsMap[key] = val
	}

	// 2. Apply ALL mountFlags to fuseOptionsMap. Every entry in mountFlags becomes
	//    a FUSE kernel mount option; on key conflict with defaults, mountFlags win.
	for _, f := range mountFlags {
		setFuseOption(fuseOptionsMap, f)
	}

	// 3. Process options:
	//    - If the key already exists in fuseOptionsMap (defaults ∪ mountFlags),
	//      options take precedence; a warning is logged because FUSE-kernel options
	//      are normally configured via defaults or pv.spec.mountOptions, not via
	//      volumeAttributes.otherOpts.
	//    - Otherwise the entry becomes a daemon option.
	for _, o := range options {
		key, ok := parseOptionKey(o)
		if !ok {
			continue
		}
		if existing, present := fuseOptionsMap[key]; present {
			existingStr := existing
			if existing == nullVal {
				existingStr = "<flag>"
			}
			klog.Warningf("FUSE option %q in volumeAttributes.otherOpts overrides existing fuseOption (was %q); "+
				"FUSE-kernel options should normally be configured via defaults or pv.spec.mountOptions",
				o, existingStr)
			setFuseOption(fuseOptionsMap, o)
			continue
		}
		// NOTE: Daemon options may include FUSE-type mount parameters (e.g. exec/noexec)
		// that the client should either ignore or explicitly reject.
		// CSI does not validate these; the client is responsible for handling them.
		daemonOptionSet.Insert(o)
	}

	// Build fuse options list
	for k, v := range fuseOptionsMap {
		if v == nullVal {
			fuseOptions = append(fuseOptions, k)
		} else {
			fuseOptions = append(fuseOptions, k+"="+v)
		}
	}

	daemonOptions = daemonOptionSet.UnsortedList()
	return
}

// parseOptionKey extracts the key from an option string ("key" or "key=value").
// Returns ("", false) for empty / blank-key inputs and "rw" (the default mode,
// silently ignored to avoid conflicts with "ro").
func parseOptionKey(opt string) (string, bool) {
	kv := strings.SplitN(opt, "=", 2)
	if len(kv) == 0 || kv[0] == "" {
		return "", false
	}
	if kv[0] == "rw" {
		return "", false
	}
	return kv[0], true
}

// setFuseOption stores opt into fuseOptionsMap. If opt is "key=value", the value is
// recorded; if opt is "key" (flag-style), the value is set to nullVal. Empty entries
// and "rw" are silently skipped.
func setFuseOption(fuseOptionsMap map[string]string, opt string) {
	key, ok := parseOptionKey(opt)
	if !ok {
		return
	}
	kv := strings.SplitN(opt, "=", 2)
	if len(kv) == 2 {
		fuseOptionsMap[key] = kv[1]
	} else {
		fuseOptionsMap[key] = nullVal
	}
}
