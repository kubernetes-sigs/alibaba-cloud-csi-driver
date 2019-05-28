FROM ubuntu:18.10
LABEL maintainers="Alibaba Cloud Authors"
LABEL description="Alibaba Cloud CSI Plugin"

RUN apt update && apt upgrade -y && apt install -y ca-certificates file tzdata
COPY nsenter /nsenter

COPY plugin.csi.alibabacloud.com /bin/plugin.csi.alibabacloud.com
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /bin/plugin.csi.alibabacloud.com && chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
