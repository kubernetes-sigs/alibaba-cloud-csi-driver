FROM registry.cn-hangzhou.aliyuncs.com/acs/alinux:3-update
ARG TARGETPLATFORM
COPY install-ossfs.sh .
RUN /bin/sh install-ossfs.sh
ENTRYPOINT ["ossfs"]
