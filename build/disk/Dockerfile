FROM ubuntu:18.10
LABEL maintainers="Alibaba Cloud Authors"
LABEL description="Alibaba Cloud CSI Plugin"

RUN apt update && apt install -y ca-certificates  file

COPY ./zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
COPY csi-diskplugin /bin/csi-diskplugin
RUN chmod +x /bin/csi-diskplugin

ENTRYPOINT ["/bin/csi-diskplugin"]
