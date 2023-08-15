FROM registry.cn-hangzhou.aliyuncs.com/acs/alinux:3-update
RUN yum install -y util-linux && \
    curl -o /tmp/ossfs.rpm http://ack-csiplugin.oss-cn-hangzhou.aliyuncs.com/pre/ossfs/ossfs_1.88.1_centos8_x86_64.rpm && \
    rpm -i /tmp/ossfs.rpm && \
    yum clean all && rm -f /tmp/ossfs.rpm
ENTRYPOINT ["ossfs"]
