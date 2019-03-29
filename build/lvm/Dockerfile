FROM ubuntu:18.10
LABEL maintainers="Alibaba Cloud Authors"
LABEL description="Alibaba Cloud CSI Plugin"

RUN apt update && apt install -y ca-certificates  file
COPY nsenter /nsenter

COPY ./zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
COPY lvmplugin.csi.alibabacloud.com /bin/lvmplugin.csi.alibabacloud.com
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /bin/lvmplugin.csi.alibabacloud.com && chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
