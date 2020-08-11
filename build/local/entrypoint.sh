#!/bin/sh


lvmLine=`/nsenter --mount=/proc/1/ns/mnt rpm -qa | grep lvm2 | wc -l`
if [ "$lvmLine" = "0" ]; then
    /nsenter --mount=/proc/1/ns/mnt systemctl restart lvm2-lvmetad.service
    echo "install lvm and starting..."
fi

#if [ "$lvmLine" = "0" ]; then
#    /nsenter --mount=/proc/1/ns/mnt yum install lvm2 -y
#    /nsenter --mount=/proc/1/ns/mnt sed -i 's/udev_sync\ =\ 0/udev_sync\ =\ 1/g' /etc/lvm/lvm.conf
#    /nsenter --mount=/proc/1/ns/mnt sed -i 's/udev_rules\ =\ 0/udev_rules\ =\ 1/g' /etc/lvm/lvm.conf
#    /nsenter --mount=/proc/1/ns/mnt systemctl restart lvm2-lvmetad.service
#    echo "install lvm and starting..."
#else
#    udevLine=`/nsenter --mount=/proc/1/ns/mnt cat /etc/lvm/lvm.conf | grep "udev_sync = 0" | wc -l`
#    if [ "$udevLine" != "0" ]; then
#        /nsenter --mount=/proc/1/ns/mnt sed -i 's/udev_sync\ =\ 0/udev_sync\ =\ 1/g' /etc/lvm/lvm.conf
#        /nsenter --mount=/proc/1/ns/mnt sed -i 's/udev_rules\ =\ 0/udev_rules\ =\ 1/g' /etc/lvm/lvm.conf
#        /nsenter --mount=/proc/1/ns/mnt systemctl restart lvm2-lvmetad.service
#        echo "update lvm.conf file: udev_sync from 0 to 1, udev_rules from 0 to 1"
#    fi
#fi

/bin/plugin.csi.alibabacloud.com $@
