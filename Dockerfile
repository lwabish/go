FROM centos:7

ADD bin/lwabish-linux-amd64 /lwabish

RUN chmod +x /lwabish

ENTRYPOINT ["/lwabish"]
