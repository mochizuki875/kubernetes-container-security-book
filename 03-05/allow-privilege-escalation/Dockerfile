FROM ubuntu:22.04
RUN apt update && \
    apt install -y sudo && \
    useradd -m -U -u 1000 user01 && \
    groupadd wheel && \
    usermod -aG wheel user01 && \
    echo "%wheel ALL=NOPASSWD: ALL" >> /etc/sudoers && \
    apt clean && \
    rm -rf /var/lib/apt/lists/*

USER user01
WORKDIR /home/user01
CMD ["/bin/sh", "-c", "while :; do sleep 10; done"]