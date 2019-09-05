#!/bin/sh

echo "Starting deploy cpfs csi-plugin...."

mkdir /host/etc/cpfs-tools

# cpfs version
cpfsversion="2.10.8-202"
dkmsversion="1.0.0-202"
kernelversion74="3.10.0-693.2.2"
kernelversion75="3.10.0-862.14.4"
kernelversion76="3.10.0-957.5.1"
kernelversion=${kernelversion74}


/acs/nsenter --mount=/proc/1/ns/mnt yum install -y make gcc libyaml-devel libtool zlib-devel glibc-headers dkms expect

# Outside:: install kernel-devel, support 3.10 kernel
kernelInfo=`/acs/nsenter --mount=/proc/1/ns/mnt uname -a | awk '{print $3}'`
if [ ${kernelInfo} = ${kernelversion74}".el7.x86_64" ]; then
    kernelversion=${kernelversion74}
elif [ ${kernelInfo} = ${kernelversion75}".el7.x86_64" ]; then
    kernelversion=${kernelversion75}
elif [ ${kernelInfo} = ${kernelversion76}".el7.x86_64" ]; then
    kernelversion=${kernelversion76}
fi
if [ `/acs/nsenter --mount=/proc/1/ns/mnt rpm -qa | grep kernel-devel-${kernelversion} | wc -l` = "0" ]; then
    cp /acs/kernel-devel-${kernelversion}.el7.x86_64.rpm /host/etc/cpfs-tools/
    echo "Starting to install kernel-devel-${kernelversion}.el7.x86_64.rpm ...."
    /acs/nsenter --mount=/proc/1/ns/mnt yum install -y /etc/cpfs-tools/kernel-devel-${kernelversion}.el7.x86_64.rpm
fi

# Outside:: install/upgrade cpfs client dkms
cpfs_dkms_num=`/acs/nsenter --mount=/proc/1/ns/mnt rpm -qa | grep cpfs-client-dkms | wc -l`
cpfsdkmsver_num=`/acs/nsenter --mount=/proc/1/ns/mnt rpm -qa | grep cpfs-client-dkms-$dkmsversion | wc -l`
if [ "${cpfsdkmsver_num}" = "0" ]; then
    # remove old version
    if [ "${cpfs_dkms_num}" != "0" ]; then
        echo "Starting to remove cpfs-client-dkms ...."
        /acs/nsenter --mount=/proc/1/ns/mnt yum remove -y cpfs-client-dkms
    fi
    # install cpfs client
    cp /acs/cpfs-client-dkms-${dkmsversion}.el7.noarch.rpm /host/etc/cpfs-tools/
    echo "Starting to install cpfs-client-dkms-${dkmsversion}.el7.noarch.rpm ...."
    /acs/nsenter --mount=/proc/1/ns/mnt yum install -y /etc/cpfs-tools/cpfs-client-dkms-${dkmsversion}.el7.noarch.rpm
fi


# Inside:: install kernel-devel
if [ `rpm -qa | grep kernel-devel-${kernelversion} | wc -l` = "0" ]; then
    echo "Starting to install kernel-devel-${kernelversion}.el7.x86_64.rpm ...."
    yum install -y /acs/kernel-devel-${kernelversion}.el7.x86_64.rpm
fi

# Inside:: install/upgrade cpfs client dkms
cpfs_dkms_num=`rpm -qa | grep cpfs-client-dkms | wc -l`
cpfsdkmsver_num=`rpm -qa | grep cpfs-client-dkms-$dkmsversion | wc -l`
if [ "${cpfsdkmsver_num}" = "0" ]; then
    # install cpfs client
    echo "Starting to install cpfs-client-dkms-${dkmsversion}.el7.noarch.rpm ...."
    yum install -y /acs/cpfs-client-dkms-${dkmsversion}.el7.noarch.rpm
fi

# Inside:: install/upgrade cpfs client
cpfs_num=`rpm -qa | grep cpfs-client- | grep -v cpfs-client-dkms | wc -l`
cpfsver_num=`rpm -qa | grep cpfs-client-$cpfsversion | wc -l`
if [ "${cpfsver_num}" = "0" ]; then
    # install cpfs client
    echo "Starting to install cpfs-client-${cpfsversion}.el7.x86_64.rpm ....."
    yum install -y /acs/cpfs-client-${cpfsversion}.el7.x86_64.rpm
fi

# start daemon
/bin/plugin.csi.alibabacloud.com $@
