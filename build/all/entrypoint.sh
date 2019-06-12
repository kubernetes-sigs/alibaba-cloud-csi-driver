#!/bin/sh

run_oss="false"

mkdir -p /var/log/alicloud/

for item in $@;
do
  if [ "$item" = "--driver=ossplugin.csi.alibabacloud.com" ]; then
      echo "Running oss plugin...."
      run_oss="true"
  elif [ "$item" = "--driver=diskplugin.csi.alibabacloud.com" ]; then
      echo "Running disk plugin...."
  elif [ "$item" = "--driver=nasplugin.csi.alibabacloud.com" ]; then
      echo "Running nas plugin...."
  fi
done

if [ "$run_oss" = "true" ]; then
    echo "Starting deploy oss plugin...."

    # Get System version
    ossfsVer="1.80.3"
    # install OSSFS
    mkdir -p /host/etc/csi-tool/
    if [ ! `/nsenter --mount=/proc/1/ns/mnt which ossfs` ]; then
        echo "First install ossfs...."
        cp /root/ossfs_${ossfsVer}_centos7.0_x86_64.rpm /host/etc/csi-tool/
        /nsenter --mount=/proc/1/ns/mnt yum localinstall -y /etc/csi-tool/ossfs_${ossfsVer}_centos7.0_x86_64.rpm
    # update OSSFS
    else
        echo "Prepare Upgrade ossfs...."
        oss_info=`/nsenter --mount=/proc/1/ns/mnt ossfs --version`
        vers_conut=`echo $oss_info | grep ${ossfsVer} | wc -l`
        if [ "$vers_conut" = "0" ]; then
            echo "Upgrade ossfs...."
            /nsenter --mount=/proc/1/ns/mnt yum remove -y ossfs
            cp /root/ossfs_${ossfsVer}_centos7.0_x86_64.rpm /host/etc/csi-tool/
            /nsenter --mount=/proc/1/ns/mnt yum localinstall -y /etc/csi-tool/ossfs_${ossfsVer}_centos7.0_x86_64.rpm
        fi
    fi

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

    # install csiplugin connector service
    updateConnectorService="true"
    if [ -f "/host/usr/lib/systemd/system/csiplugin-connector.service" ];then
        echo "prepare install csiplugin-connector.service...."
        oldmd5=`md5sum /host/usr/lib/systemd/system/csiplugin-connector.service | awk '{print $1}'`
        newmd5=`md5sum /bin/csiplugin-connector.service | awk '{print $1}'`
        if [ "$oldmd5" = "$newmd5" ]; then
            updateConnectorService="false"
        else
            rm -rf /host/usr/lib/systemd/system/csiplugin-connector.service
        fi
    fi

    if [ "$updateConnectorService" = "true" ]; then
        echo "install csiplugin-connector...."
        cp /bin/csiplugin-connector.service /host/usr/lib/systemd/system/csiplugin-connector.service
    fi

    #/nsenter --mount=/proc/1/ns/mnt service csiplugin-connector-svc restart
    rm -rf /var/log/alicloud/connector.pid
    /nsenter --mount=/proc/1/ns/mnt systemctl enable csiplugin-connector.service
    /nsenter --mount=/proc/1/ns/mnt systemctl restart csiplugin-connector.service
fi

# start daemon
/bin/plugin.csi.alibabacloud.com $@
