# 表示依赖 alpine 最新版
FROM alpine:latest
MAINTAINER Yapi<scncys.cn>
ENV VERSION 1.0
# 在容器根目录 创建一个 apps 目录
WORKDIR /apps
# 拷贝当前目录下 可以执行文件
COPY cmd/vending /apps/vending
COPY config.yml /apps/config.yml
# 设置时区为上海
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone
# 设置编码
ENV LANG C.UTF-8
# 暴露端口
EXPOSE 8080
# 运行golang程序的命令s
ENTRYPOINT ["/apps/vending","-conf", "./config.yml"]