#!/bin/bash
set -e

# This directory is distroless specific, and recognized by syft
mkdir -p /staging-node/var/lib/dpkg/status.d

DEPS=(
    /etc/netconfig
    /etc/mke2fs.conf /sbin/{fsck,mkfs,mount,umount}.{ext{2,3,4},xfs,nfs}
    /usr/bin/{mount,umount,lspci,lsof,chmod,grep,tail,nsenter}
    /usr/sbin/{fsck,mkfs,sfdisk,losetup,blockdev}
    /sbin/dumpe2fs /sbin/resize2fs
    /usr/sbin/xfs_io /usr/sbin/xfs_growfs
)

declare -A FILE_PACKAGES
for line in $(dpkg-query --show --showformat 'PKG_NAME:${Package}\n${db-fsys:Files}'); do
    if [[ "$line" = PKG_NAME:* ]]; then
        pkg=${line#PKG_NAME:}
    else
        FILE_PACKAGES[$line]=$pkg
    fi
done
echo "indexed ${#FILE_PACKAGES[@]} files"

MTIME=0
gather_dep() {
    local source target t pkg copyright
    # resolve all but last component of symlink
    source=$(realpath "$(dirname "$1")")/$(basename "$1")

    # find the package that contains the source
    pkg=${FILE_PACKAGES[$source]}
    if [ -z "$pkg" ] && [[ "$source" = /usr/* ]]; then
        # retry without /usr prefix
        # use source path matching dpkg for SBOM to work, because /lib is not linked to /usr/lib in distroless
        source="${source#/usr}"
        pkg=${FILE_PACKAGES[$source]}
    fi
    if [ -z "$pkg" ]; then
        echo "failed to find package for $source"
        return 1
    fi
    if [ -e "/base$source" ]; then
        echo "$source already exist in base"
        return 0
    fi

    [ -e "/staging-node$source" ] && return 0
    if [ -h "$source" ]; then
        target=$(realpath "$source")
        echo "gathering link $source => $target"
        gather_dep "$target"
    fi
    echo "gathering dep $pkg: $source"
    t=$(stat -c '%Y' "$source")
    [ "$t" -gt "$MTIME" ] && MTIME=$t
    cp -dp --parents "$source" /staging-node

    # deb package metadata is useful for SBOM
    [ -e "/staging-node/var/lib/dpkg/status.d/$pkg" ] && return 0
    echo "installing deb package $pkg metadata"
    dpkg-query --status "$pkg" > "/staging-node/var/lib/dpkg/status.d/$pkg"
    dpkg-query --control-show "$pkg" md5sums > "/staging-node/var/lib/dpkg/status.d/$pkg.md5sums"
    copyright="/usr/share/doc/$pkg/copyright"
    if [ -e "$copyright" ]; then
        echo "installing deb package $pkg copyright"
        cp -dp --parents "$copyright" /staging-node
    fi
}
for f in "${DEPS[@]}"; do
    if ! [ -e "$f" ]; then
        echo "$f does not exist"
        continue
    fi
    gather_dep "$f"
done

mapfile -t LIBS < <(ldd /staging-node/{usr/,}{bin,sbin}/* 2>/dev/null | grep -Po '(?<= => )[^ ]+' | sort -u)
for f in "${LIBS[@]}"; do
    gather_dep "$f"
done
echo "latest mtime is $(date --date "@$MTIME" --iso-8601=seconds)"
find /staging-node -type d -exec touch --date="@$MTIME" {} +
touch --date="@$MTIME" /staging-node/var/lib/dpkg/status.d/*
