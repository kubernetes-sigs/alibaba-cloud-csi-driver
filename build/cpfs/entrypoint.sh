#!/bin/sh

echo "Starting deploy cpfs csi-plugin...."
mkdir -p /host/etc/cpfs-tools


host_os="centos-7"
/nsenter --mount=/proc/1/ns/mnt which lsb_release
lsb_release_exist=$?
if [ "$lsb_release_exist" != "0" ]; then
  /nsenter --mount=/proc/1/ns/mnt ls /etc/os-release
  os_release_exist=$?
fi
if [ "$lsb_release_exist" = "0" ]; then
    os_info=`/nsenter --mount=/proc/1/ns/mnt lsb_release -a`
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
    osId=`/nsenter --mount=/proc/1/ns/mnt cat /etc/os-release | grep "ID="`
    osVersion=`/nsenter --mount=/proc/1/ns/mnt cat /etc/os-release | grep "VERSION_ID="`

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
dkmsversion="1.0.0-204"
cpfsversion="2.10.8-202"

if [ "$host_os" = "centos-7" ] || [ "$host_os" = "alios" ] ; then
    if [ "${CPFS_VERSION}" != "" ]; then
        echo "Get env set CPFS_VERSION: "${CPFS_VERSION}
        cpfsversion=${CPFS_VERSION}
    fi
    if [ "${DKMS_VERSION}" != "" ]; then
        echo "Get env set DKMS_VERSION: "${DKMS_VERSION}
        dkmsversion=${DKMS_VERSION}
    fi

    # Aliyun Linux 2
    platform_tag="el7"
    if [ "$host_os" = "alios" ]; then
        platform_tag="al7"
    fi

    if [ "$cpfsversion" = "" ] || [ "$dkmsversion" = "" ]; then
        echo "cpfsversion, kernelversion empty, please set CPFS_VERSION/DKMS_VERSION for cpfs: ${cpfsversion}, ${dkmsversion}"
        return
    fi

    echo "Use CPFS_VERSION: "${cpfsversion}
    echo "Use DKMS_VERSION: "${dkmsversion}

        # check if need to install driver
    installDriver="true"
    cpfs_lustre_num=`/nsenter --mount=/proc/1/ns/mnt rpm -qa | grep kmod-lustre-client | wc -l`
    cpfs_dkms_ver_num=`/nsenter --mount=/proc/1/ns/mnt rpm -qa | grep cpfs-client-dkms | wc -l`
    # kmod install check
    cpfs_kmod_num=`/nsenter --mount=/proc/1/ns/mnt rpm -qa | grep kmod-cpfs-client | wc -l`
    cpfs_kmod_ver_num=`/nsenter --mount=/proc/1/ns/mnt rpm -qa | grep kmod-cpfs-client-$cpfsversion | wc -l`
    # cpfs client install check
    cpfs_num=`/nsenter --mount=/proc/1/ns/mnt rpm -qa | grep cpfs-client- | grep -v kmod-cpfs-client | wc -l`
    cpfs_ver_num=`/nsenter --mount=/proc/1/ns/mnt rpm -qa | grep -v kmod-cpfs-client | grep cpfs-client-$cpfsversion | wc -l`

    # check if install
    if [ "${cpfs_lustre_num}" != "0" ]; then
        installDriver="false"
        echo "kmod-lustre-client installed, Please upgrade by hands..."
    fi
    if [ "${cpfs_dkms_ver_num}" != "0" ]; then
        installDriver="false"
        echo "cpfs-client-dkms installed, Please upgrade by hands if needed..."
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
        # install Dependency
        /nsenter --mount=/proc/1/ns/mnt yum install -y make gcc libyaml-devel libtool zlib-devel glibc-headers dkms expect
        /nsenter --mount=/proc/1/ns/mnt yum install -y kernel-devel-`uname -r`

        # install cpfs client dkms
        if [ ! -f /acs/cpfs-client-dkms-${dkmsversion}.${platform_tag}.noarch.rpm ]; then
            if [ "${CPFS_URL_PREFIX}" != "" ]; then
                echo "Download cpfs-client-dkms-${dkmsversion}.${platform_tag}.noarch.rpm from ${CPFS_URL_PREFIX} oss..."
                wget ${CPFS_URL_PREFIX}/cpfs-client-dkms-${dkmsversion}.${platform_tag}.noarch.rpm -P /acs/
            else
                echo "Download cpfs-client-dkms-${dkmsversion}.${platform_tag}.noarch.rpm from aliyun oss..."
                wget http://cpfs-ack.oss-cn-beijing.aliyuncs.com/cpfs-client-dkms-${dkmsversion}.${platform_tag}.noarch.rpm -P /acs/
            fi
        fi
        cp /acs/cpfs-client-dkms-${dkmsversion}.${platform_tag}.noarch.rpm /host/etc/cpfs-tools
        echo "Starting to install cpfs-client-dkms-${dkmsversion}.${platform_tag}.noarch.rpm ...."
        /nsenter --mount=/proc/1/ns/mnt yum install -y /etc/cpfs-tools/cpfs-client-dkms-${dkmsversion}.${platform_tag}.noarch.rpm

        # install cpfs client
        if [ ! -f /acs/cpfs-client-${cpfsversion}.${platform_tag}.x86_64.rpm ]; then
            if [ "${CPFS_URL_PREFIX}" != "" ]; then
                echo "Download cpfs-client-${cpfsversion}.${platform_tag}.x86_64.rpm ${CPFS_URL_PREFIX} oss..."
                wget ${CPFS_URL_PREFIX}/cpfs-client-${cpfsversion}.${platform_tag}.x86_64.rpm -P /acs/
            else
                echo "Download cpfs-client-${cpfsversion}.${platform_tag}.x86_64.rpm from aliyun oss..."
                wget http://cpfs-ack.oss-cn-beijing.aliyuncs.com/cpfs-client-${cpfsversion}.${platform_tag}.x86_64.rpm -P /acs/
            fi
        fi
        cp /acs/cpfs-client-${cpfsversion}.${platform_tag}.x86_64.rpm /host/etc/cpfs-tools/
        echo "Starting to install cpfs-client-${cpfsversion}.${platform_tag}.x86_64.rpm ....."
        /nsenter --mount=/proc/1/ns/mnt yum install -y /etc/cpfs-tools/cpfs-client-${cpfsversion}.${platform_tag}.x86_64.rpm
    fi
fi

# start daemon
/bin/plugin.csi.alibabacloud.com $@
