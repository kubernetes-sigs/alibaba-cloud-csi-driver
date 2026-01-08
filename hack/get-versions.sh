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

# This script extracts version information from the VERSIONS file.
# It can be sourced by other scripts to get version variables.
#
# Usage:
#   source hack/get-versions.sh
#   echo "OSSFS version: ${OSSFS_RPM_VERSION}"

set -euo pipefail

# Get script directory, compatible with both bash and zsh
if [[ -n "${BASH_SOURCE[0]:-}" ]]; then
    SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
elif [[ -n "${(%):-%x}" ]]; then
    # zsh
    SCRIPT_DIR="$(cd "$(dirname "${(%):-%x}")" && pwd)"
else
    SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
fi
ROOT_DIR="$(cd "${SCRIPT_DIR}/.." && pwd)"

VERSIONS_FILE="${ROOT_DIR}/VERSIONS"
if [[ ! -f "${VERSIONS_FILE}" ]]; then
    echo "Error: VERSIONS file not found at ${VERSIONS_FILE}" >&2
    exit 1
fi

# Extract image tags from VERSIONS file
OSSFS_IMAGE_TAG=$(grep "^OSSFS_IMAGE_TAG=" "${VERSIONS_FILE}" | cut -d'=' -f2)
OSSFS2_IMAGE_TAG=$(grep "^OSSFS2_IMAGE_TAG=" "${VERSIONS_FILE}" | cut -d'=' -f2)

if [[ -z "${OSSFS_IMAGE_TAG}" || -z "${OSSFS2_IMAGE_TAG}" ]]; then
    echo "Error: Failed to extract image tags from VERSIONS file" >&2
    exit 1
fi

# Extract RPM versions from image tags by removing the commit hash suffix (everything after the last '-')
# Format: v<version>-<commit_hash> -> v<version>
OSSFS_RPM_VERSION=$(echo "${OSSFS_IMAGE_TAG}" | sed 's/-[^-]*$//')
OSSFS2_RPM_VERSION=$(echo "${OSSFS2_IMAGE_TAG}" | sed 's/-[^-]*$//')

# Extract ALINAS and EFC versions from VERSIONS file
ALINAS_UTILS_VERSION=$(grep "^ALINAS_UTILS_VERSION=" "${VERSIONS_FILE}" | cut -d'=' -f2)
EFC_VERSION=$(grep "^EFC_VERSION=" "${VERSIONS_FILE}" | cut -d'=' -f2)
ALINAS_UTILS_ARM_VERSION=$(grep "^ALINAS_UTILS_ARM_VERSION=" "${VERSIONS_FILE}" | cut -d'=' -f2)
EFC_ARM_VERSION=$(grep "^EFC_ARM_VERSION=" "${VERSIONS_FILE}" | cut -d'=' -f2)

if [[ -z "${ALINAS_UTILS_VERSION}" || -z "${EFC_VERSION}" || -z "${ALINAS_UTILS_ARM_VERSION}" || -z "${EFC_ARM_VERSION}" ]]; then
    echo "Error: Failed to extract ALINAS/EFC versions from VERSIONS file" >&2
    exit 1
fi

# Export variables so they can be used by sourcing scripts
export OSSFS_IMAGE_TAG
export OSSFS2_IMAGE_TAG
export OSSFS_RPM_VERSION
export OSSFS2_RPM_VERSION
export ALINAS_UTILS_VERSION
export EFC_VERSION
export ALINAS_UTILS_ARM_VERSION
export EFC_ARM_VERSION

