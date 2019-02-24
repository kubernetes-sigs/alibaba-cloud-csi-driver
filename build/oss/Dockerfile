FROM centos:7.4.1708
LABEL maintainers="Alibaba Cloud Authors"
LABEL description="Alibaba Cloud CSI Plugin"

COPY ossfs_1.80.3_centos7.0_x86_64.rpm /root/ossfs_1.80.3_centos7.0_x86_64.rpm

RUN yum -y install nfs-utils && yum -y install epel-release && yum -y install jq && yum clean all && yum localinstall -y /root/ossfs_1.80.3_centos7.0_x86_64.rpm

COPY ossplugin.csi.aliyun.com /bin/ossplugin.csi.aliyun.com
RUN chmod +x /bin/ossplugin.csi.aliyun.com

ENTRYPOINT ["/bin/ossplugin.csi.aliyun.com"]