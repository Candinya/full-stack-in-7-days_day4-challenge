FROM golang:alpine AS Builder

# 设置当前工作目录
WORKDIR /app

# 复制依赖声明文件
COPY go.mod .
COPY go.sum .

# 下载所需依赖
RUN go mod download

# 复制所有内容（包含项目代码）
COPY . .

# 构建可执行文件
RUN go build -o app .

FROM scratch AS Runner

WORKDIR /app

COPY --from=Builder /app/app /app/app

# 暴露默认端口
EXPOSE 1323/tcp

# 执行可执行文件
CMD ["/app/app"]