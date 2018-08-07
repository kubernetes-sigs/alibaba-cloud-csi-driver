FROM centos:7.4.1708
LABEL maintainers="AliCloud Authors"
LABEL description="AliCloud CSI Plugin"

RUN yum -y install nfs-utils && yum -y install epel-release && yum -y install jq && yum clean all

COPY csi-nasplugin /bin/csi-nasplugin
RUN chmod +x /bin/csi-nasplugin

ENTRYPOINT ["/bin/csi-nasplugin"]