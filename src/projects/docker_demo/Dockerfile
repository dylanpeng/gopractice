# 指定基础镜像
FROM centos:latest
LABEL authors="docker-test"

# 设置工作目录
WORKDIR /app/gateway

# 复制当前目录下的所有文件到工作目录
COPY . .

# 声明容器运行时需要暴露的端口
EXPOSE 8991 8992

# 定参
ENTRYPOINT ["./main"]
