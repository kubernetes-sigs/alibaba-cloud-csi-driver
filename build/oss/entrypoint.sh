#!/bin/sh

# Get System version
ossfsVer="1.80.3"

# install OSSFS
mkdir -p /host/etc/csi-oss-tool/
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


updateConnector="true"
if [ ! -f "/host/usr/libexec/kubernetes/kubelet-plugins/volume/exec/alicloud~csiplugin-connector-init/csiplugin-connector" ];then
    mkdir -p /host/usr/libexec/kubernetes/kubelet-plugins/volume/exec/alicloud~csiplugin-connector-init/
    echo "mkdir alicloud~csiplugin-connector-init directory..."
else
    oldmd5=`md5sum /host/usr/libexec/kubernetes/kubelet-plugins/volume/exec/alicloud~csiplugin-connector-init/csiplugin-connector | awk '{print $1}'`
    newmd5=`md5sum /bin/csiplugin-connector | awk '{print $1}'`
    if [ "$oldmd5" = "$newmd5" ]; then
        updateConnector="false"
    else
        rm -rf /host/usr/libexec/kubernetes/kubelet-plugins/volume/exec/alicloud~csiplugin-connector-init
        rm -rf /host/usr/libexec/kubernetes/connector.sock
        rm -rf /var/log/alicloud/connector.pid
        mkdir -p /host/usr/libexec/kubernetes/kubelet-plugins/volume/exec/alicloud~csiplugin-connector-init
    fi
fi

if [ "$updateConnector" = "true" ]; then
    cp /bin/csiplugin-connector /host/usr/libexec/kubernetes/kubelet-plugins/volume/exec/alicloud~csiplugin-connector-init/csiplugin-connector
    chmod 755 /host/usr/libexec/kubernetes/kubelet-plugins/volume/exec/alicloud~csiplugin-connector-init/csiplugin-connector
fi

# install csiplugin connector
updateConnectorInit="true"
if [ -f "/host/usr/libexec/kubernetes/kubelet-plugins/volume/exec/alicloud~csiplugin-connector-init/csiplugin-connector-init" ];then
    oldmd5=`md5sum /host/usr/libexec/kubernetes/kubelet-plugins/volume/exec/alicloud~csiplugin-connector-init/csiplugin-connector-init | awk '{print $1}'`
    newmd5=`md5sum /bin/csiplugin-connector-init | awk '{print $1}'`
    if [ "$oldmd5" = "$newmd5" ]; then
        updateConnectorInit="false"
    else
        rm -rf /host/usr/libexec/kubernetes/kubelet-plugins/volume/exec/alicloud~csiplugin-connector-init/csiplugin-connector-init
    fi
fi

if [ "$updateConnectorInit" = "true" ]; then
    cp /bin/csiplugin-connector-init /host/usr/libexec/kubernetes/kubelet-plugins/volume/exec/alicloud~csiplugin-connector-init/csiplugin-connector-init
    chmod 755 /host/usr/libexec/kubernetes/kubelet-plugins/volume/exec/alicloud~csiplugin-connector-init/csiplugin-connector-init
fi

/nsenter --mount=/proc/1/ns/mnt /usr/libexec/kubernetes/kubelet-plugins/volume/exec/alicloud~csiplugin-connector-init/csiplugin-connector-init init > /dev/null  2>&1

# start daemon
/bin/ossplugin.csi.alibabacloud.com $@
