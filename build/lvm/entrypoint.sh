#!/bin/sh


lvmLine=`/nsenter --mount=/proc/1/ns/mnt rpm -qa | grep lvm2 | wc -l`

if [ "$lvmLine" = "0" ]; then
    /nsenter --mount=/proc/1/ns/mnt yum install lvm2 -y
    /nsenter --mount=/proc/1/ns/mnt sed -i 's/udev_sync\ =\ 1/udev_sync\ =\ 0/g' /etc/lvm/lvm.conf
    /nsenter --mount=/proc/1/ns/mnt sed -i 's/udev_rules\ =\ 1/udev_rules\ =\ 0/g' /etc/lvm/lvm.conf
    /nsenter --mount=/proc/1/ns/mnt systemctl restart lvm2-lvmetad.service
else
    udevLine=`/nsenter --mount=/proc/1/ns/mnt cat /etc/lvm/lvm.conf | grep "udev_sync = 1" | wc -l`
    if [ "$udevLine" != "0" ]; then
        /nsenter --mount=/proc/1/ns/mnt sed -i 's/udev_sync\ =\ 1/udev_sync\ =\ 0/g' /etc/lvm/lvm.conf
        /nsenter --mount=/proc/1/ns/mnt sed -i 's/udev_rules\ =\ 1/udev_rules\ =\ 0/g' /etc/lvm/lvm.conf
        /nsenter --mount=/proc/1/ns/mnt systemctl restart lvm2-lvmetad.service
    fi
fi

/bin/lvmplugin.csi.alibabacloud.com $@
