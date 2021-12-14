FROM centos:7

RUN yum install -y epel-release
RUN yum install -y vim git gcc supervisor
RUN curl -O 'https://mirrors.ustc.edu.cn/golang/go1.13.linux-amd64.tar.gz'
RUN tar -C /usr/local -xzf go1.13.linux-amd64.tar.gz
RUN rm -f go1.13.linux-amd64.tar.gz
RUN ln -s /usr/local/go/bin/go /usr/bin/go

# RUN go get -v -u github.com/derekparker/delve/cmd/dlv
# RUN yum install -y strace tcpdump telnet lsof

RUN mkdir -p /go/src
ENV GOPATH "/go"
ENV GO111MODULE "on"
ENV GOPROXY "https://goproxy.cn,direct"
ENV GOFLAGS "-mod=vendor"
ENV GOPRIVATE "adgit.src.corp.qihoo.net"
ENV PATH $PATH:/root/go/bin

WORKDIR /go/src/mayo