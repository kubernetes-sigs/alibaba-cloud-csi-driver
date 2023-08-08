#!/bin/sh

run_oss="false"
run_disk="false"

mkdir -p /var/log/alicloud/
mkdir -p /host/etc/kubernetes/volumes/disk/uuid
mv /acs/fuse-2.9.2-11.el7.aarch64.rpm /var/lib/kubelet/fuse-2.9.2-11.el7.aarch64.rpm
mv /acs/fuse-libs-2.9.2-11.el7.aarch64.rpm /var/lib/kubelet/fuse-libs-2.9.2-11.el7.aarch64.rpm
mv /acs/which-2.20-7.el7.aarch64.rpm /var/lib/kubelet/which-2.20-7.el7.aarch64.rpm
ossfsVer="1.80.6.ack.1"
armfsVer="1.80.6"

HOST_CMD="/nsenter --mount=/proc/1/ns/mnt"

host_os="centos"
${HOST_CMD} ls /etc/os-release
os_release_exist=$?

#get host os type
if [[ "$os_release_exist" = "0" ]]; then
    osID=`${HOST_CMD} cat /etc/os-release | grep "ID=" | grep -v "VERSION_ID"`
    if [[ `echo ${osID} | grep "alinux" | wc -l` != "0" ]]; then
        osVersion=`${HOST_CMD} cat /etc/os-release | grep "VERSION_ID=" | grep "^VERSION_ID=\"3"`
        if [[ "${osVersion}" ]]; then
            host_os="alinux3"
        fi
    fi
    if [[ `echo ${osID} | grep "kylin" | wc -l` != "0" ]]; then
        host_os="kylin"
    fi
    if [[ `echo ${osID} | grep "uos" | wc -l` != "0" ]] ; then
        host_os="uos"
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

#update oss file in kylin/uos
ossPath=/usr/bin/ossfs
if [[ ${host_os} == "kylin" ]] || [[ ${host_os} == "uos" ]] || [[ ${host_os} == "anolis8" ]] || [[ ${host_os} == "alinux3" ]]; then
   ossPath=/usr/bin/ossfs-8u
fi

## check which plugin is running
for item in $@;
do
  if [ "$item" = "--driver=ossplugin.csi.alibabacloud.com" ]; then
      echo "Running oss plugin...."
      run_oss="true"
      mkdir -p /var/lib/kubelet/csi-plugins/ossplugin.csi.alibabacloud.com
      rm -rf /var/lib/kubelet/plugins/ossplugin.csi.alibabacloud.com/csi.sock
      /usr/bin/nsenter yum localinstall -y /var/lib/kubelet/which-2.20-7.el7.aarch64.rpm
			/usr/bin/nsenter yum localinstall -y /var/lib/kubelet/fuse-libs-2.9.2-11.el7.aarch64.rpm
      /usr/bin/nsenter yum localinstall -y /var/lib/kubelet/fuse-2.9.2-11.el7.aarch64.rpm
      if [ ! `/usr/bin/nsenter --mount=/proc/1/ns/mnt which ossfs` ]; then
          echo "First install ossfs...."
          cp $ossPath /host/usr/bin/
          echo "cp result -- `/usr/bin/nsenter --mount=/proc/1/ns/mnt which ossfs` --"
      else
          echo "ossfs is already on host"
      fi
  elif [ "$item" = "--driver=diskplugin.csi.alibabacloud.com" ]; then
      echo "Running disk plugin...."
			run_disk="true"
      mkdir -p /var/lib/kubelet/csi-plugins/diskplugin.csi.alibabacloud.com
      rm -rf /var/lib/kubelet/plugins/diskplugin.csi.alibabacloud.com/csi.sock
  elif [ "$item" = "--driver=nasplugin.csi.alibabacloud.com" ]; then
      echo "Running nas plugin...."
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
							/usr/bin/nsenter yum localinstall -y /var/lib/kubelet/which-2.20-7.el7.aarch64.rpm
							/usr/bin/nsenter yum localinstall -y /var/lib/kubelet/fuse-libs-2.9.2-11.el7.aarch64.rpm
              /usr/bin/nsenter yum localinstall -y /var/lib/kubelet/fuse-2.9.2-11.el7.aarch64.rpm
							if [ ! `/usr/bin/nsenter --mount=/proc/1/ns/mnt which ossfs` ]; then
									echo "First install ossfs...."
									cp $ossPath /host/usr/bin/
									echo "cp result -- `/usr/bin/nsenter --mount=/proc/1/ns/mnt which ossfs` --"
							else
									echo "ossfs is already on host"
							fi
          elif [ "$driver_type" = "disk" ]; then
              echo "Running disk plugin...."
							run_disk="true"
              mkdir -p /var/lib/kubelet/csi-plugins/diskplugin.csi.alibabacloud.com
              rm -rf /var/lib/kubelet/plugins/diskplugin.csi.alibabacloud.com/csi.sock
          elif [ "$driver_type" = "nas" ]; then
              echo "Running nas plugin...."
              mkdir -p /var/lib/kubelet/csi-plugins/nasplugin.csi.alibabacloud.com
              rm -rf /var/lib/kubelet/plugins/nasplugin.csi.alibabacloud.com/csi.sock
          fi
      done
  fi
done


## OSS plugin setup
if [ "$run_oss" = "true" ]; then
    echo "Starting deploy oss csi-plugin...."
    #update libstdc++ for arm kylin/uos
    if [[ ${host_os} == "kylin" ]] || [[ ${host_os} == "uos" ]]; then
      if [[ `strings /usr/lib64/libstdc++.so.6 | grep GLIBCXX | grep 3.4.26 | wc -l` != "1" ]]; then
        echo "update libstdc++ version, link libstdc++.so.6.0.28 to /lib64"
        /bin/cp -f /acs/libstdc++.so.6.0.28 /host/usr/lib64/libstdc++.so.6.0.28
        /nsenter --mount=/proc/1/ns/mnt unlink /lib64/libstdc++.so.6
        /nsenter --mount=/proc/1/ns/mnt ln -sf /lib64/libstdc++.so.6.0.28 /lib64/libstdc++.so.6
      fi
    fi

		if [ -e "/host/usr/lib64/libfuse.so.2" ]; then
        echo "libfuse exists in host"
    else
				echo "copy libfuse.so.2 to /host/usr/lib64"
				cp /acs/libfuse.so.2 /host/usr/lib64/
    fi

    systemdDir="/host/usr/lib/systemd/system"
    if [[ ${host_os} == "lifsea" ]]; then
        systemdDir="/host/etc/systemd/system"
    fi
    # install OSSFS
    mkdir -p /host/etc/csi-tool/
    if [ ! `/nsenter --mount=/proc/1/ns/mnt which ossfs` ]; then
        echo "First install ossfs...."
        /nsenter --mount=/proc/1/ns/mnt yum install -y ossfs
        /nsenter --mount=/proc/1/ns/mnt ln -sf $ossPath /usr/local/bin/ossfs
    # update OSSFS
    else
        echo "Check ossfs Version...."
        oss_info=`/nsenter --mount=/proc/1/ns/mnt ossfs --version | grep -E -o "V[0-9.a-z]+" | cut -d"V" -f2`
        if [ "$oss_info" != "$ossfsVer" ] || [ "$oss_info" != "$armfsVer" ]; then
            echo "Upgrade ossfs...."
            /nsenter --mount=/proc/1/ns/mnt yum remove -y ossfs
            /nsenter --mount=/proc/1/ns/mnt yum install -y ossfs
            /nsenter --mount=/proc/1/ns/mnt ln -sf $ossPath /usr/local/bin/ossfs
        fi
    fi
fi

if [ "$run_disk" = "true" ] || [ "$run_oss" = "true" ]; then
    updateConnector="true"
    if [ ! -f "/host/etc/csi-tool/csiplugin-connector" ]; then
      mkdir -p /host/etc/csi-tool/
			echo "mkdir /etc/csi-tool/ directory ...."
		else
			oldmd5=`md5sum /host/etc/csi-tool/csiplugin-connector | awk '{print $1}'`
			newmd5=`md5sum /bin/csiplugin-connector | awk '{print $1}'`
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
    if [ -f "$systemdDir/csiplugin-connector.service" ];then
        echo "Check csiplugin-connector.service...."
        oldmd5=`md5sum $systemdDir/csiplugin-connector.service | awk '{print $1}'`
        newmd5=`md5sum /bin/csiplugin-connector.service | awk '{print $1}'`
        if [ "$oldmd5" = "$newmd5" ]; then
            updateConnectorService="false"
        else
            rm -rf $systemdDir/csiplugin-connector.service
        fi
    fi

    if [ "$updateConnectorService" = "true" ]; then
        echo "Install csiplugin connector service...."
        cp /bin/csiplugin-connector.service $systemdDir/csiplugin-connector.service
        /nsenter --mount=/proc/1/ns/mnt systemctl daemon-reload
    fi

    rm -rf /var/log/alicloud/connector.pid
    rm -rf /var/run/csiplugin/connector.pid
    /nsenter --mount=/proc/1/ns/mnt systemctl enable csiplugin-connector.service
    /nsenter --mount=/proc/1/ns/mnt systemctl restart csiplugin-connector.service

fi


# start daemon
/bin/plugin.csi.alibabacloud.com $@
