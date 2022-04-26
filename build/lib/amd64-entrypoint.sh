#!/bin/sh

run_oss="false"
run_disk="false"
run_nas="false"

mkdir -p /var/log/alicloud/
mkdir -p /host/etc/kubernetes/volumes/disk/uuid

cp /csi/nsenter /nsenter
HOST_CMD="/nsenter --mount=/proc/1/ns/mnt"
zone_id=`${HOST_CMD} curl http://100.100.100.200/latest/meta-data/zone-id`


host_os="centos"
${HOST_CMD} ls /etc/os-release
os_release_exist=$?

if [[ "$os_release_exist" = "0" ]]; then
    osID=`${HOST_CMD} cat /etc/os-release | grep "ID=" | grep -v "VERSION_ID"`
    osVersion=`${HOST_CMD} cat /etc/os-release | grep "VERSION_ID=" | grep "^VERSION_ID=\"3"`
    if [[ `echo ${osID} | grep "alinux" | wc -l` != "0" ]] && [[ "${osVersion}" ]]; then
        host_os="alinux3"
    fi
		if [[ `echo ${osID} | grep "lifsea" | wc -l` != "0" ]]; then
        host_os="lifsea"
    fi
fi

## check which plugin is running
for item in $@;
do
    if [ "$item" = "--driver=ossplugin.csi.alibabacloud.com" ]; then
        echo "Running oss plugin...."
        run_oss="true"
        mkdir -p /var/lib/kubelet/csi-plugins/ossplugin.csi.alibabacloud.com
        rm -rf /var/lib/kubelet/plugins/ossplugin.csi.alibabacloud.com/csi.sock
    elif [ "$item" = "--driver=diskplugin.csi.alibabacloud.com" ]; then
        echo "Running disk plugin...."
				run_disk="true"
        mkdir -p /var/lib/kubelet/csi-plugins/diskplugin.csi.alibabacloud.com
        rm -rf /var/lib/kubelet/plugins/diskplugin.csi.alibabacloud.com/csi.sock
    elif [ "$item" = "--driver=nasplugin.csi.alibabacloud.com" ]; then
        echo "Running nas plugin...."
        run_nas="true"
        mkdir -p /var/lib/kubelet/csi-plugins/nasplugin.csi.alibabacloud.com
        rm -rf /var/lib/kubelet/plugins/nasplugin.csi.alibabacloud.com/csi.sock
    elif [[ $item==*--driver=* ]]; then
        tmp=${item}
        driver_types=${tmp#*--driver=}
        driver_type_array=(${driver_types//,/ })
        for driver_type in ${driver_type_array[@]};
        do
            if [ "$driver_type" = "oss" ]; then
                echo "Running oss plugin...."
                run_oss="true"
                mkdir -p /var/lib/kubelet/csi-plugins/ossplugin.csi.alibabacloud.com
                rm -rf /var/lib/kubelet/plugins/ossplugin.csi.alibabacloud.com/csi.sock
            elif [ "$driver_type" = "disk" ]; then
                echo "Running disk plugin...."
								run_disk="true"
                mkdir -p /var/lib/kubelet/csi-plugins/diskplugin.csi.alibabacloud.com
                rm -rf /var/lib/kubelet/plugins/diskplugin.csi.alibabacloud.com/csi.sock
            elif [ "$driver_type" = "nas" ]; then
                echo "Running nas plugin...."
                run_nas="true"
                mkdir -p /var/lib/kubelet/csi-plugins/nasplugin.csi.alibabacloud.com
                rm -rf /var/lib/kubelet/plugins/nasplugin.csi.alibabacloud.com/csi.sock
            fi
        done
    fi
done


## OSS plugin setup
if [ "$run_oss" = "true" ]; then

    ossfsVer="1.80.6.ack.1"
    if [ "$USE_UPDATE_OSSFS" == "" ]; then
        ossfsVer="1.86.4"
    fi

    ossfsArch="centos7.0"
    if [[ ${host_os} == "alinux3" ]]; then
        ${HOST_CMD} yum install -y libcurl-devel libxml2-devel fuse-devel openssl-devel
        ossfsArch="centos8"
    fi

		if [[ ${host_os} == "lifsea" ]]; then
        ossfsArch="centos8"
    fi

    echo "Starting deploy oss csi-plugin..."
    echo "osHost:"${host_os}
    echo "ossfsVersion:"${ossfsVer}
    echo "ossfsArch:"${ossfsArch}

    # install OSSFS
    mkdir -p /host/etc/csi-tool/
		reconcileOssFS="skip"
    if [ ! `/nsenter --mount=/proc/1/ns/mnt which ossfs` ]; then
        echo "First install ossfs, ossfsVersion: $ossfsVer"
        cp /root/ossfs_${ossfsVer}_${ossfsArch}_x86_64.rpm /host/etc/csi-tool/
        # /nsenter --mount=/proc/1/ns/mnt yum localinstall -y /etc/csi-tool/ossfs_${ossfsVer}_${ossfsArch}_x86_64.rpm
				reconcileOssFS="install"
    # update OSSFS
    else
        echo "Check ossfs Version...."
        oss_info=`/nsenter --mount=/proc/1/ns/mnt ossfs --version | grep -E -o "V[0-9.a-z]+" | cut -d"V" -f2`
        if [ "$oss_info" != "$ossfsVer" ]; then
            echo "Upgrade ossfs, ossfsVersion: $ossfsVer"
            /nsenter --mount=/proc/1/ns/mnt yum remove -y ossfs
            cp /root/ossfs_${ossfsVer}_${ossfsArch}_x86_64.rpm /host/etc/csi-tool/
            # /nsenter --mount=/proc/1/ns/mnt yum localinstall -y /etc/csi-tool/ossfs_${ossfsVer}_${ossfsArch}_x86_64.rpm
						reconcileOssFS="upgrade"
        fi
    fi

		if [[ ${reconcileOssFS} == "install" ]]; then
      if [[ ${host_os} == "lifsea" ]]; then
          rpm2cpio /root/ossfs_${ossfsVer}_${ossfsArch}_x86_64.rpm | cpio -idmv
          cp ./usr/local/bin/ossfs /host/etc/csi-tool/
          /nsenter --mount=/proc/1/ns/mnt cp /etc/csi-tool/ossfs /usr/local/bin/ossfs
      else
          /nsenter --mount=/proc/1/ns/mnt yum localinstall -y /etc/csi-tool/ossfs_${ossfsVer}_${ossfsArch}_x86_64.rpm
      fi
    fi

    if [[ ${reconcileOssFS} == "upgrade" ]]; then
      if [[ ${host_os} == "lifsea" ]]; then
          /nsenter --mount=/proc/1/ns/mnt rm /usr/local/bin/ossfs
          rpm2cpio /root/ossfs_${ossfsVer}_${ossfsArch}_x86_64.rpm | cpio -idmv
          cp ./usr/local/bin/ossfs /host/etc/csi-tool/
          /nsenter --mount=/proc/1/ns/mnt cp /etc/csi-tool/ossfs /usr/local/bin/ossfs
      else
          /nsenter --mount=/proc/1/ns/mnt yum remove -y ossfs
          /nsenter --mount=/proc/1/ns/mnt yum localinstall -y /etc/csi-tool/ossfs_${ossfsVer}_${ossfsArch}_x86_64.rpm
      fi
    fi

    # install Jindofs
    if [ ! -f "/host/etc/jindofs-tool/jindofs-fuse" ];then
        mkdir -p /host/etc/jindofs-tool/
        cp /jindofs-fuse /host/etc/jindofs-tool/jindofs-fuse
        echo "install jindofs..."
    else
        oldmd5=`md5sum /host/etc/jindofs-tool/jindofs-fuse | awk '{print $1}'`
        newmd5=`md5sum /jindofs-fuse | awk '{print $1}'`
        if [ "$oldmd5" != "$newmd5" ]; then
            rm -rf /host/etc/jindofs-tool/jindofs-fuse
            cp /jindofs-fuse /host/etc/jindofs-tool/jindofs-fuse
            echo "upgrade jindofs..."
        fi
    fi

fi

if [ "$run_oss" = "true" ] || ["$run_disk" = "true"]; then
    ## install/update csi connector
    updateConnector="true"
		systemdDir="/host/usr/lib/systemd/system"
    if [[ ${host_os} == "lifsea" ]]; then
        systemdDir="/host/etc/systemd/system"
    fi
    if [ ! -f "/host/etc/csi-tool/csiplugin-connector" ];then
        mkdir -p /host/etc/csi-tool/
        echo "mkdir /etc/csi-tool/ directory..."
    else
        oldmd5=`md5sum /host/etc/csi-tool/csiplugin-connector | awk '{print $1}'`
        newmd5=`md5sum /csi/csiplugin-connector | awk '{print $1}'`
        if [ "$oldmd5" = "$newmd5" ]; then
            updateConnector="false"
        else
            rm -rf /host/etc/csi-tool/
            rm -rf /host/etc/csi-tool/connector.sock
            rm -rf /var/log/alicloud/connector.pid
            mkdir -p /host/etc/csi-tool/
        fi
    fi
		cp /csi/freezefs.sh /host/etc/csi-tool/freezefs.sh
    if [ "$updateConnector" = "true" ]; then
        echo "Install csiplugin-connector...."
        cp /csi/csiplugin-connector /host/etc/csi-tool/csiplugin-connector
        chmod 755 /host/etc/csi-tool/csiplugin-connector
    fi


    # install/update csiplugin connector service
    updateConnectorService="true"
    if [[ ! -z "${PLUGINS_SOCKETS}" ]];then
        sed -i 's/Restart=always/Restart=on-failure/g' /csi/csiplugin-connector.service
        sed -i '/^\[Service\]/a Environment=\"WATCHDOG_SOCKETS_PATH='"${PLUGINS_SOCKETS}"'\"' /csi/csiplugin-connector.service
        sed -i '/ExecStop=\/bin\/kill -s QUIT $MAINPID/d' /csi/csiplugin-connector.service
        sed -i '/^\[Service\]/a ExecStop=sh -xc "if [ x$MAINPID != x ]; then /bin/kill -s QUIT $MAINPID; fi"' /csi/csiplugin-connector.service
    fi
    if [ -f "$systemdDir/csiplugin-connector.service" ];then
        echo "Check csiplugin-connector.service...."
        oldmd5=`md5sum $systemdDir/csiplugin-connector.service | awk '{print $1}'`
        newmd5=`md5sum /csi/csiplugin-connector.service | awk '{print $1}'`
        if [ "$oldmd5" = "$newmd5" ]; then
            updateConnectorService="false"
        else
            rm -rf $systemdDir/csiplugin-connector.service
        fi
    fi

    if [ "$updateConnectorService" = "true" ]; then
        echo "Install csiplugin connector service...."
        cp /csi/csiplugin-connector.service $systemdDir/csiplugin-connector.service
        /nsenter --mount=/proc/1/ns/mnt systemctl daemon-reload
    fi

    rm -rf /var/log/alicloud/connector.pid
    /nsenter --mount=/proc/1/ns/mnt systemctl enable csiplugin-connector.service
    /nsenter --mount=/proc/1/ns/mnt systemctl restart csiplugin-connector.service
fi

## CPFS-NAS plugin setup
if [ "$run_nas" = "true" ]; then
    # cpfs-nas rpm
    cp /csi/aliyun-alinas-utils-1.1-2.al7.noarch.rpm /host/etc/csi-tool/
    # nas-rich-client rpm
    cp /csi/aliyun-alinas-utils-1.1-2.alios7.noarch.rpm /host/etc/csi-tool/
    cp /csi/alinas-unas-1.1-1.alios7.x86_64.rpm /host/etc/csi-tool/
fi


# start daemon
cp /csi/plugin.csi.alibabacloud.com /bin/plugin.csi.alibabacloud.com
/bin/plugin.csi.alibabacloud.com $@
