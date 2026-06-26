#!/usr/bin/env bash
# Copyright 2024 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd "${SCRIPT_DIR}/.." && pwd)"

# Source version extraction script
source "${SCRIPT_DIR}/get-versions.sh"

echo "Updating versions from VERSIONS file..."
echo "OSSFS_RPM_VERSION: ${OSSFS_RPM_VERSION}"
echo "OSSFS2_RPM_VERSION: ${OSSFS2_RPM_VERSION}"
echo "OSSFS_IMAGE_TAG: ${OSSFS_IMAGE_TAG}"
echo "OSSFS2_IMAGE_TAG: ${OSSFS2_IMAGE_TAG}"
echo "ALINAS_UTILS_VERSION: ${ALINAS_UTILS_VERSION}"
echo "EFC_VERSION: ${EFC_VERSION}"
echo "ALINAS_RPM_BASE_URL: ${ALINAS_RPM_BASE_URL}"
echo ""

# Function to escape special characters for sed
escape_sed() {
    echo "$1" | sed 's/[[\.*^$()+?{|]/\\&/g'
}

# Update Go code
GO_FILE="${ROOT_DIR}/pkg/mounter/fuse_pod_manager/oss/utils.go"
echo "Updating ${GO_FILE}..."
OSSFS_IMAGE_TAG_ESC=$(escape_sed "${OSSFS_IMAGE_TAG}")
OSSFS2_IMAGE_TAG_ESC=$(escape_sed "${OSSFS2_IMAGE_TAG}")
sed -i '' "s/defaultOssfsUpdatedImageTag = \".*\"/defaultOssfsUpdatedImageTag = \"${OSSFS_IMAGE_TAG_ESC}\"/" "${GO_FILE}"
sed -i '' "s/defaultOssfs2ImageTag.*= \".*\"/defaultOssfs2ImageTag       = \"${OSSFS2_IMAGE_TAG_ESC}\"/" "${GO_FILE}"

# Update mount-proxy Dockerfile
MOUNT_PROXY_DOCKERFILE="${ROOT_DIR}/build/mount-proxy/Dockerfile"
echo "Updating ${MOUNT_PROXY_DOCKERFILE}..."
OSSFS_RPM_VERSION_ESC=$(escape_sed "${OSSFS_RPM_VERSION}")
OSSFS2_RPM_VERSION_ESC=$(escape_sed "${OSSFS2_RPM_VERSION}")
ALINAS_UTILS_VERSION_ESC=$(escape_sed "${ALINAS_UTILS_VERSION}")
EFC_VERSION_ESC=$(escape_sed "${EFC_VERSION}")
ALINAS_RPM_BASE_URL_ESC=$(escape_sed "${ALINAS_RPM_BASE_URL}")

# Update OSSFS_VERSION for ossfs stage
sed -i '' "/^FROM.*AS ossfs$/,/^FROM.*AS ossfs-1.88$/ s/^ARG OSSFS_VERSION=.*/ARG OSSFS_VERSION=${OSSFS_RPM_VERSION_ESC}/" "${MOUNT_PROXY_DOCKERFILE}"
# Update OSSFS_VERSION for ossfs2 stage
sed -i '' "/^FROM.*AS ossfs2$/,/^FROM.*AS alinas$/ s/^ARG OSSFS_VERSION=.*/ARG OSSFS_VERSION=${OSSFS2_RPM_VERSION_ESC}/" "${MOUNT_PROXY_DOCKERFILE}"
# Update OSSFS_VERSION and OSSFS2_VERSION for aio stage
sed -i '' "/^FROM.*AS aio$/,\$ s/^ARG OSSFS_VERSION=.*/ARG OSSFS_VERSION=${OSSFS_RPM_VERSION_ESC}/" "${MOUNT_PROXY_DOCKERFILE}"
sed -i '' "/^FROM.*AS aio$/,\$ s/^ARG OSSFS2_VERSION=.*/ARG OSSFS2_VERSION=${OSSFS2_RPM_VERSION_ESC}/" "${MOUNT_PROXY_DOCKERFILE}"
# Update alinas/efc versions (all occurrences)
sed -i '' "s|^ARG ALINAS_RPM_BASE_URL=.*|ARG ALINAS_RPM_BASE_URL=${ALINAS_RPM_BASE_URL_ESC}|" "${MOUNT_PROXY_DOCKERFILE}"
sed -i '' "s/^ARG ALINAS_UTILS_VERSION=.*/ARG ALINAS_UTILS_VERSION=${ALINAS_UTILS_VERSION_ESC}/" "${MOUNT_PROXY_DOCKERFILE}"
sed -i '' "s/^ARG EFC_VERSION=.*/ARG EFC_VERSION=${EFC_VERSION_ESC}/" "${MOUNT_PROXY_DOCKERFILE}"

# Update multi Dockerfile
MULTI_DOCKERFILE="${ROOT_DIR}/build/multi/Dockerfile.multi"
echo "Updating ${MULTI_DOCKERFILE}..."
sed -i '' "s/^ARG OSSFS_IMAGE_TAG=.*/ARG OSSFS_IMAGE_TAG=${OSSFS_IMAGE_TAG_ESC}/" "${MULTI_DOCKERFILE}"
sed -i '' "s/^ARG OSSFS2_IMAGE_TAG=.*/ARG OSSFS2_IMAGE_TAG=${OSSFS2_IMAGE_TAG_ESC}/" "${MULTI_DOCKERFILE}"

# Update multi.asi Dockerfile
MULTI_ASI_DOCKERFILE="${ROOT_DIR}/build/multi/Dockerfile.multi.asi"
echo "Updating ${MULTI_ASI_DOCKERFILE}..."
sed -i '' "s/^ARG OSSFS_IMAGE_TAG=.*/ARG OSSFS_IMAGE_TAG=${OSSFS_IMAGE_TAG_ESC}/" "${MULTI_ASI_DOCKERFILE}"
sed -i '' "s/^ARG OSSFS2_IMAGE_TAG=.*/ARG OSSFS2_IMAGE_TAG=${OSSFS2_IMAGE_TAG_ESC}/" "${MULTI_ASI_DOCKERFILE}"

# Update Helm values.yaml
HELM_VALUES="${ROOT_DIR}/deploy/charts/alibaba-cloud-csi-driver/values.yaml"
echo "Updating ${HELM_VALUES}..."
# Update ossfs tag - match the line with "tag:" under "ossfs:"
sed -i '' "/^  ossfs:$/,/^  [a-z]/ s|^    tag: \".*\"|    tag: \"${OSSFS_IMAGE_TAG_ESC}\"|" "${HELM_VALUES}"
# Update ossfs2 tag - match the line with "tag:" under "ossfs2:"
sed -i '' "/^  ossfs2:$/,/^  [a-z]/ s|^    tag: \".*\"|    tag: \"${OSSFS2_IMAGE_TAG_ESC}\"|" "${HELM_VALUES}"

echo ""
echo "Versions updated successfully! ✓"
echo "Running gofmt to fix formatting..."

# Run gofmt on the Go file
if command -v gofmt >/dev/null 2>&1; then
    gofmt -w "${GO_FILE}"
    echo "Formatted ${GO_FILE}"
else
    echo "Warning: gofmt not found, skipping Go file formatting"
fi

echo ""
echo "All versions have been updated from VERSIONS file."
