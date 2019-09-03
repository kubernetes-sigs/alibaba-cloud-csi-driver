#!/bin/sh

echo "Starting deploy cpfs csi-plugin...."

# start daemon
/bin/plugin.csi.alibabacloud.com $@
