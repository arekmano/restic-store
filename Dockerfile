FROM alpine:latest

RUN apk add --update --no-cache \
    ca-certificates \
    fuse \
    openssh-client \
    bash
RUN wget https://github.com/restic/restic/releases/download/v0.9.6/restic_0.9.6_linux_amd64.bz2 && \
    bunzip2 restic_0.9.6_linux_amd64.bz2 && \
    mv restic_0.9.6_linux_amd64 /usr/bin/restic && \
    chmod 700 /usr/bin/restic

COPY ./restic-secret-store /usr/bin/restic-secret-store

ENTRYPOINT [ "/usr/bin/restic-secret-store" ]