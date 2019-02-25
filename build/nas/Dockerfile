FROM centos:7.4.1708
LABEL maintainers="Alibaba Cloud Authors"
LABEL description="Alibaba Cloud CSI Plugin"

RUN yum -y install nfs-utils && yum -y install epel-release && yum -y install jq && yum clean all

COPY nasplugin.csi.alibabacloud.com /bin/nasplugin.csi.alibabacloud.com
RUN chmod +x /bin/nasplugin.csi.alibabacloud.com

ENTRYPOINT ["/bin/nasplugin.csi.alibabacloud.com"]