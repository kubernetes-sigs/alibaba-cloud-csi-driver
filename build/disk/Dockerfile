FROM ubuntu:18.10
LABEL maintainers="Alibaba Cloud Authors"
LABEL description="Alibaba Cloud CSI Plugin"

RUN apt update && apt install -y ca-certificates  file

COPY ./zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
COPY diskplugin.csi.alibabacloud.com /bin/diskplugin.csi.alibabacloud.com
RUN chmod +x /bin/diskplugin.csi.alibabacloud.com

ENTRYPOINT ["/bin/diskplugin.csi.alibabacloud.com"]
