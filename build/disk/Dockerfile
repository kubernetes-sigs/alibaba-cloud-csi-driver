FROM ubuntu:18.10
LABEL maintainers="Alibaba Cloud Authors"
LABEL description="Alibaba Cloud CSI Plugin"

RUN apt update && apt upgrade -y && apt install -y ca-certificates file tzdata

COPY diskplugin.csi.alibabacloud.com /bin/diskplugin.csi.alibabacloud.com
RUN chmod +x /bin/diskplugin.csi.alibabacloud.com

ENTRYPOINT ["/bin/diskplugin.csi.alibabacloud.com"]
