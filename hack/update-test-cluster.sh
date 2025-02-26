#!/bin/bash

set -e

read -r k8s containerd alinux3 alinux3arm containeros < \
    <(aliyun cs GET /api/v1/metadata/versions --Region 'cn-beijing' --ClusterType ManagedKubernetes | \
        jq -r '.[0] | [
            (.version | tojson),
            (.runtimes.[] | select(.name=="containerd") | .version | tojson),
            (.images.[] | select(.image_type=="AliyunLinux3") | .image_id),
            (.images.[] | select(.image_type=="AliyunLinux3Arm64") | .image_id),
            (.images.[] | select(.image_type=="ContainerOS") | .image_id)
        ] | @tsv')

echo "     cluster: $k8s"
echo "  containerd: $containerd"
echo "    ALinux 3: $alinux3"
echo "ALinux 3 ARM: $alinux3arm"
echo " ContainerOS: $containeros"

sed -i '' "s/kubernetes_version:.*/kubernetes_version: $k8s,/;
           s/runtime_version:.*/runtime_version: $containerd,/;" \
    test/cluster/cluster-template.jsonnet

sed -i '' "s/\${OS_IMAGE_ALINUX3:.*}/\${OS_IMAGE_ALINUX3:-$alinux3}/;
           s/\${OS_IMAGE_ALINUX3_ARM64:.*}/\${OS_IMAGE_ALINUX3_ARM64:-$alinux3arm}/;
           s/\${OS_IMAGE_CONTAINEROS3:.*}/\${OS_IMAGE_CONTAINEROS3:-$containeros}/;" \
    test/cluster/cluster_setup.sh
