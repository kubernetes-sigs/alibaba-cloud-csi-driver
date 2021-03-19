#!/bin/sh

if uname -a | grep -i -q ubuntu; then
  lvmLine=`nsenter --mount=/proc/1/ns/mnt dpkg --get-selections lvm2 | grep install -w -i | wc -l`
  if [ "$lvmLine" = "0" ]; then
    nsenter --mount=/proc/1/ns/mnt apt install lvm2 -y
  fi
else
  lvmLine=`nsenter --mount=/proc/1/ns/mnt rpm -qa lvm2 | wc -l`
  if [ "$lvmLine" = "0" ]; then
    nsenter --mount=/proc/1/ns/mnt yum install lvm2 -y
  fi
fi

if [ "$lvmLine" = "0" ]; then
    nsenter --mount=/proc/1/ns/mnt sed -i 's/udev_sync\ =\ 0/udev_sync\ =\ 1/g' /etc/lvm/lvm.conf
    nsenter --mount=/proc/1/ns/mnt sed -i 's/udev_rules\ =\ 0/udev_rules\ =\ 1/g' /etc/lvm/lvm.conf
    nsenter --mount=/proc/1/ns/mnt systemctl restart lvm2-lvmetad.service
    echo "install lvm and starting..."
else
    udevLine=`nsenter --mount=/proc/1/ns/mnt cat /etc/lvm/lvm.conf | grep "udev_sync = 0" | wc -l`
    if [ "$udevLine" != "0" ]; then
        nsenter --mount=/proc/1/ns/mnt sed -i 's/udev_sync\ =\ 0/udev_sync\ =\ 1/g' /etc/lvm/lvm.conf
        nsenter --mount=/proc/1/ns/mnt sed -i 's/udev_rules\ =\ 0/udev_rules\ =\ 1/g' /etc/lvm/lvm.conf
        nsenter --mount=/proc/1/ns/mnt systemctl restart lvm2-lvmetad.service
        echo "update lvm.conf file: udev_sync from 0 to 1, udev_rules from 0 to 1"
    fi
fi

# edge node, set incluster config endpoint ENV
if [ -e /etc/kubernetes/manifests/edge-hub-dp.yaml ]; then
  if ! echo ${APISERVER_ENDPOINT} | grep -E "https://*";  then
    echo "apiserver endpoint ${APISERVER_ENDPOINT} is malformed "
    exit 1
  fi
  export KUBERNETES_SERVICE_HOST=`echo ${APISERVER_ENDPOINT} | sed 's/https:\/\///' | awk -F ":" '{print $1}'`
  export KUBERNETES_SERVICE_PORT=`echo ${APISERVER_ENDPOINT} | sed 's/https:\/\///' | awk -F ":" '{print $2}'`
fi


/bin/plugin.csi.alibabacloud.com $@
