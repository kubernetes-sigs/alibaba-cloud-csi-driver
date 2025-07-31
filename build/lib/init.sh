#!/bin/sh
# shellcheck shell=busybox

# skip all the setup if running in provisioner mode
if [ "$SERVICE_TYPE" = "provisioner" ]; then
    exit 0
fi

if [ -z "$KUBELET_ROOT_DIR" ]; then
    KUBELET_ROOT_DIR="/var/lib/kubelet"
fi

run_nas="false"

mkdir -p /var/log/alicloud/
mkdir -p /host/etc/kubernetes/volumes/disk/uuid

host_cmd() { nsenter --mount=/proc/1/ns/mnt "$@"; }
HOST_ROOT="/proc/1/root"

OS_RELEASE=/etc/os-release
if ! [ -f $HOST_ROOT$OS_RELEASE ]; then
    OS_RELEASE=/usr/lib/os-release
fi

OS_RELEASE_ID=$(. $HOST_ROOT$OS_RELEASE; echo "$ID")
OS_RELEASE_VERSION_ID=$(. $HOST_ROOT$OS_RELEASE; echo "$VERSION_ID")
ARCH=$(uname -m)
echo "detected host OS: $OS_RELEASE_ID, version: $OS_RELEASE_VERSION_ID, arch: $ARCH"

cleanup_old_staging_path() (
    echo cleaning up old volume staging path.  # kubelet will mount the new path at startup.
    cd "$OLD_STAGING_PATH" || return 1
    for i in *; do
        echo "$i"
        umount "$i/globalmount"
        rmdir "$i/globalmount"
        rm -f "$i/vol_data.json"
        rmdir "$i"
    done
)

OLD_STAGING_PATH=${KUBELET_ROOT_DIR}/plugins/kubernetes.io/csi/pv
if [ -d "$OLD_STAGING_PATH" ]; then
    cleanup_old_staging_path
    rmdir "$OLD_STAGING_PATH" && echo "removed $OLD_STAGING_PATH, finished cleanup."
fi

## check which plugin is running
for item in "$@";
do
    if [ "$item" = "--driver=ossplugin.csi.alibabacloud.com" ]; then
        echo "Running oss plugin...."
        mkdir -p "$KUBELET_ROOT_DIR/csi-plugins/ossplugin.csi.alibabacloud.com"
        rm -rf "$KUBELET_ROOT_DIR/plugins/ossplugin.csi.alibabacloud.com/csi.sock"
    elif [ "$item" = "--driver=diskplugin.csi.alibabacloud.com" ]; then
        echo "Running disk plugin...."
        mkdir -p "$KUBELET_ROOT_DIR/csi-plugins/diskplugin.csi.alibabacloud.com"
        rm -rf "$KUBELET_ROOT_DIR/plugins/diskplugin.csi.alibabacloud.com/csi.sock"
    elif [ "$item" = "--driver=nasplugin.csi.alibabacloud.com" ]; then
        echo "Running nas plugin...."
        run_nas="true"
        mkdir -p "$KUBELET_ROOT_DIR/csi-plugins/nasplugin.csi.alibabacloud.com"
        rm -rf "$KUBELET_ROOT_DIR/plugins/nasplugin.csi.alibabacloud.com/csi.sock"
    elif [[ "$item" = *--driver=* ]]; then
        IFS=,
        for driver_type in ${item#*--driver=};
        do
            if [ "$driver_type" = "oss" ]; then
                echo "Running oss plugin...."
                mkdir -p "$KUBELET_ROOT_DIR/csi-plugins/ossplugin.csi.alibabacloud.com"
                rm -rf "$KUBELET_ROOT_DIR/plugins/ossplugin.csi.alibabacloud.com/csi.sock"
            elif [ "$driver_type" = "disk" ]; then
                echo "Running disk plugin...."
                mkdir -p "$KUBELET_ROOT_DIR/csi-plugins/diskplugin.csi.alibabacloud.com"
                rm -rf "$KUBELET_ROOT_DIR/plugins/diskplugin.csi.alibabacloud.com/csi.sock"
            elif [ "$driver_type" = "nas" ]; then
                echo "Running nas plugin...."
                run_nas="true"
                mkdir -p "$KUBELET_ROOT_DIR/csi-plugins/nasplugin.csi.alibabacloud.com"
                rm -rf "$KUBELET_ROOT_DIR/plugins/nasplugin.csi.alibabacloud.com/csi.sock"
            elif [ "$driver_type" = "bmcpfs" ]; then
                echo "Running bmcpfs plugin...."
                host_cmd modprobe fuse
                echo "modprobe fuse returned: $?"
            elif [ "$driver_type" = "pov" ]; then
                echo "Running pov plugin...."
                run_pov="true"
            fi
        done
        unset IFS
    fi
done

# config /etc/updatedb.config if needed
if [ "$SKIP_UPDATEDB_CONFIG" != "true" ]; then
    ## check cron.daily dir
    if [ -f /host/etc/cron.daily/mlocate ]; then
        if [ -f /host/etc/updatedb.conf ]; then
            sed -i 's/PRUNE_BIND_MOUNTS.*$/PRUNE_BIND_MOUNTS = \"yes\"/g' /host/etc/updatedb.conf
            if ! grep "^PRUNEFS" /host/etc/updatedb.conf | grep -q -F "fuse.ossfs"; then
                PRUNEFS="fuse.ossfs"
                echo "add PRUNEFS: $PRUNEFS to /etc/updatedb.conf"
                sed -i "s|PRUNEFS\s*=\s*\"|&${PRUNEFS//|/\\|} |" /host/etc/updatedb.conf
            fi
            if ! grep "^PRUNEPATHS" /host/etc/updatedb.conf | grep -q -F "$KUBELET_ROOT_DIR"; then
                PRUNEPATHS="$KUBELET_ROOT_DIR"
                if [ "$KUBELET_ROOT_DIR" = "/var/lib/kubelet" ]; then
                    PRUNEPATHS="$PRUNEPATHS /var/lib/container/kubelet"
                fi
                echo "add PRUNEPATHS: $PRUNEPATHS to /etc/updatedb.conf"
                sed -i "s|PRUNEPATHS\s*=\s*\"|&${PRUNEPATHS//|/\\|} |" /host/etc/updatedb.conf
            fi
        fi
    fi
fi

# skip installing csiplugin-connector when DISABLE_CSIPLUGIN_CONNECTOR=true
if [ "$DISABLE_CSIPLUGIN_CONNECTOR" != "true" ] && [ "$run_nas" = "true" ]; then
    ## install/update csi connector
    updateConnector="true"
	systemdDir="/host/usr/lib/systemd/system"
    if [ "$OS_RELEASE_ID" = "lifsea" ]; then
        systemdDir="/host/etc/systemd/system"
    fi
    if [ ! -f "/host/etc/csi-tool/csiplugin-connector" ]; then
        mkdir -p /host/etc/csi-tool/
        echo "mkdir /etc/csi-tool/ directory..."
    else
        oldmd5=$(md5sum /host/etc/csi-tool/csiplugin-connector | awk '{print $1}')
        newmd5=$(md5sum /csi/csiplugin-connector | awk '{print $1}')
        if [ "$oldmd5" = "$newmd5" ]; then
            updateConnector="false"
        else
            rm -rf /host/etc/csi-tool/
            rm -rf /host/etc/csi-tool/connector.sock
            rm -rf /host/etc/csi-tool/diskconnector.sock
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
    if [ -n "${PLUGINS_SOCKETS}" ];then
        sed -i 's/Restart=always/Restart=on-failure/g' /csi/csiplugin-connector.service
        sed -i '/^\[Service\]/a Environment=\"WATCHDOG_SOCKETS_PATH='"${PLUGINS_SOCKETS}"'\"' /csi/csiplugin-connector.service
    fi
    if [ -f "$systemdDir/csiplugin-connector.service" ];then
        echo "Check csiplugin-connector.service...."
        oldmd5=$(md5sum $systemdDir/csiplugin-connector.service | awk '{print $1}')
        newmd5=$(md5sum /csi/csiplugin-connector.service | awk '{print $1}')
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
        for i in $(seq 10);
        do
            if host_cmd systemctl daemon-reload; then
                break
            else
                echo "Starting retry again systemctl daemon-reload.retry count:$i"
                sleep 2
            fi
        done
    fi

    rm -rf /var/log/alicloud/connector.pid
    echo "Starting systemctl enable csiplugin-connector.service."
    for i in $(seq 5);
    do
        if host_cmd systemctl enable csiplugin-connector.service; then
            break
        else
            echo "Starting retry again systemctl enable csiplugin-connector.service.retry count:$i"
            sleep 2
        fi
    done

    echo "Starting systemctl restart csiplugin-connector.service."
    for i in $(seq 5);
    do
        if host_cmd systemctl restart csiplugin-connector.service; then
            break
        else
            echo "Starting retry again systemctl restart csiplugin-connector.service.retry count:$i"
            sleep 2
        fi
    done
fi

echo "Start checking if the host system packages needs to be installed"

dnf_install() {
    host_cmd dnf install -y "$@"
}

yum_install() {
    # yum install is not idempotent, so we need to check if the package is already installed.
    # yum also fails if the already installed version is newer, so we need to downgrade.
    # fail if both install and downgrade fail
    # shellcheck disable=SC2016 # the expansion should happen on host, not in container.
    host_cmd sh -c 'if rpm --query $(rpm --query --package "$1"); then
        echo "$1 already installed";
    else
        yum install -y "$1" || yum downgrade -y "$1";
    fi' install_pkg "$1"
}

apt_install() {
    host_cmd apt-get install -y --allow-downgrades "$@"
}

if host_cmd dnf --version; then
    echo "dnf is available"
    PKG_MGR=dnf
    SUPPORT_RPM=true
elif host_cmd yum --version; then
    echo "yum is available"
    PKG_MGR=yum
    SUPPORT_RPM=true
elif host_cmd apt-get --version; then
    echo "apt-get is available"
    PKG_MGR=apt
fi

# idempotent package installation
# downgrade if the already installed version is newer than the one we want to install
install_package() {
    cp "/root/$1" /host/etc/csi-tool/
    if [ -z "$PKG_MGR" ]; then
        echo "no package manager found"
        return 1
    fi
    "${PKG_MGR}_install" "/etc/csi-tool/$1" || {
        echo "failed to install $1"
        return 1
    }
    rm -f "/host/etc/csi-tool/$1"
}

## CPFS-NAS plugin setup
if [ "$ARCH" = "x86_64" ] && [ "$run_nas" = "true" ]; then
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
    if test -e $client_properties && grep -qE "alinas-utils=true" $client_properties; then
        install_utils="true"
    fi

    if [ $install_utils = "true" ]; then
        # cpfs-nas nas-rich-client common rpm
        echo "installing aliyun-alinas-utils"
        if [ "$SUPPORT_RPM" = "true" ]; then
            PKG=aliyun-alinas-utils-1.1-8.20240527201444.2012cc.al7.noarch.rpm
        elif [ "$PKG_MGR" = "apt" ]; then
            PKG=aliyun-alinas-utils-1.1-8.deb
        else
            echo "WARN: no supported package manager found, skip install aliyun-alinas-utils"
            install_utils="false"
        fi
        if [ $install_utils = "true" ]; then
            install_package "$PKG" || exit $?
        fi
    fi

    if { [ "$OS_RELEASE_ID" != "alinux" ] || [[ "$OS_RELEASE_VERSION_ID" != 2.* ]]; } && [ $install_efc = "true" ]; then
        echo "WARN: skip install efc because host os is not alinux2"
        install_efc="false"
    fi
    if [ $install_efc = "true" ]; then
        # nas-rich-client rpm
        echo "installing alinas-efc"
        cp /root/alinas-efc-1.2-3.x86_64.rpm /host/etc/csi-tool/
        host_cmd yum install -y /etc/csi-tool/alinas-efc-1.2-3.x86_64.rpm

        echo "checking alinas-efc-1.2-3.x86_64 installed"
        host_cmd rpm -q alinas-efc-1.2-3.x86_64 || exit 1
        echo "starting aliyun-alinas-mount-watchdog"
        host_cmd systemctl start aliyun-alinas-mount-watchdog || exit 1
    fi
fi

# place it here to remove leftover from previous version
rm -rf /host/etc/csi-tool/*.rpm /host/etc/csi-tool/*.deb
