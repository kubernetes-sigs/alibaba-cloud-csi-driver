#!/bin/sh


echo "Starting deploy oss csi-plugin...."

ossfsVer="1.80.6.ack.1"
# install OSSFS
mkdir -p /host/etc/csi-tool/
if [ ! `/nsenter --mount=/proc/1/ns/mnt which ossfs` ]; then
    echo "First install ossfs...."
    cp /root/ossfs_${ossfsVer}_centos7.0_x86_64.rpm /host/etc/csi-tool/
    /nsenter --mount=/proc/1/ns/mnt yum localinstall -y /etc/csi-tool/ossfs_${ossfsVer}_centos7.0_x86_64.rpm
# update OSSFS
else
    echo "Check ossfs Version...."
    oss_info=`/nsenter --mount=/proc/1/ns/mnt ossfs --version | grep -E -o "V[0-9.a-z]+" | cut -d"V" -f2`
    #vers_conut=`echo $oss_info | grep ${ossfsVer} | wc -l`
    #if [ "$vers_conut" = "0" ]; then
    if [ "$oss_info" != "$ossfsVer" ]; then
        echo "Upgrade ossfs...."
        /nsenter --mount=/proc/1/ns/mnt yum remove -y ossfs
        cp /root/ossfs_${ossfsVer}_centos7.0_x86_64.rpm /host/etc/csi-tool/
        /nsenter --mount=/proc/1/ns/mnt yum localinstall -y /etc/csi-tool/ossfs_${ossfsVer}_centos7.0_x86_64.rpm
    fi
fi

## install/update csi connector
updateConnector="true"
if [ ! -f "/host/etc/csi-tool/csiplugin-connector" ];then
    mkdir -p /host/etc/csi-tool/
    echo "mkdir /etc/csi-tool/ directory..."
else
    oldmd5=`md5sum /host/etc/csi-tool/csiplugin-connector | awk '{print $1}'`
    newmd5=`md5sum /bin/csiplugin-connector | awk '{print $1}'`
    if [ "$oldmd5" = "$newmd5" ]; then
        updateConnector="false"
    else
        rm -rf /host/etc/csi-tool/
        rm -rf /host/etc/csi-tool/connector.sock
        rm -rf /var/log/alicloud/connector.pid
        mkdir -p /host/etc/csi-tool/
    fi
fi

if [ "$updateConnector" = "true" ]; then
    echo "copy csiplugin-connector...."
    cp /bin/csiplugin-connector /host/etc/csi-tool/csiplugin-connector
    chmod 755 /host/etc/csi-tool/csiplugin-connector
fi


# install/update csiplugin connector service
updateConnectorService="true"
if [ -f "/host/usr/lib/systemd/system/csiplugin-connector.service" ];then
    echo "check upgrade csiplugin-connector.service...."
    oldmd5=`md5sum /host/usr/lib/systemd/system/csiplugin-connector.service | awk '{print $1}'`
    newmd5=`md5sum /bin/csiplugin-connector.service | awk '{print $1}'`
    if [ "$oldmd5" = "$newmd5" ]; then
        updateConnectorService="false"
    else
        rm -rf /host/usr/lib/systemd/system/csiplugin-connector.service
    fi
fi

if [ "$updateConnectorService" = "true" ]; then
    echo "install csiplugin connector service...."
    cp /bin/csiplugin-connector.service /host/usr/lib/systemd/system/csiplugin-connector.service
    /nsenter --mount=/proc/1/ns/mnt systemctl daemon-reload
fi

rm -rf /var/log/alicloud/connector.pid
/nsenter --mount=/proc/1/ns/mnt systemctl enable csiplugin-connector.service
/nsenter --mount=/proc/1/ns/mnt systemctl restart csiplugin-connector.service

# start daemon
/bin/plugin.csi.alibabacloud.com $@
