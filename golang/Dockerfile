FROM golang:1.20.1

LABEL maintainer="lxy <longxinyu2021@163.com>"

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

# 设置工作目录
WORKDIR /usr/src/code

# 安装必要的软件包和依赖包
# USER root
# RUN sed -i 's/deb.debian.org/mirrors.tuna.tsinghua.edu.cn/' /etc/apt/sources.list && \
#     sed -i 's/security.debian.org/mirrors.tuna.tsinghua.edu.cn/' /etc/apt/sources.list && \
#     sed -i 's/security-cdn.debian.org/mirrors.tuna.tsinghua.edu.cn/' /etc/apt/sources.list && \
#     apt-get update && \
#     apt-get upgrade -y && \
#     apt-get install -y --no-install-recommends \
#     curl \
#     zip \
#     unzip \
#     git \
#     vim

# $GOPATH/bin 添加到环境变量中
ENV PATH $GOPATH/bin:$PATH

# 清理垃圾
# USER root
# RUN apt-get clean && \
#     rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* && \
#     rm /var/log/lastlog /var/log/faillog

# RUN go build -o ./build/main ./main.go
EXPOSE 8000
# ENTRYPOINT  ["go build -o ./build/main ./main.go"]
# CMD ["./build/main"]