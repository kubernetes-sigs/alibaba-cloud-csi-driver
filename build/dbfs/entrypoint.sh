#!/bin/sh

mkdir -p /var/log/alicloud/

# start daemon
/bin/plugin.csi.alibabacloud.com $@
