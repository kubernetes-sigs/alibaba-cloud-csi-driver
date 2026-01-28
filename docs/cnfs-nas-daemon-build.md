### Introduction

cnfs-nas-daemon is a DaemonSet component specifically designed for Compute-Optimized CPFS, standard CPFS, and NAS storage systems to provide declarative mounting services. Its core design enables seamless integration of storage resources with containerized workloads through multi-layer containerization and dynamic configuration management. This document mainly introduces how to build the cnfs-nas-daemon image.

### Environment Preparation

#### Step 1: Install Docker Engine

1.1 Update System Packages

First, update your system's package index to ensure all installed software is up to date.

```shell
sudo yum update -y
```

1.2 Install Required Dependencies

Install tools such as yum-utils which help you manage YUM repositories.

```shell
sudo yum install -y yum-utils device-mapper-persistent-data lvm2
```

1.3 Add Docker's Official Repository

Add Docker's official repository to your system. This allows you to install and update Docker directly through the yum command.

```shell
sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
```

1.4 Install Docker Engine

Now install Docker Community Edition (CE) and its related components from the repository you just added. This command will install:
● docker-ce: Docker engine core.
● docker-ce-cli: Docker command-line tool.
● containerd.io: A standalone container runtime.
● docker-buildx-plugin: Official plugin for BuildKit that supports advanced build features.
● docker-compose-plugin: Official Docker Compose plugin.

```shell
sudo yum install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin -y
```

1.5 Start Docker Service

```shell
sudo systemctl start docker
```

1.6 Verify Docker Installation

Run a simple hello-world container to verify Docker is correctly installed and working.

```shell
sudo docker run hello-world
```

If the installation is successful, you will see a welcome message containing "Hello from Docker!".

#### Step 2: Enable and Verify BuildKit

BuildKit is an efficient, modern image building engine that is typically enabled by default in newer versions of Docker.

2.1 Check BuildKit Status

You can check if the BuildKit plugin is correctly installed and active using the docker buildx command.

```shell
docker buildx version
```

This command should display the installed buildx version information, indicating that BuildKit is working.

2.2 (Optional) Explicitly Enable BuildKit

Although usually enabled by default, you can also force-enable BuildKit for your current session by setting the environment variable DOCKER_BUILDKIT=1.

```shell
export DOCKER_BUILDKIT=1
```

Tip: To make this setting permanently effective for your user, you can add this line to your shell configuration file (such as ~/.bashrc or ~/.zshrc), then run source ~/.bashrc or re-login.

### Create Working Directory and Initialize Code Repository

```shell
mkdir -p /workspace/output && \
git clone https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver.git /workspace/output/
```

2. Build Environment Setup: Toolchain and Dependency Configuration

Enter the project root directory

```shell
cd /workspace/output/alibaba-cloud-csi-driver
```

### Building

Implement incremental builds and cross-platform compilation through Docker BuildKit:

```shell
export DOCKER_BUILDKIT=1
export CONTEXT_DIR=$(pwd)
export DOCKERFILE_PATH="build/mount-proxy/Dockerfile"
export IMAGE_REPO="registry.cn-hangzhou.aliyuncs.com/my-namespace/csi-alinas"
export IMAGE_TAG="v1.2.3-g1a2b3c4"
export TARGET_STAGE="alinas"
```

Related image_repo and image_tag need to be replaced with your own image address.

#### Start Build Process

```shell
docker buildx build \
  --file "${DOCKERFILE_PATH}" \
  --platform "linux/amd64" \
  --tag "${IMAGE_REPO}:${IMAGE_TAG}" \
  --target "${TARGET_STAGE}" \
  --cache-to=type=inline,mode=max \
  --cache-from=type=registry,ref="${IMAGE_REPO}:${IMAGE_TAG}" \
  --push \
  "${CONTEXT_DIR}"
```

### Deployment

Use kubectl command to update the cnfs-nas-daemon image in the cluster:

```shell
kubectl set image daemonset cnfs-nas-daemon -n cnfs-system mount-proxy=registry.cn-hangzhou.aliyuncs.com/my-namespace/csi-alinas:v1.2.3-g1a2b3c4
```