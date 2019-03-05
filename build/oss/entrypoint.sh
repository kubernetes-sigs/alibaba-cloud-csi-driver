#!/bin/sh

# Get System version
ossfsVer="1.80.3"

# install OSSFS
mkdir /host/etc/csi-oss-tool/
if [ ! `/nsenter --mount=/proc/1/ns/mnt which ossfs` ]; then
    cp /root/ossfs_${ossfsVer}_centos7.0_x86_64.rpm /host/etc/csi-oss-tool/
    /nsenter --mount=/proc/1/ns/mnt yum localinstall -y /host/etc/csi-oss-tool/ossfs_${ossfsVer}_centos7.0_x86_64.rpm
# update OSSFS
else
    oss_info=`/nsenter --mount=/proc/1/ns/mnt ossfs --version`
    vers_conut=`echo $oss_info | grep ${ossfsVer} | wc -l`
    if [ "$vers_conut" = "0" ]; then
        /nsenter --mount=/proc/1/ns/mnt yum remove -y ossfs
        cp /root/ossfs_${ossfsVer}_centos7.0_x86_64.rpm /host/etc/csi-oss-tool/
        /nsenter --mount=/proc/1/ns/mnt yum localinstall -y /host/etc/csi-oss-tool/ossfs_${ossfsVer}_centos7.0_x86_64.rpm
    fi
fi

/bin/ossplugin.csi.alibabacloud.com $@
