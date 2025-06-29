FROM alpine:3

ADD bin/lwabish-linux-amd64 /lwabish

RUN chmod +x /lwabish

ENTRYPOINT ["/lwabish"]
