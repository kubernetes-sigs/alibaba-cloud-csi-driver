#!/bin/sh

echo "Starting deploy cpfs csi-plugin...."
mkdir -p /host/etc/cpfs-tools


host_os="centos-7"
/acs/nsenter --mount=/proc/1/ns/mnt which lsb_release
lsb_release_exist=$?
if [ "$lsb_release_exist" != "0" ]; then
  /acs/nsenter --mount=/proc/1/ns/mnt ls /etc/os-release
  os_release_exist=$?
fi
if [ "$lsb_release_exist" = "0" ]; then
    os_info=`/acs/nsenter --mount=/proc/1/ns/mnt lsb_release -a`
    if [ `echo $os_info | grep CentOS | grep 7. | wc -l` != "0" ]; then
        host_os="centos-7"
    elif [ `echo $os_info | grep 14.04 | wc -l` != "0" ]; then
        host_os="ubuntu-1404"
    elif [ `echo $os_info | grep 16.04 | wc -l` != "0" ]; then
        host_os="ubuntu-1604"
    elif [ `echo $os_info | grep Aliyun | wc -l` != "0" ]; then
        host_os="alios"
    elif [ `echo $os_info | grep Alibaba | wc -l` != "0" ]; then
        host_os="alios"
    else
        echo "OS is not ubuntu 1604/1404, Centos7"
        echo "system information: "$os_info
        exit 1
    fi

elif [ "$os_release_exist" = "0" ]; then
    osId=`/acs/nsenter --mount=/proc/1/ns/mnt cat /etc/os-release | grep "ID="`
    osVersion=`/acs/nsenter --mount=/proc/1/ns/mnt cat /etc/os-release | grep "VERSION_ID="`

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
          host_os="ubuntu-1404"
        elif [ `echo $osVersion | grep "16.04" | wc -l` != "0" ]; then
          host_os="ubuntu-1604"
        fi
    else
        echo "OS is not ubuntu-1604/1404/Centos7"
        echo "system information: "$osId
        exit 1
    fi
fi


# cpfs version
cpfsversion="2.10.8-202"
kernelversion74="3.10.0-693.2.2"
kernelversion75="3.10.0-862.14.4"
kernelversion76="3.10.0-957.5.1"
kernelversion763="3.10.0-957.21.3"
kernelversion=${kernelversion763}


if [ "$host_os" = "centos-7" ] || [ "$host_os" = "alios" ] ; then
    # install kernel-devel
    kernelInfo=`/acs/nsenter --mount=/proc/1/ns/mnt uname -a | awk '{print $3}'`
    if [ ${kernelInfo} = ${kernelversion74}".el7.x86_64" ]; then
        kernelversion=${kernelversion74}
    elif [ ${kernelInfo} = ${kernelversion75}".el7.x86_64" ]; then
        kernelversion=${kernelversion75}
    elif [ ${kernelInfo} = ${kernelversion76}".el7.x86_64" ]; then
        kernelversion=${kernelversion76}
    elif [ ${kernelInfo} = ${kernelversion763}".el7.x86_64" ]; then
        kernelversion=${kernelversion763}
    fi

    # check if need to install driver
    installDriver="true"
    cpfs_lustre_num=`/acs/nsenter --mount=/proc/1/ns/mnt rpm -qa | grep kmod-lustre-client | wc -l`
    cpfs_dkms_ver_num=`/acs/nsenter --mount=/proc/1/ns/mnt rpm -qa | grep cpfs-client-dkms | wc -l`
    # kmod install check
    cpfs_kmod_num=`/acs/nsenter --mount=/proc/1/ns/mnt rpm -qa | grep kmod-cpfs-client | wc -l`
    cpfs_kmod_ver_num=`/acs/nsenter --mount=/proc/1/ns/mnt rpm -qa | grep kmod-cpfs-client-$cpfsversion | wc -l`
    # cpfs client install check
    cpfs_num=`/acs/nsenter --mount=/proc/1/ns/mnt rpm -qa | grep cpfs-client- | grep -v kmod-cpfs-client | wc -l`
    cpfs_ver_num=`/acs/nsenter --mount=/proc/1/ns/mnt rpm -qa | grep -v kmod-cpfs-client | grep cpfs-client-$cpfsversion | wc -l`

    # check if install
    if [ "${cpfs_lustre_num}" != "0" ]; then
        installDriver="false"
        echo "kmod-lustre-client installed, Please upgrade cpfs client by hands..."
    fi
    if [ "${cpfs_dkms_ver_num}" != "0" ]; then
        installDriver="false"
        echo "cpfs-client-dkms installed, Please upgrade cpfs client by hands..."
    fi
    if [ "${cpfs_kmod_num}" != "0" ]; then
        installDriver="false"
        if [ "${cpfs_kmod_ver_num}" = "0" ]; then
            echo "kmod-cpfs-client installed, Please upgrade by hands..."
        else
            echo "kmod-cpfs-client-$cpfsversion already installed, no need update..."
        fi
    fi
    if [ "${cpfs_num}" != "0" ]; then
        installDriver="false"
        if [ "${cpfs_ver_num}" = "0" ]; then
            echo "cpfs-client installed, Please upgrade by hands..."
        else
            echo "cpfs-client-$cpfsversion already installed, no need update..."
        fi
    fi

    if [ "${installDriver}" = "true" ] ; then
        # install cpfs client kmod
        cp /acs/kmod-cpfs-client-${cpfsversion}-${kernelversion}.el7.x86_64.rpm /host/etc/cpfs-tools/
        echo "Starting to install kmod-cpfs-client-${cpfsversion}-${kernelversion}.el7.x86_64.rpm ...."
        /acs/nsenter --mount=/proc/1/ns/mnt yum install -y /etc/cpfs-tools/kmod-cpfs-client-${cpfsversion}-${kernelversion}.el7.x86_64.rpm

        # install cpfs client
        cp /acs/cpfs-client-${cpfsversion}-${kernelversion}.el7.x86_64.rpm /host/etc/cpfs-tools/
        echo "Starting to install cpfs-client-${cpfsversion}-${kernelversion}.el7.x86_64.rpm ....."
        /acs/nsenter --mount=/proc/1/ns/mnt yum install -y /etc/cpfs-tools/cpfs-client-${cpfsversion}-${kernelversion}.el7.x86_64.rpm
    fi
fi

# start daemon
/bin/plugin.csi.alibabacloud.com $@
