# 解决镜像太大的问题
FROM alpine:latest

# 参数设置
# 项目名称
ARG appName=proproc
# 设置项目环境变量
ENV SYS_NAME $SYS_NAME
ENV SYS_NODE $SYS_NODE
ENV SYS_ENV  $SYS_ENV
# 暴露端口
# 需要与配置文件一致
EXPOSE 9040

# 安装 /bin/bash 方便调试
RUN echo "https://mirror.tuna.tsinghua.edu.cn/alpine/v3.4/main/" > /etc/apk/repositories
RUN apk update \
        && apk upgrade \
        && apk add --no-cache bash \
        bash-doc \
        bash-completion \
        && rm -rf /var/cache/apk/* \
        && /bin/bash

# need install glib, if ont get error: exec user process caused: no such file or director
ARG GLIBC_MIRROR=https://github.com/sgerrand/alpine-pkg-glibc
ARG GLIBC_VER=2.31-r0
RUN  apk add --no-cache --virtual=.build-deps curl && \
 curl -Lo /etc/apk/keys/sgerrand.rsa.pub "${GLIBC_MIRROR}/master/sgerrand.rsa.pub" && \
 curl -Lo /glibc.apk "${GLIBC_MIRROR}/releases/download/${GLIBC_VER}/glibc-${GLIBC_VER}.apk" && \
 curl -Lo /glibc-bin.apk "${GLIBC_MIRROR}/releases/download/${GLIBC_VER}/glibc-bin-${GLIBC_VER}.apk" \
 && ls -l \
 && apk add --no-cache --allow-untrusted \
   /glibc.apk \
   /glibc-bin.apk \
 && rm /glibc.apk \
 && rm /glibc-bin.apk

# 设置目录
WORKDIR /app/preproc

# 把app文件从 "宿主机" 中拷贝到本级的当前目录
COPY ./${appName} .
COPY ./build/docs/config ./config

# 项目默认启动命令
CMD ["./preproc"]
