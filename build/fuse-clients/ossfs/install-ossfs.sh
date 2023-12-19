#!/bin/sh
echo "TARGETPLATFORM: $TARGETPLATFORM"

echo "installing ossfs"
case $TARGETPLATFORM in
    linux/amd64)
        yum install -y https://ack-csiplugin.oss-cn-hangzhou.aliyuncs.com/pre/ossfs/ossfs_1.88.4_centos8.0_x86_64.rpm
        ;;
    linux/arm64)
        yum install -y \
            https://ack-csiplugin.oss-cn-hangzhou.aliyuncs.com/multi-private/arm64-ossfs-v1.80.6/fuse-libs-2.9.2-11.el7.aarch64.rpm \
            https://ack-csiplugin.oss-cn-hangzhou.aliyuncs.com/multi-private/arm64-ossfs-v1.80.6/fuse-2.9.2-11.el7.aarch64.rpm
        curl -o /usr/lib64/libfuse.so.2 https://ack-csiplugin.oss-cn-hangzhou.aliyuncs.com/multi-private/arm64-ossfs-v1.80.6/libfuse.so.2
        curl -o /usr/local/bin/ossfs https://ack-csiplugin.oss-cn-hangzhou.aliyuncs.com/multi-private/arm64-ossfs-v1.80.6/ossfs-8u
        chmod +x /usr/local/bin/ossfs
        ;;
    *)
        echo "unknown arch"
        exit 1
        ;;
esac

echo "installing util-linux" # mount, findmnt ...
yum install -y util-linux mailcap && yum clean all 
mv /etc/mime.types /etc/csi-mime.types