FROM registry.cn-hangzhou.aliyuncs.com/acs/alinux:3-update
ARG TARGETPLATFORM
RUN --mount=target=/build /bin/sh /build/install-ossfs.sh
ADD entrypoint.sh /
ENTRYPOINT ["/entrypoint.sh"]
