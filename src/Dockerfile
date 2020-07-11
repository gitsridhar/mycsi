FROM registry.access.redhat.com/ubi8:latest
LABEL name="My CSI driver" \
      vendor="sridhar" \
      version="1.0.0" \
      release="1" \
      run='docker run my-driver' \
      summary="An implementation of CSI Plugin for the Sridhar product."\
      description="CSI Plugin for Sridhar"\
      maintainers="Sridhar"

RUN mkdir -p /var/lib/kubelet/plugins/my-csi-driver/controller

RUN mkdir -p /etc/config
COPY my-csi-driver /my-csi-driver
RUN chmod +x /my-csi-driver
RUN yum -y update
RUN yum -y install sudo
ENTRYPOINT ["/my-csi-driver"]
