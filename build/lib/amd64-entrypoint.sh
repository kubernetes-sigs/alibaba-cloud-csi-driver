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
    if [[ `echo ${osID} | grep "alinux" | wc -l` != "0" ]]; then
        host_os="alinux$(eval $($HOST_CMD cat /etc/os-release); echo ${VERSION_ID} | cut -d'.' -f1)"
    fi
		if [[ `echo ${osID} | grep "lifsea" | wc -l` != "0" ]]; then
        host_os="lifsea"
    fi
    if [[ `echo ${osID} | grep "anolis" | wc -l` != "0" ]]; then
        osVersion=`${HOST_CMD} cat /etc/os-release | grep "VERSION_ID=" | grep "^VERSION_ID=\"8"`
        if [[ "${osVersion}" ]]; then
            host_os="anolis8"
        fi
    fi
fi
echo "detected host os: $host_os"

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
            elif [ "$driver_type" = "pov" ]; then
                echo "Running pov plugin...."
                run_pov="true"
            fi
        done
    fi
done

# config /etc/updatedb.config if needed
if [ -z "$KUBELET_ROOT_DIR" ]; then
    KUBELET_ROOT_DIR="/var/lib/kubelet"
fi
if [ "$SKIP_UPDATEDB_CONFIG" != "true" ]; then
    ## check cron.daily dir
    if [ -f /host/etc/cron.daily/mlocate ]; then
        if [ -f /host/etc/updatedb.conf ]; then
            PRUNEFS=`sed '/^PRUNEFS =/!d; s/.*= //' /host/etc/updatedb.conf | sed 's/\"//g'`
            PRUNEPATHS=`sed '/^PRUNEPATHS =/!d; s/.*= //' /host/etc/updatedb.conf | sed 's/\"//g'`

            sed -i 's/PRUNE_BIND_MOUNTS.*$/PRUNE_BIND_MOUNTS = \"yes\"/g' /host/etc/updatedb.conf
            if [[ `echo ${PRUNEFS} | grep "fuse.ossfs" | wc -l` == "0" ]]; then
                echo "original PRUNEFS=${PRUNEFS}"
                echo "add PRUNEFS:fuse.ossfs in /etc/updatedb.conf"
                PRUNEFS="\"${PRUNEFS} fuse.ossfs\""
                sed -i 's/PRUNEFS.*$/PRUNEFS = '"${PRUNEFS}"'/g' /host/etc/updatedb.conf
            fi
            if [[ `echo ${PRUNEPATHS} | grep ${KUBELET_ROOT_DIR} | wc -l` == "0" ]]; then
                echo "original PRUNEPATHS=${PRUNEPATHS}"
                if [ "$KUBELET_ROOT_DIR" == "/var/lib/kubelet" ]; then
                    PRUNEPATHS="\"${PRUNEPATHS} /var/lib/kubelet /var/lib/container/kubelet\""
                    echo "add PRUNEPATHS:/var/lib/kubelet /var/lib/container/kubelet in /etc/updatedb.conf"
                else
                    PRUNEPATHS="\"${PRUNEPATHS} ${KUBELET_ROOT_DIR}\""
                    echo "add PRUNEPATHS:${KUBELET_ROOT_DIR} in /etc/updatedb.conf"
                fi
                sed -i 's#PRUNEPATHS.*$#PRUNEPATHS = '"${PRUNEPATHS}"'#g' /host/etc/updatedb.conf
            fi
        fi
    fi
fi

# skip installing csiplugin-connector when DISABLE_CSIPLUGIN_CONNECTOR=true
if [ "$DISABLE_CSIPLUGIN_CONNECTOR" != "true" ] && ([ "$run_oss" = "true" ] || [ "$run_disk" = "true" ] || [ "$run_nas" = "true" ]); then
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
            rm -rf /var/run/csiplugin/connector.pid
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
    rm -rf /var/run/csiplugin/connector.pid
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
if ([ "$DISK_BDF_ENABLE" = "true" ] && [ "$run_disk" = "true" ]) || [ "$run_pov" = "true" ]; then
    isbdf="false"
    for i in $(${HOST_CMD} lspci -D | grep "storage controller" | grep "1ded\|Alibaba" | awk '{print $1}' |  sed -n '/0$/p'); 
    do 
        out=`${HOST_CMD} lspci -s $i -v`;
        if [[ $out == *"Single Root I/O Virtualization"* ]]; then
            isbdf="true"
            break
        fi
    done
    export IS_BDF="$isbdf"
    echo "isbdf node: $isbdf"
    if [ $isbdf = "true" ]; then
        echo "start install vfhp"
        ${HOST_CMD} yum install -y "http://yum.tbsite.net/taobao/7/x86_64/current/iohub-vfhp-helper/iohub-vfhp-helper-0.1.8-20230811073719.x86_64.rpm"
        if [ $? -ne 0 ]; then
            ${HOST_CMD} yum install -y "https://iohub-vfhp-helper.oss-rg-china-mainland.aliyuncs.com/iohub-vfhp-helper-0.1.8-20230811073719.x86_64.rpm"
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
    cp /root/aliyun-alinas-utils-1.1-6.al7.noarch.rpm /host/etc/csi-tool/
    # nas-rich-client rpm
    cp /root/alinas-efc-1.2-3.x86_64.rpm /host/etc/csi-tool/

    install_utils="false"
    install_efc="false"

    if test -e /etc/csi-plugin/config/cpfs-nas-enable; then
        install_utils="true"
    fi

    client_properties="/etc/csi-plugin/config/cnfs-client-properties"
    if test -e $client_properties && grep -qE "enable=true|efc=true" $client_properties; then
        install_utils="true"
        install_efc="true"
    fi

    if [ $install_utils = "true" ]; then
        if ${HOST_CMD} rpm -q aliyun-alinas-utils; then 
            echo "upgrade aliyun-alinas-utils"
            ${HOST_CMD} rpm -Uvh /etc/csi-tool/aliyun-alinas-utils-1.1-6.al7.noarch.rpm
        else
            echo "install aliyun-alinas-utils"
            ${HOST_CMD} yum install -y /etc/csi-tool/aliyun-alinas-utils-1.1-6.al7.noarch.rpm
        fi
        if ! ${HOST_CMD} rpm -q aliyun-alinas-utils | grep -E "^aliyun-alinas-utils-1.1-6" >/dev/null; then
            echo "failed to install aliyun-alinas-utils"
            exit 1
        fi
    fi

    if [ $host_os != "alinux2" ] && [ $install_efc = "true" ]; then
        echo "WARN: skip install efc because host os is not alinux2"
        install_efc="false"
    fi
    if [ $install_efc = "true" ]; then
        echo "installing alinas-efc"
        ${HOST_CMD} rpm -Uvh /etc/csi-tool/alinas-efc-1.2-3.x86_64.rpm
        echo "checking alinas-efc-1.2-3.x86_64 installed"
        ${HOST_CMD} rpm -q alinas-efc-1.2-3.x86_64 || exit 1
        echo "starting aliyun-alinas-mount-watchdog"
        ${HOST_CMD} systemctl start aliyun-alinas-mount-watchdog || exit 1
    fi
fi

## Jindofs plugin setup
#if [ "$run_oss" = "true" ]; then
#    # jindofs common rpm
#    ${HOST_CMD} yum install -y fuse3 fuse3-devel
#fi


# start daemon
/bin/plugin.csi.alibabacloud.com $@
