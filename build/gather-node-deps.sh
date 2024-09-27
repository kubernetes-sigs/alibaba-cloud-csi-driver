#!/bin/bash
set -e
mkdir /staging-node
DEPS=(
    /etc/mke2fs.conf
    /usr/bin/{mount,umount,lspci,mkdir,chmod,grep,tail,nsenter}
    /usr/sbin/{mount.*,umount.*,fsck,fsck.*,mkfs,mkfs.*,*2fs,xfs_*,sfdisk,losetup,blockdev}
)
MTIME=0
gather_dep() {
    if [ -e "/base$1" ]; then
        echo "$1 already exist in base"
        return 0
    fi
    local source target t
    # resolve all but last component of symlink
    source=$(realpath "$(dirname "$1")")/$(basename "$1")
    [ -e "/staging-node$source" ] && return 0
    if [ -h "$source" ]; then
        target=$(realpath "$source")
        echo "gathering link $source => $target"
        gather_dep "$target"
    fi
    echo "gathering dep $source"
    t=$(stat -c '%Y' "$source")
    [ "$t" -gt "$MTIME" ] && MTIME=$t
    cp -dp --parents "$source" /staging-node
}
for f in "${DEPS[@]}"; do
    gather_dep "$f"
done

mapfile -t LIBS < <(ldd /staging-node/usr/{bin,sbin}/* 2>/dev/null | grep -Po '(?<= => )[^ ]+' | sort | uniq)
for f in "${LIBS[@]}"; do
    gather_dep "$f"
done
echo "latest mtime is $(date --date "@$MTIME" --iso-8601=seconds)"
find /staging-node -type d -exec touch --date="@$MTIME" {} +
