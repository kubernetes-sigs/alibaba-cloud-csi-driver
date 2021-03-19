#!/usr/bin/env bash
set -e

# usage:
#  ./build.sh ARCH=amd64 VERSION=v1.0

CSIPLUGIN_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd -P)"
CSIPLUGIN_BUILD_DIR=${CSIPLUGIN_ROOT}/build/local-edge
CSIPLUGIN_OUTPUT_DIR=${CSIPLUGIN_BUILD_DIR}/_output
CSIPLUGIN_BUILD_IMAGE="golang:1.13.3-alpine"

GIT_VERSION="v1.14.8"
GIT_VERSION=(${VERSION:-${GIT_VERSION}})
GIT_SHA=`git rev-parse --short HEAD || echo "HEAD"`
GIT_BRANCH=`git rev-parse --abbrev-ref HEAD 2>/dev/null`
BUILD_TIME=`date "+%Y-%m-%d-%H:%M:%S"`

IMG_REPO="registry.cn-hangzhou.aliyuncs.com/acs/csi-plugin"
IMG_VERSION=${GIT_VERSION}-${GIT_SHA}

readonly -a SUPPORTED_ARCH=(
    amd64
    arm
    arm64
)

readonly -a target_arch=(${ARCH:-${SUPPORTED_ARCH[@]}})

function build_multi_arch_binaries() {
    local docker_run_opts=(
        "-i"
        "--rm"
        "--network host"
        "-v ${CSIPLUGIN_ROOT}:/opt/src"
        "--env CGO_ENABLED=0"
        "--env GOOS=linux"
        "--env GIT_VERSION=${GIT_VERSION}"
        "--env GIT_SHA=${GIT_SHA}"
        "--env GIT_BRANCH=${GIT_BRANCH}"
        "--env BUILD_TIME=${BUILD_TIME}"
    )
    # use goproxy if build from inside mainland China
    [[ $region == "cn" ]] && docker_run_opts+=("--env GOPROXY=https://goproxy.cn")

    local docker_run_cmd=(
        "/bin/sh"
        "-xe"
        "-c"
    )

    local sub_commands="sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories; \
        apk --no-cache add bash; \
        cd /opt/src; "
    for arch in ${target_arch[@]}; do
        sub_commands+="CGO_ENABLED=0 GOOS=linux GOARCH='${arch}' go build \
  -ldflags '-X main._BRANCH_=${GIT_BRANCH} -X main._VERSION_=${IMG_VERSION} -X main._BUILDTIME_=${BUILD_TIME}' -mod vendor -o build/local-edge/_output/plugin.csi.alibabacloud.com.${arch} main.go; "
    done

    docker run ${docker_run_opts[@]} ${CSIPLUGIN_BUILD_IMAGE} ${docker_run_cmd[@]} "${sub_commands}"
}

function build_images() {
    for arch in ${target_arch[@]}; do
        local docker_file_path=${CSIPLUGIN_BUILD_DIR}/Dockerfile.$arch
        local docker_build_path=${CSIPLUGIN_BUILD_DIR}
        local csiplugin_image=$IMG_REPO:${IMG_VERSION}.${arch}
        local base_image
        case $arch in
            amd64)
                base_image="amd64/alpine:3.12"
                ;;
            arm64)
                base_image="arm64v8/alpine:3.12"
                ;;
            arm)
                base_image="arm32v7/alpine:3.12"
                ;;
            *)
                echo unknown arch $arch
                exit 1
        esac
        cat << EOF > $docker_file_path
FROM ${base_image}
LABEL maintainers="Alibaba Cloud Authors"
LABEL description="Alibaba Cloud CSI Plugin"

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update && apk add --no-cache ca-certificates file tzdata lvm2 e2fsprogs blkid util-linux
RUN cp /usr/bin/nsenter /nsenter
COPY _output/plugin.csi.alibabacloud.com.${arch} /bin/plugin.csi.alibabacloud.com
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /bin/plugin.csi.alibabacloud.com && chmod +x /entrypoint.sh \
 && sed -i 's/use_lvmetad\ =\ 1/use_lvmetad\ =\ 0/g' /etc/lvm/lvm.conf

ENTRYPOINT ["/entrypoint.sh"]
EOF
        docker build --no-cache -t $csiplugin_image -f $docker_file_path $docker_build_path
        docker save $csiplugin_image > ${CSIPLUGIN_OUTPUT_DIR}/csi-plugin-${arch}.tar
    done
}

rm -rf ${CSIPLUGIN_OUTPUT_DIR}
mkdir -p ${CSIPLUGIN_OUTPUT_DIR}
build_multi_arch_binaries
build_images
