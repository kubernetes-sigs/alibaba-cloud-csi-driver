#!/bin/sh

echo "Preparing deploy jindofs csi-plugin...."

HOST_CMD="/nsenter --mount=/proc/1/ns/mnt"

echo "Starting deploy jindofs csi-plugin...."

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

# install Jindofs-fuse
if [ ! -f "/host/etc/csi-tool/jindofs-fuse" ];then
    mkdir -p /host/etc/csi-tool/
    cp /bin/jindofs-fuse /host/etc/csi-tool/
    echo "install jindofs to /etc/csi-tool/ ..."
else
    oldmd5=`md5sum /host/etc/csi-tool/jindofs-fuse | awk '{print $1}'`
    newmd5=`md5sum /bin/jindofs-fuse | awk '{print $1}'`
    if [ "$oldmd5" != "$newmd5" ]; then
        rm -rf /host/etc/csi-tool/jindofs-fuse
        cp /bin/jindofs-fuse /host/etc/csi-tool/
        echo "upgrade jindofs to /etc/csi-tool/ ..."
    fi
fi

if [ "$updateConnector" = "true" ]; then
    echo "copy csiplugin-connector...."
    cp /bin/csiplugin-connector /host/etc/csi-tool/csiplugin-connector
    chmod 755 /host/etc/csi-tool/csiplugin-connector
fi


# install/update csiplugin connector service
updateConnectorService="true"
if [[ ! -z "${PLUGINS_SOCKETS}" ]];then
    sed -i 's/Restart=always/Restart=on-failure/g' /bin/csiplugin-connector.service
    sed -i '/^\[Service\]/a Environment=\"WATCHDOG_SOCKETS_PATH='"${PLUGINS_SOCKETS}"'\"' /bin/csiplugin-connector.service
    sed -i '/ExecStop=\/bin\/kill -s QUIT $MAINPID/d' /bin/csiplugin-connector.service
    sed -i '/^\[Service\]/a ExecStop=sh -xc "if [ x$MAINPID != x ]; then /bin/kill -s QUIT $MAINPID; fi"' /bin/csiplugin-connector.service
fi
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
