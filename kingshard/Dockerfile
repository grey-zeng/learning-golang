FROM golang:latest

# 安装最新版的kingshard
RUN git clone https://github.com/flike/kingshard.git $GOPATH/src/github.com/flike/kingshard
WORKDIR $GOPATH/src/github.com/flike/kingshard
RUN /bin/bash -c "source ./dev.sh" \
    && make

COPY ks.yaml $GOPATH/src/github.com/flike/kingshard/etc/

ADD . $GOPATH/src/github.com/flike/kingshard/

RUN ["chmod", "+x", "./start.sh"]

EXPOSE 9696 9797
