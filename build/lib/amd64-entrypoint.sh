#!/bin/sh

run_oss="false"
run_disk="false"
run_nas="false"

mkdir -p /var/log/alicloud/
mkdir -p /host/etc/kubernetes/volumes/disk/uuid

HOST_CMD="/nsenter --mount=/proc/1/ns/mnt"


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
    if [[ `echo ${osID} | grep "anolis" | wc -l` != "0" ]]; then
        host_os="anolis"
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
        ossfsVer="1.88.0"
    fi

    ossfsArch="centos7.0"
    if [[ ${host_os} == "alinux3" ]]; then
        for((i=1;i<=10;i++));
        do
            ${HOST_CMD} yum install -y libcurl-devel libxml2-devel fuse-devel openssl-devel
            if [ $? -eq 0 ]; then
                break
            else
                echo "Starting retry again yum install -y libcurl-devel libxml2-devel fuse-devel openssl-devel.retry count:$i"
                sleep 2
            fi
        done
        ossfsArch="centos8"
    fi

		if [[ ${host_os} == "lifsea" ]] || [[ ${host_os} == "anolis" ]]; then
        ossfsArch="centos8"
    fi

    echo "Starting deploy oss csi-plugin..."
    echo "osHost:"${host_os}
    echo "ossfsVersion:"${ossfsVer}
    echo "ossfsArch:"${ossfsArch}

    # install OSSFS
    mkdir -p /host/etc/csi-tool/
		reconcileOssFS="skip"
    if [ ! `${HOST_CMD}  which ossfs` ]; then
        echo "First install ossfs, ossfsVersion: $ossfsVer"
        cp /root/ossfs_${ossfsVer}_${ossfsArch}_x86_64.rpm /host/etc/csi-tool/
				reconcileOssFS="install"
    # update OSSFS
    else
        echo "Check ossfs Version...."
        oss_info=`${HOST_CMD}  ossfs --version | grep -E -o "V[0-9.a-z]+" | cut -d"V" -f2`
        if [ "$oss_info" != "$ossfsVer" ]; then
            echo "Upgrade ossfs, ossfsVersion: $ossfsVer"
            ${HOST_CMD}  yum remove -y ossfs
            cp /root/ossfs_${ossfsVer}_${ossfsArch}_x86_64.rpm /host/etc/csi-tool/
						reconcileOssFS="upgrade"
        fi
    fi

	if [[ ${reconcileOssFS} == "install" ]]; then
      if [[ ${host_os} == "lifsea" ]]; then
          rpm2cpio /root/ossfs_${ossfsVer}_${ossfsArch}_x86_64.rpm | cpio -idmv
          cp ./usr/local/bin/ossfs /host/etc/csi-tool/
          ${HOST_CMD} cp /etc/csi-tool/ossfs /usr/local/bin/ossfs
      else
          ${HOST_CMD} rpm -i /etc/csi-tool/ossfs_${ossfsVer}_${ossfsArch}_x86_64.rpm
      fi
    fi

    if [[ ${reconcileOssFS} == "upgrade" ]]; then
      if [[ ${host_os} == "lifsea" ]]; then
          ${HOST_CMD}  rm /usr/local/bin/ossfs
          rpm2cpio /root/ossfs_${ossfsVer}_${ossfsArch}_x86_64.rpm | cpio -idmv
          cp ./usr/local/bin/ossfs /host/etc/csi-tool/
          ${HOST_CMD}  cp /etc/csi-tool/ossfs /usr/local/bin/ossfs
      else
          ${HOST_CMD}  yum remove -y ossfs
          ${HOST_CMD}  rpm -i /etc/csi-tool/ossfs_${ossfsVer}_${ossfsArch}_x86_64.rpm
      fi
    fi

    # install Jindofs
    #if [ ! -f "/host/etc/jindofs-tool/jindo-fuse" ];then
    #    mkdir -p /host/etc/jindofs-tool/
    #    cp /jindo-fuse /host/etc/jindofs-tool/jindo-fuse
    #    echo "install jindofs..."
    #else
    #    oldmd5=`md5sum /host/etc/jindofs-tool/jindo-fuse | awk '{print $1}'`
    #    newmd5=`md5sum /jindo-fuse | awk '{print $1}'`
    #    if [ "$oldmd5" != "$newmd5" ]; then
    #        rm -rf /host/etc/jindofs-tool/jindo-fuse
    #        cp /jindo-fuse /host/etc/jindofs-tool/jindo-fuse
    #        echo "upgrade jindofs..."
    #    fi
    #fi
fi

# skip installing csiplugin-connector when DISABLE_CSIPLUGIN_CONNECTOR=true
if [ "$DISABLE_CSIPLUGIN_CONNECTOR" != "true" ] && ([ "$run_oss" = "true" ] || [ "$run_disk" = "true" ]); then
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
    cp /freezefs.sh /host/etc/csi-tool/freezefs.sh
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
        echo "Starting systemctl daemon-reload."
        for((i=1;i<=10;i++));
        do
            ${HOST_CMD} systemctl daemon-reload
            if [ $? -eq 0 ]; then
                break
            else
                echo "Starting retry again systemctl daemon-reload.retry count:$i"
                sleep 2
            fi
        done
    fi

    rm -rf /var/log/alicloud/connector.pid
    echo "Starting systemctl enable csiplugin-connector.service."
    for((i=1;i<=5;i++));
    do
        ${HOST_CMD} systemctl enable csiplugin-connector.service
        if [ $? -eq 0 ]; then
            break
        else
            echo "Starting retry again systemctl enable csiplugin-connector.service.retry count:$i"
            sleep 2
        fi
    done

    echo "Starting systemctl restart csiplugin-connector.service."
    for((i=1;i<=5;i++));
    do
        ${HOST_CMD} systemctl restart csiplugin-connector.service
        if [ $? -eq 0 ]; then
            break
        else
            echo "Starting retry again systemctl restart csiplugin-connector.service.retry count:$i"
            sleep 2
        fi
    done
fi

echo "Start checking if the rpm package needs to be installed"
if [ "$DISK_BDF_ENABLE" = "true" ] && [ "$run_disk" = "true" ]; then
    isbdf="false"
    for i in $(${HOST_CMD} lspci -D | grep "storage controller" | grep "1ded" | awk '{print $1}' |  sed -n '/0$/p'); 
    do 
        out=`${HOST_CMD} lspci -s $i -v`;
        if [[ $out == *"Single Root I/O Virtualization"* ]]; then
            isbdf="true"
            break
        fi
    done
    echo "isbdf node: $isbdf"
    if [ $isbdf = "true" ]; then
        echo "start install vfhp"
        ${HOST_CMD} yum install -y "http://yum.tbsite.net/taobao/7/x86_64/current/iohub-vfhp-helper/iohub-vfhp-helper-0.1.3-20230417103419.x86_64.rpm"
        if [ $? -ne 0 ]; then
            ${HOST_CMD} yum install -y "https://iohub-vfhp-helper.oss-rg-china-mainland.aliyuncs.com/iohub-vfhp-helper-0.1.3-20230417103419.x86_64.rpm"
        fi
        # takes 10s
        output=`${HOST_CMD} iohub-vfhp-helper -s`
        if [[ $output == *"backend support auto vf hotplug."* ]]; then
            echo "backend support auto vf hotplugin"
            ${HOST_CMD} sudo service iohub-vfhp-helper start
        else
            echo "backend not support auto vf hotplugin"
        fi
    fi
fi

## CPFS-NAS plugin setup
if [ "$run_nas" = "true" ]; then
    # cpfs-nas nas-rich-client common rpm
    cp /root/aliyun-alinas-utils-1.1-5.al7.noarch.rpm /host/etc/csi-tool/
    # nas-rich-client rpm
    cp /root/alinas-efc-1.2-2.x86_64.rpm /host/etc/csi-tool/
fi

## Jindofs plugin setup
#if [ "$run_oss" = "true" ]; then
#    # jindofs common rpm
#    ${HOST_CMD} yum install -y fuse3 fuse3-devel
#fi


# start daemon
/bin/plugin.csi.alibabacloud.com $@
