#!/bin/sh

echo "Preparing deploy oss csi-plugin...."

HOST_CMD="/nsenter --mount=/proc/1/ns/mnt"

host_os="centos-7"
$HOST_CMD which lsb_release
lsb_release_exist=$?
echo $CSI_OSS_PLUGIN_TYPE
if [ "$CSI_OSS_PLUGIN_TYPE" != "provisioner" ]; then
    /bin/plugin.csi.alibabacloud.com $@
    exit 0
fi

if [ "$lsb_release_exist" != "0" ]; then
    $HOST_CMD ls /etc/os-release
    os_release_exist=$?
fi

if [ "$lsb_release_exist" = "0" ]; then
    os_info=`${HOST_CMD} lsb_release -a`
    if [ `echo $os_info | grep CentOS | grep 7. | wc -l` != "0" ]; then
        host_os="centos-7"
    elif [ `echo $os_info | grep 14.04 | wc -l` != "0" ]; then
        : # host_os="ubuntu-1404"
    elif [ `echo $os_info | grep 16.04 | wc -l` != "0" ]; then
        : # host_os="ubuntu-1604"
    elif [ `echo $os_info | grep 18.04 | wc -l` != "0" ]; then
        host_os="ubuntu-1804"
    elif [ `echo $os_info | grep Aliyun | wc -l` != "0" ]; then
        : # host_os="alios"
    elif [ `echo $os_info | grep Alibaba | wc -l` != "0" ]; then
        : # host_os="alios"
    fi

elif [ "$os_release_exist" = "0" ]; then
    osId=`$HOST_CMD cat /etc/os-release | grep "ID="`
    osVersion=`$HOST_CMD cat /etc/os-release | grep "VERSION_ID="`

    if [ `echo $osId | grep "centos" | wc -l` != "0" ]; then
        if [ `echo $osVersion | grep "7" | wc -l` = "1" ]; then
            host_os="centos-7"
        fi
    elif [ `echo $osId | grep "alios" | wc -l` != "0" ];then
        if [ `echo $osVersion | grep "7" | wc -l` = "1" ]; then
            host_os="centos-7"
        fi
    elif [ `echo $osId | grep "ubuntu" | wc -l` != "0" ]; then
        if [ `echo $osVersion | grep "14.04" | wc -l` != "0" ]; then
            : # host_os="ubuntu-1404"
        elif [ `echo $osVersion | grep "16.04" | wc -l` != "0" ]; then
            : # host_os="ubuntu-1604"
        elif [ `echo $osVersion | grep "16.04" | wc -l` != "0" ]; then
            host_os="ubuntu-1804"
        fi
    fi
fi


case $host_os in

    ubuntu-1804)
        ossfsVer="1.80.6"
        ossfsPackage=ossfs_${ossfsVer}_ubuntu18.04_amd64.deb
        cmd_localinstall="$HOST_CMD apt-get install -y"
        cmd_localremove="$HOST_CMD apt-get remove -y" 
    ;;

    *)
        ossfsVer="1.80.6.ack.1"
        ossfsPackage=ossfs_${ossfsVer}_centos7.0_x86_64.rpm
        cmd_localinstall="$HOST_CMD yum localinstall -y"
        cmd_localremove="$HOST_CMD yum remove -y"
    ;;

esac


echo "Starting deploy oss csi-plugin...."

# install OSSFS
mkdir -p /host/etc/csi-tool/
if [ ! `$HOST_CMD which ossfs` ]; then
    echo "First install ossfs...."
    cp /root/$ossfsPackage /host/etc/csi-tool/
    $cmd_localinstall /etc/csi-tool/$ossfsPackage
# update OSSFS
else
    echo "Check ossfs Version...."
    oss_info=`$HOST_CMD ossfs --version | grep -E -o "V[0-9.a-z]+" | cut -d"V" -f2`
    #vers_conut=`echo $oss_info | grep ${ossfsVer} | wc -l`
    #if [ "$vers_conut" = "0" ]; then
    if [ "$oss_info" != "$ossfsVer" ]; then
        echo "Upgrade ossfs...."
        $cmd_localremove ossfs
        cp /root/$ossfsPackage /host/etc/csi-tool/
        $cmd_localinstall /etc/csi-tool/$ossfsPackage
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
    mkdir -p /host/usr/lib/systemd/system
    cp /bin/csiplugin-connector.service /host/usr/lib/systemd/system/csiplugin-connector.service
    $HOST_CMD systemctl daemon-reload
fi

rm -rf /var/log/alicloud/connector.pid
$HOST_CMD systemctl enable csiplugin-connector.service
$HOST_CMD systemctl restart csiplugin-connector.service

# start daemon
/bin/plugin.csi.alibabacloud.com $@
