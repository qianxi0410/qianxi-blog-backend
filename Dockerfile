# 拉取 Go 语言最新的基础镜像
FROM golang:1.17 as Builder


# 设置 GOPROXY 环境变量
ENV GOPROXY="https://goproxy.cn"
# 开启go.mod
ENV GO111MODULE on

# 在容器内设置 /app 为当前工作目录
WORKDIR /app

# 把文件复制到当前工作目录
COPY . .

# 下载全部依赖项
RUN go mod download

# 编译项目
RUN GOOS=linux CGO_ENABLED=0 go build -o blog_api service/blog/api/blog.go

FROM alpine:latest

# 设置时区
RUN apk update && apk add tzdata

# 暴露 9998 端口
EXPOSE 9998

WORKDIR /app

COPY --from=Builder /app/blog_api .
COPY --from=Builder /app/service/blog/api/etc/blog-api.yaml ./etc/

# 执行可执行文件
CMD ["./blog_api", "-f", "./etc/blog-api.yaml"]