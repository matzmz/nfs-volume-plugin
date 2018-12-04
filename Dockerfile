FROM alpine

RUN apk update && apk add sshfs

RUN mkdir -p /run/docker/plugins /mnt/state /mnt/volumes

COPY nfs-volume-plugin.bin nfs-volume-plugin

COPY init.sh /

CMD ["/init.sh"]

