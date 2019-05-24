FROM ubuntu:18.10
LABEL maintainers="Alibaba Cloud Authors"
LABEL description="Alibaba Cloud CSI Plugin"

RUN apt update && apt upgrade -y && apt install -y ca-certificates file tzdata

COPY plugin.csi.alibabacloud.com /bin/plugin.csi.alibabacloud.com
RUN chmod +x /bin/plugin.csi.alibabacloud.com

ENTRYPOINT ["/bin/plugin.csi.alibabacloud.com"]
