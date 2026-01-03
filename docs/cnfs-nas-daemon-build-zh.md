### 背景介绍
cnfs-nas-daemon 是一个 DaemonSet 组件，专为智算版 CPFS（Cloud Parallel File System）， 普通版 CPFS 及 NAS 存储系统提供声明式挂载服务。其设计核心在于通过多层容器化封装与动态配置管理，实现存储资源与容器化工作负载的无缝集成。 本文主要介绍如何构建 cnfs-nas-daemon 镜像

### 环境准备
#### 第 1 步：安装 Docker Engine
1.1 更新系统包
首先，更新您系统的包索引，确保所有已安装的软件都是最新版本。
```shell
sudo yum update -y
```

1.2 安装必要的依赖包
安装 yum-utils 等工具，它们可以帮助您管理 YUM 仓库。
```shell
sudo yum install -y yum-utils device-mapper-persistent-data lvm2
```

1.3 添加 Docker 官方软件源
将 Docker 的官方软件源添加到您的系统中。这样，您就可以直接通过 yum 命令安装和更新 Docker。
```shell
sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
```

1.4 安装 Docker Engine
现在，从您刚刚添加的软件源中安装 Docker 社区版（CE）及其相关组件。这个命令会同时安装：
● docker-ce: Docker 引擎核心。
● docker-ce-cli: Docker 命令行工具。
● containerd.io: 一个独立的容器运行时。
● docker-buildx-plugin: BuildKit 的官方插件，用于支持高级构建功能。
● docker-compose-plugin: Docker Compose 的官方插件。
```shell
sudo yum install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin -y
```

1.5 启动docker服务
```shell
sudo systemctl start docker
```
1.6 验证 Docker 安装
运行一个简单的 hello-world 容器来验证 Docker 是否已正确安装并可以正常工作。
```shell
sudo docker run hello-world
```

如果安装成功，您将看到一条包含 "Hello from Docker!" 的欢迎信息。

#### 第 2 步：启用并验证 BuildKit

BuildKit 是一个高效、现代化的镜像构建引擎，它在较新版本的 Docker 中通常是默认启用的。

2.1 确认 BuildKit 状态
您可以通过 docker buildx 命令来检查 BuildKit 插件是否已正确安装并处于活动状态。
```shell
docker buildx version
```

此命令应显示已安装的 buildx 版本信息，表明 BuildKit 正在工作。

2.2 (可选) 显式启用 BuildKit
虽然通常是默认开启的，但您也可以通过设置环境变量 DOCKER_BUILDKIT=1 来强制、显式地为您的当前会话启用 BuildKit。
```shell
export DOCKER_BUILDKIT=1
```

提示：为了让这个设置永久对您的用户生效，可以将这行命令添加到您的 shell 配置文件中（如 ~/.bashrc 或 ~/.zshrc），然后运行 source ~/.bashrc 或重新登录。


### 创建工作目录并初始化代码仓库
```shell
mkdir -p /workspace/output && \
git clone https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver.git /workspace/output/
```
2. 构建环境准备：工具链与依赖配置
进入项目根目录
```shell
cd /workspace/output/alibaba-cloud-csi-driver
```


### 构建
通过 Docker BuildKit 实现增量构建与多平台交叉编译：
```shell
export DOCKER_BUILDKIT=1
export CONTEXT_DIR=$(pwd)
export DOCKERFILE_PATH="build/mount-proxy/Dockerfile"
export IMAGE_REPO="registry.cn-hangzhou.aliyuncs.com/my-namespace/csi-alinas"
export IMAGE_TAG="v1.2.3-g1a2b3c4"
export TARGET_STAGE="alinas"
```
相关 image_repo 和 image_tag 需要替换成自己的镜像地址

#### 启动构建流程
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

### 部署

使用 kubectl 命令 更新 集群里面 cnfs-nas-daemon 镜像

```shell
kubectl set image daemonset cnfs-nas-daemon -n cnfs-system mount-proxy=registry.cn-hangzhou.aliyuncs.com/my-namespace/csi-alinas:v1.2.3-g1a2b3c4
```
